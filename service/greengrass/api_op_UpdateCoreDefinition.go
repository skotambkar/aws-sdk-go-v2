// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opUpdateCoreDefinition = "UpdateCoreDefinition"

// UpdateCoreDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Updates a core definition.
//
//    // Example sending a request using UpdateCoreDefinitionRequest.
//    req := client.UpdateCoreDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateCoreDefinition
func (c *Client) UpdateCoreDefinitionRequest(input *types.UpdateCoreDefinitionInput) UpdateCoreDefinitionRequest {
	op := &aws.Operation{
		Name:       opUpdateCoreDefinition,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/definition/cores/{CoreDefinitionId}",
	}

	if input == nil {
		input = &types.UpdateCoreDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.UpdateCoreDefinitionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateCoreDefinitionMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateCoreDefinitionRequest{Request: req, Input: input, Copy: c.UpdateCoreDefinitionRequest}
}

// UpdateCoreDefinitionRequest is the request type for the
// UpdateCoreDefinition API operation.
type UpdateCoreDefinitionRequest struct {
	*aws.Request
	Input *types.UpdateCoreDefinitionInput
	Copy  func(*types.UpdateCoreDefinitionInput) UpdateCoreDefinitionRequest
}

// Send marshals and sends the UpdateCoreDefinition API request.
func (r UpdateCoreDefinitionRequest) Send(ctx context.Context) (*UpdateCoreDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCoreDefinitionResponse{
		UpdateCoreDefinitionOutput: r.Request.Data.(*types.UpdateCoreDefinitionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCoreDefinitionResponse is the response type for the
// UpdateCoreDefinition API operation.
type UpdateCoreDefinitionResponse struct {
	*types.UpdateCoreDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCoreDefinition request.
func (r *UpdateCoreDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
