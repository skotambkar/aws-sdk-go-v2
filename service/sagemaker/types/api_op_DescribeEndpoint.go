// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/enums"
)

type DescribeEndpointInput struct {
	_ struct{} `type:"structure"`

	// The name of the endpoint.
	//
	// EndpointName is a required field
	EndpointName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEndpointInput"}

	if s.EndpointName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEndpointOutput struct {
	_ struct{} `type:"structure"`

	// A timestamp that shows when the endpoint was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// The Amazon Resource Name (ARN) of the endpoint.
	//
	// EndpointArn is a required field
	EndpointArn *string `min:"20" type:"string" required:"true"`

	// The name of the endpoint configuration associated with this endpoint.
	//
	// EndpointConfigName is a required field
	EndpointConfigName *string `type:"string" required:"true"`

	// Name of the endpoint.
	//
	// EndpointName is a required field
	EndpointName *string `type:"string" required:"true"`

	// The status of the endpoint.
	//
	//    * OutOfService: Endpoint is not available to take incoming requests.
	//
	//    * Creating: CreateEndpoint is executing.
	//
	//    * Updating: UpdateEndpoint or UpdateEndpointWeightsAndCapacities is executing.
	//
	//    * SystemUpdating: Endpoint is undergoing maintenance and cannot be updated
	//    or deleted or re-scaled until it has completed. This maintenance operation
	//    does not change any customer-specified values such as VPC config, KMS
	//    encryption, model, instance type, or instance count.
	//
	//    * RollingBack: Endpoint fails to scale up or down or change its variant
	//    weight and is in the process of rolling back to its previous configuration.
	//    Once the rollback completes, endpoint returns to an InService status.
	//    This transitional status only applies to an endpoint that has autoscaling
	//    enabled and is undergoing variant weight or capacity changes as part of
	//    an UpdateEndpointWeightsAndCapacities call or when the UpdateEndpointWeightsAndCapacities
	//    operation is called explicitly.
	//
	//    * InService: Endpoint is available to process incoming requests.
	//
	//    * Deleting: DeleteEndpoint is executing.
	//
	//    * Failed: Endpoint could not be created, updated, or re-scaled. Use DescribeEndpointOutput$FailureReason
	//    for information about the failure. DeleteEndpoint is the only operation
	//    that can be performed on a failed endpoint.
	//
	// EndpointStatus is a required field
	EndpointStatus enums.EndpointStatus `type:"string" required:"true" enum:"true"`

	// If the status of the endpoint is Failed, the reason why it failed.
	FailureReason *string `type:"string"`

	// A timestamp that shows when the endpoint was last modified.
	//
	// LastModifiedTime is a required field
	LastModifiedTime *time.Time `type:"timestamp" required:"true"`

	// An array of ProductionVariantSummary objects, one for each model hosted behind
	// this endpoint.
	ProductionVariants []ProductionVariantSummary `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}
