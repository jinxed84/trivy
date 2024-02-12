package trules

import (
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/accessanalyzer"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/apigateway"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/athena"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/cloudfront"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/cloudtrail"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/cloudwatch"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/codebuild"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/config"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/documentdb"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/dynamodb"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/ec2"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/ecr"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/ecs"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/efs"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/eks"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/elasticache"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/elasticsearch"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/elb"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/emr"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/iam"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/kinesis"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/kms"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/lambda"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/mq"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/msk"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/neptune"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/rds"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/redshift"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/s3"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/sam"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/sns"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/sqs"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/ssm"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/aws/workspaces"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/appservice"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/authorization"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/compute"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/container"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/database"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/datafactory"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/datalake"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/keyvault"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/monitor"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/network"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/securitycenter"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/storage"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/azure/synapse"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/cloudstack/compute"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/digitalocean/compute"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/digitalocean/spaces"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/github/actions"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/github/branch_protections"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/github/repositories"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/google/bigquery"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/google/compute"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/google/dns"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/google/gke"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/google/iam"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/google/kms"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/google/sql"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/google/storage"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/nifcloud/computing"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/nifcloud/dns"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/nifcloud/nas"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/nifcloud/network"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/nifcloud/rdb"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/nifcloud/sslcertificate"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/openstack/compute"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/openstack/networking"
	_ "github.com/aquasecurity/trivy-policies/checks/cloud/oracle/compute"
	_ "github.com/aquasecurity/trivy-policies/checks/kubernetes/network"
	//trules "github.com/aquasecurity/trivy-policies/pkg/trules"
	"github.com/aquasecurity/trivy/pkg/scan"
)

func init() {
	for _, r := range GetRules() {
		Register(r)
	}
}

var trules []scan.Rule

//
//func Register(r scan.Rule, f scan.CheckFunc) scan.Rule {
//	r.Check = f
//	trules = append(trules, r)
//
//	return r
//}

func GetRules() []scan.Rule {
	return trules
}