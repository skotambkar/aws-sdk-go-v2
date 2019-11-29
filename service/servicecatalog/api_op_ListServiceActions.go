// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opListServiceActions = "ListServiceActions"

// ListServiceActionsRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists all self-service actions.
//
//    // Example sending a request using ListServiceActionsRequest.
//    req := client.ListServiceActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListServiceActions
func (c *Client) ListServiceActionsRequest(input *types.ListServiceActionsInput) ListServiceActionsRequest {
	op := &aws.Operation{
		Name:       opListServiceActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListServiceActionsInput{}
	}

	req := c.newRequest(op, input, &types.ListServiceActionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListServiceActionsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListServiceActionsRequest{Request: req, Input: input, Copy: c.ListServiceActionsRequest}
}

// ListServiceActionsRequest is the request type for the
// ListServiceActions API operation.
type ListServiceActionsRequest struct {
	*aws.Request
	Input *types.ListServiceActionsInput
	Copy  func(*types.ListServiceActionsInput) ListServiceActionsRequest
}

// Send marshals and sends the ListServiceActions API request.
func (r ListServiceActionsRequest) Send(ctx context.Context) (*ListServiceActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServiceActionsResponse{
		ListServiceActionsOutput: r.Request.Data.(*types.ListServiceActionsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServiceActionsRequestPaginator returns a paginator for ListServiceActions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServiceActionsRequest(input)
//   p := servicecatalog.NewListServiceActionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServiceActionsPaginator(req ListServiceActionsRequest) ListServiceActionsPaginator {
	return ListServiceActionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListServiceActionsInput
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

// ListServiceActionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServiceActionsPaginator struct {
	aws.Pager
}

func (p *ListServiceActionsPaginator) CurrentPage() *types.ListServiceActionsOutput {
	return p.Pager.CurrentPage().(*types.ListServiceActionsOutput)
}

// ListServiceActionsResponse is the response type for the
// ListServiceActions API operation.
type ListServiceActionsResponse struct {
	*types.ListServiceActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServiceActions request.
func (r *ListServiceActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
