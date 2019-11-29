// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeVpcClassicLink = "DescribeVpcClassicLink"

// DescribeVpcClassicLinkRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the ClassicLink status of one or more VPCs.
//
//    // Example sending a request using DescribeVpcClassicLinkRequest.
//    req := client.DescribeVpcClassicLinkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcClassicLink
func (c *Client) DescribeVpcClassicLinkRequest(input *types.DescribeVpcClassicLinkInput) DescribeVpcClassicLinkRequest {
	op := &aws.Operation{
		Name:       opDescribeVpcClassicLink,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeVpcClassicLinkInput{}
	}

	req := c.newRequest(op, input, &types.DescribeVpcClassicLinkOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DescribeVpcClassicLinkMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeVpcClassicLinkRequest{Request: req, Input: input, Copy: c.DescribeVpcClassicLinkRequest}
}

// DescribeVpcClassicLinkRequest is the request type for the
// DescribeVpcClassicLink API operation.
type DescribeVpcClassicLinkRequest struct {
	*aws.Request
	Input *types.DescribeVpcClassicLinkInput
	Copy  func(*types.DescribeVpcClassicLinkInput) DescribeVpcClassicLinkRequest
}

// Send marshals and sends the DescribeVpcClassicLink API request.
func (r DescribeVpcClassicLinkRequest) Send(ctx context.Context) (*DescribeVpcClassicLinkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpcClassicLinkResponse{
		DescribeVpcClassicLinkOutput: r.Request.Data.(*types.DescribeVpcClassicLinkOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVpcClassicLinkResponse is the response type for the
// DescribeVpcClassicLink API operation.
type DescribeVpcClassicLinkResponse struct {
	*types.DescribeVpcClassicLinkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpcClassicLink request.
func (r *DescribeVpcClassicLinkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
