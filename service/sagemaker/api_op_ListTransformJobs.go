// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opListTransformJobs = "ListTransformJobs"

// ListTransformJobsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists transform jobs.
//
//    // Example sending a request using ListTransformJobsRequest.
//    req := client.ListTransformJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListTransformJobs
func (c *Client) ListTransformJobsRequest(input *types.ListTransformJobsInput) ListTransformJobsRequest {
	op := &aws.Operation{
		Name:       opListTransformJobs,
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
		input = &types.ListTransformJobsInput{}
	}

	req := c.newRequest(op, input, &types.ListTransformJobsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListTransformJobsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListTransformJobsRequest{Request: req, Input: input, Copy: c.ListTransformJobsRequest}
}

// ListTransformJobsRequest is the request type for the
// ListTransformJobs API operation.
type ListTransformJobsRequest struct {
	*aws.Request
	Input *types.ListTransformJobsInput
	Copy  func(*types.ListTransformJobsInput) ListTransformJobsRequest
}

// Send marshals and sends the ListTransformJobs API request.
func (r ListTransformJobsRequest) Send(ctx context.Context) (*ListTransformJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTransformJobsResponse{
		ListTransformJobsOutput: r.Request.Data.(*types.ListTransformJobsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTransformJobsRequestPaginator returns a paginator for ListTransformJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTransformJobsRequest(input)
//   p := sagemaker.NewListTransformJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTransformJobsPaginator(req ListTransformJobsRequest) ListTransformJobsPaginator {
	return ListTransformJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListTransformJobsInput
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

// ListTransformJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTransformJobsPaginator struct {
	aws.Pager
}

func (p *ListTransformJobsPaginator) CurrentPage() *types.ListTransformJobsOutput {
	return p.Pager.CurrentPage().(*types.ListTransformJobsOutput)
}

// ListTransformJobsResponse is the response type for the
// ListTransformJobs API operation.
type ListTransformJobsResponse struct {
	*types.ListTransformJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTransformJobs request.
func (r *ListTransformJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
