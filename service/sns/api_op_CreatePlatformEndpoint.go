// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const opCreatePlatformEndpoint = "CreatePlatformEndpoint"

// CreatePlatformEndpointRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Creates an endpoint for a device and mobile app on one of the supported push
// notification services, such as GCM and APNS. CreatePlatformEndpoint requires
// the PlatformApplicationArn that is returned from CreatePlatformApplication.
// The EndpointArn that is returned when using CreatePlatformEndpoint can then
// be used by the Publish action to send a message to a mobile app or by the
// Subscribe action for subscription to a topic. The CreatePlatformEndpoint
// action is idempotent, so if the requester already owns an endpoint with the
// same device token and attributes, that endpoint's ARN is returned without
// creating a new endpoint. For more information, see Using Amazon SNS Mobile
// Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
//
// When using CreatePlatformEndpoint with Baidu, two attributes must be provided:
// ChannelId and UserId. The token field must also contain the ChannelId. For
// more information, see Creating an Amazon SNS Endpoint for Baidu (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePushBaiduEndpoint.html).
//
//    // Example sending a request using CreatePlatformEndpointRequest.
//    req := client.CreatePlatformEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/CreatePlatformEndpoint
func (c *Client) CreatePlatformEndpointRequest(input *types.CreatePlatformEndpointInput) CreatePlatformEndpointRequest {
	op := &aws.Operation{
		Name:       opCreatePlatformEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreatePlatformEndpointInput{}
	}

	req := c.newRequest(op, input, &types.CreatePlatformEndpointOutput{})
	return CreatePlatformEndpointRequest{Request: req, Input: input, Copy: c.CreatePlatformEndpointRequest}
}

// CreatePlatformEndpointRequest is the request type for the
// CreatePlatformEndpoint API operation.
type CreatePlatformEndpointRequest struct {
	*aws.Request
	Input *types.CreatePlatformEndpointInput
	Copy  func(*types.CreatePlatformEndpointInput) CreatePlatformEndpointRequest
}

// Send marshals and sends the CreatePlatformEndpoint API request.
func (r CreatePlatformEndpointRequest) Send(ctx context.Context) (*CreatePlatformEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePlatformEndpointResponse{
		CreatePlatformEndpointOutput: r.Request.Data.(*types.CreatePlatformEndpointOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePlatformEndpointResponse is the response type for the
// CreatePlatformEndpoint API operation.
type CreatePlatformEndpointResponse struct {
	*types.CreatePlatformEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePlatformEndpoint request.
func (r *CreatePlatformEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
