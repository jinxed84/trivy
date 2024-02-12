package github

import (
	"github.com/aquasecurity/trivy/internal/adapters/terraform/github/branch_protections"
	"github.com/aquasecurity/trivy/internal/adapters/terraform/github/repositories"
	"github.com/aquasecurity/trivy/internal/adapters/terraform/github/secrets"
	"github.com/aquasecurity/trivy/pkg/providers/github"
	"github.com/aquasecurity/trivy/pkg/terraform"
)

func Adapt(modules terraform.Modules) github.GitHub {
	return github.GitHub{
		Repositories:       repositories.Adapt(modules),
		EnvironmentSecrets: secrets.Adapt(modules),
		BranchProtections:  branch_protections.Adapt(modules),
	}
}