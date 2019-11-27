// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeNetworkInterfaces = "DescribeNetworkInterfaces"

// DescribeNetworkInterfacesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your network interfaces.
//
//    // Example sending a request using DescribeNetworkInterfacesRequest.
//    req := client.DescribeNetworkInterfacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeNetworkInterfaces
func (c *Client) DescribeNetworkInterfacesRequest(input *types.DescribeNetworkInterfacesInput) DescribeNetworkInterfacesRequest {
	op := &aws.Operation{
		Name:       opDescribeNetworkInterfaces,
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
		input = &types.DescribeNetworkInterfacesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeNetworkInterfacesOutput{})
	return DescribeNetworkInterfacesRequest{Request: req, Input: input, Copy: c.DescribeNetworkInterfacesRequest}
}

// DescribeNetworkInterfacesRequest is the request type for the
// DescribeNetworkInterfaces API operation.
type DescribeNetworkInterfacesRequest struct {
	*aws.Request
	Input *types.DescribeNetworkInterfacesInput
	Copy  func(*types.DescribeNetworkInterfacesInput) DescribeNetworkInterfacesRequest
}

// Send marshals and sends the DescribeNetworkInterfaces API request.
func (r DescribeNetworkInterfacesRequest) Send(ctx context.Context) (*DescribeNetworkInterfacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNetworkInterfacesResponse{
		DescribeNetworkInterfacesOutput: r.Request.Data.(*types.DescribeNetworkInterfacesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeNetworkInterfacesRequestPaginator returns a paginator for DescribeNetworkInterfaces.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeNetworkInterfacesRequest(input)
//   p := ec2.NewDescribeNetworkInterfacesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeNetworkInterfacesPaginator(req DescribeNetworkInterfacesRequest) DescribeNetworkInterfacesPaginator {
	return DescribeNetworkInterfacesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeNetworkInterfacesInput
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

// DescribeNetworkInterfacesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeNetworkInterfacesPaginator struct {
	aws.Pager
}

func (p *DescribeNetworkInterfacesPaginator) CurrentPage() *types.DescribeNetworkInterfacesOutput {
	return p.Pager.CurrentPage().(*types.DescribeNetworkInterfacesOutput)
}

// DescribeNetworkInterfacesResponse is the response type for the
// DescribeNetworkInterfaces API operation.
type DescribeNetworkInterfacesResponse struct {
	*types.DescribeNetworkInterfacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNetworkInterfaces request.
func (r *DescribeNetworkInterfacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
