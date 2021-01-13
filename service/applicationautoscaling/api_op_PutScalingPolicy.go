// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationautoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a scaling policy for an Application Auto Scaling scalable
// target. Each scalable target is identified by a service namespace, resource ID,
// and scalable dimension. A scaling policy applies to the scalable target
// identified by those three attributes. You cannot create a scaling policy until
// you have registered the resource as a scalable target. Multiple scaling policies
// can be in force at the same time for the same scalable target. You can have one
// or more target tracking scaling policies, one or more step scaling policies, or
// both. However, there is a chance that multiple policies could conflict,
// instructing the scalable target to scale out or in at the same time. Application
// Auto Scaling gives precedence to the policy that provides the largest capacity
// for both scale out and scale in. For example, if one policy increases capacity
// by 3, another policy increases capacity by 200 percent, and the current capacity
// is 10, Application Auto Scaling uses the policy with the highest calculated
// capacity (200% of 10 = 20) and scales out to 30. We recommend caution, however,
// when using target tracking scaling policies with step scaling policies because
// conflicts between these policies can cause undesirable behavior. For example, if
// the step scaling policy initiates a scale-in activity before the target tracking
// policy is ready to scale in, the scale-in activity will not be blocked. After
// the scale-in activity completes, the target tracking policy could instruct the
// scalable target to scale out again. For more information, see Target Tracking
// Scaling Policies
// (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html)
// and Step Scaling Policies
// (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html)
// in the Application Auto Scaling User Guide. If a scalable target is
// deregistered, the scalable target is no longer available to execute scaling
// policies. Any scaling policies that were specified for the scalable target are
// deleted.
func (c *Client) PutScalingPolicy(ctx context.Context, params *PutScalingPolicyInput, optFns ...func(*Options)) (*PutScalingPolicyOutput, error) {
	if params == nil {
		params = &PutScalingPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutScalingPolicy", params, optFns, addOperationPutScalingPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutScalingPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutScalingPolicyInput struct {

	// The name of the scaling policy.
	//
	// This member is required.
	PolicyName *string

	// The identifier of the resource associated with the scaling policy. This string
	// consists of the resource type and unique identifier.
	//
	// * ECS service - The
	// resource type is service and the unique identifier is the cluster name and
	// service name. Example: service/default/sample-webapp.
	//
	// * Spot Fleet request -
	// The resource type is spot-fleet-request and the unique identifier is the Spot
	// Fleet request ID. Example:
	// spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	// * EMR cluster -
	// The resource type is instancegroup and the unique identifier is the cluster ID
	// and instance group ID. Example:
	// instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	// * AppStream 2.0 fleet - The
	// resource type is fleet and the unique identifier is the fleet name. Example:
	// fleet/sample-fleet.
	//
	// * DynamoDB table - The resource type is table and the
	// unique identifier is the table name. Example: table/my-table.
	//
	// * DynamoDB global
	// secondary index - The resource type is index and the unique identifier is the
	// index name. Example: table/my-table/index/my-table-index.
	//
	// * Aurora DB cluster -
	// The resource type is cluster and the unique identifier is the cluster name.
	// Example: cluster:my-db-cluster.
	//
	// * Amazon SageMaker endpoint variant - The
	// resource type is variant and the unique identifier is the resource ID. Example:
	// endpoint/my-end-point/variant/KMeansClustering.
	//
	// * Custom resources are not
	// supported with a resource type. This parameter must specify the OutputValue from
	// the CloudFormation template stack used to access the resources. The unique
	// identifier is defined by the service provider. More information is available in
	// our GitHub repository
	// (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	// * Amazon Comprehend
	// document classification endpoint - The resource type and unique identifier are
	// specified using the endpoint ARN. Example:
	// arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.
	//
	// *
	// Amazon Comprehend entity recognizer endpoint - The resource type and unique
	// identifier are specified using the endpoint ARN. Example:
	// arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.
	//
	// *
	// Lambda provisioned concurrency - The resource type is function and the unique
	// identifier is the function name with a function version or alias name suffix
	// that is not $LATEST. Example: function:my-function:prod or
	// function:my-function:1.
	//
	// * Amazon Keyspaces table - The resource type is table
	// and the unique identifier is the table name. Example:
	// keyspace/mykeyspace/table/mytable.
	//
	// * Amazon MSK cluster - The resource type and
	// unique identifier are specified using the cluster ARN. Example:
	// arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.
	//
	// This member is required.
	ResourceId *string

	// The scalable dimension. This string consists of the service namespace, resource
	// type, and scaling property.
	//
	// * ecs:service:DesiredCount - The desired task count
	// of an ECS service.
	//
	// * ec2:spot-fleet-request:TargetCapacity - The target
	// capacity of a Spot Fleet request.
	//
	// *
	// elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR
	// Instance Group.
	//
	// * appstream:fleet:DesiredCapacity - The desired capacity of an
	// AppStream 2.0 fleet.
	//
	// * dynamodb:table:ReadCapacityUnits - The provisioned read
	// capacity for a DynamoDB table.
	//
	// * dynamodb:table:WriteCapacityUnits - The
	// provisioned write capacity for a DynamoDB table.
	//
	// *
	// dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB
	// global secondary index.
	//
	// * dynamodb:index:WriteCapacityUnits - The provisioned
	// write capacity for a DynamoDB global secondary index.
	//
	// *
	// rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB
	// cluster. Available for Aurora MySQL-compatible edition and Aurora
	// PostgreSQL-compatible edition.
	//
	// * sagemaker:variant:DesiredInstanceCount - The
	// number of EC2 instances for an Amazon SageMaker model endpoint variant.
	//
	// *
	// custom-resource:ResourceType:Property - The scalable dimension for a custom
	// resource provided by your own application or service.
	//
	// *
	// comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of
	// inference units for an Amazon Comprehend document classification endpoint.
	//
	// *
	// comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of
	// inference units for an Amazon Comprehend entity recognizer endpoint.
	//
	// *
	// lambda:function:ProvisionedConcurrency - The provisioned concurrency for a
	// Lambda function.
	//
	// * cassandra:table:ReadCapacityUnits - The provisioned read
	// capacity for an Amazon Keyspaces table.
	//
	// * cassandra:table:WriteCapacityUnits -
	// The provisioned write capacity for an Amazon Keyspaces table.
	//
	// *
	// kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for
	// brokers in an Amazon MSK cluster.
	//
	// This member is required.
	ScalableDimension types.ScalableDimension

	// The namespace of the AWS service that provides the resource. For a resource
	// provided by your own application or service, use custom-resource instead.
	//
	// This member is required.
	ServiceNamespace types.ServiceNamespace

	// The policy type. This parameter is required if you are creating a scaling
	// policy. The following policy types are supported: TargetTrackingScaling—Not
	// supported for Amazon EMR StepScaling—Not supported for DynamoDB, Amazon
	// Comprehend, Lambda, Amazon Keyspaces (for Apache Cassandra), or Amazon MSK. For
	// more information, see Target Tracking Scaling Policies
	// (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html)
	// and Step Scaling Policies
	// (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html)
	// in the Application Auto Scaling User Guide.
	PolicyType types.PolicyType

	// A step scaling policy. This parameter is required if you are creating a policy
	// and the policy type is StepScaling.
	StepScalingPolicyConfiguration *types.StepScalingPolicyConfiguration

	// A target tracking scaling policy. Includes support for predefined or customized
	// metrics. This parameter is required if you are creating a policy and the policy
	// type is TargetTrackingScaling.
	TargetTrackingScalingPolicyConfiguration *types.TargetTrackingScalingPolicyConfiguration
}

type PutScalingPolicyOutput struct {

	// The Amazon Resource Name (ARN) of the resulting scaling policy.
	//
	// This member is required.
	PolicyARN *string

	// The CloudWatch alarms created for the target tracking scaling policy.
	Alarms []types.Alarm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutScalingPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutScalingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutScalingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutScalingPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutScalingPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutScalingPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "application-autoscaling",
		OperationName: "PutScalingPolicy",
	}
}
