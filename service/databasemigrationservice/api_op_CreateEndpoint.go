// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opCreateEndpoint = "CreateEndpoint"

// CreateEndpointRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Creates an endpoint using the provided settings.
//
//    // Example sending a request using CreateEndpointRequest.
//    req := client.CreateEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/CreateEndpoint
func (c *Client) CreateEndpointRequest(input *types.CreateEndpointInput) CreateEndpointRequest {
	op := &aws.Operation{
		Name:       opCreateEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateEndpointInput{}
	}

	req := c.newRequest(op, input, &types.CreateEndpointOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateEndpointMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateEndpointRequest{Request: req, Input: input, Copy: c.CreateEndpointRequest}
}

// CreateEndpointRequest is the request type for the
// CreateEndpoint API operation.
type CreateEndpointRequest struct {
	*aws.Request
	Input *types.CreateEndpointInput
	Copy  func(*types.CreateEndpointInput) CreateEndpointRequest
}

// Send marshals and sends the CreateEndpoint API request.
func (r CreateEndpointRequest) Send(ctx context.Context) (*CreateEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEndpointResponse{
		CreateEndpointOutput: r.Request.Data.(*types.CreateEndpointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEndpointResponse is the response type for the
// CreateEndpoint API operation.
type CreateEndpointResponse struct {
	*types.CreateEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEndpoint request.
func (r *CreateEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
