// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
)

const opListPipelineExecutions = "ListPipelineExecutions"

// ListPipelineExecutionsRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Gets a summary of the most recent executions for a pipeline.
//
//    // Example sending a request using ListPipelineExecutionsRequest.
//    req := client.ListPipelineExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/ListPipelineExecutions
func (c *Client) ListPipelineExecutionsRequest(input *types.ListPipelineExecutionsInput) ListPipelineExecutionsRequest {
	op := &aws.Operation{
		Name:       opListPipelineExecutions,
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
		input = &types.ListPipelineExecutionsInput{}
	}

	req := c.newRequest(op, input, &types.ListPipelineExecutionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListPipelineExecutionsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListPipelineExecutionsRequest{Request: req, Input: input, Copy: c.ListPipelineExecutionsRequest}
}

// ListPipelineExecutionsRequest is the request type for the
// ListPipelineExecutions API operation.
type ListPipelineExecutionsRequest struct {
	*aws.Request
	Input *types.ListPipelineExecutionsInput
	Copy  func(*types.ListPipelineExecutionsInput) ListPipelineExecutionsRequest
}

// Send marshals and sends the ListPipelineExecutions API request.
func (r ListPipelineExecutionsRequest) Send(ctx context.Context) (*ListPipelineExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPipelineExecutionsResponse{
		ListPipelineExecutionsOutput: r.Request.Data.(*types.ListPipelineExecutionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPipelineExecutionsRequestPaginator returns a paginator for ListPipelineExecutions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPipelineExecutionsRequest(input)
//   p := codepipeline.NewListPipelineExecutionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPipelineExecutionsPaginator(req ListPipelineExecutionsRequest) ListPipelineExecutionsPaginator {
	return ListPipelineExecutionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPipelineExecutionsInput
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

// ListPipelineExecutionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPipelineExecutionsPaginator struct {
	aws.Pager
}

func (p *ListPipelineExecutionsPaginator) CurrentPage() *types.ListPipelineExecutionsOutput {
	return p.Pager.CurrentPage().(*types.ListPipelineExecutionsOutput)
}

// ListPipelineExecutionsResponse is the response type for the
// ListPipelineExecutions API operation.
type ListPipelineExecutionsResponse struct {
	*types.ListPipelineExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPipelineExecutions request.
func (r *ListPipelineExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
