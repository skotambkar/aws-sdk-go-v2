// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opListObjectAttributes = "ListObjectAttributes"

// ListObjectAttributesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists all attributes that are associated with an object.
//
//    // Example sending a request using ListObjectAttributesRequest.
//    req := client.ListObjectAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectAttributes
func (c *Client) ListObjectAttributesRequest(input *types.ListObjectAttributesInput) ListObjectAttributesRequest {
	op := &aws.Operation{
		Name:       opListObjectAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/attributes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListObjectAttributesInput{}
	}

	req := c.newRequest(op, input, &types.ListObjectAttributesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListObjectAttributesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListObjectAttributesRequest{Request: req, Input: input, Copy: c.ListObjectAttributesRequest}
}

// ListObjectAttributesRequest is the request type for the
// ListObjectAttributes API operation.
type ListObjectAttributesRequest struct {
	*aws.Request
	Input *types.ListObjectAttributesInput
	Copy  func(*types.ListObjectAttributesInput) ListObjectAttributesRequest
}

// Send marshals and sends the ListObjectAttributes API request.
func (r ListObjectAttributesRequest) Send(ctx context.Context) (*ListObjectAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectAttributesResponse{
		ListObjectAttributesOutput: r.Request.Data.(*types.ListObjectAttributesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectAttributesRequestPaginator returns a paginator for ListObjectAttributes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectAttributesRequest(input)
//   p := clouddirectory.NewListObjectAttributesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectAttributesPaginator(req ListObjectAttributesRequest) ListObjectAttributesPaginator {
	return ListObjectAttributesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListObjectAttributesInput
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

// ListObjectAttributesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectAttributesPaginator struct {
	aws.Pager
}

func (p *ListObjectAttributesPaginator) CurrentPage() *types.ListObjectAttributesOutput {
	return p.Pager.CurrentPage().(*types.ListObjectAttributesOutput)
}

// ListObjectAttributesResponse is the response type for the
// ListObjectAttributes API operation.
type ListObjectAttributesResponse struct {
	*types.ListObjectAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjectAttributes request.
func (r *ListObjectAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
