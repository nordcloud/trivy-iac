package github

import (
	"github.com/aquasecurity/defsec/pkg/providers/github"
	"github.com/aquasecurity/defsec/pkg/terraform"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/github/branch_protections"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/github/repositories"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/github/secrets"
)

func Adapt(modules terraform.Modules) github.GitHub {
	return github.GitHub{
		Repositories:       repositories.Adapt(modules),
		EnvironmentSecrets: secrets.Adapt(modules),
		BranchProtections:  branch_protections.Adapt(modules),
	}
}
