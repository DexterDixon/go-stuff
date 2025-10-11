package tests

import (
	"go-stuff/msg-util/internal"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestGetClientset(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		restCfg *rest.Config
		want    *kubernetes.Clientset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := internal.GetClientset(tt.restCfg)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetClientset() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetClientset() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetClientset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSecret(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		client     *kubernetes.Clientset
		namespace  string
		secretName string
		want       *corev1.Secret
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := internal.GetSecret(tt.client, tt.namespace, tt.secretName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetSecret() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetSecret() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRestConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		kubeconfigPath string
		want           *rest.Config
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := internal.GetRestConfig(tt.kubeconfigPath)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetRestConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetRestConfig() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetRestConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
