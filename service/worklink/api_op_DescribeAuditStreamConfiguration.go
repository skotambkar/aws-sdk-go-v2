// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
)

const opDescribeAuditStreamConfiguration = "DescribeAuditStreamConfiguration"

// DescribeAuditStreamConfigurationRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Describes the configuration for delivering audit streams to the customer
// account.
//
//    // Example sending a request using DescribeAuditStreamConfigurationRequest.
//    req := client.DescribeAuditStreamConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeAuditStreamConfiguration
func (c *Client) DescribeAuditStreamConfigurationRequest(input *types.DescribeAuditStreamConfigurationInput) DescribeAuditStreamConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeAuditStreamConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/describeAuditStreamConfiguration",
	}

	if input == nil {
		input = &types.DescribeAuditStreamConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.DescribeAuditStreamConfigurationOutput{})
	return DescribeAuditStreamConfigurationRequest{Request: req, Input: input, Copy: c.DescribeAuditStreamConfigurationRequest}
}

// DescribeAuditStreamConfigurationRequest is the request type for the
// DescribeAuditStreamConfiguration API operation.
type DescribeAuditStreamConfigurationRequest struct {
	*aws.Request
	Input *types.DescribeAuditStreamConfigurationInput
	Copy  func(*types.DescribeAuditStreamConfigurationInput) DescribeAuditStreamConfigurationRequest
}

// Send marshals and sends the DescribeAuditStreamConfiguration API request.
func (r DescribeAuditStreamConfigurationRequest) Send(ctx context.Context) (*DescribeAuditStreamConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAuditStreamConfigurationResponse{
		DescribeAuditStreamConfigurationOutput: r.Request.Data.(*types.DescribeAuditStreamConfigurationOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAuditStreamConfigurationResponse is the response type for the
// DescribeAuditStreamConfiguration API operation.
type DescribeAuditStreamConfigurationResponse struct {
	*types.DescribeAuditStreamConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAuditStreamConfiguration request.
func (r *DescribeAuditStreamConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
