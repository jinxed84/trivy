package compute

import (
	"github.com/aquasecurity/trivy/pkg/providers/google/compute"
	"github.com/aquasecurity/trivy/pkg/terraform"
)

func Adapt(modules terraform.Modules) compute.Compute {
	return compute.Compute{
		ProjectMetadata: adaptProjectMetadata(modules),
		Instances:       adaptInstances(modules),
		Disks:           adaptDisks(modules),
		Networks:        adaptNetworks(modules),
		SSLPolicies:     adaptSSLPolicies(modules),
	}
}