/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/timebertt/kubernetes-controller-sharding/pkg/shard/controller"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	corev1 "github.com/tbhardwaj20/notifier-controller/api/v1"
)

// NotifierReconciler reconciles a Notifier object
type NotifierReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	ShardName string 
}

// +kubebuilder:rbac:groups=core.tbh.dev,resources=notifiers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core.tbh.dev,resources=notifiers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core.tbh.dev,resources=notifiers/finalizers,verbs=update

func (r *NotifierReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	var notifier corev1.Notifier
	if err := r.Get(ctx, req.NamespacedName, &notifier); err != nil {
		log.Error(err, "unable to fetch Notifier")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("Reconciling Notifier",
		"name", notifier.Name,
		"message", notifier.Spec.Message,
		"target", notifier.Spec.Target,
		"type", notifier.Spec.Type,
	)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NotifierReconciler) SetupWithManager(mgr ctrl.Manager) error {
	controllerRing := "notifier-ring" // must match ControllerRing CR

	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Notifier{}, builder.WithPredicates(
			controller.Predicate(controllerRing, r.ShardName, predicate.ResourceVersionChangedPredicate{}),
		)).
		Named("notifier").
		Complete(
			controller.NewShardedReconciler(mgr).
				For(&corev1.Notifier{}).
				InControllerRing(controllerRing).
				WithShardName(r.ShardName).
				MustBuild(r),
		)
}
