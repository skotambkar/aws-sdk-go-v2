// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opPutEmailIdentityFeedbackAttributes = "PutEmailIdentityFeedbackAttributes"

// PutEmailIdentityFeedbackAttributesRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Used to enable or disable feedback forwarding for an identity. This setting
// determines what happens when an identity is used to send an email that results
// in a bounce or complaint event.
//
// When you enable feedback forwarding, Amazon Pinpoint sends you email notifications
// when bounce or complaint events occur. Amazon Pinpoint sends this notification
// to the address that you specified in the Return-Path header of the original
// email.
//
// When you disable feedback forwarding, Amazon Pinpoint sends notifications
// through other mechanisms, such as by notifying an Amazon SNS topic. You're
// required to have a method of tracking bounces and complaints. If you haven't
// set up another mechanism for receiving bounce or complaint notifications,
// Amazon Pinpoint sends an email notification when these events occur (even
// if this setting is disabled).
//
//    // Example sending a request using PutEmailIdentityFeedbackAttributesRequest.
//    req := client.PutEmailIdentityFeedbackAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutEmailIdentityFeedbackAttributes
func (c *Client) PutEmailIdentityFeedbackAttributesRequest(input *types.PutEmailIdentityFeedbackAttributesInput) PutEmailIdentityFeedbackAttributesRequest {
	op := &aws.Operation{
		Name:       opPutEmailIdentityFeedbackAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/identities/{EmailIdentity}/feedback",
	}

	if input == nil {
		input = &types.PutEmailIdentityFeedbackAttributesInput{}
	}

	req := c.newRequest(op, input, &types.PutEmailIdentityFeedbackAttributesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.PutEmailIdentityFeedbackAttributesMarshaler{Input: input}.GetNamedBuildHandler())

	return PutEmailIdentityFeedbackAttributesRequest{Request: req, Input: input, Copy: c.PutEmailIdentityFeedbackAttributesRequest}
}

// PutEmailIdentityFeedbackAttributesRequest is the request type for the
// PutEmailIdentityFeedbackAttributes API operation.
type PutEmailIdentityFeedbackAttributesRequest struct {
	*aws.Request
	Input *types.PutEmailIdentityFeedbackAttributesInput
	Copy  func(*types.PutEmailIdentityFeedbackAttributesInput) PutEmailIdentityFeedbackAttributesRequest
}

// Send marshals and sends the PutEmailIdentityFeedbackAttributes API request.
func (r PutEmailIdentityFeedbackAttributesRequest) Send(ctx context.Context) (*PutEmailIdentityFeedbackAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEmailIdentityFeedbackAttributesResponse{
		PutEmailIdentityFeedbackAttributesOutput: r.Request.Data.(*types.PutEmailIdentityFeedbackAttributesOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEmailIdentityFeedbackAttributesResponse is the response type for the
// PutEmailIdentityFeedbackAttributes API operation.
type PutEmailIdentityFeedbackAttributesResponse struct {
	*types.PutEmailIdentityFeedbackAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEmailIdentityFeedbackAttributes request.
func (r *PutEmailIdentityFeedbackAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
