// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opListTypes = "ListTypes"

// ListTypesRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns summary information about types that have been registered with CloudFormation.
//
//    // Example sending a request using ListTypesRequest.
//    req := client.ListTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListTypes
func (c *Client) ListTypesRequest(input *types.ListTypesInput) ListTypesRequest {
	op := &aws.Operation{
		Name:       opListTypes,
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
		input = &types.ListTypesInput{}
	}

	req := c.newRequest(op, input, &types.ListTypesOutput{})
	return ListTypesRequest{Request: req, Input: input, Copy: c.ListTypesRequest}
}

// ListTypesRequest is the request type for the
// ListTypes API operation.
type ListTypesRequest struct {
	*aws.Request
	Input *types.ListTypesInput
	Copy  func(*types.ListTypesInput) ListTypesRequest
}

// Send marshals and sends the ListTypes API request.
func (r ListTypesRequest) Send(ctx context.Context) (*ListTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTypesResponse{
		ListTypesOutput: r.Request.Data.(*types.ListTypesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTypesRequestPaginator returns a paginator for ListTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTypesRequest(input)
//   p := cloudformation.NewListTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTypesPaginator(req ListTypesRequest) ListTypesPaginator {
	return ListTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListTypesInput
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

// ListTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTypesPaginator struct {
	aws.Pager
}

func (p *ListTypesPaginator) CurrentPage() *types.ListTypesOutput {
	return p.Pager.CurrentPage().(*types.ListTypesOutput)
}

// ListTypesResponse is the response type for the
// ListTypes API operation.
type ListTypesResponse struct {
	*types.ListTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTypes request.
func (r *ListTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
