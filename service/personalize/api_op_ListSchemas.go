// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/personalize/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
)

const opListSchemas = "ListSchemas"

// ListSchemasRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Returns the list of schemas associated with the account. The response provides
// the properties for each schema, including the Amazon Resource Name (ARN).
// For more information on schemas, see CreateSchema.
//
//    // Example sending a request using ListSchemasRequest.
//    req := client.ListSchemasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/ListSchemas
func (c *Client) ListSchemasRequest(input *types.ListSchemasInput) ListSchemasRequest {
	op := &aws.Operation{
		Name:       opListSchemas,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListSchemasInput{}
	}

	req := c.newRequest(op, input, &types.ListSchemasOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListSchemasMarshaler{Input: input}.GetNamedBuildHandler())

	return ListSchemasRequest{Request: req, Input: input, Copy: c.ListSchemasRequest}
}

// ListSchemasRequest is the request type for the
// ListSchemas API operation.
type ListSchemasRequest struct {
	*aws.Request
	Input *types.ListSchemasInput
	Copy  func(*types.ListSchemasInput) ListSchemasRequest
}

// Send marshals and sends the ListSchemas API request.
func (r ListSchemasRequest) Send(ctx context.Context) (*ListSchemasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSchemasResponse{
		ListSchemasOutput: r.Request.Data.(*types.ListSchemasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSchemasRequestPaginator returns a paginator for ListSchemas.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSchemasRequest(input)
//   p := personalize.NewListSchemasRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSchemasPaginator(req ListSchemasRequest) ListSchemasPaginator {
	return ListSchemasPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListSchemasInput
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

// ListSchemasPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSchemasPaginator struct {
	aws.Pager
}

func (p *ListSchemasPaginator) CurrentPage() *types.ListSchemasOutput {
	return p.Pager.CurrentPage().(*types.ListSchemasOutput)
}

// ListSchemasResponse is the response type for the
// ListSchemas API operation.
type ListSchemasResponse struct {
	*types.ListSchemasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSchemas request.
func (r *ListSchemasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
