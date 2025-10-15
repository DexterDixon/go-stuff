package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Returns Kubenetes rest config using a kubeconfig file
// GetRestConfig returns Kubernetes rest config using a kubeconfig file
func GetRestConfig(kubeconfigPath string, inCluster bool) (*rest.Config, error) {
	if inCluster {
		inClusterConfig, err := rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("cannot load in-cluster config: %w", err)
		}
		log.Printf("Authenticated as: %s", inClusterConfig.Username)
		log.Printf("Cluster host: %s", inClusterConfig.Host)
		return inClusterConfig, nil
	}

	if kubeconfigPath == "" {
		if env := os.Getenv("KUBECONFIG"); env != "" {
			kubeconfigPath = env
		} else {
			home := homedir.HomeDir()
			if home == "" {
				return nil, fmt.Errorf("cannot determine kubeconfig path. set KUBECONFIG environment variable or  kubeConfigPath explicitly")
			}
			kubeconfigPath = filepath.Join(home, ".kube", "config")
		}
	}
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("cannot load kubeconfig: %w", err)
	}

	log.Printf("Authenticated as: %s", restConfig.Username)
	log.Printf("Cluster host: %s", restConfig.Host)

	return restConfig, nil
}

// GetClientset returns a kubernetes.Clientset for the provided rest.Config
func GetClientset(restCfg *rest.Config) (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return nil, fmt.Errorf("new clientset: %w", err)
	}
	return clientset, nil
}

// GetSecret retrieves a secret from the given namespace
// Parameters:
// - client: Kubernetes clientset
// - namespace: Namespace where the secret is located
// - secretName: Name of the secret to retrieve
// Returns:
// - *corev1.Secret: The retrieved secret object
// - error: Error if any occurred during retrieval
func GetSecret(client *kubernetes.Clientset, namespace string, secretName string) (*corev1.Secret, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	secret, err := client.CoreV1().Secrets(namespace).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get secret %s/%s: %w", namespace, secretName, err)
	}

	fmt.Printf("Secret %s/%s (type=%s) with %d keys\n", namespace, secret.Name, secret.Type, len(secret.Data))
	return secret, nil
}
