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

	z0v1 "pod-operator/api/v1"

	"sigs.k8s.io/controller-runtime/pkg/log"
	"k8s.io/apimachinery/pkg/types"
    apierrors "k8s.io/apimachinery/pkg/api/errors"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
    // "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// DeployMakerv0Reconciler reconciles a DeployMakerv0 object
type DeployMakerv0Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=z0.v0.com,resources=deploymakerv0s,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=z0.v0.com,resources=deploymakerv0s/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=z0.v0.com,resources=deploymakerv0s/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DeployMakerv0 object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *DeployMakerv0Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx)

	// TODO(user): your logic here

	log := log.FromContext(ctx)

    var deployMaker z0v1.DeployMakerv0
    if err := r.Get(ctx, req.NamespacedName, &deployMaker); err != nil {
        log.Error(err, "unable to fetch DeployMakerv0")
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    replicas := int32(1)
    if deployMaker.Spec.Replicas != nil && *deployMaker.Spec.Replicas > 0 {
        replicas = *deployMaker.Spec.Replicas
    }

    deployment := &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{
            Name:      deployMaker.Name,
            Namespace: req.Namespace,
        },
        Spec: appsv1.DeploymentSpec{
            Replicas: &replicas,
            Selector: &metav1.LabelSelector{
                MatchLabels: map[string]string{"app": deployMaker.Name},
            },
            Template: corev1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: map[string]string{"app": deployMaker.Name},
                },
                Spec: corev1.PodSpec{
                    Containers: []corev1.Container{{
                        Name:  "main",
                        Image: deployMaker.Spec.Image,
                    }},
                },
            },
        },
    }

    existing := &appsv1.Deployment{}
    err := r.Get(ctx, types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, existing)
    if err != nil && apierrors.IsNotFound(err) {
        log.Info("Creating Deployment", "name", deployment.Name)
        if err := r.Create(ctx, deployment); err != nil {
            log.Error(err, "failed to create Deployment")
            return ctrl.Result{}, err
        }
    } else if err != nil {
        log.Error(err, "failed to get Deployment")
        return ctrl.Result{}, err
    }

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DeployMakerv0Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&z0v1.DeployMakerv0{}).
		Named("deploymakerv0").
		Complete(r)
}
