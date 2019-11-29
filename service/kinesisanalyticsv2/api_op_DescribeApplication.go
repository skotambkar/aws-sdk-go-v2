// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
)

const opDescribeApplication = "DescribeApplication"

// DescribeApplicationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Returns information about a specific Amazon Kinesis Data Analytics application.
//
// If you want to retrieve a list of all applications in your account, use the
// ListApplications operation.
//
//    // Example sending a request using DescribeApplicationRequest.
//    req := client.DescribeApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/DescribeApplication
func (c *Client) DescribeApplicationRequest(input *types.DescribeApplicationInput) DescribeApplicationRequest {
	op := &aws.Operation{
		Name:       opDescribeApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeApplicationInput{}
	}

	req := c.newRequest(op, input, &types.DescribeApplicationOutput{})
	return DescribeApplicationRequest{Request: req, Input: input, Copy: c.DescribeApplicationRequest}
}

// DescribeApplicationRequest is the request type for the
// DescribeApplication API operation.
type DescribeApplicationRequest struct {
	*aws.Request
	Input *types.DescribeApplicationInput
	Copy  func(*types.DescribeApplicationInput) DescribeApplicationRequest
}

// Send marshals and sends the DescribeApplication API request.
func (r DescribeApplicationRequest) Send(ctx context.Context) (*DescribeApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeApplicationResponse{
		DescribeApplicationOutput: r.Request.Data.(*types.DescribeApplicationOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeApplicationResponse is the response type for the
// DescribeApplication API operation.
type DescribeApplicationResponse struct {
	*types.DescribeApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeApplication request.
func (r *DescribeApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
