// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateSecurityGroup = "CreateSecurityGroup"

// CreateSecurityGroupRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a security group.
//
// A security group acts as a virtual firewall for your instance to control
// inbound and outbound traffic. For more information, see Amazon EC2 Security
// Groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html)
// in the Amazon Elastic Compute Cloud User Guide and Security Groups for Your
// VPC (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
// in the Amazon Virtual Private Cloud User Guide.
//
// When you create a security group, you specify a friendly name of your choice.
// You can have a security group for use in EC2-Classic with the same name as
// a security group for use in a VPC. However, you can't have two security groups
// for use in EC2-Classic with the same name or two security groups for use
// in a VPC with the same name.
//
// You have a default security group for use in EC2-Classic and a default security
// group for use in your VPC. If you don't specify a security group when you
// launch an instance, the instance is launched into the appropriate default
// security group. A default security group includes a default rule that grants
// instances unrestricted network access to each other.
//
// You can add or remove rules from your security groups using AuthorizeSecurityGroupIngress,
// AuthorizeSecurityGroupEgress, RevokeSecurityGroupIngress, and RevokeSecurityGroupEgress.
//
// For more information about VPC security group limits, see Amazon VPC Limits
// (https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html).
//
//    // Example sending a request using CreateSecurityGroupRequest.
//    req := client.CreateSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateSecurityGroup
func (c *Client) CreateSecurityGroupRequest(input *types.CreateSecurityGroupInput) CreateSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opCreateSecurityGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateSecurityGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.CreateSecurityGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateSecurityGroupRequest{Request: req, Input: input, Copy: c.CreateSecurityGroupRequest}
}

// CreateSecurityGroupRequest is the request type for the
// CreateSecurityGroup API operation.
type CreateSecurityGroupRequest struct {
	*aws.Request
	Input *types.CreateSecurityGroupInput
	Copy  func(*types.CreateSecurityGroupInput) CreateSecurityGroupRequest
}

// Send marshals and sends the CreateSecurityGroup API request.
func (r CreateSecurityGroupRequest) Send(ctx context.Context) (*CreateSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSecurityGroupResponse{
		CreateSecurityGroupOutput: r.Request.Data.(*types.CreateSecurityGroupOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSecurityGroupResponse is the response type for the
// CreateSecurityGroup API operation.
type CreateSecurityGroupResponse struct {
	*types.CreateSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSecurityGroup request.
func (r *CreateSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
