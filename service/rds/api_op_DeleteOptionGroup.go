// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/rds/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opDeleteOptionGroup = "DeleteOptionGroup"

// DeleteOptionGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Deletes an existing option group.
//
//    // Example sending a request using DeleteOptionGroupRequest.
//    req := client.DeleteOptionGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteOptionGroup
func (c *Client) DeleteOptionGroupRequest(input *types.DeleteOptionGroupInput) DeleteOptionGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteOptionGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteOptionGroupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteOptionGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteOptionGroupMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteOptionGroupRequest{Request: req, Input: input, Copy: c.DeleteOptionGroupRequest}
}

// DeleteOptionGroupRequest is the request type for the
// DeleteOptionGroup API operation.
type DeleteOptionGroupRequest struct {
	*aws.Request
	Input *types.DeleteOptionGroupInput
	Copy  func(*types.DeleteOptionGroupInput) DeleteOptionGroupRequest
}

// Send marshals and sends the DeleteOptionGroup API request.
func (r DeleteOptionGroupRequest) Send(ctx context.Context) (*DeleteOptionGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteOptionGroupResponse{
		DeleteOptionGroupOutput: r.Request.Data.(*types.DeleteOptionGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteOptionGroupResponse is the response type for the
// DeleteOptionGroup API operation.
type DeleteOptionGroupResponse struct {
	*types.DeleteOptionGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteOptionGroup request.
func (r *DeleteOptionGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
