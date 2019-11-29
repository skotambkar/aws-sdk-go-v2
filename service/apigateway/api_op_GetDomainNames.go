// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetDomainNames = "GetDomainNames"

// GetDomainNamesRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Represents a collection of DomainName resources.
//
//    // Example sending a request using GetDomainNamesRequest.
//    req := client.GetDomainNamesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetDomainNamesRequest(input *types.GetDomainNamesInput) GetDomainNamesRequest {
	op := &aws.Operation{
		Name:       opGetDomainNames,
		HTTPMethod: "GET",
		HTTPPath:   "/domainnames",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"position"},
			OutputTokens:    []string{"position"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.GetDomainNamesInput{}
	}

	req := c.newRequest(op, input, &types.GetDomainNamesOutput{})
	return GetDomainNamesRequest{Request: req, Input: input, Copy: c.GetDomainNamesRequest}
}

// GetDomainNamesRequest is the request type for the
// GetDomainNames API operation.
type GetDomainNamesRequest struct {
	*aws.Request
	Input *types.GetDomainNamesInput
	Copy  func(*types.GetDomainNamesInput) GetDomainNamesRequest
}

// Send marshals and sends the GetDomainNames API request.
func (r GetDomainNamesRequest) Send(ctx context.Context) (*GetDomainNamesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainNamesResponse{
		GetDomainNamesOutput: r.Request.Data.(*types.GetDomainNamesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetDomainNamesRequestPaginator returns a paginator for GetDomainNames.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetDomainNamesRequest(input)
//   p := apigateway.NewGetDomainNamesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetDomainNamesPaginator(req GetDomainNamesRequest) GetDomainNamesPaginator {
	return GetDomainNamesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetDomainNamesInput
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

// GetDomainNamesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetDomainNamesPaginator struct {
	aws.Pager
}

func (p *GetDomainNamesPaginator) CurrentPage() *types.GetDomainNamesOutput {
	return p.Pager.CurrentPage().(*types.GetDomainNamesOutput)
}

// GetDomainNamesResponse is the response type for the
// GetDomainNames API operation.
type GetDomainNamesResponse struct {
	*types.GetDomainNamesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainNames request.
func (r *GetDomainNamesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
