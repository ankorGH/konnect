package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Config struct {
	d map[string]interface{}
}

func (c *Config) GetString(path string) string {
	if str, ok := c.d[path]; ok {
		return str.(string)
	}
	return ""
}

func (c *Config) GetTime(path string) time.Time {
	t, err := time.Parse(time.RFC3339, c.GetString(path))
	if err != nil {
		return time.Now()
	}
	return t
}

func (c *Config) GetDuration(path string) time.Duration {
	t, err := time.ParseDuration(c.d[path].(string))
	if err != nil {
		return 0
	}
	return t
}

func (c *Config) GetStringSlice(path string) []string {
	if v, ok := c.d[path]; ok {
		if sslice, ok := v.([]string); ok {
			return sslice
		}

		if vslice, ok := v.([]interface{}); ok {
			slice := make([]string, len(vslice))
			for idx, item := range vslice {
				slice[idx] = fmt.Sprintf("%v", item)
			}
			return slice
		}
	}

	return []string{}
}

func (c *Config) GetInt(path string) int {
	if v, ok := c.d[path].(int); ok {
		return v
	}
	return 0
}

func GetConfig() (*Config, error) {
	directory, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("getting current directory failed: %v", err)
	}
	configPath := directory + os.Getenv("CONFIG_PATH")
	fi, err := os.Stat(configPath)
	if err != nil {
		return nil, fmt.Errorf("file not found: %v", err)
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("config is not a directory")
	}
	f, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	data := make([]byte, 0)
	reader := bufio.NewReader(f)
	for {
		d := make([]byte, 1024)
		size, err := reader.Read(d)
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
		}
		data = append(data, d[:size]...)
	}
	var cfg map[string]interface{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &Config{
		d: cfg,
	}, nil
}
