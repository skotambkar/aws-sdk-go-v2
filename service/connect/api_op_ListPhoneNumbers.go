// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
)

const opListPhoneNumbers = "ListPhoneNumbers"

// ListPhoneNumbersRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Provides information about the phone numbers for the specified Amazon Connect
// instance.
//
//    // Example sending a request using ListPhoneNumbersRequest.
//    req := client.ListPhoneNumbersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/ListPhoneNumbers
func (c *Client) ListPhoneNumbersRequest(input *types.ListPhoneNumbersInput) ListPhoneNumbersRequest {
	op := &aws.Operation{
		Name:       opListPhoneNumbers,
		HTTPMethod: "GET",
		HTTPPath:   "/phone-numbers-summary/{InstanceId}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListPhoneNumbersInput{}
	}

	req := c.newRequest(op, input, &types.ListPhoneNumbersOutput{})
	return ListPhoneNumbersRequest{Request: req, Input: input, Copy: c.ListPhoneNumbersRequest}
}

// ListPhoneNumbersRequest is the request type for the
// ListPhoneNumbers API operation.
type ListPhoneNumbersRequest struct {
	*aws.Request
	Input *types.ListPhoneNumbersInput
	Copy  func(*types.ListPhoneNumbersInput) ListPhoneNumbersRequest
}

// Send marshals and sends the ListPhoneNumbers API request.
func (r ListPhoneNumbersRequest) Send(ctx context.Context) (*ListPhoneNumbersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPhoneNumbersResponse{
		ListPhoneNumbersOutput: r.Request.Data.(*types.ListPhoneNumbersOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPhoneNumbersRequestPaginator returns a paginator for ListPhoneNumbers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPhoneNumbersRequest(input)
//   p := connect.NewListPhoneNumbersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPhoneNumbersPaginator(req ListPhoneNumbersRequest) ListPhoneNumbersPaginator {
	return ListPhoneNumbersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPhoneNumbersInput
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

// ListPhoneNumbersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPhoneNumbersPaginator struct {
	aws.Pager
}

func (p *ListPhoneNumbersPaginator) CurrentPage() *types.ListPhoneNumbersOutput {
	return p.Pager.CurrentPage().(*types.ListPhoneNumbersOutput)
}

// ListPhoneNumbersResponse is the response type for the
// ListPhoneNumbers API operation.
type ListPhoneNumbersResponse struct {
	*types.ListPhoneNumbersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPhoneNumbers request.
func (r *ListPhoneNumbersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
