// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

const opStartLabelDetection = "StartLabelDetection"

// StartLabelDetectionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Starts asynchronous detection of labels in a stored video.
//
// Amazon Rekognition Video can detect labels in a video. Labels are instances
// of real-world entities. This includes objects like flower, tree, and table;
// events like wedding, graduation, and birthday party; concepts like landscape,
// evening, and nature; and activities like a person getting out of a car or
// a person skiing.
//
// The video must be stored in an Amazon S3 bucket. Use Video to specify the
// bucket name and the filename of the video. StartLabelDetection returns a
// job identifier (JobId) which you use to get the results of the operation.
// When label detection is finished, Amazon Rekognition Video publishes a completion
// status to the Amazon Simple Notification Service topic that you specify in
// NotificationChannel.
//
// To get the results of the label detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetLabelDetection and pass the job identifier (JobId) from the initial call
// to StartLabelDetection.
//
//    // Example sending a request using StartLabelDetectionRequest.
//    req := client.StartLabelDetectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartLabelDetectionRequest(input *types.StartLabelDetectionInput) StartLabelDetectionRequest {
	op := &aws.Operation{
		Name:       opStartLabelDetection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartLabelDetectionInput{}
	}

	req := c.newRequest(op, input, &types.StartLabelDetectionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.StartLabelDetectionMarshaler{Input: input}.GetNamedBuildHandler())

	return StartLabelDetectionRequest{Request: req, Input: input, Copy: c.StartLabelDetectionRequest}
}

// StartLabelDetectionRequest is the request type for the
// StartLabelDetection API operation.
type StartLabelDetectionRequest struct {
	*aws.Request
	Input *types.StartLabelDetectionInput
	Copy  func(*types.StartLabelDetectionInput) StartLabelDetectionRequest
}

// Send marshals and sends the StartLabelDetection API request.
func (r StartLabelDetectionRequest) Send(ctx context.Context) (*StartLabelDetectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartLabelDetectionResponse{
		StartLabelDetectionOutput: r.Request.Data.(*types.StartLabelDetectionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartLabelDetectionResponse is the response type for the
// StartLabelDetection API operation.
type StartLabelDetectionResponse struct {
	*types.StartLabelDetectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartLabelDetection request.
func (r *StartLabelDetectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
