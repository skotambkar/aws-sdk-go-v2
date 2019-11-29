// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeTransitGateways = "DescribeTransitGateways"

// DescribeTransitGatewaysRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more transit gateways. By default, all transit gateways
// are described. Alternatively, you can filter the results.
//
//    // Example sending a request using DescribeTransitGatewaysRequest.
//    req := client.DescribeTransitGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTransitGateways
func (c *Client) DescribeTransitGatewaysRequest(input *types.DescribeTransitGatewaysInput) DescribeTransitGatewaysRequest {
	op := &aws.Operation{
		Name:       opDescribeTransitGateways,
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
		input = &types.DescribeTransitGatewaysInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTransitGatewaysOutput{})
	return DescribeTransitGatewaysRequest{Request: req, Input: input, Copy: c.DescribeTransitGatewaysRequest}
}

// DescribeTransitGatewaysRequest is the request type for the
// DescribeTransitGateways API operation.
type DescribeTransitGatewaysRequest struct {
	*aws.Request
	Input *types.DescribeTransitGatewaysInput
	Copy  func(*types.DescribeTransitGatewaysInput) DescribeTransitGatewaysRequest
}

// Send marshals and sends the DescribeTransitGateways API request.
func (r DescribeTransitGatewaysRequest) Send(ctx context.Context) (*DescribeTransitGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTransitGatewaysResponse{
		DescribeTransitGatewaysOutput: r.Request.Data.(*types.DescribeTransitGatewaysOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTransitGatewaysRequestPaginator returns a paginator for DescribeTransitGateways.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTransitGatewaysRequest(input)
//   p := ec2.NewDescribeTransitGatewaysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTransitGatewaysPaginator(req DescribeTransitGatewaysRequest) DescribeTransitGatewaysPaginator {
	return DescribeTransitGatewaysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeTransitGatewaysInput
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

// DescribeTransitGatewaysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTransitGatewaysPaginator struct {
	aws.Pager
}

func (p *DescribeTransitGatewaysPaginator) CurrentPage() *types.DescribeTransitGatewaysOutput {
	return p.Pager.CurrentPage().(*types.DescribeTransitGatewaysOutput)
}

// DescribeTransitGatewaysResponse is the response type for the
// DescribeTransitGateways API operation.
type DescribeTransitGatewaysResponse struct {
	*types.DescribeTransitGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTransitGateways request.
func (r *DescribeTransitGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
