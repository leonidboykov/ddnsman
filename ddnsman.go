package ddnsman

import (
	"context"
	"fmt"
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
	currentIP string // strings are comparable.
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
			return nil
		}
	}
}

func (u *Updater) process(ctx context.Context) error {
	externalIP, err := u.consensus.ExternalIP()
	if err != nil {
		return fmt.Errorf("unable to fetch external ip: %w", err)
	}
	if externalIP.String() == u.currentIP {
		// Skip check if external IP hasn't changed.
		return nil
	}

	wg, ctx := errgroup.WithContext(ctx)
	for _, setting := range u.config.Settings {
		setting := setting
		wg.Go(func() error {
			return u.checkRecord(ctx, externalIP.String(), setting)
		})
	}
	if err := wg.Wait(); err != nil {
		return err
	}
	u.currentIP = externalIP.String()
	return nil
}

func (u *Updater) checkRecord(ctx context.Context, externalIP string, setting Setting) error {
	providerRecords, err := setting.provider.GetRecords(ctx, setting.Domain)
	if err != nil {
		return fmt.Errorf("getting records for zone %q: %w", setting.Domain, err)
	}
	var records []libdns.Record
	for _, providerRecord := range providerRecords {
		for _, targetRecord := range setting.Records {
			providerName := libdns.RelativeName(providerRecord.Name, setting.Domain)

			if providerName == targetRecord && providerRecord.Type == "A" && providerRecord.Value != externalIP {
				providerRecord.Value = externalIP
				records = append(records, providerRecord)
			}
		}
	}
	u.sender.Send(fmt.Sprintf("Setting IP `%s` for zone `%s`.", externalIP, setting.Domain), nil)
	if _, err := setting.provider.SetRecords(ctx, setting.Domain, records); err != nil {
		return fmt.Errorf("setting records: %w", err)
	}
	return nil
}
