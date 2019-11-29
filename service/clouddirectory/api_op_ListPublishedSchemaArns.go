// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opListPublishedSchemaArns = "ListPublishedSchemaArns"

// ListPublishedSchemaArnsRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists the major version families of each published schema. If a major version
// ARN is provided as SchemaArn, the minor version revisions in that family
// are listed instead.
//
//    // Example sending a request using ListPublishedSchemaArnsRequest.
//    req := client.ListPublishedSchemaArnsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListPublishedSchemaArns
func (c *Client) ListPublishedSchemaArnsRequest(input *types.ListPublishedSchemaArnsInput) ListPublishedSchemaArnsRequest {
	op := &aws.Operation{
		Name:       opListPublishedSchemaArns,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema/published",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListPublishedSchemaArnsInput{}
	}

	req := c.newRequest(op, input, &types.ListPublishedSchemaArnsOutput{})
	return ListPublishedSchemaArnsRequest{Request: req, Input: input, Copy: c.ListPublishedSchemaArnsRequest}
}

// ListPublishedSchemaArnsRequest is the request type for the
// ListPublishedSchemaArns API operation.
type ListPublishedSchemaArnsRequest struct {
	*aws.Request
	Input *types.ListPublishedSchemaArnsInput
	Copy  func(*types.ListPublishedSchemaArnsInput) ListPublishedSchemaArnsRequest
}

// Send marshals and sends the ListPublishedSchemaArns API request.
func (r ListPublishedSchemaArnsRequest) Send(ctx context.Context) (*ListPublishedSchemaArnsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPublishedSchemaArnsResponse{
		ListPublishedSchemaArnsOutput: r.Request.Data.(*types.ListPublishedSchemaArnsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPublishedSchemaArnsRequestPaginator returns a paginator for ListPublishedSchemaArns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPublishedSchemaArnsRequest(input)
//   p := clouddirectory.NewListPublishedSchemaArnsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPublishedSchemaArnsPaginator(req ListPublishedSchemaArnsRequest) ListPublishedSchemaArnsPaginator {
	return ListPublishedSchemaArnsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPublishedSchemaArnsInput
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

// ListPublishedSchemaArnsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPublishedSchemaArnsPaginator struct {
	aws.Pager
}

func (p *ListPublishedSchemaArnsPaginator) CurrentPage() *types.ListPublishedSchemaArnsOutput {
	return p.Pager.CurrentPage().(*types.ListPublishedSchemaArnsOutput)
}

// ListPublishedSchemaArnsResponse is the response type for the
// ListPublishedSchemaArns API operation.
type ListPublishedSchemaArnsResponse struct {
	*types.ListPublishedSchemaArnsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPublishedSchemaArns request.
func (r *ListPublishedSchemaArnsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
