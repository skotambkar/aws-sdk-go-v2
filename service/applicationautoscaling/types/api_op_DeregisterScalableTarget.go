// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/enums"
)

type DeregisterScalableTargetInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the resource associated with the scalable target. This
	// string consists of the resource type and unique identifier.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot Fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * EMR cluster - The resource type is instancegroup and the unique identifier
	//    is the cluster ID and instance group ID. Example: instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	//    * AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//    is the fleet name. Example: fleet/sample-fleet.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	//    * Amazon SageMaker endpoint variants - The resource type is variant and
	//    the unique identifier is the resource ID. Example: endpoint/my-end-point/variant/KMeansClustering.
	//
	//    * Custom resources are not supported with a resource type. This parameter
	//    must specify the OutputValue from the CloudFormation template stack used
	//    to access the resources. The unique identifier is defined by the service
	//    provider. More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The scalable dimension associated with the scalable target. This string consists
	// of the service namespace, resource type, and scaling property.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    Fleet request.
	//
	//    * elasticmapreduce:instancegroup:InstanceCount - The instance count of
	//    an EMR Instance Group.
	//
	//    * appstream:fleet:DesiredCapacity - The desired capacity of an AppStream
	//    2.0 fleet.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	//    * sagemaker:variant:DesiredInstanceCount - The number of EC2 instances
	//    for an Amazon SageMaker model endpoint variant.
	//
	//    * custom-resource:ResourceType:Property - The scalable dimension for a
	//    custom resource provided by your own application or service.
	//
	// ScalableDimension is a required field
	ScalableDimension enums.ScalableDimension `type:"string" required:"true" enum:"true"`

	// The namespace of the AWS service that provides the resource or custom-resource
	// for a resource provided by your own application or service. For more information,
	// see AWS Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	// in the Amazon Web Services General Reference.
	//
	// ServiceNamespace is a required field
	ServiceNamespace enums.ServiceNamespace `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DeregisterScalableTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterScalableTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterScalableTargetInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}
	if len(s.ScalableDimension) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ScalableDimension"))
	}
	if len(s.ServiceNamespace) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ServiceNamespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterScalableTargetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterScalableTargetOutput) String() string {
	return awsutil.Prettify(s)
}
