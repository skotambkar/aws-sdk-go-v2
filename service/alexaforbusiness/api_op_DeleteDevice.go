// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opDeleteDevice = "DeleteDevice"

// DeleteDeviceRequest returns a request value for making API operation for
// Alexa For Business.
//
// Removes a device from Alexa For Business.
//
//    // Example sending a request using DeleteDeviceRequest.
//    req := client.DeleteDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteDevice
func (c *Client) DeleteDeviceRequest(input *types.DeleteDeviceInput) DeleteDeviceRequest {
	op := &aws.Operation{
		Name:       opDeleteDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteDeviceInput{}
	}

	req := c.newRequest(op, input, &types.DeleteDeviceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteDeviceMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteDeviceRequest{Request: req, Input: input, Copy: c.DeleteDeviceRequest}
}

// DeleteDeviceRequest is the request type for the
// DeleteDevice API operation.
type DeleteDeviceRequest struct {
	*aws.Request
	Input *types.DeleteDeviceInput
	Copy  func(*types.DeleteDeviceInput) DeleteDeviceRequest
}

// Send marshals and sends the DeleteDevice API request.
func (r DeleteDeviceRequest) Send(ctx context.Context) (*DeleteDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDeviceResponse{
		DeleteDeviceOutput: r.Request.Data.(*types.DeleteDeviceOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDeviceResponse is the response type for the
// DeleteDevice API operation.
type DeleteDeviceResponse struct {
	*types.DeleteDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDevice request.
func (r *DeleteDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
