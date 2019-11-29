// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const opDeletePlatformApplication = "DeletePlatformApplication"

// DeletePlatformApplicationRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Deletes a platform application object for one of the supported push notification
// services, such as APNS and GCM. For more information, see Using Amazon SNS
// Mobile Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
//
//    // Example sending a request using DeletePlatformApplicationRequest.
//    req := client.DeletePlatformApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/DeletePlatformApplication
func (c *Client) DeletePlatformApplicationRequest(input *types.DeletePlatformApplicationInput) DeletePlatformApplicationRequest {
	op := &aws.Operation{
		Name:       opDeletePlatformApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeletePlatformApplicationInput{}
	}

	req := c.newRequest(op, input, &types.DeletePlatformApplicationOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePlatformApplicationRequest{Request: req, Input: input, Copy: c.DeletePlatformApplicationRequest}
}

// DeletePlatformApplicationRequest is the request type for the
// DeletePlatformApplication API operation.
type DeletePlatformApplicationRequest struct {
	*aws.Request
	Input *types.DeletePlatformApplicationInput
	Copy  func(*types.DeletePlatformApplicationInput) DeletePlatformApplicationRequest
}

// Send marshals and sends the DeletePlatformApplication API request.
func (r DeletePlatformApplicationRequest) Send(ctx context.Context) (*DeletePlatformApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePlatformApplicationResponse{
		DeletePlatformApplicationOutput: r.Request.Data.(*types.DeletePlatformApplicationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePlatformApplicationResponse is the response type for the
// DeletePlatformApplication API operation.
type DeletePlatformApplicationResponse struct {
	*types.DeletePlatformApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePlatformApplication request.
func (r *DeletePlatformApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
