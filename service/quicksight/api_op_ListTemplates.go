// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opListTemplates = "ListTemplates"

// ListTemplatesRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists all the templates in the QuickSight account.
//
// CLI syntax:
//
// aws quicksight list-templates --aws-account-id 111122223333 --max-results
// 1 —next-token AYADeJuxwOypAndSoOn
//
//    // Example sending a request using ListTemplatesRequest.
//    req := client.ListTemplatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListTemplates
func (c *Client) ListTemplatesRequest(input *types.ListTemplatesInput) ListTemplatesRequest {
	op := &aws.Operation{
		Name:       opListTemplates,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/templates",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListTemplatesInput{}
	}

	req := c.newRequest(op, input, &types.ListTemplatesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListTemplatesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListTemplatesRequest{Request: req, Input: input, Copy: c.ListTemplatesRequest}
}

// ListTemplatesRequest is the request type for the
// ListTemplates API operation.
type ListTemplatesRequest struct {
	*aws.Request
	Input *types.ListTemplatesInput
	Copy  func(*types.ListTemplatesInput) ListTemplatesRequest
}

// Send marshals and sends the ListTemplates API request.
func (r ListTemplatesRequest) Send(ctx context.Context) (*ListTemplatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTemplatesResponse{
		ListTemplatesOutput: r.Request.Data.(*types.ListTemplatesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTemplatesRequestPaginator returns a paginator for ListTemplates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTemplatesRequest(input)
//   p := quicksight.NewListTemplatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTemplatesPaginator(req ListTemplatesRequest) ListTemplatesPaginator {
	return ListTemplatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListTemplatesInput
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

// ListTemplatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTemplatesPaginator struct {
	aws.Pager
}

func (p *ListTemplatesPaginator) CurrentPage() *types.ListTemplatesOutput {
	return p.Pager.CurrentPage().(*types.ListTemplatesOutput)
}

// ListTemplatesResponse is the response type for the
// ListTemplates API operation.
type ListTemplatesResponse struct {
	*types.ListTemplatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTemplates request.
func (r *ListTemplatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
