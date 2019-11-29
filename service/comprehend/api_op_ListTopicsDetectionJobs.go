// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

const opListTopicsDetectionJobs = "ListTopicsDetectionJobs"

// ListTopicsDetectionJobsRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets a list of the topic detection jobs that you have submitted.
//
//    // Example sending a request using ListTopicsDetectionJobsRequest.
//    req := client.ListTopicsDetectionJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/ListTopicsDetectionJobs
func (c *Client) ListTopicsDetectionJobsRequest(input *types.ListTopicsDetectionJobsInput) ListTopicsDetectionJobsRequest {
	op := &aws.Operation{
		Name:       opListTopicsDetectionJobs,
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
		input = &types.ListTopicsDetectionJobsInput{}
	}

	req := c.newRequest(op, input, &types.ListTopicsDetectionJobsOutput{})
	return ListTopicsDetectionJobsRequest{Request: req, Input: input, Copy: c.ListTopicsDetectionJobsRequest}
}

// ListTopicsDetectionJobsRequest is the request type for the
// ListTopicsDetectionJobs API operation.
type ListTopicsDetectionJobsRequest struct {
	*aws.Request
	Input *types.ListTopicsDetectionJobsInput
	Copy  func(*types.ListTopicsDetectionJobsInput) ListTopicsDetectionJobsRequest
}

// Send marshals and sends the ListTopicsDetectionJobs API request.
func (r ListTopicsDetectionJobsRequest) Send(ctx context.Context) (*ListTopicsDetectionJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTopicsDetectionJobsResponse{
		ListTopicsDetectionJobsOutput: r.Request.Data.(*types.ListTopicsDetectionJobsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTopicsDetectionJobsRequestPaginator returns a paginator for ListTopicsDetectionJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTopicsDetectionJobsRequest(input)
//   p := comprehend.NewListTopicsDetectionJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTopicsDetectionJobsPaginator(req ListTopicsDetectionJobsRequest) ListTopicsDetectionJobsPaginator {
	return ListTopicsDetectionJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListTopicsDetectionJobsInput
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

// ListTopicsDetectionJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTopicsDetectionJobsPaginator struct {
	aws.Pager
}

func (p *ListTopicsDetectionJobsPaginator) CurrentPage() *types.ListTopicsDetectionJobsOutput {
	return p.Pager.CurrentPage().(*types.ListTopicsDetectionJobsOutput)
}

// ListTopicsDetectionJobsResponse is the response type for the
// ListTopicsDetectionJobs API operation.
type ListTopicsDetectionJobsResponse struct {
	*types.ListTopicsDetectionJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTopicsDetectionJobs request.
func (r *ListTopicsDetectionJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
