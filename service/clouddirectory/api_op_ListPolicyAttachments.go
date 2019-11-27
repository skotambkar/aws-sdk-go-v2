// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opListPolicyAttachments = "ListPolicyAttachments"

// ListPolicyAttachmentsRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns all of the ObjectIdentifiers to which a given policy is attached.
//
//    // Example sending a request using ListPolicyAttachmentsRequest.
//    req := client.ListPolicyAttachmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListPolicyAttachments
func (c *Client) ListPolicyAttachmentsRequest(input *types.ListPolicyAttachmentsInput) ListPolicyAttachmentsRequest {
	op := &aws.Operation{
		Name:       opListPolicyAttachments,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/policy/attachment",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListPolicyAttachmentsInput{}
	}

	req := c.newRequest(op, input, &types.ListPolicyAttachmentsOutput{})
	return ListPolicyAttachmentsRequest{Request: req, Input: input, Copy: c.ListPolicyAttachmentsRequest}
}

// ListPolicyAttachmentsRequest is the request type for the
// ListPolicyAttachments API operation.
type ListPolicyAttachmentsRequest struct {
	*aws.Request
	Input *types.ListPolicyAttachmentsInput
	Copy  func(*types.ListPolicyAttachmentsInput) ListPolicyAttachmentsRequest
}

// Send marshals and sends the ListPolicyAttachments API request.
func (r ListPolicyAttachmentsRequest) Send(ctx context.Context) (*ListPolicyAttachmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPolicyAttachmentsResponse{
		ListPolicyAttachmentsOutput: r.Request.Data.(*types.ListPolicyAttachmentsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPolicyAttachmentsRequestPaginator returns a paginator for ListPolicyAttachments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPolicyAttachmentsRequest(input)
//   p := clouddirectory.NewListPolicyAttachmentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPolicyAttachmentsPaginator(req ListPolicyAttachmentsRequest) ListPolicyAttachmentsPaginator {
	return ListPolicyAttachmentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPolicyAttachmentsInput
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

// ListPolicyAttachmentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPolicyAttachmentsPaginator struct {
	aws.Pager
}

func (p *ListPolicyAttachmentsPaginator) CurrentPage() *types.ListPolicyAttachmentsOutput {
	return p.Pager.CurrentPage().(*types.ListPolicyAttachmentsOutput)
}

// ListPolicyAttachmentsResponse is the response type for the
// ListPolicyAttachments API operation.
type ListPolicyAttachmentsResponse struct {
	*types.ListPolicyAttachmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicyAttachments request.
func (r *ListPolicyAttachmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
