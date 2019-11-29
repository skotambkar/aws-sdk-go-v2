// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateVpcEndpointConnectionNotification = "CreateVpcEndpointConnectionNotification"

// CreateVpcEndpointConnectionNotificationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a connection notification for a specified VPC endpoint or VPC endpoint
// service. A connection notification notifies you of specific endpoint events.
// You must create an SNS topic to receive notifications. For more information,
// see Create a Topic (https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html)
// in the Amazon Simple Notification Service Developer Guide.
//
// You can create a connection notification for interface endpoints only.
//
//    // Example sending a request using CreateVpcEndpointConnectionNotificationRequest.
//    req := client.CreateVpcEndpointConnectionNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpcEndpointConnectionNotification
func (c *Client) CreateVpcEndpointConnectionNotificationRequest(input *types.CreateVpcEndpointConnectionNotificationInput) CreateVpcEndpointConnectionNotificationRequest {
	op := &aws.Operation{
		Name:       opCreateVpcEndpointConnectionNotification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateVpcEndpointConnectionNotificationInput{}
	}

	req := c.newRequest(op, input, &types.CreateVpcEndpointConnectionNotificationOutput{})
	return CreateVpcEndpointConnectionNotificationRequest{Request: req, Input: input, Copy: c.CreateVpcEndpointConnectionNotificationRequest}
}

// CreateVpcEndpointConnectionNotificationRequest is the request type for the
// CreateVpcEndpointConnectionNotification API operation.
type CreateVpcEndpointConnectionNotificationRequest struct {
	*aws.Request
	Input *types.CreateVpcEndpointConnectionNotificationInput
	Copy  func(*types.CreateVpcEndpointConnectionNotificationInput) CreateVpcEndpointConnectionNotificationRequest
}

// Send marshals and sends the CreateVpcEndpointConnectionNotification API request.
func (r CreateVpcEndpointConnectionNotificationRequest) Send(ctx context.Context) (*CreateVpcEndpointConnectionNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpcEndpointConnectionNotificationResponse{
		CreateVpcEndpointConnectionNotificationOutput: r.Request.Data.(*types.CreateVpcEndpointConnectionNotificationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpcEndpointConnectionNotificationResponse is the response type for the
// CreateVpcEndpointConnectionNotification API operation.
type CreateVpcEndpointConnectionNotificationResponse struct {
	*types.CreateVpcEndpointConnectionNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpcEndpointConnectionNotification request.
func (r *CreateVpcEndpointConnectionNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
