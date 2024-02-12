package iam

import (
	"testing"

	defsecTypes "github.com/aquasecurity/trivy/pkg/types"

	"github.com/aquasecurity/trivy/pkg/providers/aws/iam"

	"github.com/aquasecurity/trivy/internal/adapters/terraform/tftestutil"
	"github.com/aquasecurity/trivy/test/testutil"
)

func Test_adaptPasswordPolicy(t *testing.T) {
	tests := []struct {
		name      string
		terraform string
		expected  iam.PasswordPolicy
	}{
		{
			name: "basic",
			terraform: `
			resource "aws_iam_account_password_policy" "strict" {
				minimum_password_length        = 8
				require_lowercase_characters   = true
				require_numbers                = true
				require_uppercase_characters   = true
				require_symbols                = true
				allow_users_to_change_password = true
				max_password_age               = 90
				password_reuse_prevention      = 3
			  }
`,
			expected: iam.PasswordPolicy{
				Metadata:             defsecTypes.NewTestMisconfigMetadata(),
				ReusePreventionCount: defsecTypes.Int(3, defsecTypes.NewTestMisconfigMetadata()),
				RequireLowercase:     defsecTypes.Bool(true, defsecTypes.NewTestMisconfigMetadata()),
				RequireUppercase:     defsecTypes.Bool(true, defsecTypes.NewTestMisconfigMetadata()),
				RequireNumbers:       defsecTypes.Bool(true, defsecTypes.NewTestMisconfigMetadata()),
				RequireSymbols:       defsecTypes.Bool(true, defsecTypes.NewTestMisconfigMetadata()),
				MaxAgeDays:           defsecTypes.Int(90, defsecTypes.NewTestMisconfigMetadata()),
				MinimumLength:        defsecTypes.Int(8, defsecTypes.NewTestMisconfigMetadata()),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := tftestutil.CreateModulesFromSource(t, test.terraform, ".tf")
			adapted := adaptPasswordPolicy(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}