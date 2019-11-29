// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
)

const opListUsers = "ListUsers"

// ListUsersRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Returns summaries of the organization's users.
//
//    // Example sending a request using ListUsersRequest.
//    req := client.ListUsersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ListUsers
func (c *Client) ListUsersRequest(input *types.ListUsersInput) ListUsersRequest {
	op := &aws.Operation{
		Name:       opListUsers,
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
		input = &types.ListUsersInput{}
	}

	req := c.newRequest(op, input, &types.ListUsersOutput{})
	return ListUsersRequest{Request: req, Input: input, Copy: c.ListUsersRequest}
}

// ListUsersRequest is the request type for the
// ListUsers API operation.
type ListUsersRequest struct {
	*aws.Request
	Input *types.ListUsersInput
	Copy  func(*types.ListUsersInput) ListUsersRequest
}

// Send marshals and sends the ListUsers API request.
func (r ListUsersRequest) Send(ctx context.Context) (*ListUsersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUsersResponse{
		ListUsersOutput: r.Request.Data.(*types.ListUsersOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListUsersRequestPaginator returns a paginator for ListUsers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListUsersRequest(input)
//   p := workmail.NewListUsersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListUsersPaginator(req ListUsersRequest) ListUsersPaginator {
	return ListUsersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListUsersInput
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

// ListUsersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListUsersPaginator struct {
	aws.Pager
}

func (p *ListUsersPaginator) CurrentPage() *types.ListUsersOutput {
	return p.Pager.CurrentPage().(*types.ListUsersOutput)
}

// ListUsersResponse is the response type for the
// ListUsers API operation.
type ListUsersResponse struct {
	*types.ListUsersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUsers request.
func (r *ListUsersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
