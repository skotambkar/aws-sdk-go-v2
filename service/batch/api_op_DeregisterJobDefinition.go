// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/batch/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
)

const opDeregisterJobDefinition = "DeregisterJobDefinition"

// DeregisterJobDefinitionRequest returns a request value for making API operation for
// AWS Batch.
//
// Deregisters an AWS Batch job definition.
//
//    // Example sending a request using DeregisterJobDefinitionRequest.
//    req := client.DeregisterJobDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DeregisterJobDefinition
func (c *Client) DeregisterJobDefinitionRequest(input *types.DeregisterJobDefinitionInput) DeregisterJobDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeregisterJobDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/deregisterjobdefinition",
	}

	if input == nil {
		input = &types.DeregisterJobDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.DeregisterJobDefinitionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeregisterJobDefinitionMarshaler{Input: input}.GetNamedBuildHandler())

	return DeregisterJobDefinitionRequest{Request: req, Input: input, Copy: c.DeregisterJobDefinitionRequest}
}

// DeregisterJobDefinitionRequest is the request type for the
// DeregisterJobDefinition API operation.
type DeregisterJobDefinitionRequest struct {
	*aws.Request
	Input *types.DeregisterJobDefinitionInput
	Copy  func(*types.DeregisterJobDefinitionInput) DeregisterJobDefinitionRequest
}

// Send marshals and sends the DeregisterJobDefinition API request.
func (r DeregisterJobDefinitionRequest) Send(ctx context.Context) (*DeregisterJobDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterJobDefinitionResponse{
		DeregisterJobDefinitionOutput: r.Request.Data.(*types.DeregisterJobDefinitionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterJobDefinitionResponse is the response type for the
// DeregisterJobDefinition API operation.
type DeregisterJobDefinitionResponse struct {
	*types.DeregisterJobDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterJobDefinition request.
func (r *DeregisterJobDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
