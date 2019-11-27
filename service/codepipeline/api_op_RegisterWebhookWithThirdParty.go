// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
)

const opRegisterWebhookWithThirdParty = "RegisterWebhookWithThirdParty"

// RegisterWebhookWithThirdPartyRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Configures a connection between the webhook that was created and the external
// tool with events to be detected.
//
//    // Example sending a request using RegisterWebhookWithThirdPartyRequest.
//    req := client.RegisterWebhookWithThirdPartyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/RegisterWebhookWithThirdParty
func (c *Client) RegisterWebhookWithThirdPartyRequest(input *types.RegisterWebhookWithThirdPartyInput) RegisterWebhookWithThirdPartyRequest {
	op := &aws.Operation{
		Name:       opRegisterWebhookWithThirdParty,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterWebhookWithThirdPartyInput{}
	}

	req := c.newRequest(op, input, &types.RegisterWebhookWithThirdPartyOutput{})
	return RegisterWebhookWithThirdPartyRequest{Request: req, Input: input, Copy: c.RegisterWebhookWithThirdPartyRequest}
}

// RegisterWebhookWithThirdPartyRequest is the request type for the
// RegisterWebhookWithThirdParty API operation.
type RegisterWebhookWithThirdPartyRequest struct {
	*aws.Request
	Input *types.RegisterWebhookWithThirdPartyInput
	Copy  func(*types.RegisterWebhookWithThirdPartyInput) RegisterWebhookWithThirdPartyRequest
}

// Send marshals and sends the RegisterWebhookWithThirdParty API request.
func (r RegisterWebhookWithThirdPartyRequest) Send(ctx context.Context) (*RegisterWebhookWithThirdPartyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterWebhookWithThirdPartyResponse{
		RegisterWebhookWithThirdPartyOutput: r.Request.Data.(*types.RegisterWebhookWithThirdPartyOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterWebhookWithThirdPartyResponse is the response type for the
// RegisterWebhookWithThirdParty API operation.
type RegisterWebhookWithThirdPartyResponse struct {
	*types.RegisterWebhookWithThirdPartyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterWebhookWithThirdParty request.
func (r *RegisterWebhookWithThirdPartyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
