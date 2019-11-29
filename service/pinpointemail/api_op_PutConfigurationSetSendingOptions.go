// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opPutConfigurationSetSendingOptions = "PutConfigurationSetSendingOptions"

// PutConfigurationSetSendingOptionsRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Enable or disable email sending for messages that use a particular configuration
// set in a specific AWS Region.
//
//    // Example sending a request using PutConfigurationSetSendingOptionsRequest.
//    req := client.PutConfigurationSetSendingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutConfigurationSetSendingOptions
func (c *Client) PutConfigurationSetSendingOptionsRequest(input *types.PutConfigurationSetSendingOptionsInput) PutConfigurationSetSendingOptionsRequest {
	op := &aws.Operation{
		Name:       opPutConfigurationSetSendingOptions,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/configuration-sets/{ConfigurationSetName}/sending",
	}

	if input == nil {
		input = &types.PutConfigurationSetSendingOptionsInput{}
	}

	req := c.newRequest(op, input, &types.PutConfigurationSetSendingOptionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.PutConfigurationSetSendingOptionsMarshaler{Input: input}.GetNamedBuildHandler())

	return PutConfigurationSetSendingOptionsRequest{Request: req, Input: input, Copy: c.PutConfigurationSetSendingOptionsRequest}
}

// PutConfigurationSetSendingOptionsRequest is the request type for the
// PutConfigurationSetSendingOptions API operation.
type PutConfigurationSetSendingOptionsRequest struct {
	*aws.Request
	Input *types.PutConfigurationSetSendingOptionsInput
	Copy  func(*types.PutConfigurationSetSendingOptionsInput) PutConfigurationSetSendingOptionsRequest
}

// Send marshals and sends the PutConfigurationSetSendingOptions API request.
func (r PutConfigurationSetSendingOptionsRequest) Send(ctx context.Context) (*PutConfigurationSetSendingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConfigurationSetSendingOptionsResponse{
		PutConfigurationSetSendingOptionsOutput: r.Request.Data.(*types.PutConfigurationSetSendingOptionsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConfigurationSetSendingOptionsResponse is the response type for the
// PutConfigurationSetSendingOptions API operation.
type PutConfigurationSetSendingOptionsResponse struct {
	*types.PutConfigurationSetSendingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConfigurationSetSendingOptions request.
func (r *PutConfigurationSetSendingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
