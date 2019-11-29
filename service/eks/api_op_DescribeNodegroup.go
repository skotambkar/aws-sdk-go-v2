// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

const opDescribeNodegroup = "DescribeNodegroup"

// DescribeNodegroupRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Returns descriptive information about an Amazon EKS node group.
//
//    // Example sending a request using DescribeNodegroupRequest.
//    req := client.DescribeNodegroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/DescribeNodegroup
func (c *Client) DescribeNodegroupRequest(input *types.DescribeNodegroupInput) DescribeNodegroupRequest {
	op := &aws.Operation{
		Name:       opDescribeNodegroup,
		HTTPMethod: "GET",
		HTTPPath:   "/clusters/{name}/node-groups/{nodegroupName}",
	}

	if input == nil {
		input = &types.DescribeNodegroupInput{}
	}

	req := c.newRequest(op, input, &types.DescribeNodegroupOutput{})
	return DescribeNodegroupRequest{Request: req, Input: input, Copy: c.DescribeNodegroupRequest}
}

// DescribeNodegroupRequest is the request type for the
// DescribeNodegroup API operation.
type DescribeNodegroupRequest struct {
	*aws.Request
	Input *types.DescribeNodegroupInput
	Copy  func(*types.DescribeNodegroupInput) DescribeNodegroupRequest
}

// Send marshals and sends the DescribeNodegroup API request.
func (r DescribeNodegroupRequest) Send(ctx context.Context) (*DescribeNodegroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNodegroupResponse{
		DescribeNodegroupOutput: r.Request.Data.(*types.DescribeNodegroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeNodegroupResponse is the response type for the
// DescribeNodegroup API operation.
type DescribeNodegroupResponse struct {
	*types.DescribeNodegroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNodegroup request.
func (r *DescribeNodegroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
