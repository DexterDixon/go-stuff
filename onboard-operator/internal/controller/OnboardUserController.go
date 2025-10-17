package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	batchv1 "tutorial.kubebuilder.io/project/api/v1"
)

// OnboardUserReconciler reconciles a OnboardUser object
type OnboardUserReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// `Reconcile` actually performs the reconciling for a single named object.
func (r *OnboardUserReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	var onboardUser batchv1.OnboardUser
	log.Info("Fetching OnboardUser...")
	if err := r.Get(ctx, req.NamespacedName, &onboardUser); err != nil {
		log.Error(err, "unable to fetch OnboardUser")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("Onboarding", "username", onboardUser.Spec.Username)

	op, err := createOrUpdateNamespace(ctx, r.Client, onboardUser.Spec.Username, map[string]string{"role": onboardUser.Spec.Role})
	if err != nil {
		log.Error(err, "failed to create or update namespace", "namespace", onboardUser.Spec.Username)
		return ctrl.Result{}, err
	}
	log.Info("Namespace ensured", "operation", op, "namespace", onboardUser.Spec.Username)

	log.Info("Applying Permisions for", "role", onboardUser.Spec.Role)

	log.Info("Applying Configurations for", "textEditor", onboardUser.Spec.TextEditor)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OnboardUserReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.OnboardUser{}).
		Complete(r)
}

// createOrUpdateNamespace creates or updates a Namespace with the given name and labels.
func createOrUpdateNamespace(ctx context.Context, c client.Client, name string, labels map[string]string) (controllerutil.OperationResult, error) {
	var op controllerutil.OperationResult

	err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		ns := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		}

		var innerErr error
		op, innerErr = controllerutil.CreateOrUpdate(ctx, c, ns, func() error {
			if ns.Labels == nil {
				ns.Labels = map[string]string{}
			}
			for k, v := range labels {
				ns.Labels[k] = v
			}
			return nil
		})

		// Return the error so RetryOnConflict can decide whether to retry.
		return innerErr
	})

	return op, err
}
