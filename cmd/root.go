package cmd

import (
	"github.com/lakhansamani/create-go-graphql-server/internal/server"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	// RootCmd is the root (and only) command of this service
	// TODO change this to your docker image name
	RootCmd = &cobra.Command{
		Use:   "api",
		Short: "The api Service",
		Run:   runRootCmd,
	}

	rootArgs struct {
		// Version of the service
		version string
		// Log level
		logLevel string
		// Server configuration
		server server.Config
	}
)

// SetVersion stores the given version
func SetVersion(version, build string) {
	rootArgs.version = version
}

func init() {
	f := RootCmd.Flags()
	// Logging flags
	f.StringVar(&rootArgs.logLevel, "log-level", "debug", "Minimum log level")
	// Server flags
	f.IntVar(&rootArgs.server.Port, "http-port", 3000, "Port to listen on for HTTP requests")
}

func runRootCmd(cmd *cobra.Command, args []string) {
	// Setup logging
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.MessageFieldName = "msg"
	zerolog.TimestampFieldName = "time"
	log := zerolog.New(cmd.OutOrStderr()).With().Timestamp().Logger()
	// Set log level
	logLevel, err := zerolog.ParseLevel(rootArgs.logLevel)
	if err != nil {
		// Default to debug if the log level is invalid
		logLevel = zerolog.DebugLevel
	}
	log.Level(logLevel)
	// Setup server
	srv, err := server.New(log, rootArgs.server)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create server")
	}
	// Run server
	ctx := cmd.Context()
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error { return srv.Run(ctx) })
	if err := g.Wait(); err != nil {
		log.Fatal().Err(err).Msg("Failed to run server")
	}

}
