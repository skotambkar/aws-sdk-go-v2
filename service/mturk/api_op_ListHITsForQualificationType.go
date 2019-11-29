// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
)

const opListHITsForQualificationType = "ListHITsForQualificationType"

// ListHITsForQualificationTypeRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The ListHITsForQualificationType operation returns the HITs that use the
// given Qualification type for a Qualification requirement. The operation returns
// HITs of any status, except for HITs that have been deleted with the DeleteHIT
// operation or that have been auto-deleted.
//
//    // Example sending a request using ListHITsForQualificationTypeRequest.
//    req := client.ListHITsForQualificationTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/ListHITsForQualificationType
func (c *Client) ListHITsForQualificationTypeRequest(input *types.ListHITsForQualificationTypeInput) ListHITsForQualificationTypeRequest {
	op := &aws.Operation{
		Name:       opListHITsForQualificationType,
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
		input = &types.ListHITsForQualificationTypeInput{}
	}

	req := c.newRequest(op, input, &types.ListHITsForQualificationTypeOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListHITsForQualificationTypeMarshaler{Input: input}.GetNamedBuildHandler())

	return ListHITsForQualificationTypeRequest{Request: req, Input: input, Copy: c.ListHITsForQualificationTypeRequest}
}

// ListHITsForQualificationTypeRequest is the request type for the
// ListHITsForQualificationType API operation.
type ListHITsForQualificationTypeRequest struct {
	*aws.Request
	Input *types.ListHITsForQualificationTypeInput
	Copy  func(*types.ListHITsForQualificationTypeInput) ListHITsForQualificationTypeRequest
}

// Send marshals and sends the ListHITsForQualificationType API request.
func (r ListHITsForQualificationTypeRequest) Send(ctx context.Context) (*ListHITsForQualificationTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHITsForQualificationTypeResponse{
		ListHITsForQualificationTypeOutput: r.Request.Data.(*types.ListHITsForQualificationTypeOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListHITsForQualificationTypeRequestPaginator returns a paginator for ListHITsForQualificationType.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListHITsForQualificationTypeRequest(input)
//   p := mturk.NewListHITsForQualificationTypeRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListHITsForQualificationTypePaginator(req ListHITsForQualificationTypeRequest) ListHITsForQualificationTypePaginator {
	return ListHITsForQualificationTypePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListHITsForQualificationTypeInput
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

// ListHITsForQualificationTypePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListHITsForQualificationTypePaginator struct {
	aws.Pager
}

func (p *ListHITsForQualificationTypePaginator) CurrentPage() *types.ListHITsForQualificationTypeOutput {
	return p.Pager.CurrentPage().(*types.ListHITsForQualificationTypeOutput)
}

// ListHITsForQualificationTypeResponse is the response type for the
// ListHITsForQualificationType API operation.
type ListHITsForQualificationTypeResponse struct {
	*types.ListHITsForQualificationTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHITsForQualificationType request.
func (r *ListHITsForQualificationTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
