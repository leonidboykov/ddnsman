package ddnsman

import (
	"context"
	"fmt"
	"log/slog"
	"net/netip"
	"time"

	"github.com/containrrr/shoutrrr/pkg/router"
	externalip "github.com/glendc/go-external-ip"
	"github.com/libdns/libdns"
	"golang.org/x/sync/errgroup"
)

type Updater struct {
	config    *Configuration
	consensus *externalip.Consensus
	sender    *router.ServiceRouter
	currentIP netip.Addr
}

func New(config *Configuration) (*Updater, error) {
	// Use consensus to acquire external IP.
	consensus := externalip.DefaultConsensus(nil, nil)
	consensus.UseIPProtocol(4)

	// Use shoutrrr to send notifications.
	urls, err := config.shoutrrrURLs()
	if err != nil {
		return nil, err
	}
	sender, err := router.New(nil, urls...)
	if err != nil {
		return nil, fmt.Errorf("create shoutrrr router: %w", err)
	}

	return &Updater{
		config:    config,
		consensus: consensus,
		sender:    sender,
	}, nil
}

func (u *Updater) Start(ctx context.Context) error {
	slog.Info("start DDNS updater", slog.Duration("interval", u.config.Interval))
	u.sender.Send(fmt.Sprintf("Watching records every %s", u.config.Interval), nil)
	if err := u.process(ctx); err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(u.config.Interval))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := u.process(ctx); err != nil {
				u.sender.Send(fmt.Sprintf("Error occured: `%s`. Shutting down.", err), nil)
				return err
			}
		case <-ctx.Done():
			u.sender.Send("Shutting down.", nil)
			slog.Info("shutting down")
			return nil
		}
	}
}

func (u *Updater) process(ctx context.Context) error {
	externalIP, err := u.consensus.ExternalIP()
	if err != nil {
		return fmt.Errorf("fetch external IP: %w", err)
	}
	externalIPAddr, err := netip.ParseAddr(externalIP.String())
	if err != nil {
		return fmt.Errorf("parse external IP: %w", err)
	}
	if externalIPAddr == u.currentIP {
		// Skip check if external IP hasn't changed.
		return nil
	}

	wg, ctx := errgroup.WithContext(ctx)
	for _, setting := range u.config.Settings {
		wg.Go(func() error {
			return u.checkRecord(ctx, externalIPAddr, setting)
		})
	}
	if err := wg.Wait(); err != nil {
		return err
	}
	u.currentIP = externalIPAddr
	return nil
}

func (u *Updater) checkRecord(ctx context.Context, externalIP netip.Addr, setting Setting) error {
	logger := slog.Default().With("provider", setting.Provider.Name)

	providerRecords, err := setting.provider.GetRecords(ctx, setting.Domain)
	if err != nil {
		return fmt.Errorf("getting records for zone %q: %w", setting.Domain, err)
	}
	var records []libdns.Record
	for _, providerRecord := range providerRecords {
		for _, targetRecord := range setting.Records {
			providerAddressRecord, ok := providerRecord.(libdns.Address)
			if !ok {
				continue
			}

			providerName := libdns.RelativeName(providerAddressRecord.Name, setting.Domain)

			if providerName == targetRecord && providerAddressRecord.IP != externalIP {
				logger.Info("IP address mismatch",
					slog.String("record", providerName),
					slog.String("record IP", providerAddressRecord.IP.String()),
					slog.Any("external IP", externalIP),
				)
				providerAddressRecord.IP = externalIP
				records = append(records, providerAddressRecord)
			}
		}
	}
	if len(records) == 0 {
		return nil
	}
	u.sender.Send(fmt.Sprintf("Setting IP `%s` for zone `%s`.", externalIP, setting.Domain), nil)
	if _, err := setting.provider.SetRecords(ctx, setting.Domain, records); err != nil {
		return fmt.Errorf("setting records: %w", err)
	}
	return nil
}
