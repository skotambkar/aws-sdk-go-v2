// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

const opListCollections = "ListCollections"

// ListCollectionsRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Returns list of collection IDs in your account. If the result is truncated,
// the response also provides a NextToken that you can use in the subsequent
// request to fetch the next set of collection IDs.
//
// For an example, see Listing Collections in the Amazon Rekognition Developer
// Guide.
//
// This operation requires permissions to perform the rekognition:ListCollections
// action.
//
//    // Example sending a request using ListCollectionsRequest.
//    req := client.ListCollectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListCollectionsRequest(input *types.ListCollectionsInput) ListCollectionsRequest {
	op := &aws.Operation{
		Name:       opListCollections,
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
		input = &types.ListCollectionsInput{}
	}

	req := c.newRequest(op, input, &types.ListCollectionsOutput{})
	return ListCollectionsRequest{Request: req, Input: input, Copy: c.ListCollectionsRequest}
}

// ListCollectionsRequest is the request type for the
// ListCollections API operation.
type ListCollectionsRequest struct {
	*aws.Request
	Input *types.ListCollectionsInput
	Copy  func(*types.ListCollectionsInput) ListCollectionsRequest
}

// Send marshals and sends the ListCollections API request.
func (r ListCollectionsRequest) Send(ctx context.Context) (*ListCollectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCollectionsResponse{
		ListCollectionsOutput: r.Request.Data.(*types.ListCollectionsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCollectionsRequestPaginator returns a paginator for ListCollections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCollectionsRequest(input)
//   p := rekognition.NewListCollectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCollectionsPaginator(req ListCollectionsRequest) ListCollectionsPaginator {
	return ListCollectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListCollectionsInput
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

// ListCollectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCollectionsPaginator struct {
	aws.Pager
}

func (p *ListCollectionsPaginator) CurrentPage() *types.ListCollectionsOutput {
	return p.Pager.CurrentPage().(*types.ListCollectionsOutput)
}

// ListCollectionsResponse is the response type for the
// ListCollections API operation.
type ListCollectionsResponse struct {
	*types.ListCollectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCollections request.
func (r *ListCollectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
