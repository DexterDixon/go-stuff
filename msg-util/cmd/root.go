/*
Copyright © 2025 Dexter Dixon Dexterdixon561@gmail.com
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	inCluster bool = false
	namespace string
)

var rootCmd = &cobra.Command{
	Use:   "msg-util",
	Short: "A K8S Kafka Messaging Utility",
	Long: `A Kubernetes based messaging utility used to interact with Kafka clusters.
The utility allows users to send and receive messages to/from specified topics within a Kafka cluster.:
You can use your Kubernetes context to connect to different Kafka clusters.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("please provide a command :/")
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.messaging-server.yaml)")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "namespace containing the tls secret")
	rootCmd.PersistentFlags().BoolVarP(&inCluster, "inCluster", "c", false, "use to specify if running inside a k8s cluster")
}
