/*
Copyright © 2025 Dexter Dixon dexterdixon561@gmail.com
*/
package cmd

import (
	"context"
	"fmt"
	"go-stuff/msg-util/internal"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	port int16
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Creates a local Kafka client server for testing",
	Long: `Used to create a local Kafka client that can be used to test Kafka communications.
			The server has two endpoints v1/consumer and v1/producer. These endpoints can be
			used to create, list, update, or delete kafka consumers and producers respectivly:`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		mux := http.NewServeMux()

		//Producer endpoints
		mux.HandleFunc("GET /v1/producers", internal.GetProducers)
		mux.HandleFunc("POST /v1/producers", internal.CreateProducers)
		mux.HandleFunc("PUT /v1/producers", internal.UpdateProducers)
		mux.HandleFunc("DELETE /v1/producers", internal.DeleteProducers)

		//Consumer endpoints
		mux.HandleFunc("GET /v1/consumers", internal.GetConsumers)
		mux.HandleFunc("POST /v1/consumers", internal.CreateConsumers)
		mux.HandleFunc("PUT /v1/consumers", internal.UpdateConsumers)
		mux.HandleFunc("PATCH /v1/consumers", internal.AddConsumerTopic)
		mux.HandleFunc("DELETE /v1/consumers", internal.DeleteConsumers)

		//Utility endpoints
		mux.HandleFunc("GET /v1/healthz", internal.GetHealth)
		mux.HandleFunc("GET /v1/echo/{msg...}", internal.Echo)
		mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		}))

		addr := fmt.Sprintf(":%d", port)
		srv := &http.Server{
			Addr:         addr,
			Handler:      internal.Logging(mux),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		}

		go func() {
			fmt.Printf("listening on http://localhost%s\n", srv.Addr)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				fmt.Printf("Server error: %v\n", err)
			}
		}()

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		<-stop

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("graceful shutdown error: %v", err)
		}
		fmt.Println("server stopped")

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Int16VarP(&port, "port", "p", 8080, "Port for the server to bind to")
}
