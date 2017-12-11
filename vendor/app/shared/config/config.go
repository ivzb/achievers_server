package config

import (
	"app/shared/database"
	"app/shared/token"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// *****************************************************************************
// Application Settings
// *****************************************************************************

// configuration contains the application settings
type Config struct {
	Database database.Info `json:"Database"`
	// Email    email.SMTPInfo  `json:"Email"`
	// Server   server.Server   `json:"Server"`
	Token token.TokenInfo `json:"Token"`
}

// ParseJSON unmarshals bytes to structs
func (c *Config) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// Load the JSON config file
func Load(file string) (*Config, error) {
	var err error
	var input = io.ReadCloser(os.Stdin)
	if input, err = os.Open(file); err != nil {
		return nil, err
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	// Parse the config
	if err := conf.ParseJSON(jsonBytes); err != nil {
		return nil, errors.New(fmt.Sprintf("Could not parse %q: %v", file, err))
	}

	return conf, nil
}
