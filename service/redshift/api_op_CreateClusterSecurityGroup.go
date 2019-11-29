// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opCreateClusterSecurityGroup = "CreateClusterSecurityGroup"

// CreateClusterSecurityGroupRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Creates a new Amazon Redshift security group. You use security groups to
// control access to non-VPC clusters.
//
// For information about managing security groups, go to Amazon Redshift Cluster
// Security Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using CreateClusterSecurityGroupRequest.
//    req := client.CreateClusterSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/CreateClusterSecurityGroup
func (c *Client) CreateClusterSecurityGroupRequest(input *types.CreateClusterSecurityGroupInput) CreateClusterSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opCreateClusterSecurityGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateClusterSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateClusterSecurityGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateClusterSecurityGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateClusterSecurityGroupRequest{Request: req, Input: input, Copy: c.CreateClusterSecurityGroupRequest}
}

// CreateClusterSecurityGroupRequest is the request type for the
// CreateClusterSecurityGroup API operation.
type CreateClusterSecurityGroupRequest struct {
	*aws.Request
	Input *types.CreateClusterSecurityGroupInput
	Copy  func(*types.CreateClusterSecurityGroupInput) CreateClusterSecurityGroupRequest
}

// Send marshals and sends the CreateClusterSecurityGroup API request.
func (r CreateClusterSecurityGroupRequest) Send(ctx context.Context) (*CreateClusterSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClusterSecurityGroupResponse{
		CreateClusterSecurityGroupOutput: r.Request.Data.(*types.CreateClusterSecurityGroupOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClusterSecurityGroupResponse is the response type for the
// CreateClusterSecurityGroup API operation.
type CreateClusterSecurityGroupResponse struct {
	*types.CreateClusterSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateClusterSecurityGroup request.
func (r *CreateClusterSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
