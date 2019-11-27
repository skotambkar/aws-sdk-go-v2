// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
)

const opListMeshes = "ListMeshes"

// ListMeshesRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Returns a list of existing service meshes.
//
//    // Example sending a request using ListMeshesRequest.
//    req := client.ListMeshesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListMeshes
func (c *Client) ListMeshesRequest(input *types.ListMeshesInput) ListMeshesRequest {
	op := &aws.Operation{
		Name:       opListMeshes,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListMeshesInput{}
	}

	req := c.newRequest(op, input, &types.ListMeshesOutput{})
	return ListMeshesRequest{Request: req, Input: input, Copy: c.ListMeshesRequest}
}

// ListMeshesRequest is the request type for the
// ListMeshes API operation.
type ListMeshesRequest struct {
	*aws.Request
	Input *types.ListMeshesInput
	Copy  func(*types.ListMeshesInput) ListMeshesRequest
}

// Send marshals and sends the ListMeshes API request.
func (r ListMeshesRequest) Send(ctx context.Context) (*ListMeshesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMeshesResponse{
		ListMeshesOutput: r.Request.Data.(*types.ListMeshesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMeshesRequestPaginator returns a paginator for ListMeshes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMeshesRequest(input)
//   p := appmesh.NewListMeshesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMeshesPaginator(req ListMeshesRequest) ListMeshesPaginator {
	return ListMeshesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListMeshesInput
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

// ListMeshesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMeshesPaginator struct {
	aws.Pager
}

func (p *ListMeshesPaginator) CurrentPage() *types.ListMeshesOutput {
	return p.Pager.CurrentPage().(*types.ListMeshesOutput)
}

// ListMeshesResponse is the response type for the
// ListMeshes API operation.
type ListMeshesResponse struct {
	*types.ListMeshesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMeshes request.
func (r *ListMeshesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
