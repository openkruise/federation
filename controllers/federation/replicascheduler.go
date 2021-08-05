package federation

import (
	//runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	genericclient "sigs.k8s.io/kubefed/pkg/client/generic"
	ctlutil "sigs.k8s.io/kubefed/pkg/controller/util"
	scheduletype "sigs.k8s.io/kubefed/pkg/schedulingtypes"
	//fedv1b1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
)

type ReplicaScheduler struct {
	controllerConfig *ctlutil.ControllerConfig

	eventHandlers scheduletype.SchedulerEventHandlers

	plugins *ctlutil.SafeMap

	client      genericclient.Client
	podInformer ctlutil.FederatedInformer
}

//Reconcile distribute replica according to typecofig
//func (s *ReplicaScheduler) Reconcile(obj runtimeclient.Object, qualifiedName ctlutil.QualifiedName) ctlutil.ReconciliationStatus {
//
//}

//Plan tries its best to assign each cluster the preferred number of replicas
// * a map that contains information how many replicas will be possible to run in a cluster.
// * a map that contains information how many extra replicas would be nice to schedule in a cluster
//func Plan(availableClusters []string, currentReplicaCount map[string]int64,
//	estimatedCapacity map[string]int64, replicaSetKey string) (map[string]int64, map[string]int64, error) {
//
//	}

// name selector
//func (s *ReplicaScheduler) clusterNames(name string,clusters []*fedv1b1.KubeFedCluster) []string {
//
//}

// label selector
//func (s *ReplicaScheduler) clusterLabel(label string,clusters []*fedv1b1.KubeFedCluster) []string {
//
//}
