// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
)

const opDescribeRobotApplication = "DescribeRobotApplication"

// DescribeRobotApplicationRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Describes a robot application.
//
//    // Example sending a request using DescribeRobotApplicationRequest.
//    req := client.DescribeRobotApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeRobotApplication
func (c *Client) DescribeRobotApplicationRequest(input *types.DescribeRobotApplicationInput) DescribeRobotApplicationRequest {
	op := &aws.Operation{
		Name:       opDescribeRobotApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/describeRobotApplication",
	}

	if input == nil {
		input = &types.DescribeRobotApplicationInput{}
	}

	req := c.newRequest(op, input, &types.DescribeRobotApplicationOutput{})
	return DescribeRobotApplicationRequest{Request: req, Input: input, Copy: c.DescribeRobotApplicationRequest}
}

// DescribeRobotApplicationRequest is the request type for the
// DescribeRobotApplication API operation.
type DescribeRobotApplicationRequest struct {
	*aws.Request
	Input *types.DescribeRobotApplicationInput
	Copy  func(*types.DescribeRobotApplicationInput) DescribeRobotApplicationRequest
}

// Send marshals and sends the DescribeRobotApplication API request.
func (r DescribeRobotApplicationRequest) Send(ctx context.Context) (*DescribeRobotApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRobotApplicationResponse{
		DescribeRobotApplicationOutput: r.Request.Data.(*types.DescribeRobotApplicationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRobotApplicationResponse is the response type for the
// DescribeRobotApplication API operation.
type DescribeRobotApplicationResponse struct {
	*types.DescribeRobotApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRobotApplication request.
func (r *DescribeRobotApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
