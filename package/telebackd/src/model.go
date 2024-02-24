package main

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	AdminID  int64  `json:"admin_id" env:"ADMIN_ID" env-default:"0"`
	BotToken string `json:"bot_token" env:"BOT_TOKEN" env-default:""`
	Debug    bool   `json:"debug" env:"DEBUG" env-default:"false"`
}

func (c *config) Load() error {
	if err := cleanenv.ReadEnv(c); err != nil {
		return fmt.Errorf("config load error: %s", err)
	}

	if c.Debug {
		log.Printf("Config: %+v", c)
	}

	if c.BotToken == "" {
		return fmt.Errorf("missing required fields: BOT_TOKEN")
	}

	if c.AdminID == 0 {
		return fmt.Errorf("missing required fields: ADMIN_ID")
	}

	return nil
}
