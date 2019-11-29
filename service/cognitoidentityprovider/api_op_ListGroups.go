// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opListGroups = "ListGroups"

// ListGroupsRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists the groups associated with a user pool.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using ListGroupsRequest.
//    req := client.ListGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListGroups
func (c *Client) ListGroupsRequest(input *types.ListGroupsInput) ListGroupsRequest {
	op := &aws.Operation{
		Name:       opListGroups,
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
		input = &types.ListGroupsInput{}
	}

	req := c.newRequest(op, input, &types.ListGroupsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListGroupsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListGroupsRequest{Request: req, Input: input, Copy: c.ListGroupsRequest}
}

// ListGroupsRequest is the request type for the
// ListGroups API operation.
type ListGroupsRequest struct {
	*aws.Request
	Input *types.ListGroupsInput
	Copy  func(*types.ListGroupsInput) ListGroupsRequest
}

// Send marshals and sends the ListGroups API request.
func (r ListGroupsRequest) Send(ctx context.Context) (*ListGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupsResponse{
		ListGroupsOutput: r.Request.Data.(*types.ListGroupsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGroupsRequestPaginator returns a paginator for ListGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGroupsRequest(input)
//   p := cognitoidentityprovider.NewListGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGroupsPaginator(req ListGroupsRequest) ListGroupsPaginator {
	return ListGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListGroupsInput
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

// ListGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGroupsPaginator struct {
	aws.Pager
}

func (p *ListGroupsPaginator) CurrentPage() *types.ListGroupsOutput {
	return p.Pager.CurrentPage().(*types.ListGroupsOutput)
}

// ListGroupsResponse is the response type for the
// ListGroups API operation.
type ListGroupsResponse struct {
	*types.ListGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroups request.
func (r *ListGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
