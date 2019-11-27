// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opListNotebookInstances = "ListNotebookInstances"

// ListNotebookInstancesRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns a list of the Amazon SageMaker notebook instances in the requester's
// account in an AWS Region.
//
//    // Example sending a request using ListNotebookInstancesRequest.
//    req := client.ListNotebookInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListNotebookInstances
func (c *Client) ListNotebookInstancesRequest(input *types.ListNotebookInstancesInput) ListNotebookInstancesRequest {
	op := &aws.Operation{
		Name:       opListNotebookInstances,
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
		input = &types.ListNotebookInstancesInput{}
	}

	req := c.newRequest(op, input, &types.ListNotebookInstancesOutput{})
	return ListNotebookInstancesRequest{Request: req, Input: input, Copy: c.ListNotebookInstancesRequest}
}

// ListNotebookInstancesRequest is the request type for the
// ListNotebookInstances API operation.
type ListNotebookInstancesRequest struct {
	*aws.Request
	Input *types.ListNotebookInstancesInput
	Copy  func(*types.ListNotebookInstancesInput) ListNotebookInstancesRequest
}

// Send marshals and sends the ListNotebookInstances API request.
func (r ListNotebookInstancesRequest) Send(ctx context.Context) (*ListNotebookInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNotebookInstancesResponse{
		ListNotebookInstancesOutput: r.Request.Data.(*types.ListNotebookInstancesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListNotebookInstancesRequestPaginator returns a paginator for ListNotebookInstances.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListNotebookInstancesRequest(input)
//   p := sagemaker.NewListNotebookInstancesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListNotebookInstancesPaginator(req ListNotebookInstancesRequest) ListNotebookInstancesPaginator {
	return ListNotebookInstancesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListNotebookInstancesInput
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

// ListNotebookInstancesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListNotebookInstancesPaginator struct {
	aws.Pager
}

func (p *ListNotebookInstancesPaginator) CurrentPage() *types.ListNotebookInstancesOutput {
	return p.Pager.CurrentPage().(*types.ListNotebookInstancesOutput)
}

// ListNotebookInstancesResponse is the response type for the
// ListNotebookInstances API operation.
type ListNotebookInstancesResponse struct {
	*types.ListNotebookInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNotebookInstances request.
func (r *ListNotebookInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
