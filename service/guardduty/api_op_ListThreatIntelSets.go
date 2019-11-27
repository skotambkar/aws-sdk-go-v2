// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
)

const opListThreatIntelSets = "ListThreatIntelSets"

// ListThreatIntelSetsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Lists the ThreatIntelSets of the GuardDuty service specified by the detector
// ID.
//
//    // Example sending a request using ListThreatIntelSetsRequest.
//    req := client.ListThreatIntelSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListThreatIntelSets
func (c *Client) ListThreatIntelSetsRequest(input *types.ListThreatIntelSetsInput) ListThreatIntelSetsRequest {
	op := &aws.Operation{
		Name:       opListThreatIntelSets,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/threatintelset",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListThreatIntelSetsInput{}
	}

	req := c.newRequest(op, input, &types.ListThreatIntelSetsOutput{})
	return ListThreatIntelSetsRequest{Request: req, Input: input, Copy: c.ListThreatIntelSetsRequest}
}

// ListThreatIntelSetsRequest is the request type for the
// ListThreatIntelSets API operation.
type ListThreatIntelSetsRequest struct {
	*aws.Request
	Input *types.ListThreatIntelSetsInput
	Copy  func(*types.ListThreatIntelSetsInput) ListThreatIntelSetsRequest
}

// Send marshals and sends the ListThreatIntelSets API request.
func (r ListThreatIntelSetsRequest) Send(ctx context.Context) (*ListThreatIntelSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListThreatIntelSetsResponse{
		ListThreatIntelSetsOutput: r.Request.Data.(*types.ListThreatIntelSetsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListThreatIntelSetsRequestPaginator returns a paginator for ListThreatIntelSets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListThreatIntelSetsRequest(input)
//   p := guardduty.NewListThreatIntelSetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListThreatIntelSetsPaginator(req ListThreatIntelSetsRequest) ListThreatIntelSetsPaginator {
	return ListThreatIntelSetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListThreatIntelSetsInput
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

// ListThreatIntelSetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListThreatIntelSetsPaginator struct {
	aws.Pager
}

func (p *ListThreatIntelSetsPaginator) CurrentPage() *types.ListThreatIntelSetsOutput {
	return p.Pager.CurrentPage().(*types.ListThreatIntelSetsOutput)
}

// ListThreatIntelSetsResponse is the response type for the
// ListThreatIntelSets API operation.
type ListThreatIntelSetsResponse struct {
	*types.ListThreatIntelSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListThreatIntelSets request.
func (r *ListThreatIntelSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
