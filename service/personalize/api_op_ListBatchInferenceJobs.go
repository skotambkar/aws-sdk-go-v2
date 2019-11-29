// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
)

const opListBatchInferenceJobs = "ListBatchInferenceJobs"

// ListBatchInferenceJobsRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Gets a list of the batch inference jobs that have been performed off of a
// solution version.
//
//    // Example sending a request using ListBatchInferenceJobsRequest.
//    req := client.ListBatchInferenceJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/ListBatchInferenceJobs
func (c *Client) ListBatchInferenceJobsRequest(input *types.ListBatchInferenceJobsInput) ListBatchInferenceJobsRequest {
	op := &aws.Operation{
		Name:       opListBatchInferenceJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListBatchInferenceJobsInput{}
	}

	req := c.newRequest(op, input, &types.ListBatchInferenceJobsOutput{})
	return ListBatchInferenceJobsRequest{Request: req, Input: input, Copy: c.ListBatchInferenceJobsRequest}
}

// ListBatchInferenceJobsRequest is the request type for the
// ListBatchInferenceJobs API operation.
type ListBatchInferenceJobsRequest struct {
	*aws.Request
	Input *types.ListBatchInferenceJobsInput
	Copy  func(*types.ListBatchInferenceJobsInput) ListBatchInferenceJobsRequest
}

// Send marshals and sends the ListBatchInferenceJobs API request.
func (r ListBatchInferenceJobsRequest) Send(ctx context.Context) (*ListBatchInferenceJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBatchInferenceJobsResponse{
		ListBatchInferenceJobsOutput: r.Request.Data.(*types.ListBatchInferenceJobsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListBatchInferenceJobsRequestPaginator returns a paginator for ListBatchInferenceJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListBatchInferenceJobsRequest(input)
//   p := personalize.NewListBatchInferenceJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListBatchInferenceJobsPaginator(req ListBatchInferenceJobsRequest) ListBatchInferenceJobsPaginator {
	return ListBatchInferenceJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListBatchInferenceJobsInput
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

// ListBatchInferenceJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListBatchInferenceJobsPaginator struct {
	aws.Pager
}

func (p *ListBatchInferenceJobsPaginator) CurrentPage() *types.ListBatchInferenceJobsOutput {
	return p.Pager.CurrentPage().(*types.ListBatchInferenceJobsOutput)
}

// ListBatchInferenceJobsResponse is the response type for the
// ListBatchInferenceJobs API operation.
type ListBatchInferenceJobsResponse struct {
	*types.ListBatchInferenceJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBatchInferenceJobs request.
func (r *ListBatchInferenceJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
