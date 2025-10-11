/*
Copyright © 2025 Dexter Dixon dexterdixon561@gmail.com
*/
package cmd

import (
	"context"
	"fmt"
	"go-stuff/msg-util/internal"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	groupID string
	topics  []string
	brokers []string
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Creates a Kafka consumer to read messages from a topic",
	Long:  `Used to create a Kafka consumer that can read messages from one or more topics within a Kafka cluster.:`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Consumer Yum Yum Yum :O")

		fmt.Printf("Configuring TLS")
		restConfig, err := internal.GetRestConfig("", inCluster)
		if err != nil {
			panic(err)
		}
		clientset, err := internal.GetClientset(restConfig)
		if err != nil {
			panic(err)
		}
		secret, err := internal.GetSecret(clientset, namespace, args[0])
		if err != nil {
			panic(err)
		}

		kafkaConfig, err := internal.ConfigTLS(nil, secret.Data)
		if err != nil {
			fmt.Printf("Error configuring TLS: %v\n", err)
			panic(err)
		}

		fmt.Printf("Configuring Consumer")
		consumer := internal.CreateKafkaConsumer(brokers, groupID, kafkaConfig)
		defer func() { _ = consumer.Close() }()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Stop on interrupt
		go func() {
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
			<-sig
			cancel()
		}()

		handler := internal.ConsumerGroupHandler{}

		for {
			if err := consumer.Consume(ctx, topics, handler); err != nil {
				fmt.Printf("consumer error: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)

	// group ID
	consumerCmd.Flags().StringVar(&groupID, "group", "example-group", "consumer group id")

	// brokers can be provided as: --brokers host1:9092,host2:9092
	consumerCmd.Flags().StringSliceVar(&brokers, "brokers", []string{"localhost:9092"}, "comma-separated list of Kafka brokers")

	// topics can be provided as: --topics a,b
	consumerCmd.Flags().StringSliceVarP(&topics, "topics", "t", []string{}, "comma-separated list of topics to consume")
}
