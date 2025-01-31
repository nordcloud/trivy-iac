package aws

import (
	"github.com/aquasecurity/defsec/pkg/providers/aws"
	"github.com/aquasecurity/defsec/pkg/terraform"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/apigateway"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/athena"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/cloudfront"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/cloudtrail"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/cloudwatch"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/codebuild"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/config"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/documentdb"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/dynamodb"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/ec2"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/ecr"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/ecs"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/efs"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/eks"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/elasticache"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/elasticsearch"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/elb"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/emr"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/iam"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/kinesis"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/kms"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/lambda"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/mq"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/msk"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/neptune"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/rds"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/redshift"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/s3"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/sns"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/sqs"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/ssm"
	"github.com/nordcloud/trivy-iac/internal/adapters/terraform/aws/workspaces"
)

func Adapt(modules terraform.Modules) aws.AWS {
	return aws.AWS{
		APIGateway:    apigateway.Adapt(modules),
		Athena:        athena.Adapt(modules),
		Cloudfront:    cloudfront.Adapt(modules),
		CloudTrail:    cloudtrail.Adapt(modules),
		CloudWatch:    cloudwatch.Adapt(modules),
		CodeBuild:     codebuild.Adapt(modules),
		Config:        config.Adapt(modules),
		DocumentDB:    documentdb.Adapt(modules),
		DynamoDB:      dynamodb.Adapt(modules),
		EC2:           ec2.Adapt(modules),
		ECR:           ecr.Adapt(modules),
		ECS:           ecs.Adapt(modules),
		EFS:           efs.Adapt(modules),
		EKS:           eks.Adapt(modules),
		ElastiCache:   elasticache.Adapt(modules),
		Elasticsearch: elasticsearch.Adapt(modules),
		ELB:           elb.Adapt(modules),
		EMR:           emr.Adapt(modules),
		IAM:           iam.Adapt(modules),
		Kinesis:       kinesis.Adapt(modules),
		KMS:           kms.Adapt(modules),
		Lambda:        lambda.Adapt(modules),
		MQ:            mq.Adapt(modules),
		MSK:           msk.Adapt(modules),
		Neptune:       neptune.Adapt(modules),
		RDS:           rds.Adapt(modules),
		Redshift:      redshift.Adapt(modules),
		S3:            s3.Adapt(modules),
		SNS:           sns.Adapt(modules),
		SQS:           sqs.Adapt(modules),
		SSM:           ssm.Adapt(modules),
		WorkSpaces:    workspaces.Adapt(modules),
	}
}
