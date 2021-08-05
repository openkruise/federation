/*
Copyright 2021 The Kruise Authors.

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

package federation

import (
	"context"
	
	appsv1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	genericclient "sigs.k8s.io/kubefed/pkg/client/generic"
	ctlutil "sigs.k8s.io/kubefed/pkg/controller/util"

	federationv1alpha1 "github.com/openkruise/federation/apis/federation/v1alpha1"
	"github.com/openkruise/federation/options"
)

var CloneSetResource = &metav1.APIResource{
	Name:       "clonesets",
	Group:      appsv1alpha1.SchemeGroupVersion.Group,
	Version:    appsv1alpha1.SchemeGroupVersion.Version,
	Kind:       "CloneSet",
	Namespaced: true,
}

// FederatedCloneSetReconciler reconciles a FederatedCloneSet object
type FederatedCloneSetReconciler struct {
	client.Client
	Scheme         *runtime.Scheme
	targetInformer ctlutil.FederatedInformer
}

//+kubebuilder:rbac:groups=federation.kruise.io,resources=federatedclonesets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=federation.kruise.io,resources=federatedclonesets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=federation.kruise.io,resources=federatedclonesets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the FederatedCloneSet object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *FederatedCloneSetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FederatedCloneSetReconciler) SetupWithManager(opts *options.Options, mgr ctrl.Manager) error {
	var err error
	fedClient := genericclient.NewForConfigOrDieWithUserAgent(mgr.GetConfig(), "federated-cloneset")
	r.targetInformer, err = ctlutil.NewFederatedInformer(
		&ctlutil.ControllerConfig{KubeFedNamespaces: opts.KubeFedNamespaces},
		fedClient,
		CloneSetResource,
		nil, // TODO
		nil, // TODO
	)
	if err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&federationv1alpha1.FederatedCloneSet{}).
		Complete(r)
}
