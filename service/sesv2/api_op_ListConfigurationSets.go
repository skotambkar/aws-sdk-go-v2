// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const opListConfigurationSets = "ListConfigurationSets"

// ListConfigurationSetsRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// List all of the configuration sets associated with your account in the current
// region.
//
// Configuration sets are groups of rules that you can apply to the emails you
// send. You apply a configuration set to an email by including a reference
// to the configuration set in the headers of the email. When you apply a configuration
// set to an email, all of the rules in that configuration set are applied to
// the email.
//
//    // Example sending a request using ListConfigurationSetsRequest.
//    req := client.ListConfigurationSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/ListConfigurationSets
func (c *Client) ListConfigurationSetsRequest(input *types.ListConfigurationSetsInput) ListConfigurationSetsRequest {
	op := &aws.Operation{
		Name:       opListConfigurationSets,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/email/configuration-sets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListConfigurationSetsInput{}
	}

	req := c.newRequest(op, input, &types.ListConfigurationSetsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListConfigurationSetsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListConfigurationSetsRequest{Request: req, Input: input, Copy: c.ListConfigurationSetsRequest}
}

// ListConfigurationSetsRequest is the request type for the
// ListConfigurationSets API operation.
type ListConfigurationSetsRequest struct {
	*aws.Request
	Input *types.ListConfigurationSetsInput
	Copy  func(*types.ListConfigurationSetsInput) ListConfigurationSetsRequest
}

// Send marshals and sends the ListConfigurationSets API request.
func (r ListConfigurationSetsRequest) Send(ctx context.Context) (*ListConfigurationSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListConfigurationSetsResponse{
		ListConfigurationSetsOutput: r.Request.Data.(*types.ListConfigurationSetsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListConfigurationSetsRequestPaginator returns a paginator for ListConfigurationSets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListConfigurationSetsRequest(input)
//   p := sesv2.NewListConfigurationSetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListConfigurationSetsPaginator(req ListConfigurationSetsRequest) ListConfigurationSetsPaginator {
	return ListConfigurationSetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListConfigurationSetsInput
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

// ListConfigurationSetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListConfigurationSetsPaginator struct {
	aws.Pager
}

func (p *ListConfigurationSetsPaginator) CurrentPage() *types.ListConfigurationSetsOutput {
	return p.Pager.CurrentPage().(*types.ListConfigurationSetsOutput)
}

// ListConfigurationSetsResponse is the response type for the
// ListConfigurationSets API operation.
type ListConfigurationSetsResponse struct {
	*types.ListConfigurationSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListConfigurationSets request.
func (r *ListConfigurationSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
