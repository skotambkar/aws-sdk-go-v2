// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/rds/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opCreateDBSecurityGroup = "CreateDBSecurityGroup"

// CreateDBSecurityGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new DB security group. DB security groups control access to a DB
// instance.
//
// A DB security group controls access to EC2-Classic DB instances that are
// not in a VPC.
//
//    // Example sending a request using CreateDBSecurityGroupRequest.
//    req := client.CreateDBSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBSecurityGroup
func (c *Client) CreateDBSecurityGroupRequest(input *types.CreateDBSecurityGroupInput) CreateDBSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opCreateDBSecurityGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDBSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateDBSecurityGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateDBSecurityGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDBSecurityGroupRequest{Request: req, Input: input, Copy: c.CreateDBSecurityGroupRequest}
}

// CreateDBSecurityGroupRequest is the request type for the
// CreateDBSecurityGroup API operation.
type CreateDBSecurityGroupRequest struct {
	*aws.Request
	Input *types.CreateDBSecurityGroupInput
	Copy  func(*types.CreateDBSecurityGroupInput) CreateDBSecurityGroupRequest
}

// Send marshals and sends the CreateDBSecurityGroup API request.
func (r CreateDBSecurityGroupRequest) Send(ctx context.Context) (*CreateDBSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBSecurityGroupResponse{
		CreateDBSecurityGroupOutput: r.Request.Data.(*types.CreateDBSecurityGroupOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBSecurityGroupResponse is the response type for the
// CreateDBSecurityGroup API operation.
type CreateDBSecurityGroupResponse struct {
	*types.CreateDBSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBSecurityGroup request.
func (r *CreateDBSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
