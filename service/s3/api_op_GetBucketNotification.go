// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opGetBucketNotification = "GetBucketNotification"

// GetBucketNotificationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// No longer used, see the GetBucketNotificationConfiguration operation.
//
//    // Example sending a request using GetBucketNotificationRequest.
//    req := client.GetBucketNotificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketNotification
func (c *Client) GetBucketNotificationRequest(input *types.GetBucketNotificationInput) GetBucketNotificationRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, GetBucketNotification, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opGetBucketNotification,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?notification",
	}

	if input == nil {
		input = &types.GetBucketNotificationInput{}
	}

	req := c.newRequest(op, input, &types.GetBucketNotificationOutput{})
	return GetBucketNotificationRequest{Request: req, Input: input, Copy: c.GetBucketNotificationRequest}
}

// GetBucketNotificationRequest is the request type for the
// GetBucketNotification API operation.
type GetBucketNotificationRequest struct {
	*aws.Request
	Input *types.GetBucketNotificationInput
	Copy  func(*types.GetBucketNotificationInput) GetBucketNotificationRequest
}

// Send marshals and sends the GetBucketNotification API request.
func (r GetBucketNotificationRequest) Send(ctx context.Context) (*GetBucketNotificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketNotificationResponse{
		GetBucketNotificationOutput: r.Request.Data.(*types.GetBucketNotificationOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketNotificationResponse is the response type for the
// GetBucketNotification API operation.
type GetBucketNotificationResponse struct {
	*types.GetBucketNotificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketNotification request.
func (r *GetBucketNotificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
