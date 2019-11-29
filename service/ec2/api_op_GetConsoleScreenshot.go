// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opGetConsoleScreenshot = "GetConsoleScreenshot"

// GetConsoleScreenshotRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Retrieve a JPG-format screenshot of a running instance to help with troubleshooting.
//
// The returned content is Base64-encoded.
//
//    // Example sending a request using GetConsoleScreenshotRequest.
//    req := client.GetConsoleScreenshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetConsoleScreenshot
func (c *Client) GetConsoleScreenshotRequest(input *types.GetConsoleScreenshotInput) GetConsoleScreenshotRequest {
	op := &aws.Operation{
		Name:       opGetConsoleScreenshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetConsoleScreenshotInput{}
	}

	req := c.newRequest(op, input, &types.GetConsoleScreenshotOutput{})
	return GetConsoleScreenshotRequest{Request: req, Input: input, Copy: c.GetConsoleScreenshotRequest}
}

// GetConsoleScreenshotRequest is the request type for the
// GetConsoleScreenshot API operation.
type GetConsoleScreenshotRequest struct {
	*aws.Request
	Input *types.GetConsoleScreenshotInput
	Copy  func(*types.GetConsoleScreenshotInput) GetConsoleScreenshotRequest
}

// Send marshals and sends the GetConsoleScreenshot API request.
func (r GetConsoleScreenshotRequest) Send(ctx context.Context) (*GetConsoleScreenshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConsoleScreenshotResponse{
		GetConsoleScreenshotOutput: r.Request.Data.(*types.GetConsoleScreenshotOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConsoleScreenshotResponse is the response type for the
// GetConsoleScreenshot API operation.
type GetConsoleScreenshotResponse struct {
	*types.GetConsoleScreenshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConsoleScreenshot request.
func (r *GetConsoleScreenshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
