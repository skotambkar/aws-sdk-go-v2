// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opPutRepositoryTriggers = "PutRepositoryTriggers"

// PutRepositoryTriggersRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Replaces all triggers for a repository. This can be used to create or delete
// triggers.
//
//    // Example sending a request using PutRepositoryTriggersRequest.
//    req := client.PutRepositoryTriggersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PutRepositoryTriggers
func (c *Client) PutRepositoryTriggersRequest(input *types.PutRepositoryTriggersInput) PutRepositoryTriggersRequest {
	op := &aws.Operation{
		Name:       opPutRepositoryTriggers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutRepositoryTriggersInput{}
	}

	req := c.newRequest(op, input, &types.PutRepositoryTriggersOutput{})
	return PutRepositoryTriggersRequest{Request: req, Input: input, Copy: c.PutRepositoryTriggersRequest}
}

// PutRepositoryTriggersRequest is the request type for the
// PutRepositoryTriggers API operation.
type PutRepositoryTriggersRequest struct {
	*aws.Request
	Input *types.PutRepositoryTriggersInput
	Copy  func(*types.PutRepositoryTriggersInput) PutRepositoryTriggersRequest
}

// Send marshals and sends the PutRepositoryTriggers API request.
func (r PutRepositoryTriggersRequest) Send(ctx context.Context) (*PutRepositoryTriggersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRepositoryTriggersResponse{
		PutRepositoryTriggersOutput: r.Request.Data.(*types.PutRepositoryTriggersOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRepositoryTriggersResponse is the response type for the
// PutRepositoryTriggers API operation.
type PutRepositoryTriggersResponse struct {
	*types.PutRepositoryTriggersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRepositoryTriggers request.
func (r *PutRepositoryTriggersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
