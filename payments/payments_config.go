package payments

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host       string
	Database   string
	Collection string
	APIPort    int64
}

func (c *Config) LoadConfig() {
	if _, err := toml.DecodeFile("~/.payments.toml", &c); err != nil {
		log.Println("Configuration file not found in home directory...")
		if _, err := toml.DecodeFile("/etc/payments.toml", &c); err != nil {
			log.Println("Configuration file not found in /etc directory...")
		} else {
			return
		}
	} else {
		return
	}
	log.Println("Unable to find any configuration file in standard locations. Loading defaults.")
	c.Host = "localhost"
	c.Database = "payments_api"
	c.Collection = "payments"
	c.APIPort = 9000
}
