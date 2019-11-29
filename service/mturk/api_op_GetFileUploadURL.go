// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
)

const opGetFileUploadURL = "GetFileUploadURL"

// GetFileUploadURLRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The GetFileUploadURL operation generates and returns a temporary URL. You
// use the temporary URL to retrieve a file uploaded by a Worker as an answer
// to a FileUploadAnswer question for a HIT. The temporary URL is generated
// the instant the GetFileUploadURL operation is called, and is valid for 60
// seconds. You can get a temporary file upload URL any time until the HIT is
// disposed. After the HIT is disposed, any uploaded files are deleted, and
// cannot be retrieved. Pending Deprecation on December 12, 2017. The Answer
// Specification structure will no longer support the FileUploadAnswer element
// to be used for the QuestionForm data structure. Instead, we recommend that
// Requesters who want to create HITs asking Workers to upload files to use
// Amazon S3.
//
//    // Example sending a request using GetFileUploadURLRequest.
//    req := client.GetFileUploadURLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/GetFileUploadURL
func (c *Client) GetFileUploadURLRequest(input *types.GetFileUploadURLInput) GetFileUploadURLRequest {
	op := &aws.Operation{
		Name:       opGetFileUploadURL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetFileUploadURLInput{}
	}

	req := c.newRequest(op, input, &types.GetFileUploadURLOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetFileUploadURLMarshaler{Input: input}.GetNamedBuildHandler())

	return GetFileUploadURLRequest{Request: req, Input: input, Copy: c.GetFileUploadURLRequest}
}

// GetFileUploadURLRequest is the request type for the
// GetFileUploadURL API operation.
type GetFileUploadURLRequest struct {
	*aws.Request
	Input *types.GetFileUploadURLInput
	Copy  func(*types.GetFileUploadURLInput) GetFileUploadURLRequest
}

// Send marshals and sends the GetFileUploadURL API request.
func (r GetFileUploadURLRequest) Send(ctx context.Context) (*GetFileUploadURLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFileUploadURLResponse{
		GetFileUploadURLOutput: r.Request.Data.(*types.GetFileUploadURLOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFileUploadURLResponse is the response type for the
// GetFileUploadURL API operation.
type GetFileUploadURLResponse struct {
	*types.GetFileUploadURLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFileUploadURL request.
func (r *GetFileUploadURLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
