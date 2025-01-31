package aws

import (
	"github.com/aquasecurity/defsec/pkg/providers/aws"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/apigateway"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/athena"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/cloudfront"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/cloudtrail"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/cloudwatch"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/codebuild"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/config"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/documentdb"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/dynamodb"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/ec2"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/ecr"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/ecs"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/efs"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/eks"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/elasticache"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/elasticsearch"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/elb"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/iam"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/kinesis"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/lambda"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/mq"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/msk"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/neptune"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/rds"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/redshift"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/s3"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/sam"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/sns"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/sqs"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/ssm"
	"github.com/nordcloud/trivy-iac/internal/adapters/cloudformation/aws/workspaces"
	"github.com/nordcloud/trivy-iac/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) aws.AWS {
	return aws.AWS{
		APIGateway:    apigateway.Adapt(cfFile),
		Athena:        athena.Adapt(cfFile),
		Cloudfront:    cloudfront.Adapt(cfFile),
		CloudTrail:    cloudtrail.Adapt(cfFile),
		CloudWatch:    cloudwatch.Adapt(cfFile),
		CodeBuild:     codebuild.Adapt(cfFile),
		Config:        config.Adapt(cfFile),
		DocumentDB:    documentdb.Adapt(cfFile),
		DynamoDB:      dynamodb.Adapt(cfFile),
		EC2:           ec2.Adapt(cfFile),
		ECR:           ecr.Adapt(cfFile),
		ECS:           ecs.Adapt(cfFile),
		EFS:           efs.Adapt(cfFile),
		IAM:           iam.Adapt(cfFile),
		EKS:           eks.Adapt(cfFile),
		ElastiCache:   elasticache.Adapt(cfFile),
		Elasticsearch: elasticsearch.Adapt(cfFile),
		ELB:           elb.Adapt(cfFile),
		MSK:           msk.Adapt(cfFile),
		MQ:            mq.Adapt(cfFile),
		Kinesis:       kinesis.Adapt(cfFile),
		Lambda:        lambda.Adapt(cfFile),
		Neptune:       neptune.Adapt(cfFile),
		RDS:           rds.Adapt(cfFile),
		Redshift:      redshift.Adapt(cfFile),
		S3:            s3.Adapt(cfFile),
		SAM:           sam.Adapt(cfFile),
		SNS:           sns.Adapt(cfFile),
		SQS:           sqs.Adapt(cfFile),
		SSM:           ssm.Adapt(cfFile),
		WorkSpaces:    workspaces.Adapt(cfFile),
	}
}
