package main

import (
	"errors"
	"log"
	"os"
)

type config struct {
	Port        string // yes, is properly a int but we always use it as a str so...
	HipchatKey  string
	HipchatRoom string
}

func loadConfig(c config) (config, error) {
	inError := false

	c.Port = os.Getenv("PORT")
	if c.Port == "" {
		inError = true
		log.Println("config error: PORT not set")
	}

	c.HipchatKey = os.Getenv("HIPCHAT_KEY")
	if c.HipchatKey == "" {
		inError = true
		log.Println("config error: HIPCHAT_KEY not set")
	}

	c.HipchatRoom = os.Getenv("HIPCHAT_ROOM")
	if c.HipchatRoom == "" {
		inError = true
		log.Println("config error: HIPCHAT_ROOM not set")
	}

	if inError {
		return c, errors.New("config env incomplete")
	}

	return c, nil
}
