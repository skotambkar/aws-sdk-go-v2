// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opListOfferings = "ListOfferings"

// ListOfferingsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Returns a list of products or offerings that the user can manage through
// the API. Each offering record indicates the recurring price per unit and
// the frequency for that offering. The API returns a NotEligible error if the
// user is not permitted to invoke the operation. Please contact aws-devicefarm-support@amazon.com
// (mailto:aws-devicefarm-support@amazon.com) if you believe that you should
// be able to invoke this operation.
//
//    // Example sending a request using ListOfferingsRequest.
//    req := client.ListOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListOfferings
func (c *Client) ListOfferingsRequest(input *types.ListOfferingsInput) ListOfferingsRequest {
	op := &aws.Operation{
		Name:       opListOfferings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListOfferingsInput{}
	}

	req := c.newRequest(op, input, &types.ListOfferingsOutput{})
	return ListOfferingsRequest{Request: req, Input: input, Copy: c.ListOfferingsRequest}
}

// ListOfferingsRequest is the request type for the
// ListOfferings API operation.
type ListOfferingsRequest struct {
	*aws.Request
	Input *types.ListOfferingsInput
	Copy  func(*types.ListOfferingsInput) ListOfferingsRequest
}

// Send marshals and sends the ListOfferings API request.
func (r ListOfferingsRequest) Send(ctx context.Context) (*ListOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOfferingsResponse{
		ListOfferingsOutput: r.Request.Data.(*types.ListOfferingsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListOfferingsRequestPaginator returns a paginator for ListOfferings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListOfferingsRequest(input)
//   p := devicefarm.NewListOfferingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListOfferingsPaginator(req ListOfferingsRequest) ListOfferingsPaginator {
	return ListOfferingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListOfferingsInput
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

// ListOfferingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListOfferingsPaginator struct {
	aws.Pager
}

func (p *ListOfferingsPaginator) CurrentPage() *types.ListOfferingsOutput {
	return p.Pager.CurrentPage().(*types.ListOfferingsOutput)
}

// ListOfferingsResponse is the response type for the
// ListOfferings API operation.
type ListOfferingsResponse struct {
	*types.ListOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOfferings request.
func (r *ListOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
