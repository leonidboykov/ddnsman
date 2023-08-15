package ddnsman

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/miekg/dns"
)

type Configuration struct {
	Interval      time.Duration           `json:"interval"`
	Settings      []Setting               `json:"settings"`
	ShoutrrrAddrs []ShoutrrrNotifications `json:"shoutrrr_notifications"`
}

type Setting struct {
	Domain   string   `json:"domain"`
	Records  []string `json:"records"`
	Provider struct {
		Name     string          `json:"name"`
		Settings json.RawMessage `json:"settings"`
	} `json:"provider"`
	provider Provider
}

type ShoutrrrNotifications struct {
	URL      string            `json:"url"`
	Settings map[string]string `json:"settings"`
}

func (c Configuration) shoutrrrURLs() ([]string, error) {
	urls := make([]string, 0, len(c.ShoutrrrAddrs))
	for _, addr := range c.ShoutrrrAddrs {
		u, err := url.Parse(addr.URL)
		if err != nil {
			return nil, fmt.Errorf("parse url: %w", err)
		}
		vals := url.Values{}
		for k, v := range addr.Settings {
			vals.Add(k, v)
		}
		u.RawQuery = vals.Encode()
		urls = append(urls, u.String())
	}
	return urls, nil
}

func LoadConfiguration() (*Configuration, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "ddnsman.json"
	}

	config, err := readConfiguration(configPath)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	if config.Interval == 0 {
		config.Interval = 5 * time.Minute
	}
	processConfiguration(config)
	return config, nil
}

func readConfiguration(configPath string) (*Configuration, error) {
	config := new(Configuration)
	f, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read configuration file: %w", err)
	}

	if err := yaml.UnmarshalWithOptions(f, config, yaml.UseJSONUnmarshaler()); err != nil {
		return nil, fmt.Errorf("parse configuration file: %w", err)
	}
	return config, nil
}

func processConfiguration(config *Configuration) error {
	for idx, setting := range config.Settings {
		provider, err := newProvider(setting.Provider.Name, setting.Provider.Settings)
		if err != nil {
			return fmt.Errorf("unable to create a new provider: %w", err)
		}
		config.Settings[idx].provider = provider
		config.Settings[idx].Domain = dns.Fqdn(config.Settings[idx].Domain)
	}
	return nil
}
