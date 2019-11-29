// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/appsync/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
)

const opStartSchemaCreation = "StartSchemaCreation"

// StartSchemaCreationRequest returns a request value for making API operation for
// AWS AppSync.
//
// Adds a new schema to your GraphQL API.
//
// This operation is asynchronous. Use to determine when it has completed.
//
//    // Example sending a request using StartSchemaCreationRequest.
//    req := client.StartSchemaCreationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/StartSchemaCreation
func (c *Client) StartSchemaCreationRequest(input *types.StartSchemaCreationInput) StartSchemaCreationRequest {
	op := &aws.Operation{
		Name:       opStartSchemaCreation,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/schemacreation",
	}

	if input == nil {
		input = &types.StartSchemaCreationInput{}
	}

	req := c.newRequest(op, input, &types.StartSchemaCreationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.StartSchemaCreationMarshaler{Input: input}.GetNamedBuildHandler())

	return StartSchemaCreationRequest{Request: req, Input: input, Copy: c.StartSchemaCreationRequest}
}

// StartSchemaCreationRequest is the request type for the
// StartSchemaCreation API operation.
type StartSchemaCreationRequest struct {
	*aws.Request
	Input *types.StartSchemaCreationInput
	Copy  func(*types.StartSchemaCreationInput) StartSchemaCreationRequest
}

// Send marshals and sends the StartSchemaCreation API request.
func (r StartSchemaCreationRequest) Send(ctx context.Context) (*StartSchemaCreationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartSchemaCreationResponse{
		StartSchemaCreationOutput: r.Request.Data.(*types.StartSchemaCreationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartSchemaCreationResponse is the response type for the
// StartSchemaCreation API operation.
type StartSchemaCreationResponse struct {
	*types.StartSchemaCreationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartSchemaCreation request.
func (r *StartSchemaCreationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
