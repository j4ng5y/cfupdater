package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is a struct that holds all data necessary to run cfupdater
//
// Example YAML file:
// ---
// cloudflare:
//   general:
// 	   api_token: "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*"
//   dns_record:
// 	   zone_id: "abcdefghijklmnopqrstuvwxyz1234567890"
// 	   record_name: "example.mydomain.com"
// ipinfo:
//   general:
//     api_token: "abcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*"
type Config struct {
	CloudFlare struct {
		General struct {
			APIToken string `yaml:"api_token"`
		} `yaml:"general"`
		DNSRecord struct {
			ZoneID     string `yaml:"zone_id"`
			RecordName string `yaml:"record_name"`
		} `yaml:"dns_record"`
	} `yaml:"cloudflare"`
	IPInfo struct {
		General struct {
			APIToken string `yaml:"api_token"`
		} `yaml:"general"`
	} `yaml:"ipinfo"`
}

// New creates a new Config instance from a provided file
//
// Arguments:
//     filePath (string): The path the configuration file to parse
//
// Returns:
//     (*Config): A populated Config instance, or nil if there was an error
//     (error):   An error if one exists, nil otherwise
func New(filePath string) (*Config, error) {
	config := new(Config)

	s, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		return nil, fmt.Errorf("%s is a directory, not a normal file", filePath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, config); err != nil {
		return nil, err
	}

	return config, nil
}

// Write creates a configuration file to be consumed
//
// Arguments:
//    None
//
// Returns:
//    (error): An error if one exists, nil otherwise
func (C *Config) Write() error {
	defaultConfigFile := "/opt/cfupdater/config/config.yaml"

	f, err := os.OpenFile(defaultConfigFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	defer f.Close()

	y, err := yaml.Marshal(C)
	if err != nil {
		return err
	}

	_, err = f.Write(y)
	if err != nil {
		return err
	}

	return nil
}
