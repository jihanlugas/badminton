package cmd

import (
	"context"
	"fmt"
	"github.com/jihanlugas/badminton/config"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/router"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Long:  "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func runServer() {
	var err error

	r := router.Init()

	_, closeConn := db.GetConnection()
	defer closeConn()

	if err != nil {
		r.Logger.Fatal(err)
	}

	// Start server
	go func() {
		var err error
		err = r.Start(fmt.Sprintf(":%s", config.ListenTo.Port))
		if err != nil && err != http.ErrServerClosed {
			r.Logger.Fatal("Shutting down the server ", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Shutdown(ctx)
	if err != nil {
		r.Logger.Fatal(err)
	}
}
