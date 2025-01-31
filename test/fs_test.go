package test

import (
	"context"
	"os"
	"testing"

	"github.com/aquasecurity/defsec/pkg/scanners/options"
	"github.com/nordcloud/trivy-iac/pkg/scanners/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_OS_FS(t *testing.T) {
	s := terraform.New(
		options.ScannerWithDebug(os.Stderr),
	)
	results, err := s.ScanFS(context.TODO(), os.DirFS("tf"), "fail")
	require.NoError(t, err)
	assert.Greater(t, len(results.GetFailed()), 0)
}
