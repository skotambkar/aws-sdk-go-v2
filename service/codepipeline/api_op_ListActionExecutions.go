// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
)

const opListActionExecutions = "ListActionExecutions"

// ListActionExecutionsRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Lists the action executions that have occurred in a pipeline.
//
//    // Example sending a request using ListActionExecutionsRequest.
//    req := client.ListActionExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/ListActionExecutions
func (c *Client) ListActionExecutionsRequest(input *types.ListActionExecutionsInput) ListActionExecutionsRequest {
	op := &aws.Operation{
		Name:       opListActionExecutions,
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
		input = &types.ListActionExecutionsInput{}
	}

	req := c.newRequest(op, input, &types.ListActionExecutionsOutput{})
	return ListActionExecutionsRequest{Request: req, Input: input, Copy: c.ListActionExecutionsRequest}
}

// ListActionExecutionsRequest is the request type for the
// ListActionExecutions API operation.
type ListActionExecutionsRequest struct {
	*aws.Request
	Input *types.ListActionExecutionsInput
	Copy  func(*types.ListActionExecutionsInput) ListActionExecutionsRequest
}

// Send marshals and sends the ListActionExecutions API request.
func (r ListActionExecutionsRequest) Send(ctx context.Context) (*ListActionExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListActionExecutionsResponse{
		ListActionExecutionsOutput: r.Request.Data.(*types.ListActionExecutionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListActionExecutionsRequestPaginator returns a paginator for ListActionExecutions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListActionExecutionsRequest(input)
//   p := codepipeline.NewListActionExecutionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListActionExecutionsPaginator(req ListActionExecutionsRequest) ListActionExecutionsPaginator {
	return ListActionExecutionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListActionExecutionsInput
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

// ListActionExecutionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListActionExecutionsPaginator struct {
	aws.Pager
}

func (p *ListActionExecutionsPaginator) CurrentPage() *types.ListActionExecutionsOutput {
	return p.Pager.CurrentPage().(*types.ListActionExecutionsOutput)
}

// ListActionExecutionsResponse is the response type for the
// ListActionExecutions API operation.
type ListActionExecutionsResponse struct {
	*types.ListActionExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListActionExecutions request.
func (r *ListActionExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
