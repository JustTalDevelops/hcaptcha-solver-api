package main

import (
	"crypto/rand"
	"encoding/hex"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math"
	"os"
)

// Config represents a configuration file.
type Config struct {
	Port                int    `yaml:"port"`
	AuthorizationHeader string `yaml:"authorization_header"`
	SolveTimeout        int    `yaml:"solve_timeout"`
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
	buff := make([]byte, int(math.Round(float64(24)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str[:24]
}
