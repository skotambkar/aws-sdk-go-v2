// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opDescribeStorediSCSIVolumes = "DescribeStorediSCSIVolumes"

// DescribeStorediSCSIVolumesRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns the description of the gateway volumes specified in the request.
// The list of gateway volumes in the request must be from one gateway. In the
// response Amazon Storage Gateway returns volume information sorted by volume
// ARNs. This operation is only supported in stored volume gateway type.
//
//    // Example sending a request using DescribeStorediSCSIVolumesRequest.
//    req := client.DescribeStorediSCSIVolumesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeStorediSCSIVolumes
func (c *Client) DescribeStorediSCSIVolumesRequest(input *types.DescribeStorediSCSIVolumesInput) DescribeStorediSCSIVolumesRequest {
	op := &aws.Operation{
		Name:       opDescribeStorediSCSIVolumes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeStorediSCSIVolumesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeStorediSCSIVolumesOutput{})
	return DescribeStorediSCSIVolumesRequest{Request: req, Input: input, Copy: c.DescribeStorediSCSIVolumesRequest}
}

// DescribeStorediSCSIVolumesRequest is the request type for the
// DescribeStorediSCSIVolumes API operation.
type DescribeStorediSCSIVolumesRequest struct {
	*aws.Request
	Input *types.DescribeStorediSCSIVolumesInput
	Copy  func(*types.DescribeStorediSCSIVolumesInput) DescribeStorediSCSIVolumesRequest
}

// Send marshals and sends the DescribeStorediSCSIVolumes API request.
func (r DescribeStorediSCSIVolumesRequest) Send(ctx context.Context) (*DescribeStorediSCSIVolumesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStorediSCSIVolumesResponse{
		DescribeStorediSCSIVolumesOutput: r.Request.Data.(*types.DescribeStorediSCSIVolumesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStorediSCSIVolumesResponse is the response type for the
// DescribeStorediSCSIVolumes API operation.
type DescribeStorediSCSIVolumesResponse struct {
	*types.DescribeStorediSCSIVolumesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStorediSCSIVolumes request.
func (r *DescribeStorediSCSIVolumesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
