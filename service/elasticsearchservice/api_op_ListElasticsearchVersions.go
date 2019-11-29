// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
)

const opListElasticsearchVersions = "ListElasticsearchVersions"

// ListElasticsearchVersionsRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// List all supported Elasticsearch versions
//
//    // Example sending a request using ListElasticsearchVersionsRequest.
//    req := client.ListElasticsearchVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListElasticsearchVersionsRequest(input *types.ListElasticsearchVersionsInput) ListElasticsearchVersionsRequest {
	op := &aws.Operation{
		Name:       opListElasticsearchVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/versions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListElasticsearchVersionsInput{}
	}

	req := c.newRequest(op, input, &types.ListElasticsearchVersionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListElasticsearchVersionsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListElasticsearchVersionsRequest{Request: req, Input: input, Copy: c.ListElasticsearchVersionsRequest}
}

// ListElasticsearchVersionsRequest is the request type for the
// ListElasticsearchVersions API operation.
type ListElasticsearchVersionsRequest struct {
	*aws.Request
	Input *types.ListElasticsearchVersionsInput
	Copy  func(*types.ListElasticsearchVersionsInput) ListElasticsearchVersionsRequest
}

// Send marshals and sends the ListElasticsearchVersions API request.
func (r ListElasticsearchVersionsRequest) Send(ctx context.Context) (*ListElasticsearchVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListElasticsearchVersionsResponse{
		ListElasticsearchVersionsOutput: r.Request.Data.(*types.ListElasticsearchVersionsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListElasticsearchVersionsRequestPaginator returns a paginator for ListElasticsearchVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListElasticsearchVersionsRequest(input)
//   p := elasticsearchservice.NewListElasticsearchVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListElasticsearchVersionsPaginator(req ListElasticsearchVersionsRequest) ListElasticsearchVersionsPaginator {
	return ListElasticsearchVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListElasticsearchVersionsInput
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

// ListElasticsearchVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListElasticsearchVersionsPaginator struct {
	aws.Pager
}

func (p *ListElasticsearchVersionsPaginator) CurrentPage() *types.ListElasticsearchVersionsOutput {
	return p.Pager.CurrentPage().(*types.ListElasticsearchVersionsOutput)
}

// ListElasticsearchVersionsResponse is the response type for the
// ListElasticsearchVersions API operation.
type ListElasticsearchVersionsResponse struct {
	*types.ListElasticsearchVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListElasticsearchVersions request.
func (r *ListElasticsearchVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
