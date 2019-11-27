// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opDisassociatePhoneNumbersFromVoiceConnector = "DisassociatePhoneNumbersFromVoiceConnector"

// DisassociatePhoneNumbersFromVoiceConnectorRequest returns a request value for making API operation for
// Amazon Chime.
//
// Disassociates the specified phone number from the specified Amazon Chime
// Voice Connector.
//
//    // Example sending a request using DisassociatePhoneNumbersFromVoiceConnectorRequest.
//    req := client.DisassociatePhoneNumbersFromVoiceConnectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DisassociatePhoneNumbersFromVoiceConnector
func (c *Client) DisassociatePhoneNumbersFromVoiceConnectorRequest(input *types.DisassociatePhoneNumbersFromVoiceConnectorInput) DisassociatePhoneNumbersFromVoiceConnectorRequest {
	op := &aws.Operation{
		Name:       opDisassociatePhoneNumbersFromVoiceConnector,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}?operation=disassociate-phone-numbers",
	}

	if input == nil {
		input = &types.DisassociatePhoneNumbersFromVoiceConnectorInput{}
	}

	req := c.newRequest(op, input, &types.DisassociatePhoneNumbersFromVoiceConnectorOutput{})
	return DisassociatePhoneNumbersFromVoiceConnectorRequest{Request: req, Input: input, Copy: c.DisassociatePhoneNumbersFromVoiceConnectorRequest}
}

// DisassociatePhoneNumbersFromVoiceConnectorRequest is the request type for the
// DisassociatePhoneNumbersFromVoiceConnector API operation.
type DisassociatePhoneNumbersFromVoiceConnectorRequest struct {
	*aws.Request
	Input *types.DisassociatePhoneNumbersFromVoiceConnectorInput
	Copy  func(*types.DisassociatePhoneNumbersFromVoiceConnectorInput) DisassociatePhoneNumbersFromVoiceConnectorRequest
}

// Send marshals and sends the DisassociatePhoneNumbersFromVoiceConnector API request.
func (r DisassociatePhoneNumbersFromVoiceConnectorRequest) Send(ctx context.Context) (*DisassociatePhoneNumbersFromVoiceConnectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociatePhoneNumbersFromVoiceConnectorResponse{
		DisassociatePhoneNumbersFromVoiceConnectorOutput: r.Request.Data.(*types.DisassociatePhoneNumbersFromVoiceConnectorOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociatePhoneNumbersFromVoiceConnectorResponse is the response type for the
// DisassociatePhoneNumbersFromVoiceConnector API operation.
type DisassociatePhoneNumbersFromVoiceConnectorResponse struct {
	*types.DisassociatePhoneNumbersFromVoiceConnectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociatePhoneNumbersFromVoiceConnector request.
func (r *DisassociatePhoneNumbersFromVoiceConnectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
