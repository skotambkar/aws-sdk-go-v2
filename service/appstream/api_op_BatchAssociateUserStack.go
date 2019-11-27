// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

const opBatchAssociateUserStack = "BatchAssociateUserStack"

// BatchAssociateUserStackRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Associates the specified users with the specified stacks. Users in a user
// pool cannot be assigned to stacks with fleets that are joined to an Active
// Directory domain.
//
//    // Example sending a request using BatchAssociateUserStackRequest.
//    req := client.BatchAssociateUserStackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/BatchAssociateUserStack
func (c *Client) BatchAssociateUserStackRequest(input *types.BatchAssociateUserStackInput) BatchAssociateUserStackRequest {
	op := &aws.Operation{
		Name:       opBatchAssociateUserStack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchAssociateUserStackInput{}
	}

	req := c.newRequest(op, input, &types.BatchAssociateUserStackOutput{})
	return BatchAssociateUserStackRequest{Request: req, Input: input, Copy: c.BatchAssociateUserStackRequest}
}

// BatchAssociateUserStackRequest is the request type for the
// BatchAssociateUserStack API operation.
type BatchAssociateUserStackRequest struct {
	*aws.Request
	Input *types.BatchAssociateUserStackInput
	Copy  func(*types.BatchAssociateUserStackInput) BatchAssociateUserStackRequest
}

// Send marshals and sends the BatchAssociateUserStack API request.
func (r BatchAssociateUserStackRequest) Send(ctx context.Context) (*BatchAssociateUserStackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchAssociateUserStackResponse{
		BatchAssociateUserStackOutput: r.Request.Data.(*types.BatchAssociateUserStackOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchAssociateUserStackResponse is the response type for the
// BatchAssociateUserStack API operation.
type BatchAssociateUserStackResponse struct {
	*types.BatchAssociateUserStackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchAssociateUserStack request.
func (r *BatchAssociateUserStackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
