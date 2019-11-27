// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opGetMergeCommit = "GetMergeCommit"

// GetMergeCommitRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about a specified merge commit.
//
//    // Example sending a request using GetMergeCommitRequest.
//    req := client.GetMergeCommitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetMergeCommit
func (c *Client) GetMergeCommitRequest(input *types.GetMergeCommitInput) GetMergeCommitRequest {
	op := &aws.Operation{
		Name:       opGetMergeCommit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetMergeCommitInput{}
	}

	req := c.newRequest(op, input, &types.GetMergeCommitOutput{})
	return GetMergeCommitRequest{Request: req, Input: input, Copy: c.GetMergeCommitRequest}
}

// GetMergeCommitRequest is the request type for the
// GetMergeCommit API operation.
type GetMergeCommitRequest struct {
	*aws.Request
	Input *types.GetMergeCommitInput
	Copy  func(*types.GetMergeCommitInput) GetMergeCommitRequest
}

// Send marshals and sends the GetMergeCommit API request.
func (r GetMergeCommitRequest) Send(ctx context.Context) (*GetMergeCommitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMergeCommitResponse{
		GetMergeCommitOutput: r.Request.Data.(*types.GetMergeCommitOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMergeCommitResponse is the response type for the
// GetMergeCommit API operation.
type GetMergeCommitResponse struct {
	*types.GetMergeCommitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMergeCommit request.
func (r *GetMergeCommitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
