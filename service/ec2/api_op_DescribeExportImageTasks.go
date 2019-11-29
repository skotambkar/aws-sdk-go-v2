// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeExportImageTasks = "DescribeExportImageTasks"

// DescribeExportImageTasksRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified export image tasks or all your export image tasks.
//
//    // Example sending a request using DescribeExportImageTasksRequest.
//    req := client.DescribeExportImageTasksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeExportImageTasks
func (c *Client) DescribeExportImageTasksRequest(input *types.DescribeExportImageTasksInput) DescribeExportImageTasksRequest {
	op := &aws.Operation{
		Name:       opDescribeExportImageTasks,
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
		input = &types.DescribeExportImageTasksInput{}
	}

	req := c.newRequest(op, input, &types.DescribeExportImageTasksOutput{})
	return DescribeExportImageTasksRequest{Request: req, Input: input, Copy: c.DescribeExportImageTasksRequest}
}

// DescribeExportImageTasksRequest is the request type for the
// DescribeExportImageTasks API operation.
type DescribeExportImageTasksRequest struct {
	*aws.Request
	Input *types.DescribeExportImageTasksInput
	Copy  func(*types.DescribeExportImageTasksInput) DescribeExportImageTasksRequest
}

// Send marshals and sends the DescribeExportImageTasks API request.
func (r DescribeExportImageTasksRequest) Send(ctx context.Context) (*DescribeExportImageTasksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeExportImageTasksResponse{
		DescribeExportImageTasksOutput: r.Request.Data.(*types.DescribeExportImageTasksOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeExportImageTasksRequestPaginator returns a paginator for DescribeExportImageTasks.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeExportImageTasksRequest(input)
//   p := ec2.NewDescribeExportImageTasksRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeExportImageTasksPaginator(req DescribeExportImageTasksRequest) DescribeExportImageTasksPaginator {
	return DescribeExportImageTasksPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeExportImageTasksInput
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

// DescribeExportImageTasksPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeExportImageTasksPaginator struct {
	aws.Pager
}

func (p *DescribeExportImageTasksPaginator) CurrentPage() *types.DescribeExportImageTasksOutput {
	return p.Pager.CurrentPage().(*types.DescribeExportImageTasksOutput)
}

// DescribeExportImageTasksResponse is the response type for the
// DescribeExportImageTasks API operation.
type DescribeExportImageTasksResponse struct {
	*types.DescribeExportImageTasksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeExportImageTasks request.
func (r *DescribeExportImageTasksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
