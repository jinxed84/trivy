package cloudstack

import (
	"github.com/aquasecurity/trivy/pkg/providers/cloudstack/compute"
)

type CloudStack struct {
	Compute compute.Compute
}