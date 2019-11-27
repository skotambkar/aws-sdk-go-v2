// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package simpledb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/simpledb/types"
)

const opSelect = "Select"

// SelectRequest returns a request value for making API operation for
// Amazon SimpleDB.
//
// The Select operation returns a set of attributes for ItemNames that match
// the select expression. Select is similar to the standard SQL SELECT statement.
//
// The total size of the response cannot exceed 1 MB in total size. Amazon SimpleDB
// automatically adjusts the number of items returned per page to enforce this
// limit. For example, if the client asks to retrieve 2500 items, but each individual
// item is 10 kB in size, the system returns 100 items and an appropriate NextToken
// so the client can access the next page of results.
//
// For information on how to construct select expressions, see Using Select
// to Create Amazon SimpleDB Queries in the Developer Guide.
//
//    // Example sending a request using SelectRequest.
//    req := client.SelectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) SelectRequest(input *types.SelectInput) SelectRequest {
	op := &aws.Operation{
		Name:       opSelect,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.SelectInput{}
	}

	req := c.newRequest(op, input, &types.SelectOutput{})
	return SelectRequest{Request: req, Input: input, Copy: c.SelectRequest}
}

// SelectRequest is the request type for the
// Select API operation.
type SelectRequest struct {
	*aws.Request
	Input *types.SelectInput
	Copy  func(*types.SelectInput) SelectRequest
}

// Send marshals and sends the Select API request.
func (r SelectRequest) Send(ctx context.Context) (*SelectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SelectResponse{
		SelectOutput: r.Request.Data.(*types.SelectOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSelectRequestPaginator returns a paginator for Select.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SelectRequest(input)
//   p := simpledb.NewSelectRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSelectPaginator(req SelectRequest) SelectPaginator {
	return SelectPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.SelectInput
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

// SelectPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SelectPaginator struct {
	aws.Pager
}

func (p *SelectPaginator) CurrentPage() *types.SelectOutput {
	return p.Pager.CurrentPage().(*types.SelectOutput)
}

// SelectResponse is the response type for the
// Select API operation.
type SelectResponse struct {
	*types.SelectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Select request.
func (r *SelectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
