// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opListObjectChildren = "ListObjectChildren"

// ListObjectChildrenRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns a paginated list of child objects that are associated with a given
// object.
//
//    // Example sending a request using ListObjectChildrenRequest.
//    req := client.ListObjectChildrenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectChildren
func (c *Client) ListObjectChildrenRequest(input *types.ListObjectChildrenInput) ListObjectChildrenRequest {
	op := &aws.Operation{
		Name:       opListObjectChildren,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/children",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListObjectChildrenInput{}
	}

	req := c.newRequest(op, input, &types.ListObjectChildrenOutput{})
	return ListObjectChildrenRequest{Request: req, Input: input, Copy: c.ListObjectChildrenRequest}
}

// ListObjectChildrenRequest is the request type for the
// ListObjectChildren API operation.
type ListObjectChildrenRequest struct {
	*aws.Request
	Input *types.ListObjectChildrenInput
	Copy  func(*types.ListObjectChildrenInput) ListObjectChildrenRequest
}

// Send marshals and sends the ListObjectChildren API request.
func (r ListObjectChildrenRequest) Send(ctx context.Context) (*ListObjectChildrenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectChildrenResponse{
		ListObjectChildrenOutput: r.Request.Data.(*types.ListObjectChildrenOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectChildrenRequestPaginator returns a paginator for ListObjectChildren.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectChildrenRequest(input)
//   p := clouddirectory.NewListObjectChildrenRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectChildrenPaginator(req ListObjectChildrenRequest) ListObjectChildrenPaginator {
	return ListObjectChildrenPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListObjectChildrenInput
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

// ListObjectChildrenPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectChildrenPaginator struct {
	aws.Pager
}

func (p *ListObjectChildrenPaginator) CurrentPage() *types.ListObjectChildrenOutput {
	return p.Pager.CurrentPage().(*types.ListObjectChildrenOutput)
}

// ListObjectChildrenResponse is the response type for the
// ListObjectChildren API operation.
type ListObjectChildrenResponse struct {
	*types.ListObjectChildrenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjectChildren request.
func (r *ListObjectChildrenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
