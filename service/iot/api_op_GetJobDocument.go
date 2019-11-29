// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opGetJobDocument = "GetJobDocument"

// GetJobDocumentRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets a job document.
//
//    // Example sending a request using GetJobDocumentRequest.
//    req := client.GetJobDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetJobDocumentRequest(input *types.GetJobDocumentInput) GetJobDocumentRequest {
	op := &aws.Operation{
		Name:       opGetJobDocument,
		HTTPMethod: "GET",
		HTTPPath:   "/jobs/{jobId}/job-document",
	}

	if input == nil {
		input = &types.GetJobDocumentInput{}
	}

	req := c.newRequest(op, input, &types.GetJobDocumentOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetJobDocumentMarshaler{Input: input}.GetNamedBuildHandler())

	return GetJobDocumentRequest{Request: req, Input: input, Copy: c.GetJobDocumentRequest}
}

// GetJobDocumentRequest is the request type for the
// GetJobDocument API operation.
type GetJobDocumentRequest struct {
	*aws.Request
	Input *types.GetJobDocumentInput
	Copy  func(*types.GetJobDocumentInput) GetJobDocumentRequest
}

// Send marshals and sends the GetJobDocument API request.
func (r GetJobDocumentRequest) Send(ctx context.Context) (*GetJobDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobDocumentResponse{
		GetJobDocumentOutput: r.Request.Data.(*types.GetJobDocumentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobDocumentResponse is the response type for the
// GetJobDocument API operation.
type GetJobDocumentResponse struct {
	*types.GetJobDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJobDocument request.
func (r *GetJobDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
