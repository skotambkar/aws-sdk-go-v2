// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opAdminSetUserSettings = "AdminSetUserSettings"

// AdminSetUserSettingsRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// This action is no longer supported. You can use it to configure only SMS
// MFA. You can't use it to configure TOTP software token MFA. To configure
// either type of MFA, use the AdminSetUserMFAPreference action instead.
//
//    // Example sending a request using AdminSetUserSettingsRequest.
//    req := client.AdminSetUserSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminSetUserSettings
func (c *Client) AdminSetUserSettingsRequest(input *types.AdminSetUserSettingsInput) AdminSetUserSettingsRequest {
	op := &aws.Operation{
		Name:       opAdminSetUserSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AdminSetUserSettingsInput{}
	}

	req := c.newRequest(op, input, &types.AdminSetUserSettingsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AdminSetUserSettingsMarshaler{Input: input}.GetNamedBuildHandler())

	return AdminSetUserSettingsRequest{Request: req, Input: input, Copy: c.AdminSetUserSettingsRequest}
}

// AdminSetUserSettingsRequest is the request type for the
// AdminSetUserSettings API operation.
type AdminSetUserSettingsRequest struct {
	*aws.Request
	Input *types.AdminSetUserSettingsInput
	Copy  func(*types.AdminSetUserSettingsInput) AdminSetUserSettingsRequest
}

// Send marshals and sends the AdminSetUserSettings API request.
func (r AdminSetUserSettingsRequest) Send(ctx context.Context) (*AdminSetUserSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminSetUserSettingsResponse{
		AdminSetUserSettingsOutput: r.Request.Data.(*types.AdminSetUserSettingsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminSetUserSettingsResponse is the response type for the
// AdminSetUserSettings API operation.
type AdminSetUserSettingsResponse struct {
	*types.AdminSetUserSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminSetUserSettings request.
func (r *AdminSetUserSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
