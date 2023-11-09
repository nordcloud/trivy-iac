package arm

import (
	"context"

	"github.com/nordcloud/trivy-iac/internal/adapters/arm/appservice"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/authorization"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/compute"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/container"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/database"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/datafactory"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/datalake"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/keyvault"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/monitor"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/network"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/securitycenter"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/storage"
	"github.com/nordcloud/trivy-iac/internal/adapters/arm/synapse"

	"github.com/aquasecurity/defsec/pkg/providers/azure"
	"github.com/aquasecurity/defsec/pkg/state"
	scanner "github.com/nordcloud/trivy-iac/pkg/scanners/azure"
)

// Adapt ...
func Adapt(ctx context.Context, deployment scanner.Deployment) *state.State {
	return &state.State{
		Azure: adaptAzure(deployment),
	}
}

func adaptAzure(deployment scanner.Deployment) azure.Azure {

	return azure.Azure{
		AppService:     appservice.Adapt(deployment),
		Authorization:  authorization.Adapt(deployment),
		Compute:        compute.Adapt(deployment),
		Container:      container.Adapt(deployment),
		Database:       database.Adapt(deployment),
		DataFactory:    datafactory.Adapt(deployment),
		DataLake:       datalake.Adapt(deployment),
		KeyVault:       keyvault.Adapt(deployment),
		Monitor:        monitor.Adapt(deployment),
		Network:        network.Adapt(deployment),
		SecurityCenter: securitycenter.Adapt(deployment),
		Storage:        storage.Adapt(deployment),
		Synapse:        synapse.Adapt(deployment),
	}

}
