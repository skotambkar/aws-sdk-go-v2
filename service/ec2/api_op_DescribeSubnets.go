// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeSubnets = "DescribeSubnets"

// DescribeSubnetsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your subnets.
//
// For more information, see Your VPC and Subnets (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using DescribeSubnetsRequest.
//    req := client.DescribeSubnetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSubnets
func (c *Client) DescribeSubnetsRequest(input *types.DescribeSubnetsInput) DescribeSubnetsRequest {
	op := &aws.Operation{
		Name:       opDescribeSubnets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeSubnetsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeSubnetsOutput{})
	return DescribeSubnetsRequest{Request: req, Input: input, Copy: c.DescribeSubnetsRequest}
}

// DescribeSubnetsRequest is the request type for the
// DescribeSubnets API operation.
type DescribeSubnetsRequest struct {
	*aws.Request
	Input *types.DescribeSubnetsInput
	Copy  func(*types.DescribeSubnetsInput) DescribeSubnetsRequest
}

// Send marshals and sends the DescribeSubnets API request.
func (r DescribeSubnetsRequest) Send(ctx context.Context) (*DescribeSubnetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSubnetsResponse{
		DescribeSubnetsOutput: r.Request.Data.(*types.DescribeSubnetsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeSubnetsRequestPaginator returns a paginator for DescribeSubnets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeSubnetsRequest(input)
//   p := ec2.NewDescribeSubnetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeSubnetsPaginator(req DescribeSubnetsRequest) DescribeSubnetsPaginator {
	return DescribeSubnetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeSubnetsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeSubnetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeSubnetsPaginator struct {
	aws.Pager
}

func (p *DescribeSubnetsPaginator) CurrentPage() *types.DescribeSubnetsOutput {
	return p.Pager.CurrentPage().(*types.DescribeSubnetsOutput)
}

// DescribeSubnetsResponse is the response type for the
// DescribeSubnets API operation.
type DescribeSubnetsResponse struct {
	*types.DescribeSubnetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSubnets request.
func (r *DescribeSubnetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
