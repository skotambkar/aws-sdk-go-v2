// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opCreateConfigurationSetEventDestination = "CreateConfigurationSetEventDestination"

// CreateConfigurationSetEventDestinationRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Create an event destination. In Amazon Pinpoint, events include message sends,
// deliveries, opens, clicks, bounces, and complaints. Event destinations are
// places that you can send information about these events to. For example,
// you can send event data to Amazon SNS to receive notifications when you receive
// bounces or complaints, or you can use Amazon Kinesis Data Firehose to stream
// data to Amazon S3 for long-term storage.
//
// A single configuration set can include more than one event destination.
//
//    // Example sending a request using CreateConfigurationSetEventDestinationRequest.
//    req := client.CreateConfigurationSetEventDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/CreateConfigurationSetEventDestination
func (c *Client) CreateConfigurationSetEventDestinationRequest(input *types.CreateConfigurationSetEventDestinationInput) CreateConfigurationSetEventDestinationRequest {
	op := &aws.Operation{
		Name:       opCreateConfigurationSetEventDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/email/configuration-sets/{ConfigurationSetName}/event-destinations",
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
