// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice/types"
)

const opCreateConfigurationSetEventDestination = "CreateConfigurationSetEventDestination"

// CreateConfigurationSetEventDestinationRequest returns a request value for making API operation for
// Amazon Pinpoint SMS and Voice Service.
//
// Create a new event destination in a configuration set.
//
//    // Example sending a request using CreateConfigurationSetEventDestinationRequest.
//    req := client.CreateConfigurationSetEventDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/CreateConfigurationSetEventDestination
func (c *Client) CreateConfigurationSetEventDestinationRequest(input *types.CreateConfigurationSetEventDestinationInput) CreateConfigurationSetEventDestinationRequest {
	op := &aws.Operation{
		Name:       opCreateConfigurationSetEventDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/sms-voice/configuration-sets/{ConfigurationSetName}/event-destinations",
	}

	if input == nil {
		input = &types.CreateConfigurationSetEventDestinationInput{}
	}

	req := c.newRequest(op, input, &types.CreateConfigurationSetEventDestinationOutput{})
	return CreateConfigurationSetEventDestinationRequest{Request: req, Input: input, Copy: c.CreateConfigurationSetEventDestinationRequest}
}

// CreateConfigurationSetEventDestinationRequest is the request type for the
// CreateConfigurationSetEventDestination API operation.
type CreateConfigurationSetEventDestinationRequest struct {
	*aws.Request
	Input *types.CreateConfigurationSetEventDestinationInput
	Copy  func(*types.CreateConfigurationSetEventDestinationInput) CreateConfigurationSetEventDestinationRequest
}

// Send marshals and sends the CreateConfigurationSetEventDestination API request.
func (r CreateConfigurationSetEventDestinationRequest) Send(ctx context.Context) (*CreateConfigurationSetEventDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConfigurationSetEventDestinationResponse{
		CreateConfigurationSetEventDestinationOutput: r.Request.Data.(*types.CreateConfigurationSetEventDestinationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConfigurationSetEventDestinationResponse is the response type for the
// CreateConfigurationSetEventDestination API operation.
type CreateConfigurationSetEventDestinationResponse struct {
	*types.CreateConfigurationSetEventDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConfigurationSetEventDestination request.
func (r *CreateConfigurationSetEventDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
