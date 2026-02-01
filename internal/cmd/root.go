package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mbndr/figlet4go"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v1gn35h7/transaction-service/internal/config"
	"github.com/v1gn35h7/transaction-service/internal/datastore/postgresql"
	"github.com/v1gn35h7/transaction-service/internal/logging"
	"github.com/v1gn35h7/transaction-service/internal/service"
	httptransport "github.com/v1gn35h7/transaction-service/internal/transport/http"
)

var (
	configPath string
	ascii      = figlet4go.NewAsciiRender()
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "transaction-service",
		Short: "Transaction service",
		Long:  "-----------------------------------------------",
		Run: func(cmd *cobra.Command, args []string) {
			printLogo()

			//Bootstrap server
			bootStrapServer()
		},
	}

	// Bind cli flags
	rootCmd.PersistentFlags().StringVar(&configPath, "conf", "", "config file path")

	return rootCmd
}

func printLogo() {
	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		// Colors can be given by default ansi color codes...
		figlet4go.ColorMagenta,
	}
	renderStr, _ := ascii.RenderOpts("Pismo", options)
	fmt.Print(renderStr)
}

func bootStrapServer() {
	//Logger setup
	logger := logging.Logger()

	// Init read config
	config.ReadConfig(configPath, logger)

	// Setup datbase
	ds, err := postgresql.NewDatastore(config.LoadPostgresqlConfig(), logger)
	if err != nil {
		logger.Error(err, "Failed to create Postgresql Datastore")
		os.Exit(1)
	}

	// Create CloudBees Train Booking Service instance
	srvc := service.New(ds, logger)

	// Mux Routes
	r := httptransport.MakeHandlers(srvc)

	port := viper.GetString("ts.server.port")

	// Start Server
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	var g group.Group
	{
		g.Add(func() error {
			logger.Info("transport", "HTTP", "addr", "localhost:8080")
			return srv.ListenAndServe()
		}, func(error) {
			srv.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received cancel signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	logger.Error(g.Run(), "Exit service")
}
