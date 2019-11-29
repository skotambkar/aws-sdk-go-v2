// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opPutFile = "PutFile"

// PutFileRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Adds or updates a file in a branch in an AWS CodeCommit repository, and generates
// a commit for the addition in the specified branch.
//
//    // Example sending a request using PutFileRequest.
//    req := client.PutFileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PutFile
func (c *Client) PutFileRequest(input *types.PutFileInput) PutFileRequest {
	op := &aws.Operation{
		Name:       opPutFile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutFileInput{}
	}

	req := c.newRequest(op, input, &types.PutFileOutput{})
	return PutFileRequest{Request: req, Input: input, Copy: c.PutFileRequest}
}

// PutFileRequest is the request type for the
// PutFile API operation.
type PutFileRequest struct {
	*aws.Request
	Input *types.PutFileInput
	Copy  func(*types.PutFileInput) PutFileRequest
}

// Send marshals and sends the PutFile API request.
func (r PutFileRequest) Send(ctx context.Context) (*PutFileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutFileResponse{
		PutFileOutput: r.Request.Data.(*types.PutFileOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutFileResponse is the response type for the
// PutFile API operation.
type PutFileResponse struct {
	*types.PutFileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutFile request.
func (r *PutFileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
