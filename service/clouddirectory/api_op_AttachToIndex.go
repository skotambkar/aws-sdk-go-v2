// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opAttachToIndex = "AttachToIndex"

// AttachToIndexRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Attaches the specified object to the specified index.
//
//    // Example sending a request using AttachToIndexRequest.
//    req := client.AttachToIndexRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/AttachToIndex
func (c *Client) AttachToIndexRequest(input *types.AttachToIndexInput) AttachToIndexRequest {
	op := &aws.Operation{
		Name:       opAttachToIndex,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/index/attach",
	}

	if input == nil {
		input = &types.AttachToIndexInput{}
	}

	req := c.newRequest(op, input, &types.AttachToIndexOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.AttachToIndexMarshaler{Input: input}.GetNamedBuildHandler())

	return AttachToIndexRequest{Request: req, Input: input, Copy: c.AttachToIndexRequest}
}

// AttachToIndexRequest is the request type for the
// AttachToIndex API operation.
type AttachToIndexRequest struct {
	*aws.Request
	Input *types.AttachToIndexInput
	Copy  func(*types.AttachToIndexInput) AttachToIndexRequest
}

// Send marshals and sends the AttachToIndex API request.
func (r AttachToIndexRequest) Send(ctx context.Context) (*AttachToIndexResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachToIndexResponse{
		AttachToIndexOutput: r.Request.Data.(*types.AttachToIndexOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachToIndexResponse is the response type for the
// AttachToIndex API operation.
type AttachToIndexResponse struct {
	*types.AttachToIndexOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachToIndex request.
func (r *AttachToIndexResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
