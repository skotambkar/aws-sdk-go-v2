// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opUpdateDeviceStatus = "UpdateDeviceStatus"

// UpdateDeviceStatusRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Updates the device status.
//
//    // Example sending a request using UpdateDeviceStatusRequest.
//    req := client.UpdateDeviceStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/UpdateDeviceStatus
func (c *Client) UpdateDeviceStatusRequest(input *types.UpdateDeviceStatusInput) UpdateDeviceStatusRequest {
	op := &aws.Operation{
		Name:       opUpdateDeviceStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateDeviceStatusInput{}
	}

	req := c.newRequest(op, input, &types.UpdateDeviceStatusOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateDeviceStatusMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateDeviceStatusRequest{Request: req, Input: input, Copy: c.UpdateDeviceStatusRequest}
}

// UpdateDeviceStatusRequest is the request type for the
// UpdateDeviceStatus API operation.
type UpdateDeviceStatusRequest struct {
	*aws.Request
	Input *types.UpdateDeviceStatusInput
	Copy  func(*types.UpdateDeviceStatusInput) UpdateDeviceStatusRequest
}

// Send marshals and sends the UpdateDeviceStatus API request.
func (r UpdateDeviceStatusRequest) Send(ctx context.Context) (*UpdateDeviceStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDeviceStatusResponse{
		UpdateDeviceStatusOutput: r.Request.Data.(*types.UpdateDeviceStatusOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDeviceStatusResponse is the response type for the
// UpdateDeviceStatus API operation.
type UpdateDeviceStatusResponse struct {
	*types.UpdateDeviceStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDeviceStatus request.
func (r *UpdateDeviceStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
