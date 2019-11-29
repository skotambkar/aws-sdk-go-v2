// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opCreateCommit = "CreateCommit"

// CreateCommitRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Creates a commit for a repository on the tip of a specified branch.
//
//    // Example sending a request using CreateCommitRequest.
//    req := client.CreateCommitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/CreateCommit
func (c *Client) CreateCommitRequest(input *types.CreateCommitInput) CreateCommitRequest {
	op := &aws.Operation{
		Name:       opCreateCommit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateCommitInput{}
	}

	req := c.newRequest(op, input, &types.CreateCommitOutput{})
	return CreateCommitRequest{Request: req, Input: input, Copy: c.CreateCommitRequest}
}

// CreateCommitRequest is the request type for the
// CreateCommit API operation.
type CreateCommitRequest struct {
	*aws.Request
	Input *types.CreateCommitInput
	Copy  func(*types.CreateCommitInput) CreateCommitRequest
}

// Send marshals and sends the CreateCommit API request.
func (r CreateCommitRequest) Send(ctx context.Context) (*CreateCommitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCommitResponse{
		CreateCommitOutput: r.Request.Data.(*types.CreateCommitOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCommitResponse is the response type for the
// CreateCommit API operation.
type CreateCommitResponse struct {
	*types.CreateCommitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCommit request.
func (r *CreateCommitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
