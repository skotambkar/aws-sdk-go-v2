// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeVolumeAttribute = "DescribeVolumeAttribute"

// DescribeVolumeAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified attribute of the specified volume. You can specify
// only one attribute at a time.
//
// For more information about EBS volumes, see Amazon EBS Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumes.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeVolumeAttributeRequest.
//    req := client.DescribeVolumeAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVolumeAttribute
func (c *Client) DescribeVolumeAttributeRequest(input *types.DescribeVolumeAttributeInput) DescribeVolumeAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeVolumeAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeVolumeAttributeInput{}
	}

	req := c.newRequest(op, input, &types.DescribeVolumeAttributeOutput{})
	return DescribeVolumeAttributeRequest{Request: req, Input: input, Copy: c.DescribeVolumeAttributeRequest}
}

// DescribeVolumeAttributeRequest is the request type for the
// DescribeVolumeAttribute API operation.
type DescribeVolumeAttributeRequest struct {
	*aws.Request
	Input *types.DescribeVolumeAttributeInput
	Copy  func(*types.DescribeVolumeAttributeInput) DescribeVolumeAttributeRequest
}

// Send marshals and sends the DescribeVolumeAttribute API request.
func (r DescribeVolumeAttributeRequest) Send(ctx context.Context) (*DescribeVolumeAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVolumeAttributeResponse{
		DescribeVolumeAttributeOutput: r.Request.Data.(*types.DescribeVolumeAttributeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVolumeAttributeResponse is the response type for the
// DescribeVolumeAttribute API operation.
type DescribeVolumeAttributeResponse struct {
	*types.DescribeVolumeAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVolumeAttribute request.
func (r *DescribeVolumeAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
