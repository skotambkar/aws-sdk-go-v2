// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
)

const opSetIdentityPoolConfiguration = "SetIdentityPoolConfiguration"

// SetIdentityPoolConfigurationRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Sets the necessary configuration for push sync.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
//
//    // Example sending a request using SetIdentityPoolConfigurationRequest.
//    req := client.SetIdentityPoolConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/SetIdentityPoolConfiguration
func (c *Client) SetIdentityPoolConfigurationRequest(input *types.SetIdentityPoolConfigurationInput) SetIdentityPoolConfigurationRequest {
	op := &aws.Operation{
		Name:       opSetIdentityPoolConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/identitypools/{IdentityPoolId}/configuration",
	}

	if input == nil {
		input = &types.SetIdentityPoolConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.SetIdentityPoolConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.SetIdentityPoolConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return SetIdentityPoolConfigurationRequest{Request: req, Input: input, Copy: c.SetIdentityPoolConfigurationRequest}
}

// SetIdentityPoolConfigurationRequest is the request type for the
// SetIdentityPoolConfiguration API operation.
type SetIdentityPoolConfigurationRequest struct {
	*aws.Request
	Input *types.SetIdentityPoolConfigurationInput
	Copy  func(*types.SetIdentityPoolConfigurationInput) SetIdentityPoolConfigurationRequest
}

// Send marshals and sends the SetIdentityPoolConfiguration API request.
func (r SetIdentityPoolConfigurationRequest) Send(ctx context.Context) (*SetIdentityPoolConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetIdentityPoolConfigurationResponse{
		SetIdentityPoolConfigurationOutput: r.Request.Data.(*types.SetIdentityPoolConfigurationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetIdentityPoolConfigurationResponse is the response type for the
// SetIdentityPoolConfiguration API operation.
type SetIdentityPoolConfigurationResponse struct {
	*types.SetIdentityPoolConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetIdentityPoolConfiguration request.
func (r *SetIdentityPoolConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
