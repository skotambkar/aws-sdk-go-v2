// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opGetConnectorDefinition = "GetConnectorDefinition"

// GetConnectorDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a connector definition.
//
//    // Example sending a request using GetConnectorDefinitionRequest.
//    req := client.GetConnectorDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetConnectorDefinition
func (c *Client) GetConnectorDefinitionRequest(input *types.GetConnectorDefinitionInput) GetConnectorDefinitionRequest {
	op := &aws.Operation{
		Name:       opGetConnectorDefinition,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/connectors/{ConnectorDefinitionId}",
	}

	if input == nil {
		input = &types.GetConnectorDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.GetConnectorDefinitionOutput{})
	return GetConnectorDefinitionRequest{Request: req, Input: input, Copy: c.GetConnectorDefinitionRequest}
}

// GetConnectorDefinitionRequest is the request type for the
// GetConnectorDefinition API operation.
type GetConnectorDefinitionRequest struct {
	*aws.Request
	Input *types.GetConnectorDefinitionInput
	Copy  func(*types.GetConnectorDefinitionInput) GetConnectorDefinitionRequest
}

// Send marshals and sends the GetConnectorDefinition API request.
func (r GetConnectorDefinitionRequest) Send(ctx context.Context) (*GetConnectorDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectorDefinitionResponse{
		GetConnectorDefinitionOutput: r.Request.Data.(*types.GetConnectorDefinitionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConnectorDefinitionResponse is the response type for the
// GetConnectorDefinition API operation.
type GetConnectorDefinitionResponse struct {
	*types.GetConnectorDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnectorDefinition request.
func (r *GetConnectorDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
