// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	infrav1 "github.com/talos-systems/cluster-api-provider-metal/api/v1alpha2"
	"github.com/talos-systems/cluster-api-provider-metal/internal/pkg/ipmi"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

// MetalMachineReconciler reconciles a MetalMachine object
type MetalMachineReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=metalmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=metalmachines/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cluster.x-k8s.io,resources=machines;machines/status,verbs=get;list;watch

func (r *MetalMachineReconciler) Reconcile(req ctrl.Request) (_ ctrl.Result, rerr error) {
	ctx := context.Background()
	logger := r.Log.WithValues("metalmachine", req.NamespacedName)

	// Fetch the metalMachine instance.
	metalMachine := &infrav1.MetalMachine{}
	err := r.Get(ctx, req.NamespacedName, metalMachine)
	if apierrors.IsNotFound(err) {
		return ctrl.Result{}, nil
	}
	if err != nil {
		return ctrl.Result{}, err
	}

	machine, err := util.GetOwnerMachine(ctx, r.Client, metalMachine.ObjectMeta)
	if err != nil {
		return ctrl.Result{}, err
	}
	if machine == nil {
		return ctrl.Result{}, fmt.Errorf("no ownerref for machine %s", metalMachine.ObjectMeta.Name)
	}

	logger = logger.WithName(fmt.Sprintf("machine=%s", machine.Name))

	cluster, err := util.GetClusterFromMetadata(ctx, r.Client, machine.ObjectMeta)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("no cluster label or cluster does not exist")
	}

	logger = logger.WithName(fmt.Sprintf("cluster=%s", cluster.Name))

	if !cluster.Status.InfrastructureReady {
		return ctrl.Result{}, fmt.Errorf("clusterInfra is not available yet")
	}

	if machine.Spec.Bootstrap.Data == nil {
		return ctrl.Result{}, fmt.Errorf("bootstrapData is not available yet")
	}

	// Initialize the patch helper
	patchHelper, err := patch.NewHelper(metalMachine, r)
	if err != nil {
		return ctrl.Result{}, err
	}
	// Always attempt to Patch the argesCluster object and status after each reconciliation.
	defer func() {
		if err := patchHelper.Patch(ctx, metalMachine); err != nil {
			logger.Error(err, "failed to patch metalMachine")
			if rerr == nil {
				rerr = err
			}
		}
	}()

	// If the metalMachine doesn't have our finalizer, add it.
	if !util.Contains(metalMachine.Finalizers, infrav1.MachineFinalizer) {
		metalMachine.Finalizers = append(metalMachine.Finalizers, infrav1.MachineFinalizer)
	}

	// Handle deleted machines
	if !metalMachine.ObjectMeta.DeletionTimestamp.IsZero() {
		logger.Info("deleting machine")
		return r.reconcileDelete(metalMachine)
	}

	ipmiClient, err := ipmi.NewClient(metalMachine.Spec.BMC)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = ipmiClient.SetPXE()
	if err != nil {
		return ctrl.Result{}, err
	}

	err = ipmiClient.PowerOn()
	if err != nil {
		return ctrl.Result{}, err
	}

	metalMachine.Status.Ready = true
	return ctrl.Result{}, nil
}

func (r *MetalMachineReconciler) reconcileDelete(metalMachine *infrav1.MetalMachine) (ctrl.Result, error) {
	//TODO(rsmitty): add in call to reset node via talos api

	ipmiClient, err := ipmi.NewClient(metalMachine.Spec.BMC)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = ipmiClient.SetPXE()
	if err != nil {
		return ctrl.Result{}, err
	}

	err = ipmiClient.PowerOff()
	if err != nil {
		return ctrl.Result{}, err
	}

	// Cluster is deleted so remove the finalizer.
	metalMachine.Finalizers = util.Filter(metalMachine.Finalizers, infrav1.MachineFinalizer)

	return ctrl.Result{}, nil
}

func (r *MetalMachineReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(&infrav1.MetalMachine{}).
		Complete(r)
}
