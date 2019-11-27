// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
)

const opDescribeAsset = "DescribeAsset"

// DescribeAssetRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Returns a description of a MediaPackage VOD Asset resource.
//
//    // Example sending a request using DescribeAssetRequest.
//    req := client.DescribeAssetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/DescribeAsset
func (c *Client) DescribeAssetRequest(input *types.DescribeAssetInput) DescribeAssetRequest {
	op := &aws.Operation{
		Name:       opDescribeAsset,
		HTTPMethod: "GET",
		HTTPPath:   "/assets/{id}",
	}

	if input == nil {
		input = &types.DescribeAssetInput{}
	}

	req := c.newRequest(op, input, &types.DescribeAssetOutput{})
	return DescribeAssetRequest{Request: req, Input: input, Copy: c.DescribeAssetRequest}
}

// DescribeAssetRequest is the request type for the
// DescribeAsset API operation.
type DescribeAssetRequest struct {
	*aws.Request
	Input *types.DescribeAssetInput
	Copy  func(*types.DescribeAssetInput) DescribeAssetRequest
}

// Send marshals and sends the DescribeAsset API request.
func (r DescribeAssetRequest) Send(ctx context.Context) (*DescribeAssetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAssetResponse{
		DescribeAssetOutput: r.Request.Data.(*types.DescribeAssetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAssetResponse is the response type for the
// DescribeAsset API operation.
type DescribeAssetResponse struct {
	*types.DescribeAssetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAsset request.
func (r *DescribeAssetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
