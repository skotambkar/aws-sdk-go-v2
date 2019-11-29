// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opGetDeviceDefinition = "GetDeviceDefinition"

// GetDeviceDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a device definition.
//
//    // Example sending a request using GetDeviceDefinitionRequest.
//    req := client.GetDeviceDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetDeviceDefinition
func (c *Client) GetDeviceDefinitionRequest(input *types.GetDeviceDefinitionInput) GetDeviceDefinitionRequest {
	op := &aws.Operation{
		Name:       opGetDeviceDefinition,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/devices/{DeviceDefinitionId}",
	}

	if input == nil {
		input = &types.GetDeviceDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.GetDeviceDefinitionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetDeviceDefinitionMarshaler{Input: input}.GetNamedBuildHandler())

	return GetDeviceDefinitionRequest{Request: req, Input: input, Copy: c.GetDeviceDefinitionRequest}
}

// GetDeviceDefinitionRequest is the request type for the
// GetDeviceDefinition API operation.
type GetDeviceDefinitionRequest struct {
	*aws.Request
	Input *types.GetDeviceDefinitionInput
	Copy  func(*types.GetDeviceDefinitionInput) GetDeviceDefinitionRequest
}

// Send marshals and sends the GetDeviceDefinition API request.
func (r GetDeviceDefinitionRequest) Send(ctx context.Context) (*GetDeviceDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeviceDefinitionResponse{
		GetDeviceDefinitionOutput: r.Request.Data.(*types.GetDeviceDefinitionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeviceDefinitionResponse is the response type for the
// GetDeviceDefinition API operation.
type GetDeviceDefinitionResponse struct {
	*types.GetDeviceDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeviceDefinition request.
func (r *GetDeviceDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
