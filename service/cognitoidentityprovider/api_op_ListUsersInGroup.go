// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opListUsersInGroup = "ListUsersInGroup"

// ListUsersInGroupRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists the users in the specified group.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using ListUsersInGroupRequest.
//    req := client.ListUsersInGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListUsersInGroup
func (c *Client) ListUsersInGroupRequest(input *types.ListUsersInGroupInput) ListUsersInGroupRequest {
	op := &aws.Operation{
		Name:       opListUsersInGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListUsersInGroupInput{}
	}

	req := c.newRequest(op, input, &types.ListUsersInGroupOutput{})
	return ListUsersInGroupRequest{Request: req, Input: input, Copy: c.ListUsersInGroupRequest}
}

// ListUsersInGroupRequest is the request type for the
// ListUsersInGroup API operation.
type ListUsersInGroupRequest struct {
	*aws.Request
	Input *types.ListUsersInGroupInput
	Copy  func(*types.ListUsersInGroupInput) ListUsersInGroupRequest
}

// Send marshals and sends the ListUsersInGroup API request.
func (r ListUsersInGroupRequest) Send(ctx context.Context) (*ListUsersInGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUsersInGroupResponse{
		ListUsersInGroupOutput: r.Request.Data.(*types.ListUsersInGroupOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListUsersInGroupRequestPaginator returns a paginator for ListUsersInGroup.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListUsersInGroupRequest(input)
//   p := cognitoidentityprovider.NewListUsersInGroupRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListUsersInGroupPaginator(req ListUsersInGroupRequest) ListUsersInGroupPaginator {
	return ListUsersInGroupPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListUsersInGroupInput
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

// ListUsersInGroupPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListUsersInGroupPaginator struct {
	aws.Pager
}

func (p *ListUsersInGroupPaginator) CurrentPage() *types.ListUsersInGroupOutput {
	return p.Pager.CurrentPage().(*types.ListUsersInGroupOutput)
}

// ListUsersInGroupResponse is the response type for the
// ListUsersInGroup API operation.
type ListUsersInGroupResponse struct {
	*types.ListUsersInGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUsersInGroup request.
func (r *ListUsersInGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
