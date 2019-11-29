// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
)

const opListServiceQuotaIncreaseRequestsInTemplate = "ListServiceQuotaIncreaseRequestsInTemplate"

// ListServiceQuotaIncreaseRequestsInTemplateRequest returns a request value for making API operation for
// Service Quotas.
//
// Returns a list of the quota increase requests in the template.
//
//    // Example sending a request using ListServiceQuotaIncreaseRequestsInTemplateRequest.
//    req := client.ListServiceQuotaIncreaseRequestsInTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/ListServiceQuotaIncreaseRequestsInTemplate
func (c *Client) ListServiceQuotaIncreaseRequestsInTemplateRequest(input *types.ListServiceQuotaIncreaseRequestsInTemplateInput) ListServiceQuotaIncreaseRequestsInTemplateRequest {
	op := &aws.Operation{
		Name:       opListServiceQuotaIncreaseRequestsInTemplate,
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
		input = &types.ListServiceQuotaIncreaseRequestsInTemplateInput{}
	}

	req := c.newRequest(op, input, &types.ListServiceQuotaIncreaseRequestsInTemplateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListServiceQuotaIncreaseRequestsInTemplateMarshaler{Input: input}.GetNamedBuildHandler())

	return ListServiceQuotaIncreaseRequestsInTemplateRequest{Request: req, Input: input, Copy: c.ListServiceQuotaIncreaseRequestsInTemplateRequest}
}

// ListServiceQuotaIncreaseRequestsInTemplateRequest is the request type for the
// ListServiceQuotaIncreaseRequestsInTemplate API operation.
type ListServiceQuotaIncreaseRequestsInTemplateRequest struct {
	*aws.Request
	Input *types.ListServiceQuotaIncreaseRequestsInTemplateInput
	Copy  func(*types.ListServiceQuotaIncreaseRequestsInTemplateInput) ListServiceQuotaIncreaseRequestsInTemplateRequest
}

// Send marshals and sends the ListServiceQuotaIncreaseRequestsInTemplate API request.
func (r ListServiceQuotaIncreaseRequestsInTemplateRequest) Send(ctx context.Context) (*ListServiceQuotaIncreaseRequestsInTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServiceQuotaIncreaseRequestsInTemplateResponse{
		ListServiceQuotaIncreaseRequestsInTemplateOutput: r.Request.Data.(*types.ListServiceQuotaIncreaseRequestsInTemplateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServiceQuotaIncreaseRequestsInTemplateRequestPaginator returns a paginator for ListServiceQuotaIncreaseRequestsInTemplate.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServiceQuotaIncreaseRequestsInTemplateRequest(input)
//   p := servicequotas.NewListServiceQuotaIncreaseRequestsInTemplateRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServiceQuotaIncreaseRequestsInTemplatePaginator(req ListServiceQuotaIncreaseRequestsInTemplateRequest) ListServiceQuotaIncreaseRequestsInTemplatePaginator {
	return ListServiceQuotaIncreaseRequestsInTemplatePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListServiceQuotaIncreaseRequestsInTemplateInput
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

// ListServiceQuotaIncreaseRequestsInTemplatePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServiceQuotaIncreaseRequestsInTemplatePaginator struct {
	aws.Pager
}

func (p *ListServiceQuotaIncreaseRequestsInTemplatePaginator) CurrentPage() *types.ListServiceQuotaIncreaseRequestsInTemplateOutput {
	return p.Pager.CurrentPage().(*types.ListServiceQuotaIncreaseRequestsInTemplateOutput)
}

// ListServiceQuotaIncreaseRequestsInTemplateResponse is the response type for the
// ListServiceQuotaIncreaseRequestsInTemplate API operation.
type ListServiceQuotaIncreaseRequestsInTemplateResponse struct {
	*types.ListServiceQuotaIncreaseRequestsInTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServiceQuotaIncreaseRequestsInTemplate request.
func (r *ListServiceQuotaIncreaseRequestsInTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
