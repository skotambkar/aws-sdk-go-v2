// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opAdminListDevices = "AdminListDevices"

// AdminListDevicesRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists devices, as an administrator.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using AdminListDevicesRequest.
//    req := client.AdminListDevicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminListDevices
func (c *Client) AdminListDevicesRequest(input *types.AdminListDevicesInput) AdminListDevicesRequest {
	op := &aws.Operation{
		Name:       opAdminListDevices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AdminListDevicesInput{}
	}

	req := c.newRequest(op, input, &types.AdminListDevicesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AdminListDevicesMarshaler{Input: input}.GetNamedBuildHandler())

	return AdminListDevicesRequest{Request: req, Input: input, Copy: c.AdminListDevicesRequest}
}

// AdminListDevicesRequest is the request type for the
// AdminListDevices API operation.
type AdminListDevicesRequest struct {
	*aws.Request
	Input *types.AdminListDevicesInput
	Copy  func(*types.AdminListDevicesInput) AdminListDevicesRequest
}

// Send marshals and sends the AdminListDevices API request.
func (r AdminListDevicesRequest) Send(ctx context.Context) (*AdminListDevicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminListDevicesResponse{
		AdminListDevicesOutput: r.Request.Data.(*types.AdminListDevicesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminListDevicesResponse is the response type for the
// AdminListDevices API operation.
type AdminListDevicesResponse struct {
	*types.AdminListDevicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminListDevices request.
func (r *AdminListDevicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
