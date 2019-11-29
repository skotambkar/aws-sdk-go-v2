// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opListJobs = "ListJobs"

// ListJobsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the names of all job resources in this AWS account, or the resources
// with the specified tag. This operation allows you to see which resources
// are available in your account, and their names.
//
// This operation takes the optional Tags field, which you can use as a filter
// on the response so that tagged resources can be retrieved as a group. If
// you choose to use tags filtering, only resources with the tag are retrieved.
//
//    // Example sending a request using ListJobsRequest.
//    req := client.ListJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/ListJobs
func (c *Client) ListJobsRequest(input *types.ListJobsInput) ListJobsRequest {
	op := &aws.Operation{
		Name:       opListJobs,
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
		input = &types.ListJobsInput{}
	}

	req := c.newRequest(op, input, &types.ListJobsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListJobsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListJobsRequest{Request: req, Input: input, Copy: c.ListJobsRequest}
}

// ListJobsRequest is the request type for the
// ListJobs API operation.
type ListJobsRequest struct {
	*aws.Request
	Input *types.ListJobsInput
	Copy  func(*types.ListJobsInput) ListJobsRequest
}

// Send marshals and sends the ListJobs API request.
func (r ListJobsRequest) Send(ctx context.Context) (*ListJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJobsResponse{
		ListJobsOutput: r.Request.Data.(*types.ListJobsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListJobsRequestPaginator returns a paginator for ListJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListJobsRequest(input)
//   p := glue.NewListJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListJobsPaginator(req ListJobsRequest) ListJobsPaginator {
	return ListJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListJobsInput
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

// ListJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListJobsPaginator struct {
	aws.Pager
}

func (p *ListJobsPaginator) CurrentPage() *types.ListJobsOutput {
	return p.Pager.CurrentPage().(*types.ListJobsOutput)
}

// ListJobsResponse is the response type for the
// ListJobs API operation.
type ListJobsResponse struct {
	*types.ListJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJobs request.
func (r *ListJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
