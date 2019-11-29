// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opListFacetNames = "ListFacetNames"

// ListFacetNamesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Retrieves the names of facets that exist in a schema.
//
//    // Example sending a request using ListFacetNamesRequest.
//    req := client.ListFacetNamesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListFacetNames
func (c *Client) ListFacetNamesRequest(input *types.ListFacetNamesInput) ListFacetNamesRequest {
	op := &aws.Operation{
		Name:       opListFacetNames,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/facet/list",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListFacetNamesInput{}
	}

	req := c.newRequest(op, input, &types.ListFacetNamesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListFacetNamesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListFacetNamesRequest{Request: req, Input: input, Copy: c.ListFacetNamesRequest}
}

// ListFacetNamesRequest is the request type for the
// ListFacetNames API operation.
type ListFacetNamesRequest struct {
	*aws.Request
	Input *types.ListFacetNamesInput
	Copy  func(*types.ListFacetNamesInput) ListFacetNamesRequest
}

// Send marshals and sends the ListFacetNames API request.
func (r ListFacetNamesRequest) Send(ctx context.Context) (*ListFacetNamesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFacetNamesResponse{
		ListFacetNamesOutput: r.Request.Data.(*types.ListFacetNamesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFacetNamesRequestPaginator returns a paginator for ListFacetNames.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFacetNamesRequest(input)
//   p := clouddirectory.NewListFacetNamesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFacetNamesPaginator(req ListFacetNamesRequest) ListFacetNamesPaginator {
	return ListFacetNamesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListFacetNamesInput
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

// ListFacetNamesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFacetNamesPaginator struct {
	aws.Pager
}

func (p *ListFacetNamesPaginator) CurrentPage() *types.ListFacetNamesOutput {
	return p.Pager.CurrentPage().(*types.ListFacetNamesOutput)
}

// ListFacetNamesResponse is the response type for the
// ListFacetNames API operation.
type ListFacetNamesResponse struct {
	*types.ListFacetNamesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFacetNames request.
func (r *ListFacetNamesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
