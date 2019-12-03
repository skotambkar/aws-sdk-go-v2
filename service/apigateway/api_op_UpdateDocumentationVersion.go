// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opUpdateDocumentationVersion = "UpdateDocumentationVersion"

// UpdateDocumentationVersionRequest returns a request value for making API operation for
// Amazon API Gateway.
//
//    // Example sending a request using UpdateDocumentationVersionRequest.
//    req := client.UpdateDocumentationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateDocumentationVersionRequest(input *types.UpdateDocumentationVersionInput) UpdateDocumentationVersionRequest {
	op := &aws.Operation{
		Name:       opUpdateDocumentationVersion,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}/documentation/versions/{doc_version}",
	}

	if input == nil {
		input = &types.UpdateDocumentationVersionInput{}
	}

	req := c.newRequest(op, input, &types.UpdateDocumentationVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateDocumentationVersionMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateDocumentationVersionRequest{Request: req, Input: input, Copy: c.UpdateDocumentationVersionRequest}
}

// UpdateDocumentationVersionRequest is the request type for the
// UpdateDocumentationVersion API operation.
type UpdateDocumentationVersionRequest struct {
	*aws.Request
	Input *types.UpdateDocumentationVersionInput
	Copy  func(*types.UpdateDocumentationVersionInput) UpdateDocumentationVersionRequest
}

// Send marshals and sends the UpdateDocumentationVersion API request.
func (r UpdateDocumentationVersionRequest) Send(ctx context.Context) (*UpdateDocumentationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDocumentationVersionResponse{
		UpdateDocumentationVersionOutput: r.Request.Data.(*types.UpdateDocumentationVersionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDocumentationVersionResponse is the response type for the
// UpdateDocumentationVersion API operation.
type UpdateDocumentationVersionResponse struct {
	*types.UpdateDocumentationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDocumentationVersion request.
func (r *UpdateDocumentationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
