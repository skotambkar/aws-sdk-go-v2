// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opListTriggers = "ListTriggers"

// ListTriggersRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves the names of all trigger resources in this AWS account, or the
// resources with the specified tag. This operation allows you to see which
// resources are available in your account, and their names.
//
// This operation takes the optional Tags field, which you can use as a filter
// on the response so that tagged resources can be retrieved as a group. If
// you choose to use tags filtering, only resources with the tag are retrieved.
//
//    // Example sending a request using ListTriggersRequest.
//    req := client.ListTriggersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/ListTriggers
func (c *Client) ListTriggersRequest(input *types.ListTriggersInput) ListTriggersRequest {
	op := &aws.Operation{
		Name:       opListTriggers,
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
		input = &types.ListTriggersInput{}
	}

	req := c.newRequest(op, input, &types.ListTriggersOutput{})
	return ListTriggersRequest{Request: req, Input: input, Copy: c.ListTriggersRequest}
}

// ListTriggersRequest is the request type for the
// ListTriggers API operation.
type ListTriggersRequest struct {
	*aws.Request
	Input *types.ListTriggersInput
	Copy  func(*types.ListTriggersInput) ListTriggersRequest
}

// Send marshals and sends the ListTriggers API request.
func (r ListTriggersRequest) Send(ctx context.Context) (*ListTriggersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTriggersResponse{
		ListTriggersOutput: r.Request.Data.(*types.ListTriggersOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTriggersRequestPaginator returns a paginator for ListTriggers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTriggersRequest(input)
//   p := glue.NewListTriggersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTriggersPaginator(req ListTriggersRequest) ListTriggersPaginator {
	return ListTriggersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListTriggersInput
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

// ListTriggersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTriggersPaginator struct {
	aws.Pager
}

func (p *ListTriggersPaginator) CurrentPage() *types.ListTriggersOutput {
	return p.Pager.CurrentPage().(*types.ListTriggersOutput)
}

// ListTriggersResponse is the response type for the
// ListTriggers API operation.
type ListTriggersResponse struct {
	*types.ListTriggersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTriggers request.
func (r *ListTriggersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
