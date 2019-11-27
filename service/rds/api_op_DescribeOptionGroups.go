// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opDescribeOptionGroups = "DescribeOptionGroups"

// DescribeOptionGroupsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Describes the available option groups.
//
//    // Example sending a request using DescribeOptionGroupsRequest.
//    req := client.DescribeOptionGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeOptionGroups
func (c *Client) DescribeOptionGroupsRequest(input *types.DescribeOptionGroupsInput) DescribeOptionGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeOptionGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeOptionGroupsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeOptionGroupsOutput{})
	return DescribeOptionGroupsRequest{Request: req, Input: input, Copy: c.DescribeOptionGroupsRequest}
}

// DescribeOptionGroupsRequest is the request type for the
// DescribeOptionGroups API operation.
type DescribeOptionGroupsRequest struct {
	*aws.Request
	Input *types.DescribeOptionGroupsInput
	Copy  func(*types.DescribeOptionGroupsInput) DescribeOptionGroupsRequest
}

// Send marshals and sends the DescribeOptionGroups API request.
func (r DescribeOptionGroupsRequest) Send(ctx context.Context) (*DescribeOptionGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOptionGroupsResponse{
		DescribeOptionGroupsOutput: r.Request.Data.(*types.DescribeOptionGroupsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeOptionGroupsRequestPaginator returns a paginator for DescribeOptionGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeOptionGroupsRequest(input)
//   p := rds.NewDescribeOptionGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeOptionGroupsPaginator(req DescribeOptionGroupsRequest) DescribeOptionGroupsPaginator {
	return DescribeOptionGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeOptionGroupsInput
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

// DescribeOptionGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeOptionGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeOptionGroupsPaginator) CurrentPage() *types.DescribeOptionGroupsOutput {
	return p.Pager.CurrentPage().(*types.DescribeOptionGroupsOutput)
}

// DescribeOptionGroupsResponse is the response type for the
// DescribeOptionGroups API operation.
type DescribeOptionGroupsResponse struct {
	*types.DescribeOptionGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOptionGroups request.
func (r *DescribeOptionGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
