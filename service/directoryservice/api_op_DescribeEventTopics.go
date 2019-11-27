// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
)

const opDescribeEventTopics = "DescribeEventTopics"

// DescribeEventTopicsRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Obtains information about which SNS topics receive status messages from the
// specified directory.
//
// If no input parameters are provided, such as DirectoryId or TopicName, this
// request describes all of the associations in the account.
//
//    // Example sending a request using DescribeEventTopicsRequest.
//    req := client.DescribeEventTopicsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DescribeEventTopics
func (c *Client) DescribeEventTopicsRequest(input *types.DescribeEventTopicsInput) DescribeEventTopicsRequest {
	op := &aws.Operation{
		Name:       opDescribeEventTopics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeEventTopicsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEventTopicsOutput{})
	return DescribeEventTopicsRequest{Request: req, Input: input, Copy: c.DescribeEventTopicsRequest}
}

// DescribeEventTopicsRequest is the request type for the
// DescribeEventTopics API operation.
type DescribeEventTopicsRequest struct {
	*aws.Request
	Input *types.DescribeEventTopicsInput
	Copy  func(*types.DescribeEventTopicsInput) DescribeEventTopicsRequest
}

// Send marshals and sends the DescribeEventTopics API request.
func (r DescribeEventTopicsRequest) Send(ctx context.Context) (*DescribeEventTopicsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventTopicsResponse{
		DescribeEventTopicsOutput: r.Request.Data.(*types.DescribeEventTopicsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEventTopicsResponse is the response type for the
// DescribeEventTopics API operation.
type DescribeEventTopicsResponse struct {
	*types.DescribeEventTopicsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventTopics request.
func (r *DescribeEventTopicsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
