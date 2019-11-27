// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
)

const opListJobsByPipeline = "ListJobsByPipeline"

// ListJobsByPipelineRequest returns a request value for making API operation for
// Amazon Elastic Transcoder.
//
// The ListJobsByPipeline operation gets a list of the jobs currently in a pipeline.
//
// Elastic Transcoder returns all of the jobs currently in the specified pipeline.
// The response body contains one element for each job that satisfies the search
// criteria.
//
//    // Example sending a request using ListJobsByPipelineRequest.
//    req := client.ListJobsByPipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListJobsByPipelineRequest(input *types.ListJobsByPipelineInput) ListJobsByPipelineRequest {
	op := &aws.Operation{
		Name:       opListJobsByPipeline,
		HTTPMethod: "GET",
		HTTPPath:   "/2012-09-25/jobsByPipeline/{PipelineId}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListJobsByPipelineInput{}
	}

	req := c.newRequest(op, input, &types.ListJobsByPipelineOutput{})
	return ListJobsByPipelineRequest{Request: req, Input: input, Copy: c.ListJobsByPipelineRequest}
}

// ListJobsByPipelineRequest is the request type for the
// ListJobsByPipeline API operation.
type ListJobsByPipelineRequest struct {
	*aws.Request
	Input *types.ListJobsByPipelineInput
	Copy  func(*types.ListJobsByPipelineInput) ListJobsByPipelineRequest
}

// Send marshals and sends the ListJobsByPipeline API request.
func (r ListJobsByPipelineRequest) Send(ctx context.Context) (*ListJobsByPipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJobsByPipelineResponse{
		ListJobsByPipelineOutput: r.Request.Data.(*types.ListJobsByPipelineOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListJobsByPipelineRequestPaginator returns a paginator for ListJobsByPipeline.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListJobsByPipelineRequest(input)
//   p := elastictranscoder.NewListJobsByPipelineRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListJobsByPipelinePaginator(req ListJobsByPipelineRequest) ListJobsByPipelinePaginator {
	return ListJobsByPipelinePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListJobsByPipelineInput
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

// ListJobsByPipelinePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListJobsByPipelinePaginator struct {
	aws.Pager
}

func (p *ListJobsByPipelinePaginator) CurrentPage() *types.ListJobsByPipelineOutput {
	return p.Pager.CurrentPage().(*types.ListJobsByPipelineOutput)
}

// ListJobsByPipelineResponse is the response type for the
// ListJobsByPipeline API operation.
type ListJobsByPipelineResponse struct {
	*types.ListJobsByPipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJobsByPipeline request.
func (r *ListJobsByPipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
