// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opListGateways = "ListGateways"

// ListGatewaysRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Lists gateways owned by an AWS account in an AWS Region specified in the
// request. The returned list is ordered by gateway Amazon Resource Name (ARN).
//
// By default, the operation returns a maximum of 100 gateways. This operation
// supports pagination that allows you to optionally reduce the number of gateways
// returned in a response.
//
// If you have more gateways than are returned in a response (that is, the response
// returns only a truncated list of your gateways), the response contains a
// marker that you can specify in your next request to fetch the next page of
// gateways.
//
//    // Example sending a request using ListGatewaysRequest.
//    req := client.ListGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/ListGateways
func (c *Client) ListGatewaysRequest(input *types.ListGatewaysInput) ListGatewaysRequest {
	op := &aws.Operation{
		Name:       opListGateways,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListGatewaysInput{}
	}

	req := c.newRequest(op, input, &types.ListGatewaysOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListGatewaysMarshaler{Input: input}.GetNamedBuildHandler())

	return ListGatewaysRequest{Request: req, Input: input, Copy: c.ListGatewaysRequest}
}

// ListGatewaysRequest is the request type for the
// ListGateways API operation.
type ListGatewaysRequest struct {
	*aws.Request
	Input *types.ListGatewaysInput
	Copy  func(*types.ListGatewaysInput) ListGatewaysRequest
}

// Send marshals and sends the ListGateways API request.
func (r ListGatewaysRequest) Send(ctx context.Context) (*ListGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGatewaysResponse{
		ListGatewaysOutput: r.Request.Data.(*types.ListGatewaysOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGatewaysRequestPaginator returns a paginator for ListGateways.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGatewaysRequest(input)
//   p := storagegateway.NewListGatewaysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGatewaysPaginator(req ListGatewaysRequest) ListGatewaysPaginator {
	return ListGatewaysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListGatewaysInput
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

// ListGatewaysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGatewaysPaginator struct {
	aws.Pager
}

func (p *ListGatewaysPaginator) CurrentPage() *types.ListGatewaysOutput {
	return p.Pager.CurrentPage().(*types.ListGatewaysOutput)
}

// ListGatewaysResponse is the response type for the
// ListGateways API operation.
type ListGatewaysResponse struct {
	*types.ListGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGateways request.
func (r *ListGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
