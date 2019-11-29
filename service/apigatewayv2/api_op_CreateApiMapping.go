// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opCreateApiMapping = "CreateApiMapping"

// CreateApiMappingRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Creates an API mapping.
//
//    // Example sending a request using CreateApiMappingRequest.
//    req := client.CreateApiMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateApiMapping
func (c *Client) CreateApiMappingRequest(input *types.CreateApiMappingInput) CreateApiMappingRequest {
	op := &aws.Operation{
		Name:       opCreateApiMapping,
		HTTPMethod: "POST",
		HTTPPath:   "/v2/domainnames/{domainName}/apimappings",
	}

	if input == nil {
		input = &types.CreateApiMappingInput{}
	}

	req := c.newRequest(op, input, &types.CreateApiMappingOutput{})
	return CreateApiMappingRequest{Request: req, Input: input, Copy: c.CreateApiMappingRequest}
}

// CreateApiMappingRequest is the request type for the
// CreateApiMapping API operation.
type CreateApiMappingRequest struct {
	*aws.Request
	Input *types.CreateApiMappingInput
	Copy  func(*types.CreateApiMappingInput) CreateApiMappingRequest
}

// Send marshals and sends the CreateApiMapping API request.
func (r CreateApiMappingRequest) Send(ctx context.Context) (*CreateApiMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApiMappingResponse{
		CreateApiMappingOutput: r.Request.Data.(*types.CreateApiMappingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApiMappingResponse is the response type for the
// CreateApiMapping API operation.
type CreateApiMappingResponse struct {
	*types.CreateApiMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApiMapping request.
func (r *CreateApiMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
