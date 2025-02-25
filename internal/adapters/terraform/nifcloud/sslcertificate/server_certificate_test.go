package sslcertificate

import (
	"testing"
	"time"

	"github.com/aquasecurity/defsec/pkg/providers/nifcloud/sslcertificate"
	defsecTypes "github.com/aquasecurity/defsec/pkg/types"

	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/tftestutil"

	"github.com/nordcloud/trivy-iac/test/testutil"
)

const certificate = `
-----BEGIN CERTIFICATE-----
MIIB0zCCAX2gAwIBAgIJAI/M7BYjwB+uMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTIwOTEyMjE1MjAyWhcNMTUwOTEyMjE1MjAyWjBF
MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50
ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANLJ
hPHhITqQbPklG3ibCVxwGMRfp/v4XqhfdQHdcVfHap6NQ5Wok/4xIA+ui35/MmNa
rtNuC+BdZ1tMuVCPFZcCAwEAAaNQME4wHQYDVR0OBBYEFJvKs8RfJaXTH08W+SGv
zQyKn0H8MB8GA1UdIwQYMBaAFJvKs8RfJaXTH08W+SGvzQyKn0H8MAwGA1UdEwQF
MAMBAf8wDQYJKoZIhvcNAQEFBQADQQBJlffJHybjDGxRMqaRmDhX0+6v02TUKZsW
r5QuVbpQhH6u+0UgcW0jp9QwpxoPTLTWGXEWBBBurxFwiCBhkQ+V
-----END CERTIFICATE-----
`

func Test_adaptServerCertificates(t *testing.T) {
	tests := []struct {
		name      string
		terraform string
		expected  []sslcertificate.ServerCertificate
	}{
		{
			name: "configured",
			terraform: `
			resource "nifcloud_ssl_certificate" "example" {
				certificate  = <<EOT` + certificate + `EOT
			}
`,
			expected: []sslcertificate.ServerCertificate{{
				Metadata: defsecTypes.NewTestMetadata(),
				Expiration: defsecTypes.Time(func(timeVal string) time.Time {
					parsed, _ := time.Parse(time.RFC3339, timeVal)
					return parsed
				}("2015-09-12T21:52:02Z"), defsecTypes.NewTestMetadata()),
			}},
		},
		{
			name: "defaults",
			terraform: `
			resource "nifcloud_ssl_certificate" "example" {
			}
`,

			expected: []sslcertificate.ServerCertificate{{
				Metadata:   defsecTypes.NewTestMetadata(),
				Expiration: defsecTypes.Time(time.Time{}, defsecTypes.NewTestMetadata()),
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := tftestutil.CreateModulesFromSource(t, test.terraform, ".tf")
			adapted := adaptServerCertificates(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}
