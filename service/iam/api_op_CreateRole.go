// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opCreateRole = "CreateRole"

// CreateRoleRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates a new role for your AWS account. For more information about roles,
// go to IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html).
// For information about limitations on role names and the number of roles you
// can create, go to Limitations on IAM Entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html)
// in the IAM User Guide.
//
//    // Example sending a request using CreateRoleRequest.
//    req := client.CreateRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateRole
func (c *Client) CreateRoleRequest(input *types.CreateRoleInput) CreateRoleRequest {
	op := &aws.Operation{
		Name:       opCreateRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateRoleInput{}
	}

	req := c.newRequest(op, input, &types.CreateRoleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateRoleMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateRoleRequest{Request: req, Input: input, Copy: c.CreateRoleRequest}
}

// CreateRoleRequest is the request type for the
// CreateRole API operation.
type CreateRoleRequest struct {
	*aws.Request
	Input *types.CreateRoleInput
	Copy  func(*types.CreateRoleInput) CreateRoleRequest
}

// Send marshals and sends the CreateRole API request.
func (r CreateRoleRequest) Send(ctx context.Context) (*CreateRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRoleResponse{
		CreateRoleOutput: r.Request.Data.(*types.CreateRoleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRoleResponse is the response type for the
// CreateRole API operation.
type CreateRoleResponse struct {
	*types.CreateRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRole request.
func (r *CreateRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
