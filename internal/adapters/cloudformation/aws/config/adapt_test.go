package config

import (
	"context"
	"testing"

	"github.com/aquasecurity/trivy/pkg/providers/aws/config"
	"github.com/aquasecurity/trivy/pkg/types"
	"github.com/stretchr/testify/require"

	"github.com/aquasecurity/trivy/pkg/scanners/cloudformation/parser"
	"github.com/aquasecurity/trivy/test/testutil"
)

func TestAdapt(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected config.Config
	}{
		{
			name: "Config aggregator with AccountAggregationSources",
			source: `AWSTemplateFormatVersion: "2010-09-09"
Resources:
  ConfigurationAggregator:
    Type: AWS::Config::ConfigurationAggregator
    Properties:
      AccountAggregationSources:
        - AllAwsRegions: "true"
`,
			expected: config.Config{
				ConfigurationAggregrator: config.ConfigurationAggregrator{
					Metadata:         types.NewTestMisconfigMetadata(),
					SourceAllRegions: types.Bool(true, types.NewTestMisconfigMetadata()),
				},
			},
		},
		{
			name: "Config aggregator with OrganizationAggregationSource",
			source: `AWSTemplateFormatVersion: "2010-09-09"
Resources:
  ConfigurationAggregator:
    Type: AWS::Config::ConfigurationAggregator
    Properties:
      OrganizationAggregationSource:
        AllAwsRegions: "true"
`,
			expected: config.Config{
				ConfigurationAggregrator: config.ConfigurationAggregrator{
					Metadata:         types.NewTestMisconfigMetadata(),
					SourceAllRegions: types.Bool(true, types.NewTestMisconfigMetadata()),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := testutil.CreateFS(t, map[string]string{
				"template.yaml": tt.source,
			})

			p := parser.New()
			fctx, err := p.ParseFile(context.TODO(), fs, "template.yaml")
			require.NoError(t, err)

			testutil.AssertDefsecEqual(t, tt.expected, Adapt(*fctx))
		})
	}

}