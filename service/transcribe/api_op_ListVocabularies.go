// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
)

const opListVocabularies = "ListVocabularies"

// ListVocabulariesRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Returns a list of vocabularies that match the specified criteria. If no criteria
// are specified, returns the entire list of vocabularies.
//
//    // Example sending a request using ListVocabulariesRequest.
//    req := client.ListVocabulariesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/ListVocabularies
func (c *Client) ListVocabulariesRequest(input *types.ListVocabulariesInput) ListVocabulariesRequest {
	op := &aws.Operation{
		Name:       opListVocabularies,
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
		input = &types.ListVocabulariesInput{}
	}

	req := c.newRequest(op, input, &types.ListVocabulariesOutput{})
	return ListVocabulariesRequest{Request: req, Input: input, Copy: c.ListVocabulariesRequest}
}

// ListVocabulariesRequest is the request type for the
// ListVocabularies API operation.
type ListVocabulariesRequest struct {
	*aws.Request
	Input *types.ListVocabulariesInput
	Copy  func(*types.ListVocabulariesInput) ListVocabulariesRequest
}

// Send marshals and sends the ListVocabularies API request.
func (r ListVocabulariesRequest) Send(ctx context.Context) (*ListVocabulariesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVocabulariesResponse{
		ListVocabulariesOutput: r.Request.Data.(*types.ListVocabulariesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListVocabulariesRequestPaginator returns a paginator for ListVocabularies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListVocabulariesRequest(input)
//   p := transcribe.NewListVocabulariesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListVocabulariesPaginator(req ListVocabulariesRequest) ListVocabulariesPaginator {
	return ListVocabulariesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListVocabulariesInput
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

// ListVocabulariesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListVocabulariesPaginator struct {
	aws.Pager
}

func (p *ListVocabulariesPaginator) CurrentPage() *types.ListVocabulariesOutput {
	return p.Pager.CurrentPage().(*types.ListVocabulariesOutput)
}

// ListVocabulariesResponse is the response type for the
// ListVocabularies API operation.
type ListVocabulariesResponse struct {
	*types.ListVocabulariesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVocabularies request.
func (r *ListVocabulariesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
