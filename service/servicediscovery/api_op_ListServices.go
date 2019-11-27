// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
)

const opListServices = "ListServices"

// ListServicesRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Lists summary information for all the services that are associated with one
// or more specified namespaces.
//
//    // Example sending a request using ListServicesRequest.
//    req := client.ListServicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/ListServices
func (c *Client) ListServicesRequest(input *types.ListServicesInput) ListServicesRequest {
	op := &aws.Operation{
		Name:       opListServices,
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
		input = &types.ListServicesInput{}
	}

	req := c.newRequest(op, input, &types.ListServicesOutput{})
	return ListServicesRequest{Request: req, Input: input, Copy: c.ListServicesRequest}
}

// ListServicesRequest is the request type for the
// ListServices API operation.
type ListServicesRequest struct {
	*aws.Request
	Input *types.ListServicesInput
	Copy  func(*types.ListServicesInput) ListServicesRequest
}

// Send marshals and sends the ListServices API request.
func (r ListServicesRequest) Send(ctx context.Context) (*ListServicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServicesResponse{
		ListServicesOutput: r.Request.Data.(*types.ListServicesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServicesRequestPaginator returns a paginator for ListServices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServicesRequest(input)
//   p := servicediscovery.NewListServicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServicesPaginator(req ListServicesRequest) ListServicesPaginator {
	return ListServicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListServicesInput
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

// ListServicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServicesPaginator struct {
	aws.Pager
}

func (p *ListServicesPaginator) CurrentPage() *types.ListServicesOutput {
	return p.Pager.CurrentPage().(*types.ListServicesOutput)
}

// ListServicesResponse is the response type for the
// ListServices API operation.
type ListServicesResponse struct {
	*types.ListServicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServices request.
func (r *ListServicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
