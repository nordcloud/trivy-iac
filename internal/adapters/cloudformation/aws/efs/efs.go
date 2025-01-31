package efs

import (
	"github.com/aquasecurity/defsec/pkg/providers/aws/efs"
	"github.com/nordcloud/trivy-iac/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) efs.EFS {
	return efs.EFS{
		FileSystems: getFileSystems(cfFile),
	}
}
