// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dax/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
)

const opDeleteSubnetGroup = "DeleteSubnetGroup"

// DeleteSubnetGroupRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Deletes a subnet group.
//
// You cannot delete a subnet group if it is associated with any DAX clusters.
//
//    // Example sending a request using DeleteSubnetGroupRequest.
//    req := client.DeleteSubnetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DeleteSubnetGroup
func (c *Client) DeleteSubnetGroupRequest(input *types.DeleteSubnetGroupInput) DeleteSubnetGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteSubnetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteSubnetGroupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteSubnetGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteSubnetGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteSubnetGroupRequest{Request: req, Input: input, Copy: c.DeleteSubnetGroupRequest}
}

// DeleteSubnetGroupRequest is the request type for the
// DeleteSubnetGroup API operation.
type DeleteSubnetGroupRequest struct {
	*aws.Request
	Input *types.DeleteSubnetGroupInput
	Copy  func(*types.DeleteSubnetGroupInput) DeleteSubnetGroupRequest
}

// Send marshals and sends the DeleteSubnetGroup API request.
func (r DeleteSubnetGroupRequest) Send(ctx context.Context) (*DeleteSubnetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSubnetGroupResponse{
		DeleteSubnetGroupOutput: r.Request.Data.(*types.DeleteSubnetGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSubnetGroupResponse is the response type for the
// DeleteSubnetGroup API operation.
type DeleteSubnetGroupResponse struct {
	*types.DeleteSubnetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSubnetGroup request.
func (r *DeleteSubnetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
