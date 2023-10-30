package azure

import (
	"github.com/aquasecurity/defsec/pkg/providers/azure"
	"github.com/aquasecurity/defsec/pkg/terraform"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/appservice"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/authorization"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/compute"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/container"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/database"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/datafactory"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/datalake"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/keyvault"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/monitor"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/network"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/securitycenter"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/storage"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/azure/synapse"
)

func Adapt(modules terraform.Modules) azure.Azure {
	return azure.Azure{
		AppService:     appservice.Adapt(modules),
		Authorization:  authorization.Adapt(modules),
		Compute:        compute.Adapt(modules),
		Container:      container.Adapt(modules),
		Database:       database.Adapt(modules),
		DataFactory:    datafactory.Adapt(modules),
		DataLake:       datalake.Adapt(modules),
		KeyVault:       keyvault.Adapt(modules),
		Monitor:        monitor.Adapt(modules),
		Network:        network.Adapt(modules),
		SecurityCenter: securitycenter.Adapt(modules),
		Storage:        storage.Adapt(modules),
		Synapse:        synapse.Adapt(modules),
	}
}
