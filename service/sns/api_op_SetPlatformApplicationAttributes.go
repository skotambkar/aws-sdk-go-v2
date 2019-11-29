// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sns/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const opSetPlatformApplicationAttributes = "SetPlatformApplicationAttributes"

// SetPlatformApplicationAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Sets the attributes of the platform application object for the supported
// push notification services, such as APNS and GCM. For more information, see
// Using Amazon SNS Mobile Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
// For information on configuring attributes for message delivery status, see
// Using Amazon SNS Application Attributes for Message Delivery Status (https://docs.aws.amazon.com/sns/latest/dg/sns-msg-status.html).
//
//    // Example sending a request using SetPlatformApplicationAttributesRequest.
//    req := client.SetPlatformApplicationAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/SetPlatformApplicationAttributes
func (c *Client) SetPlatformApplicationAttributesRequest(input *types.SetPlatformApplicationAttributesInput) SetPlatformApplicationAttributesRequest {
	op := &aws.Operation{
		Name:       opSetPlatformApplicationAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetPlatformApplicationAttributesInput{}
	}

	req := c.newRequest(op, input, &types.SetPlatformApplicationAttributesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.SetPlatformApplicationAttributesMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetPlatformApplicationAttributesRequest{Request: req, Input: input, Copy: c.SetPlatformApplicationAttributesRequest}
}

// SetPlatformApplicationAttributesRequest is the request type for the
// SetPlatformApplicationAttributes API operation.
type SetPlatformApplicationAttributesRequest struct {
	*aws.Request
	Input *types.SetPlatformApplicationAttributesInput
	Copy  func(*types.SetPlatformApplicationAttributesInput) SetPlatformApplicationAttributesRequest
}

// Send marshals and sends the SetPlatformApplicationAttributes API request.
func (r SetPlatformApplicationAttributesRequest) Send(ctx context.Context) (*SetPlatformApplicationAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetPlatformApplicationAttributesResponse{
		SetPlatformApplicationAttributesOutput: r.Request.Data.(*types.SetPlatformApplicationAttributesOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetPlatformApplicationAttributesResponse is the response type for the
// SetPlatformApplicationAttributes API operation.
type SetPlatformApplicationAttributesResponse struct {
	*types.SetPlatformApplicationAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetPlatformApplicationAttributes request.
func (r *SetPlatformApplicationAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
