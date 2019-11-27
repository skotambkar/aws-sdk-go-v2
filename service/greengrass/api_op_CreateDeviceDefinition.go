// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opCreateDeviceDefinition = "CreateDeviceDefinition"

// CreateDeviceDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a device definition. You may provide the initial version of the device
// definition now or use ''CreateDeviceDefinitionVersion'' at a later time.
//
//    // Example sending a request using CreateDeviceDefinitionRequest.
//    req := client.CreateDeviceDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateDeviceDefinition
func (c *Client) CreateDeviceDefinitionRequest(input *types.CreateDeviceDefinitionInput) CreateDeviceDefinitionRequest {
	op := &aws.Operation{
		Name:       opCreateDeviceDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/definition/devices",
	}

	if input == nil {
		input = &types.CreateDeviceDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.CreateDeviceDefinitionOutput{})
	return CreateDeviceDefinitionRequest{Request: req, Input: input, Copy: c.CreateDeviceDefinitionRequest}
}

// CreateDeviceDefinitionRequest is the request type for the
// CreateDeviceDefinition API operation.
type CreateDeviceDefinitionRequest struct {
	*aws.Request
	Input *types.CreateDeviceDefinitionInput
	Copy  func(*types.CreateDeviceDefinitionInput) CreateDeviceDefinitionRequest
}

// Send marshals and sends the CreateDeviceDefinition API request.
func (r CreateDeviceDefinitionRequest) Send(ctx context.Context) (*CreateDeviceDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDeviceDefinitionResponse{
		CreateDeviceDefinitionOutput: r.Request.Data.(*types.CreateDeviceDefinitionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDeviceDefinitionResponse is the response type for the
// CreateDeviceDefinition API operation.
type CreateDeviceDefinitionResponse struct {
	*types.CreateDeviceDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDeviceDefinition request.
func (r *CreateDeviceDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
