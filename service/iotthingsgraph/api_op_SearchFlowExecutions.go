// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
)

const opSearchFlowExecutions = "SearchFlowExecutions"

// SearchFlowExecutionsRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Searches for AWS IoT Things Graph workflow execution instances.
//
//    // Example sending a request using SearchFlowExecutionsRequest.
//    req := client.SearchFlowExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchFlowExecutions
func (c *Client) SearchFlowExecutionsRequest(input *types.SearchFlowExecutionsInput) SearchFlowExecutionsRequest {
	op := &aws.Operation{
		Name:       opSearchFlowExecutions,
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
		input = &types.SearchFlowExecutionsInput{}
	}

	req := c.newRequest(op, input, &types.SearchFlowExecutionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.SearchFlowExecutionsMarshaler{Input: input}.GetNamedBuildHandler())

	return SearchFlowExecutionsRequest{Request: req, Input: input, Copy: c.SearchFlowExecutionsRequest}
}

// SearchFlowExecutionsRequest is the request type for the
// SearchFlowExecutions API operation.
type SearchFlowExecutionsRequest struct {
	*aws.Request
	Input *types.SearchFlowExecutionsInput
	Copy  func(*types.SearchFlowExecutionsInput) SearchFlowExecutionsRequest
}

// Send marshals and sends the SearchFlowExecutions API request.
func (r SearchFlowExecutionsRequest) Send(ctx context.Context) (*SearchFlowExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchFlowExecutionsResponse{
		SearchFlowExecutionsOutput: r.Request.Data.(*types.SearchFlowExecutionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchFlowExecutionsRequestPaginator returns a paginator for SearchFlowExecutions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchFlowExecutionsRequest(input)
//   p := iotthingsgraph.NewSearchFlowExecutionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchFlowExecutionsPaginator(req SearchFlowExecutionsRequest) SearchFlowExecutionsPaginator {
	return SearchFlowExecutionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.SearchFlowExecutionsInput
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

// SearchFlowExecutionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchFlowExecutionsPaginator struct {
	aws.Pager
}

func (p *SearchFlowExecutionsPaginator) CurrentPage() *types.SearchFlowExecutionsOutput {
	return p.Pager.CurrentPage().(*types.SearchFlowExecutionsOutput)
}

// SearchFlowExecutionsResponse is the response type for the
// SearchFlowExecutions API operation.
type SearchFlowExecutionsResponse struct {
	*types.SearchFlowExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchFlowExecutions request.
func (r *SearchFlowExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
