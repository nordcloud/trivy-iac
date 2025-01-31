package athena

import (
	"github.com/aquasecurity/defsec/pkg/providers/aws/athena"
	"github.com/nordcloud/trivy-iac/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) athena.Athena {
	return athena.Athena{
		Databases:  nil,
		Workgroups: getWorkGroups(cfFile),
	}
}
