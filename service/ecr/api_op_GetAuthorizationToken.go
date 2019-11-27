// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

const opGetAuthorizationToken = "GetAuthorizationToken"

// GetAuthorizationTokenRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Retrieves a token that is valid for a specified registry for 12 hours. This
// command allows you to use the docker CLI to push and pull images with Amazon
// ECR. If you do not specify a registry, the default registry is assumed.
//
// The authorizationToken returned for each registry specified is a base64 encoded
// string that can be decoded and used in a docker login command to authenticate
// to a registry. The AWS CLI offers an aws ecr get-login command that simplifies
// the login process.
//
//    // Example sending a request using GetAuthorizationTokenRequest.
//    req := client.GetAuthorizationTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/GetAuthorizationToken
func (c *Client) GetAuthorizationTokenRequest(input *types.GetAuthorizationTokenInput) GetAuthorizationTokenRequest {
	op := &aws.Operation{
		Name:       opGetAuthorizationToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetAuthorizationTokenInput{}
	}

	req := c.newRequest(op, input, &types.GetAuthorizationTokenOutput{})
	return GetAuthorizationTokenRequest{Request: req, Input: input, Copy: c.GetAuthorizationTokenRequest}
}

// GetAuthorizationTokenRequest is the request type for the
// GetAuthorizationToken API operation.
type GetAuthorizationTokenRequest struct {
	*aws.Request
	Input *types.GetAuthorizationTokenInput
	Copy  func(*types.GetAuthorizationTokenInput) GetAuthorizationTokenRequest
}

// Send marshals and sends the GetAuthorizationToken API request.
func (r GetAuthorizationTokenRequest) Send(ctx context.Context) (*GetAuthorizationTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAuthorizationTokenResponse{
		GetAuthorizationTokenOutput: r.Request.Data.(*types.GetAuthorizationTokenOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAuthorizationTokenResponse is the response type for the
// GetAuthorizationToken API operation.
type GetAuthorizationTokenResponse struct {
	*types.GetAuthorizationTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAuthorizationToken request.
func (r *GetAuthorizationTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
