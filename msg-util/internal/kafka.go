package internal

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func CreateKafkaProducer(brokers []string, config *sarama.Config) (sarama.SyncProducer, error) {
	if config == nil {
		cfg := sarama.NewConfig()
		config = cfg
	}
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Version = sarama.V2_1_0_0

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
	return producer, nil
}

func CreateKafkaConsumer(brokers []string, groupID string, config *sarama.Config) sarama.ConsumerGroup {
	config.Version = sarama.V2_1_0_0
	config.Consumer.Return.Errors = true
	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}
	return consumerGroup
}

// GetSecureKafkaConfig retrieves Kafka configuration with TLS settings from a Kubernetes secret
//
// Parameters:
//   - inCluster: Boolean indicating if the code is running inside a Kubernetes cluster
//   - namespace: Namespace where the secret is located
//   - secretName: Name of the secret containing TLS certificates
//
// Returns:
//   - *sarama.Config: Configured Sarama Kafka configuration with TLS settings
//   - error: Error if any occurred during the process
func GetSecureKafkaConfig(inCluster bool, namespace string, secretName string) (*sarama.Config, error) {
	fmt.Printf("Configuring TLS")
	restConfig, err := GetRestConfig("", inCluster)
	if err != nil {
		return nil, fmt.Errorf("Failed to retreive RestConfig: %w", err)
	}
	clientset, err := GetClientset(restConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to retreive ClientSet: %w", err)
	}
	secret, err := GetSecret(clientset, namespace, secretName)
	if err != nil {
		return nil, fmt.Errorf("Failed to retreive Secret: %w", err)
	}

	kafkaConfig, err := ConfigTLS(nil, secret.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to configure TLS: %w", err)
	}
	return kafkaConfig, nil
}

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	fmt.Println("Setting up consumer")
	return nil
}
func (ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("Cleaning up consumer")
	return nil
}
func (ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("msg topic=%s key=%s value=%s",
			msg.Topic, string(msg.Key), string(msg.Value))
		session.MarkMessage(msg, "Processed by: msg-util")
	}
	return nil
}

func ConfigTLS(config *sarama.Config, certs map[string][]byte) (*sarama.Config, error) {
	if config == nil {
		cfg := sarama.NewConfig()
		config = cfg
	}

	certPEM, ok := certs["tls.crt"]
	if !ok || len(certPEM) == 0 {
		return nil, fmt.Errorf("missing tls.crt in certs map")
	}
	keyPEM, ok := certs["tls.key"]
	if !ok || len(keyPEM) == 0 {
		return nil, fmt.Errorf("missing tls.key in certs map")
	}
	caPEM := certs["ca.crt"] // optional

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to load certificate/key pair: %w", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	if len(caPEM) > 0 {
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(caPEM) {
			return nil, fmt.Errorf("failed to parse ca.crt")
		}
		tlsConfig.RootCAs = pool
	}

	config.Net.TLS.Enable = true
	config.Net.TLS.Config = tlsConfig

	return config, nil
}
