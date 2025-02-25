package terraform

import (
	"github.com/aquasecurity/defsec/pkg/state"
	"github.com/aquasecurity/defsec/pkg/terraform"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/cloudstack"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/digitalocean"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/github"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/google"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/kubernetes"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/nifcloud"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/openstack"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/oracle"
)

func Adapt(modules terraform.Modules) *state.State {
	return &state.State{
		AWS:          aws.Adapt(modules),
		Azure:        azure.Adapt(modules),
		CloudStack:   cloudstack.Adapt(modules),
		DigitalOcean: digitalocean.Adapt(modules),
		GitHub:       github.Adapt(modules),
		Google:       google.Adapt(modules),
		Kubernetes:   kubernetes.Adapt(modules),
		Nifcloud:     nifcloud.Adapt(modules),
		OpenStack:    openstack.Adapt(modules),
		Oracle:       oracle.Adapt(modules),
	}
}
