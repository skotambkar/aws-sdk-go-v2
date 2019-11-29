// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
)

const opDescribeLoggingOptions = "DescribeLoggingOptions"

// DescribeLoggingOptionsRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves the current settings of the AWS IoT Analytics logging options.
//
//    // Example sending a request using DescribeLoggingOptionsRequest.
//    req := client.DescribeLoggingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribeLoggingOptions
func (c *Client) DescribeLoggingOptionsRequest(input *types.DescribeLoggingOptionsInput) DescribeLoggingOptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeLoggingOptions,
		HTTPMethod: "GET",
		HTTPPath:   "/logging",
	}

	if input == nil {
		input = &types.DescribeLoggingOptionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeLoggingOptionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeLoggingOptionsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeLoggingOptionsRequest{Request: req, Input: input, Copy: c.DescribeLoggingOptionsRequest}
}

// DescribeLoggingOptionsRequest is the request type for the
// DescribeLoggingOptions API operation.
type DescribeLoggingOptionsRequest struct {
	*aws.Request
	Input *types.DescribeLoggingOptionsInput
	Copy  func(*types.DescribeLoggingOptionsInput) DescribeLoggingOptionsRequest
}

// Send marshals and sends the DescribeLoggingOptions API request.
func (r DescribeLoggingOptionsRequest) Send(ctx context.Context) (*DescribeLoggingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLoggingOptionsResponse{
		DescribeLoggingOptionsOutput: r.Request.Data.(*types.DescribeLoggingOptionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLoggingOptionsResponse is the response type for the
// DescribeLoggingOptions API operation.
type DescribeLoggingOptionsResponse struct {
	*types.DescribeLoggingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLoggingOptions request.
func (r *DescribeLoggingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
