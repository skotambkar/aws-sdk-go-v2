// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opAssociatePhoneNumbersWithVoiceConnector = "AssociatePhoneNumbersWithVoiceConnector"

// AssociatePhoneNumbersWithVoiceConnectorRequest returns a request value for making API operation for
// Amazon Chime.
//
// Associates phone numbers with the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using AssociatePhoneNumbersWithVoiceConnectorRequest.
//    req := client.AssociatePhoneNumbersWithVoiceConnectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/AssociatePhoneNumbersWithVoiceConnector
func (c *Client) AssociatePhoneNumbersWithVoiceConnectorRequest(input *types.AssociatePhoneNumbersWithVoiceConnectorInput) AssociatePhoneNumbersWithVoiceConnectorRequest {
	op := &aws.Operation{
		Name:       opAssociatePhoneNumbersWithVoiceConnector,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}?operation=associate-phone-numbers",
	}

	if input == nil {
		input = &types.AssociatePhoneNumbersWithVoiceConnectorInput{}
	}

	req := c.newRequest(op, input, &types.AssociatePhoneNumbersWithVoiceConnectorOutput{})
	return AssociatePhoneNumbersWithVoiceConnectorRequest{Request: req, Input: input, Copy: c.AssociatePhoneNumbersWithVoiceConnectorRequest}
}

// AssociatePhoneNumbersWithVoiceConnectorRequest is the request type for the
// AssociatePhoneNumbersWithVoiceConnector API operation.
type AssociatePhoneNumbersWithVoiceConnectorRequest struct {
	*aws.Request
	Input *types.AssociatePhoneNumbersWithVoiceConnectorInput
	Copy  func(*types.AssociatePhoneNumbersWithVoiceConnectorInput) AssociatePhoneNumbersWithVoiceConnectorRequest
}

// Send marshals and sends the AssociatePhoneNumbersWithVoiceConnector API request.
func (r AssociatePhoneNumbersWithVoiceConnectorRequest) Send(ctx context.Context) (*AssociatePhoneNumbersWithVoiceConnectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociatePhoneNumbersWithVoiceConnectorResponse{
		AssociatePhoneNumbersWithVoiceConnectorOutput: r.Request.Data.(*types.AssociatePhoneNumbersWithVoiceConnectorOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociatePhoneNumbersWithVoiceConnectorResponse is the response type for the
// AssociatePhoneNumbersWithVoiceConnector API operation.
type AssociatePhoneNumbersWithVoiceConnectorResponse struct {
	*types.AssociatePhoneNumbersWithVoiceConnectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociatePhoneNumbersWithVoiceConnector request.
func (r *AssociatePhoneNumbersWithVoiceConnectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
