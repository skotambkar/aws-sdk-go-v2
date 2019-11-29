// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opGetDeviceDefinitionVersion = "GetDeviceDefinitionVersion"

// GetDeviceDefinitionVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a device definition version.
//
//    // Example sending a request using GetDeviceDefinitionVersionRequest.
//    req := client.GetDeviceDefinitionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetDeviceDefinitionVersion
func (c *Client) GetDeviceDefinitionVersionRequest(input *types.GetDeviceDefinitionVersionInput) GetDeviceDefinitionVersionRequest {
	op := &aws.Operation{
		Name:       opGetDeviceDefinitionVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/devices/{DeviceDefinitionId}/versions/{DeviceDefinitionVersionId}",
	}

	if input == nil {
		input = &types.GetDeviceDefinitionVersionInput{}
	}

	req := c.newRequest(op, input, &types.GetDeviceDefinitionVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetDeviceDefinitionVersionMarshaler{Input: input}.GetNamedBuildHandler())

	return GetDeviceDefinitionVersionRequest{Request: req, Input: input, Copy: c.GetDeviceDefinitionVersionRequest}
}

// GetDeviceDefinitionVersionRequest is the request type for the
// GetDeviceDefinitionVersion API operation.
type GetDeviceDefinitionVersionRequest struct {
	*aws.Request
	Input *types.GetDeviceDefinitionVersionInput
	Copy  func(*types.GetDeviceDefinitionVersionInput) GetDeviceDefinitionVersionRequest
}

// Send marshals and sends the GetDeviceDefinitionVersion API request.
func (r GetDeviceDefinitionVersionRequest) Send(ctx context.Context) (*GetDeviceDefinitionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeviceDefinitionVersionResponse{
		GetDeviceDefinitionVersionOutput: r.Request.Data.(*types.GetDeviceDefinitionVersionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeviceDefinitionVersionResponse is the response type for the
// GetDeviceDefinitionVersion API operation.
type GetDeviceDefinitionVersionResponse struct {
	*types.GetDeviceDefinitionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeviceDefinitionVersion request.
func (r *GetDeviceDefinitionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
