// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dax/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
)

const opCreateSubnetGroup = "CreateSubnetGroup"

// CreateSubnetGroupRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Creates a new subnet group.
//
//    // Example sending a request using CreateSubnetGroupRequest.
//    req := client.CreateSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/CreateSubnetGroup
func (c *Client) CreateSubnetGroupRequest(input *types.CreateSubnetGroupInput) CreateSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateSubnetGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateSubnetGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateSubnetGroupRequest{Request: req, Input: input, Copy: c.CreateSubnetGroupRequest}
}

// CreateSubnetGroupRequest is the request type for the
// CreateSubnetGroup API operation.
type CreateSubnetGroupRequest struct {
	*aws.Request
	Input *types.CreateSubnetGroupInput
	Copy  func(*types.CreateSubnetGroupInput) CreateSubnetGroupRequest
}

// Send marshals and sends the CreateSubnetGroup API request.
func (r CreateSubnetGroupRequest) Send(ctx context.Context) (*CreateSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSubnetGroupResponse{
		CreateSubnetGroupOutput: r.Request.Data.(*types.CreateSubnetGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSubnetGroupResponse is the response type for the
// CreateSubnetGroup API operation.
type CreateSubnetGroupResponse struct {
	*types.CreateSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSubnetGroup request.
func (r *CreateSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
