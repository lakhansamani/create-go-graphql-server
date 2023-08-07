package main

import (
	"github.com/rs/zerolog/log"

	"github.com/lakhansamani/create-go-graphql-server/cmd"
)

var version string

func main() {
	cmd.SetVersion(version, "")
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to run root command")
	}
}
