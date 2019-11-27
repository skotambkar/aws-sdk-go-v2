// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opGetConnections = "GetConnections"

// GetConnectionsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a list of connection definitions from the Data Catalog.
//
//    // Example sending a request using GetConnectionsRequest.
//    req := client.GetConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetConnections
func (c *Client) GetConnectionsRequest(input *types.GetConnectionsInput) GetConnectionsRequest {
	op := &aws.Operation{
		Name:       opGetConnections,
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
		input = &types.GetConnectionsInput{}
	}

	req := c.newRequest(op, input, &types.GetConnectionsOutput{})
	return GetConnectionsRequest{Request: req, Input: input, Copy: c.GetConnectionsRequest}
}

// GetConnectionsRequest is the request type for the
// GetConnections API operation.
type GetConnectionsRequest struct {
	*aws.Request
	Input *types.GetConnectionsInput
	Copy  func(*types.GetConnectionsInput) GetConnectionsRequest
}

// Send marshals and sends the GetConnections API request.
func (r GetConnectionsRequest) Send(ctx context.Context) (*GetConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectionsResponse{
		GetConnectionsOutput: r.Request.Data.(*types.GetConnectionsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetConnectionsRequestPaginator returns a paginator for GetConnections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetConnectionsRequest(input)
//   p := glue.NewGetConnectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetConnectionsPaginator(req GetConnectionsRequest) GetConnectionsPaginator {
	return GetConnectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetConnectionsInput
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

// GetConnectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetConnectionsPaginator struct {
	aws.Pager
}

func (p *GetConnectionsPaginator) CurrentPage() *types.GetConnectionsOutput {
	return p.Pager.CurrentPage().(*types.GetConnectionsOutput)
}

// GetConnectionsResponse is the response type for the
// GetConnections API operation.
type GetConnectionsResponse struct {
	*types.GetConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnections request.
func (r *GetConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
