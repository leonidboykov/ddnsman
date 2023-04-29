package ddnsman

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

type stringDuration time.Duration

func (d stringDuration) String() string {
	return time.Duration(d).String()
}

func (d *stringDuration) UnmarshalJSON(data []byte) error {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("parse json: %w", err)
	}
	switch val := v.(type) {
	case float64:
		*(*time.Duration)(d) = time.Duration(val)
	case string:
		var err error
		*(*time.Duration)(d), err = time.ParseDuration(val)
		if err != nil {
			return fmt.Errorf("parse time: %w", err)
		}
	default:
		return errors.New("invalid duration")
	}
	return nil
}

type Configuration struct {
	Interval      stringDuration          `json:"interval"`
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

func LoadConfiguration(name string) (*Configuration, error) {
	var config Configuration
	f, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("unable to read configuration file: %w", err)
	}
	if err := json.Unmarshal(f, &config); err != nil {
		return nil, fmt.Errorf("unable to parse configuration file: %w", err)
	}
	if config.Interval == 0 {
		config.Interval = stringDuration(5 * time.Minute)
	}
	processConfiguration(&config)
	return &config, nil
}

func processConfiguration(config *Configuration) error {
	for idx, setting := range config.Settings {
		provider, err := newProvider(setting.Provider.Name, setting.Provider.Settings)
		if err != nil {
			return fmt.Errorf("unable to create a new provider: %w", err)
		}
		config.Settings[idx].provider = provider
		if !strings.HasSuffix(config.Settings[idx].Domain, ".") {
			config.Settings[idx].Domain += "."
		}
	}
	return nil
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
