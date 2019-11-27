// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
)

const opDescribeCases = "DescribeCases"

// DescribeCasesRequest returns a request value for making API operation for
// AWS Support.
//
// Returns a list of cases that you specify by passing one or more case IDs.
// In addition, you can filter the cases by date by setting values for the afterTime
// and beforeTime request parameters. You can set values for the includeResolvedCases
// and includeCommunications request parameters to control how much information
// is returned.
//
// Case data is available for 12 months after creation. If a case was created
// more than 12 months ago, a request for data might cause an error.
//
// The response returns the following in JSON format:
//
//    * One or more CaseDetails data types.
//
//    * One or more nextToken values, which specify where to paginate the returned
//    records represented by the CaseDetails objects.
//
//    // Example sending a request using DescribeCasesRequest.
//    req := client.DescribeCasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeCases
func (c *Client) DescribeCasesRequest(input *types.DescribeCasesInput) DescribeCasesRequest {
	op := &aws.Operation{
		Name:       opDescribeCases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeCasesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeCasesOutput{})
	return DescribeCasesRequest{Request: req, Input: input, Copy: c.DescribeCasesRequest}
}

// DescribeCasesRequest is the request type for the
// DescribeCases API operation.
type DescribeCasesRequest struct {
	*aws.Request
	Input *types.DescribeCasesInput
	Copy  func(*types.DescribeCasesInput) DescribeCasesRequest
}

// Send marshals and sends the DescribeCases API request.
func (r DescribeCasesRequest) Send(ctx context.Context) (*DescribeCasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCasesResponse{
		DescribeCasesOutput: r.Request.Data.(*types.DescribeCasesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCasesRequestPaginator returns a paginator for DescribeCases.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCasesRequest(input)
//   p := support.NewDescribeCasesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCasesPaginator(req DescribeCasesRequest) DescribeCasesPaginator {
	return DescribeCasesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeCasesInput
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

// DescribeCasesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCasesPaginator struct {
	aws.Pager
}

func (p *DescribeCasesPaginator) CurrentPage() *types.DescribeCasesOutput {
	return p.Pager.CurrentPage().(*types.DescribeCasesOutput)
}

// DescribeCasesResponse is the response type for the
// DescribeCases API operation.
type DescribeCasesResponse struct {
	*types.DescribeCasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCases request.
func (r *DescribeCasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
