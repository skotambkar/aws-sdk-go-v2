// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opDescribeSchemas = "DescribeSchemas"

// DescribeSchemasRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Returns information about the schema for the specified endpoint.
//
//    // Example sending a request using DescribeSchemasRequest.
//    req := client.DescribeSchemasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeSchemas
func (c *Client) DescribeSchemasRequest(input *types.DescribeSchemasInput) DescribeSchemasRequest {
	op := &aws.Operation{
		Name:       opDescribeSchemas,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeSchemasInput{}
	}

	req := c.newRequest(op, input, &types.DescribeSchemasOutput{})
	return DescribeSchemasRequest{Request: req, Input: input, Copy: c.DescribeSchemasRequest}
}

// DescribeSchemasRequest is the request type for the
// DescribeSchemas API operation.
type DescribeSchemasRequest struct {
	*aws.Request
	Input *types.DescribeSchemasInput
	Copy  func(*types.DescribeSchemasInput) DescribeSchemasRequest
}

// Send marshals and sends the DescribeSchemas API request.
func (r DescribeSchemasRequest) Send(ctx context.Context) (*DescribeSchemasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSchemasResponse{
		DescribeSchemasOutput: r.Request.Data.(*types.DescribeSchemasOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeSchemasRequestPaginator returns a paginator for DescribeSchemas.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeSchemasRequest(input)
//   p := databasemigrationservice.NewDescribeSchemasRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeSchemasPaginator(req DescribeSchemasRequest) DescribeSchemasPaginator {
	return DescribeSchemasPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeSchemasInput
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

// DescribeSchemasPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeSchemasPaginator struct {
	aws.Pager
}

func (p *DescribeSchemasPaginator) CurrentPage() *types.DescribeSchemasOutput {
	return p.Pager.CurrentPage().(*types.DescribeSchemasOutput)
}

// DescribeSchemasResponse is the response type for the
// DescribeSchemas API operation.
type DescribeSchemasResponse struct {
	*types.DescribeSchemasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSchemas request.
func (r *DescribeSchemasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
