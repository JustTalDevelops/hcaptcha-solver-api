package main

import (
	"crypto/rand"
	"encoding/base64"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math"
	"os"
)

// Config represents a configuration file.
type Config struct {
	Port int `yaml:"port"`
	AuthorizationHeader string `json:"authorization_header"`
	SolveTimeout int `json:"solve_timeout"`
}

// loadConfig loads the configuration file and returns a Config.
func loadConfig() (config Config) {
	b, err := os.ReadFile("config.yml")
	if err != nil {
		config.Port = 80
		config.AuthorizationHeader = generateAuthorizationHeader()
		config.SolveTimeout = 60

		b, err = yaml.Marshal(config)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile("config.yml", b, 0777)
		if err != nil {
			panic(err)
		}

		return
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		panic("the configuration file passed isn't formatted correctly: " + err.Error())
	}

	return
}

func generateAuthorizationHeader() string {
	buff := make([]byte, int(math.Round(float64(24)/1.33333333333)))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:1] // strip 1 extra character we get from odd length results
}