// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opListTemplateAliases = "ListTemplateAliases"

// ListTemplateAliasesRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists all the aliases of a template.
//
// CLI syntax:
//
// aws quicksight list-template-aliases --aws-account-id 111122223333 —template-id
// 'reports_test_template'
//
//    // Example sending a request using ListTemplateAliasesRequest.
//    req := client.ListTemplateAliasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListTemplateAliases
func (c *Client) ListTemplateAliasesRequest(input *types.ListTemplateAliasesInput) ListTemplateAliasesRequest {
	op := &aws.Operation{
		Name:       opListTemplateAliases,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}/aliases",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListTemplateAliasesInput{}
	}

	req := c.newRequest(op, input, &types.ListTemplateAliasesOutput{})
	return ListTemplateAliasesRequest{Request: req, Input: input, Copy: c.ListTemplateAliasesRequest}
}

// ListTemplateAliasesRequest is the request type for the
// ListTemplateAliases API operation.
type ListTemplateAliasesRequest struct {
	*aws.Request
	Input *types.ListTemplateAliasesInput
	Copy  func(*types.ListTemplateAliasesInput) ListTemplateAliasesRequest
}

// Send marshals and sends the ListTemplateAliases API request.
func (r ListTemplateAliasesRequest) Send(ctx context.Context) (*ListTemplateAliasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTemplateAliasesResponse{
		ListTemplateAliasesOutput: r.Request.Data.(*types.ListTemplateAliasesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTemplateAliasesRequestPaginator returns a paginator for ListTemplateAliases.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTemplateAliasesRequest(input)
//   p := quicksight.NewListTemplateAliasesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTemplateAliasesPaginator(req ListTemplateAliasesRequest) ListTemplateAliasesPaginator {
	return ListTemplateAliasesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListTemplateAliasesInput
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

// ListTemplateAliasesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTemplateAliasesPaginator struct {
	aws.Pager
}

func (p *ListTemplateAliasesPaginator) CurrentPage() *types.ListTemplateAliasesOutput {
	return p.Pager.CurrentPage().(*types.ListTemplateAliasesOutput)
}

// ListTemplateAliasesResponse is the response type for the
// ListTemplateAliases API operation.
type ListTemplateAliasesResponse struct {
	*types.ListTemplateAliasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTemplateAliases request.
func (r *ListTemplateAliasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
