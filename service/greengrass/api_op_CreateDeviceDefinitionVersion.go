// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opCreateDeviceDefinitionVersion = "CreateDeviceDefinitionVersion"

// CreateDeviceDefinitionVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a version of a device definition that has already been defined.
//
//    // Example sending a request using CreateDeviceDefinitionVersionRequest.
//    req := client.CreateDeviceDefinitionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateDeviceDefinitionVersion
func (c *Client) CreateDeviceDefinitionVersionRequest(input *types.CreateDeviceDefinitionVersionInput) CreateDeviceDefinitionVersionRequest {
	op := &aws.Operation{
		Name:       opCreateDeviceDefinitionVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/definition/devices/{DeviceDefinitionId}/versions",
	}

	if input == nil {
		input = &types.CreateDeviceDefinitionVersionInput{}
	}

	req := c.newRequest(op, input, &types.CreateDeviceDefinitionVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateDeviceDefinitionVersionMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDeviceDefinitionVersionRequest{Request: req, Input: input, Copy: c.CreateDeviceDefinitionVersionRequest}
}

// CreateDeviceDefinitionVersionRequest is the request type for the
// CreateDeviceDefinitionVersion API operation.
type CreateDeviceDefinitionVersionRequest struct {
	*aws.Request
	Input *types.CreateDeviceDefinitionVersionInput
	Copy  func(*types.CreateDeviceDefinitionVersionInput) CreateDeviceDefinitionVersionRequest
}

// Send marshals and sends the CreateDeviceDefinitionVersion API request.
func (r CreateDeviceDefinitionVersionRequest) Send(ctx context.Context) (*CreateDeviceDefinitionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeviceDefinitionVersionResponse{
		CreateDeviceDefinitionVersionOutput: r.Request.Data.(*types.CreateDeviceDefinitionVersionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeviceDefinitionVersionResponse is the response type for the
// CreateDeviceDefinitionVersion API operation.
type CreateDeviceDefinitionVersionResponse struct {
	*types.CreateDeviceDefinitionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeviceDefinitionVersion request.
func (r *CreateDeviceDefinitionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
