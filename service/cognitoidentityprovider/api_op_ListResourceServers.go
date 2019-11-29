// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opListResourceServers = "ListResourceServers"

// ListResourceServersRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists the resource servers for a user pool.
//
//    // Example sending a request using ListResourceServersRequest.
//    req := client.ListResourceServersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListResourceServers
func (c *Client) ListResourceServersRequest(input *types.ListResourceServersInput) ListResourceServersRequest {
	op := &aws.Operation{
		Name:       opListResourceServers,
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
		input = &types.ListResourceServersInput{}
	}

	req := c.newRequest(op, input, &types.ListResourceServersOutput{})
	return ListResourceServersRequest{Request: req, Input: input, Copy: c.ListResourceServersRequest}
}

// ListResourceServersRequest is the request type for the
// ListResourceServers API operation.
type ListResourceServersRequest struct {
	*aws.Request
	Input *types.ListResourceServersInput
	Copy  func(*types.ListResourceServersInput) ListResourceServersRequest
}

// Send marshals and sends the ListResourceServers API request.
func (r ListResourceServersRequest) Send(ctx context.Context) (*ListResourceServersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourceServersResponse{
		ListResourceServersOutput: r.Request.Data.(*types.ListResourceServersOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListResourceServersRequestPaginator returns a paginator for ListResourceServers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListResourceServersRequest(input)
//   p := cognitoidentityprovider.NewListResourceServersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListResourceServersPaginator(req ListResourceServersRequest) ListResourceServersPaginator {
	return ListResourceServersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListResourceServersInput
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

// ListResourceServersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListResourceServersPaginator struct {
	aws.Pager
}

func (p *ListResourceServersPaginator) CurrentPage() *types.ListResourceServersOutput {
	return p.Pager.CurrentPage().(*types.ListResourceServersOutput)
}

// ListResourceServersResponse is the response type for the
// ListResourceServers API operation.
type ListResourceServersResponse struct {
	*types.ListResourceServersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResourceServers request.
func (r *ListResourceServersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
