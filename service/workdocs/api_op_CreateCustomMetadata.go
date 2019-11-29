// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
)

const opCreateCustomMetadata = "CreateCustomMetadata"

// CreateCustomMetadataRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Adds one or more custom properties to the specified resource (a folder, document,
// or version).
//
//    // Example sending a request using CreateCustomMetadataRequest.
//    req := client.CreateCustomMetadataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/CreateCustomMetadata
func (c *Client) CreateCustomMetadataRequest(input *types.CreateCustomMetadataInput) CreateCustomMetadataRequest {
	op := &aws.Operation{
		Name:       opCreateCustomMetadata,
		HTTPMethod: "PUT",
		HTTPPath:   "/api/v1/resources/{ResourceId}/customMetadata",
	}

	if input == nil {
		input = &types.CreateCustomMetadataInput{}
	}

	req := c.newRequest(op, input, &types.CreateCustomMetadataOutput{})
	return CreateCustomMetadataRequest{Request: req, Input: input, Copy: c.CreateCustomMetadataRequest}
}

// CreateCustomMetadataRequest is the request type for the
// CreateCustomMetadata API operation.
type CreateCustomMetadataRequest struct {
	*aws.Request
	Input *types.CreateCustomMetadataInput
	Copy  func(*types.CreateCustomMetadataInput) CreateCustomMetadataRequest
}

// Send marshals and sends the CreateCustomMetadata API request.
func (r CreateCustomMetadataRequest) Send(ctx context.Context) (*CreateCustomMetadataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCustomMetadataResponse{
		CreateCustomMetadataOutput: r.Request.Data.(*types.CreateCustomMetadataOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCustomMetadataResponse is the response type for the
// CreateCustomMetadata API operation.
type CreateCustomMetadataResponse struct {
	*types.CreateCustomMetadataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCustomMetadata request.
func (r *CreateCustomMetadataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
