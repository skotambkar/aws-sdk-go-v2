// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/enums"
)

// Gets a specified VPC link under the caller's account in a region.
type GetVpcLinkInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the VpcLink. It is used in an Integration to
	// reference this VpcLink.
	//
	// VpcLinkId is a required field
	VpcLinkId *string `location:"uri" locationName:"vpclink_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVpcLinkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVpcLinkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVpcLinkInput"}

	if s.VpcLinkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcLinkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A API Gateway VPC link for a RestApi to access resources in an Amazon Virtual
// Private Cloud (VPC).
//
// To enable access to a resource in an Amazon Virtual Private Cloud through
// Amazon API Gateway, you, as an API developer, create a VpcLink resource targeted
// for one or more network load balancers of the VPC and then integrate an API
// method with a private integration that uses the VpcLink. The private integration
// has an integration type of HTTP or HTTP_PROXY and has a connection type of
// VPC_LINK. The integration uses the connectionId property to identify the
// VpcLink used.
type GetVpcLinkOutput struct {
	_ struct{} `type:"structure"`

	// The description of the VPC link.
	Description *string `locationName:"description" type:"string"`

	// The identifier of the VpcLink. It is used in an Integration to reference
	// this VpcLink.
	Id *string `locationName:"id" type:"string"`

	// The name used to label and identify the VPC link.
	Name *string `locationName:"name" type:"string"`

	// The status of the VPC link. The valid values are AVAILABLE, PENDING, DELETING,
	// or FAILED. Deploying an API will wait if the status is PENDING and will fail
	// if the status is DELETING.
	Status enums.VpcLinkStatus `locationName:"status" type:"string" enum:"true"`

	// A description about the VPC link status.
	StatusMessage *string `locationName:"statusMessage" type:"string"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The ARNs of network load balancers of the VPC targeted by the VPC link. The
	// network load balancers must be owned by the same AWS account of the API owner.
	TargetArns []string `locationName:"targetArns" type:"list"`
}

// String returns the string representation
func (s GetVpcLinkOutput) String() string {
	return awsutil.Prettify(s)
}
