// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opGetEmailIdentity = "GetEmailIdentity"

// GetEmailIdentityRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Provides information about a specific identity associated with your Amazon
// Pinpoint account, including the identity's verification status, its DKIM
// authentication status, and its custom Mail-From settings.
//
//    // Example sending a request using GetEmailIdentityRequest.
//    req := client.GetEmailIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetEmailIdentity
func (c *Client) GetEmailIdentityRequest(input *types.GetEmailIdentityInput) GetEmailIdentityRequest {
	op := &aws.Operation{
		Name:       opGetEmailIdentity,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/identities/{EmailIdentity}",
	}

	if input == nil {
		input = &types.GetEmailIdentityInput{}
	}

	req := c.newRequest(op, input, &types.GetEmailIdentityOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetEmailIdentityMarshaler{Input: input}.GetNamedBuildHandler())

	return GetEmailIdentityRequest{Request: req, Input: input, Copy: c.GetEmailIdentityRequest}
}

// GetEmailIdentityRequest is the request type for the
// GetEmailIdentity API operation.
type GetEmailIdentityRequest struct {
	*aws.Request
	Input *types.GetEmailIdentityInput
	Copy  func(*types.GetEmailIdentityInput) GetEmailIdentityRequest
}

// Send marshals and sends the GetEmailIdentity API request.
func (r GetEmailIdentityRequest) Send(ctx context.Context) (*GetEmailIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEmailIdentityResponse{
		GetEmailIdentityOutput: r.Request.Data.(*types.GetEmailIdentityOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEmailIdentityResponse is the response type for the
// GetEmailIdentity API operation.
type GetEmailIdentityResponse struct {
	*types.GetEmailIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEmailIdentity request.
func (r *GetEmailIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
