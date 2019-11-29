// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
)

const opListQualificationTypes = "ListQualificationTypes"

// ListQualificationTypesRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The ListQualificationTypes operation returns a list of Qualification types,
// filtered by an optional search term.
//
//    // Example sending a request using ListQualificationTypesRequest.
//    req := client.ListQualificationTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/ListQualificationTypes
func (c *Client) ListQualificationTypesRequest(input *types.ListQualificationTypesInput) ListQualificationTypesRequest {
	op := &aws.Operation{
		Name:       opListQualificationTypes,
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
		input = &types.ListQualificationTypesInput{}
	}

	req := c.newRequest(op, input, &types.ListQualificationTypesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListQualificationTypesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListQualificationTypesRequest{Request: req, Input: input, Copy: c.ListQualificationTypesRequest}
}

// ListQualificationTypesRequest is the request type for the
// ListQualificationTypes API operation.
type ListQualificationTypesRequest struct {
	*aws.Request
	Input *types.ListQualificationTypesInput
	Copy  func(*types.ListQualificationTypesInput) ListQualificationTypesRequest
}

// Send marshals and sends the ListQualificationTypes API request.
func (r ListQualificationTypesRequest) Send(ctx context.Context) (*ListQualificationTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListQualificationTypesResponse{
		ListQualificationTypesOutput: r.Request.Data.(*types.ListQualificationTypesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListQualificationTypesRequestPaginator returns a paginator for ListQualificationTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListQualificationTypesRequest(input)
//   p := mturk.NewListQualificationTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListQualificationTypesPaginator(req ListQualificationTypesRequest) ListQualificationTypesPaginator {
	return ListQualificationTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListQualificationTypesInput
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

// ListQualificationTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListQualificationTypesPaginator struct {
	aws.Pager
}

func (p *ListQualificationTypesPaginator) CurrentPage() *types.ListQualificationTypesOutput {
	return p.Pager.CurrentPage().(*types.ListQualificationTypesOutput)
}

// ListQualificationTypesResponse is the response type for the
// ListQualificationTypes API operation.
type ListQualificationTypesResponse struct {
	*types.ListQualificationTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListQualificationTypes request.
func (r *ListQualificationTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
