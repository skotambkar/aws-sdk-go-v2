// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
)

const opListPackagingGroups = "ListPackagingGroups"

// ListPackagingGroupsRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Returns a collection of MediaPackage VOD PackagingGroup resources.
//
//    // Example sending a request using ListPackagingGroupsRequest.
//    req := client.ListPackagingGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListPackagingGroups
func (c *Client) ListPackagingGroupsRequest(input *types.ListPackagingGroupsInput) ListPackagingGroupsRequest {
	op := &aws.Operation{
		Name:       opListPackagingGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/packaging_groups",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListPackagingGroupsInput{}
	}

	req := c.newRequest(op, input, &types.ListPackagingGroupsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListPackagingGroupsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListPackagingGroupsRequest{Request: req, Input: input, Copy: c.ListPackagingGroupsRequest}
}

// ListPackagingGroupsRequest is the request type for the
// ListPackagingGroups API operation.
type ListPackagingGroupsRequest struct {
	*aws.Request
	Input *types.ListPackagingGroupsInput
	Copy  func(*types.ListPackagingGroupsInput) ListPackagingGroupsRequest
}

// Send marshals and sends the ListPackagingGroups API request.
func (r ListPackagingGroupsRequest) Send(ctx context.Context) (*ListPackagingGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPackagingGroupsResponse{
		ListPackagingGroupsOutput: r.Request.Data.(*types.ListPackagingGroupsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPackagingGroupsRequestPaginator returns a paginator for ListPackagingGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPackagingGroupsRequest(input)
//   p := mediapackagevod.NewListPackagingGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPackagingGroupsPaginator(req ListPackagingGroupsRequest) ListPackagingGroupsPaginator {
	return ListPackagingGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPackagingGroupsInput
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

// ListPackagingGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPackagingGroupsPaginator struct {
	aws.Pager
}

func (p *ListPackagingGroupsPaginator) CurrentPage() *types.ListPackagingGroupsOutput {
	return p.Pager.CurrentPage().(*types.ListPackagingGroupsOutput)
}

// ListPackagingGroupsResponse is the response type for the
// ListPackagingGroups API operation.
type ListPackagingGroupsResponse struct {
	*types.ListPackagingGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPackagingGroups request.
func (r *ListPackagingGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
