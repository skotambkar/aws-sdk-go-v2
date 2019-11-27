// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribePrefixLists = "DescribePrefixLists"

// DescribePrefixListsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes available AWS services in a prefix list format, which includes
// the prefix list name and prefix list ID of the service and the IP address
// range for the service. A prefix list ID is required for creating an outbound
// security group rule that allows traffic from a VPC to access an AWS service
// through a gateway VPC endpoint. Currently, the services that support this
// action are Amazon S3 and Amazon DynamoDB.
//
//    // Example sending a request using DescribePrefixListsRequest.
//    req := client.DescribePrefixListsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribePrefixLists
func (c *Client) DescribePrefixListsRequest(input *types.DescribePrefixListsInput) DescribePrefixListsRequest {
	op := &aws.Operation{
		Name:       opDescribePrefixLists,
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
		input = &types.DescribePrefixListsInput{}
	}

	req := c.newRequest(op, input, &types.DescribePrefixListsOutput{})
	return DescribePrefixListsRequest{Request: req, Input: input, Copy: c.DescribePrefixListsRequest}
}

// DescribePrefixListsRequest is the request type for the
// DescribePrefixLists API operation.
type DescribePrefixListsRequest struct {
	*aws.Request
	Input *types.DescribePrefixListsInput
	Copy  func(*types.DescribePrefixListsInput) DescribePrefixListsRequest
}

// Send marshals and sends the DescribePrefixLists API request.
func (r DescribePrefixListsRequest) Send(ctx context.Context) (*DescribePrefixListsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePrefixListsResponse{
		DescribePrefixListsOutput: r.Request.Data.(*types.DescribePrefixListsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribePrefixListsRequestPaginator returns a paginator for DescribePrefixLists.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribePrefixListsRequest(input)
//   p := ec2.NewDescribePrefixListsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribePrefixListsPaginator(req DescribePrefixListsRequest) DescribePrefixListsPaginator {
	return DescribePrefixListsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribePrefixListsInput
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

// DescribePrefixListsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribePrefixListsPaginator struct {
	aws.Pager
}

func (p *DescribePrefixListsPaginator) CurrentPage() *types.DescribePrefixListsOutput {
	return p.Pager.CurrentPage().(*types.DescribePrefixListsOutput)
}

// DescribePrefixListsResponse is the response type for the
// DescribePrefixLists API operation.
type DescribePrefixListsResponse struct {
	*types.DescribePrefixListsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePrefixLists request.
func (r *DescribePrefixListsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
