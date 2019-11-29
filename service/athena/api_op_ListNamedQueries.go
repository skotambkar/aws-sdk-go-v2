// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

const opListNamedQueries = "ListNamedQueries"

// ListNamedQueriesRequest returns a request value for making API operation for
// Amazon Athena.
//
// Provides a list of available query IDs only for queries saved in the specified
// workgroup. Requires that you have access to the workgroup.
//
// For code samples using the AWS SDK for Java, see Examples and Code Samples
// (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in the Amazon
// Athena User Guide.
//
//    // Example sending a request using ListNamedQueriesRequest.
//    req := client.ListNamedQueriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/ListNamedQueries
func (c *Client) ListNamedQueriesRequest(input *types.ListNamedQueriesInput) ListNamedQueriesRequest {
	op := &aws.Operation{
		Name:       opListNamedQueries,
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
		input = &types.ListNamedQueriesInput{}
	}

	req := c.newRequest(op, input, &types.ListNamedQueriesOutput{})
	return ListNamedQueriesRequest{Request: req, Input: input, Copy: c.ListNamedQueriesRequest}
}

// ListNamedQueriesRequest is the request type for the
// ListNamedQueries API operation.
type ListNamedQueriesRequest struct {
	*aws.Request
	Input *types.ListNamedQueriesInput
	Copy  func(*types.ListNamedQueriesInput) ListNamedQueriesRequest
}

// Send marshals and sends the ListNamedQueries API request.
func (r ListNamedQueriesRequest) Send(ctx context.Context) (*ListNamedQueriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNamedQueriesResponse{
		ListNamedQueriesOutput: r.Request.Data.(*types.ListNamedQueriesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListNamedQueriesRequestPaginator returns a paginator for ListNamedQueries.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListNamedQueriesRequest(input)
//   p := athena.NewListNamedQueriesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListNamedQueriesPaginator(req ListNamedQueriesRequest) ListNamedQueriesPaginator {
	return ListNamedQueriesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListNamedQueriesInput
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

// ListNamedQueriesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListNamedQueriesPaginator struct {
	aws.Pager
}

func (p *ListNamedQueriesPaginator) CurrentPage() *types.ListNamedQueriesOutput {
	return p.Pager.CurrentPage().(*types.ListNamedQueriesOutput)
}

// ListNamedQueriesResponse is the response type for the
// ListNamedQueries API operation.
type ListNamedQueriesResponse struct {
	*types.ListNamedQueriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNamedQueries request.
func (r *ListNamedQueriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
