/*
Copyright © 2025 Dexter Dixon dexterdixon561@gmail.com
*/
package cmd

import (
	"fmt"
	"go-stuff/msg-util/internal"

	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
)

var (
	msg             string
	producerBrokers []string
)

// producerCmd represents the producer command
var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Sends a message to a Kafka topic",
	Long:  `Use this command to send a message to a specified Kafka topic.:`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Producer called 🗣")

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

		producer, err := internal.CreateKafkaProducer(producerBrokers, kafkaConfig)
		if err != nil {
			fmt.Printf("Error creating producer: %v\n", err)
			panic(err)
		}
		defer func() {
			if err := producer.Close(); err != nil {
				fmt.Printf("Error closing producer: %v\n", err)
			}
		}()

		kafkaMessage := &sarama.ProducerMessage{
			Topic: "test-topic",
			Value: sarama.StringEncoder(msg),
		}

		partition, offset, err := producer.SendMessage(kafkaMessage)
		if err != nil {
			fmt.Printf("Error producing message: %v\n", err)
			panic(err)
		}

		fmt.Printf("Message: %s is stored in topic(%s)/partition(%d)/offset(%d)\n", kafkaMessage.Value, kafkaMessage.Topic, partition, offset)
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)

	// message flag
	producerCmd.PersistentFlags().StringVarP(&msg, "msg", "m", "Status: Completed task", "The message to send to the topic")

	// brokers can be provided as: --brokers host1:9092,host2:9092
	producerCmd.Flags().StringSliceVar(&producerBrokers, "brokers", []string{"localhost:9092"}, "comma-separated list of Kafka brokers")

}
