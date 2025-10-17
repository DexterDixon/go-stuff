/*
Copyright © 2025 Dexter Dixon dexterdixon561@gmail.com
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	kafkav1beta1 "github.com/DexterDixon/go-stuff/msg-util-operator/api/v1beta1"
)

// KafkaProducerReconciler reconciles a KafkaProducer object
type KafkaProducerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kafka.dexter.dixon.com,resources=kafkaproducers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kafka.dexter.dixon.com,resources=kafkaproducers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kafka.dexter.dixon.com,resources=kafkaproducers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KafkaProducer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.22.1/pkg/reconcile
func (r *KafkaProducerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KafkaProducerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kafkav1beta1.KafkaProducer{}).
		Named("kafkaproducer").
		Complete(r)
}
