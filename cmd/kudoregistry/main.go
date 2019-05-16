package main

import (
	"github.com/fabianbaier/kudoregistry/pkg/kudoregistry/cmd"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Errorf("Error loading .env file")
	}

	if err := cmd.NewKudoRegistryCmd().Execute(); err != nil {
		os.Exit(-1)
	}
}
