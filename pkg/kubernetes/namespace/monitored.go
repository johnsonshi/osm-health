package namespace

import (
	"fmt"

	"k8s.io/client-go/kubernetes"

	"github.com/openservicemesh/osm-health/pkg/common"
	"github.com/openservicemesh/osm/pkg/constants"
)

// Verify interface compliance
var _ common.Runnable = (*MonitoredCheck)(nil)

// MonitoredCheck implements common.Runnable
type MonitoredCheck struct {
	client    kubernetes.Interface
	namespace string
	meshName  common.MeshName
}

// IsMonitoredBy checks whether a namespace is monitored by certain OSM Controller.
func IsMonitoredBy(client kubernetes.Interface, namespace string, meshName common.MeshName) MonitoredCheck {
	return MonitoredCheck{
		client:    client,
		namespace: namespace,
		meshName:  meshName,
	}
}

// Info implements common.Runnable
func (check MonitoredCheck) Info() string {
	return fmt.Sprintf("Checking whether namespace %s is monitored by OSM %s", check.namespace, check.meshName)
}

// Run implements common.Runnable
func (check MonitoredCheck) Run() error {
	labels, err := getLabels(check.client, check.namespace)
	if err != nil {
		return err
	}

	labelValue, ok := labels[constants.OSMKubeResourceMonitorAnnotation]
	isMonitoredByController := ok && labelValue == check.meshName.String()

	if !isMonitoredByController {
		return ErrNotMonitoredByOSMController
	}

	return nil
}

// Suggestion implements common.Runnable
func (check MonitoredCheck) Suggestion() string {
	panic("implement me")
}

// FixIt implements common.Runnable
func (check MonitoredCheck) FixIt() error {
	panic("implement me")
}
