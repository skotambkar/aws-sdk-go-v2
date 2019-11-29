// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opDeleteConnectorDefinition = "DeleteConnectorDefinition"

// DeleteConnectorDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Deletes a connector definition.
//
//    // Example sending a request using DeleteConnectorDefinitionRequest.
//    req := client.DeleteConnectorDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DeleteConnectorDefinition
func (c *Client) DeleteConnectorDefinitionRequest(input *types.DeleteConnectorDefinitionInput) DeleteConnectorDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeleteConnectorDefinition,
		HTTPMethod: "DELETE",
		HTTPPath:   "/greengrass/definition/connectors/{ConnectorDefinitionId}",
	}

	if input == nil {
		input = &types.DeleteConnectorDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteConnectorDefinitionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteConnectorDefinitionMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteConnectorDefinitionRequest{Request: req, Input: input, Copy: c.DeleteConnectorDefinitionRequest}
}

// DeleteConnectorDefinitionRequest is the request type for the
// DeleteConnectorDefinition API operation.
type DeleteConnectorDefinitionRequest struct {
	*aws.Request
	Input *types.DeleteConnectorDefinitionInput
	Copy  func(*types.DeleteConnectorDefinitionInput) DeleteConnectorDefinitionRequest
}

// Send marshals and sends the DeleteConnectorDefinition API request.
func (r DeleteConnectorDefinitionRequest) Send(ctx context.Context) (*DeleteConnectorDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConnectorDefinitionResponse{
		DeleteConnectorDefinitionOutput: r.Request.Data.(*types.DeleteConnectorDefinitionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConnectorDefinitionResponse is the response type for the
// DeleteConnectorDefinition API operation.
type DeleteConnectorDefinitionResponse struct {
	*types.DeleteConnectorDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConnectorDefinition request.
func (r *DeleteConnectorDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
