package sqs

import (
	"fmt"

	defsecTypes "github.com/aquasecurity/defsec/pkg/types"

	"github.com/nordcloud/trivy-iac/pkg/scanners/cloudformation/parser"

	"github.com/aquasecurity/defsec/pkg/providers/aws/iam"
	"github.com/aquasecurity/defsec/pkg/providers/aws/sqs"

	"github.com/liamg/iamgo"
)

func getQueues(ctx parser.FileContext) (queues []sqs.Queue) {
	for _, r := range ctx.GetResourcesByType("AWS::SQS::Queue") {
		queue := sqs.Queue{
			Metadata: r.Metadata(),
			QueueURL: defsecTypes.StringDefault("", r.Metadata()),
			Encryption: sqs.Encryption{
				Metadata:          r.Metadata(),
				ManagedEncryption: defsecTypes.Bool(false, r.Metadata()),
				KMSKeyID:          r.GetStringProperty("KmsMasterKeyId"),
			},
			Policies: []iam.Policy{},
		}
		if policy, err := getPolicy(r.ID(), ctx); err == nil {
			queue.Policies = append(queue.Policies, *policy)
		}
		queues = append(queues, queue)
	}
	return queues
}

func getPolicy(id string, ctx parser.FileContext) (*iam.Policy, error) {
	for _, policyResource := range ctx.GetResourcesByType("AWS::SQS::QueuePolicy") {
		documentProp := policyResource.GetProperty("PolicyDocument")
		if documentProp.IsNil() {
			continue
		}
		queuesProp := policyResource.GetProperty("Queues")
		if queuesProp.IsNil() {
			continue
		}
		for _, queueRef := range queuesProp.AsList() {
			if queueRef.IsString() && queueRef.AsString() == id {
				raw := documentProp.GetJsonBytes()
				parsed, err := iamgo.Parse(raw)
				if err != nil {
					continue
				}
				return &iam.Policy{
					Metadata: documentProp.Metadata(),
					Name:     defsecTypes.StringDefault("", documentProp.Metadata()),
					Document: iam.Document{
						Metadata: documentProp.Metadata(),
						Parsed:   *parsed,
					},
					Builtin: defsecTypes.Bool(false, documentProp.Metadata()),
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("no matching policy found")
}
