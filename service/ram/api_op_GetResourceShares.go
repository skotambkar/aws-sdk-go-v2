// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/ram/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
)

const opGetResourceShares = "GetResourceShares"

// GetResourceSharesRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Gets the resource shares that you own or the resource shares that are shared
// with you.
//
//    // Example sending a request using GetResourceSharesRequest.
//    req := client.GetResourceSharesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/GetResourceShares
func (c *Client) GetResourceSharesRequest(input *types.GetResourceSharesInput) GetResourceSharesRequest {
	op := &aws.Operation{
		Name:       opGetResourceShares,
		HTTPMethod: "POST",
		HTTPPath:   "/getresourceshares",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.GetResourceSharesInput{}
	}

	req := c.newRequest(op, input, &types.GetResourceSharesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetResourceSharesMarshaler{Input: input}.GetNamedBuildHandler())

	return GetResourceSharesRequest{Request: req, Input: input, Copy: c.GetResourceSharesRequest}
}

// GetResourceSharesRequest is the request type for the
// GetResourceShares API operation.
type GetResourceSharesRequest struct {
	*aws.Request
	Input *types.GetResourceSharesInput
	Copy  func(*types.GetResourceSharesInput) GetResourceSharesRequest
}

// Send marshals and sends the GetResourceShares API request.
func (r GetResourceSharesRequest) Send(ctx context.Context) (*GetResourceSharesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourceSharesResponse{
		GetResourceSharesOutput: r.Request.Data.(*types.GetResourceSharesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetResourceSharesRequestPaginator returns a paginator for GetResourceShares.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetResourceSharesRequest(input)
//   p := ram.NewGetResourceSharesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetResourceSharesPaginator(req GetResourceSharesRequest) GetResourceSharesPaginator {
	return GetResourceSharesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetResourceSharesInput
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

// GetResourceSharesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetResourceSharesPaginator struct {
	aws.Pager
}

func (p *GetResourceSharesPaginator) CurrentPage() *types.GetResourceSharesOutput {
	return p.Pager.CurrentPage().(*types.GetResourceSharesOutput)
}

// GetResourceSharesResponse is the response type for the
// GetResourceShares API operation.
type GetResourceSharesResponse struct {
	*types.GetResourceSharesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResourceShares request.
func (r *GetResourceSharesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
