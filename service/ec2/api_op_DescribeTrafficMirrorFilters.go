// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeTrafficMirrorFilters = "DescribeTrafficMirrorFilters"

// DescribeTrafficMirrorFiltersRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more Traffic Mirror filters.
//
//    // Example sending a request using DescribeTrafficMirrorFiltersRequest.
//    req := client.DescribeTrafficMirrorFiltersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTrafficMirrorFilters
func (c *Client) DescribeTrafficMirrorFiltersRequest(input *types.DescribeTrafficMirrorFiltersInput) DescribeTrafficMirrorFiltersRequest {
	op := &aws.Operation{
		Name:       opDescribeTrafficMirrorFilters,
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
		input = &types.DescribeTrafficMirrorFiltersInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTrafficMirrorFiltersOutput{})
	return DescribeTrafficMirrorFiltersRequest{Request: req, Input: input, Copy: c.DescribeTrafficMirrorFiltersRequest}
}

// DescribeTrafficMirrorFiltersRequest is the request type for the
// DescribeTrafficMirrorFilters API operation.
type DescribeTrafficMirrorFiltersRequest struct {
	*aws.Request
	Input *types.DescribeTrafficMirrorFiltersInput
	Copy  func(*types.DescribeTrafficMirrorFiltersInput) DescribeTrafficMirrorFiltersRequest
}

// Send marshals and sends the DescribeTrafficMirrorFilters API request.
func (r DescribeTrafficMirrorFiltersRequest) Send(ctx context.Context) (*DescribeTrafficMirrorFiltersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTrafficMirrorFiltersResponse{
		DescribeTrafficMirrorFiltersOutput: r.Request.Data.(*types.DescribeTrafficMirrorFiltersOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTrafficMirrorFiltersRequestPaginator returns a paginator for DescribeTrafficMirrorFilters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTrafficMirrorFiltersRequest(input)
//   p := ec2.NewDescribeTrafficMirrorFiltersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTrafficMirrorFiltersPaginator(req DescribeTrafficMirrorFiltersRequest) DescribeTrafficMirrorFiltersPaginator {
	return DescribeTrafficMirrorFiltersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeTrafficMirrorFiltersInput
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

// DescribeTrafficMirrorFiltersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTrafficMirrorFiltersPaginator struct {
	aws.Pager
}

func (p *DescribeTrafficMirrorFiltersPaginator) CurrentPage() *types.DescribeTrafficMirrorFiltersOutput {
	return p.Pager.CurrentPage().(*types.DescribeTrafficMirrorFiltersOutput)
}

// DescribeTrafficMirrorFiltersResponse is the response type for the
// DescribeTrafficMirrorFilters API operation.
type DescribeTrafficMirrorFiltersResponse struct {
	*types.DescribeTrafficMirrorFiltersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTrafficMirrorFilters request.
func (r *DescribeTrafficMirrorFiltersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
