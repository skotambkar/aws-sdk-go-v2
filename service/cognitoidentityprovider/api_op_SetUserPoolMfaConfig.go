// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opSetUserPoolMfaConfig = "SetUserPoolMfaConfig"

// SetUserPoolMfaConfigRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Set the user pool multi-factor authentication (MFA) configuration.
//
//    // Example sending a request using SetUserPoolMfaConfigRequest.
//    req := client.SetUserPoolMfaConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/SetUserPoolMfaConfig
func (c *Client) SetUserPoolMfaConfigRequest(input *types.SetUserPoolMfaConfigInput) SetUserPoolMfaConfigRequest {
	op := &aws.Operation{
		Name:       opSetUserPoolMfaConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetUserPoolMfaConfigInput{}
	}

	req := c.newRequest(op, input, &types.SetUserPoolMfaConfigOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.SetUserPoolMfaConfigMarshaler{Input: input}.GetNamedBuildHandler())

	return SetUserPoolMfaConfigRequest{Request: req, Input: input, Copy: c.SetUserPoolMfaConfigRequest}
}

// SetUserPoolMfaConfigRequest is the request type for the
// SetUserPoolMfaConfig API operation.
type SetUserPoolMfaConfigRequest struct {
	*aws.Request
	Input *types.SetUserPoolMfaConfigInput
	Copy  func(*types.SetUserPoolMfaConfigInput) SetUserPoolMfaConfigRequest
}

// Send marshals and sends the SetUserPoolMfaConfig API request.
func (r SetUserPoolMfaConfigRequest) Send(ctx context.Context) (*SetUserPoolMfaConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetUserPoolMfaConfigResponse{
		SetUserPoolMfaConfigOutput: r.Request.Data.(*types.SetUserPoolMfaConfigOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetUserPoolMfaConfigResponse is the response type for the
// SetUserPoolMfaConfig API operation.
type SetUserPoolMfaConfigResponse struct {
	*types.SetUserPoolMfaConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetUserPoolMfaConfig request.
func (r *SetUserPoolMfaConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
