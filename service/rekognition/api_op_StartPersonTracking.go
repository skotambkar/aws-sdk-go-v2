// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

const opStartPersonTracking = "StartPersonTracking"

// StartPersonTrackingRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Starts the asynchronous tracking of a person's path in a stored video.
//
// Amazon Rekognition Video can track the path of people in a video stored in
// an Amazon S3 bucket. Use Video to specify the bucket name and the filename
// of the video. StartPersonTracking returns a job identifier (JobId) which
// you use to get the results of the operation. When label detection is finished,
// Amazon Rekognition publishes a completion status to the Amazon Simple Notification
// Service topic that you specify in NotificationChannel.
//
// To get the results of the person detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetPersonTracking and pass the job identifier (JobId) from the initial call
// to StartPersonTracking.
//
//    // Example sending a request using StartPersonTrackingRequest.
//    req := client.StartPersonTrackingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartPersonTrackingRequest(input *types.StartPersonTrackingInput) StartPersonTrackingRequest {
	op := &aws.Operation{
		Name:       opStartPersonTracking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartPersonTrackingInput{}
	}

	req := c.newRequest(op, input, &types.StartPersonTrackingOutput{})
	return StartPersonTrackingRequest{Request: req, Input: input, Copy: c.StartPersonTrackingRequest}
}

// StartPersonTrackingRequest is the request type for the
// StartPersonTracking API operation.
type StartPersonTrackingRequest struct {
	*aws.Request
	Input *types.StartPersonTrackingInput
	Copy  func(*types.StartPersonTrackingInput) StartPersonTrackingRequest
}

// Send marshals and sends the StartPersonTracking API request.
func (r StartPersonTrackingRequest) Send(ctx context.Context) (*StartPersonTrackingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartPersonTrackingResponse{
		StartPersonTrackingOutput: r.Request.Data.(*types.StartPersonTrackingOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartPersonTrackingResponse is the response type for the
// StartPersonTracking API operation.
type StartPersonTrackingResponse struct {
	*types.StartPersonTrackingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartPersonTracking request.
func (r *StartPersonTrackingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
