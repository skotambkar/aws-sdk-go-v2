// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

const opListCloudFrontOriginAccessIdentities = "ListCloudFrontOriginAccessIdentities2019_03_26"

// ListCloudFrontOriginAccessIdentitiesRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Lists origin access identities.
//
//    // Example sending a request using ListCloudFrontOriginAccessIdentitiesRequest.
//    req := client.ListCloudFrontOriginAccessIdentitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListCloudFrontOriginAccessIdentities
func (c *Client) ListCloudFrontOriginAccessIdentitiesRequest(input *types.ListCloudFrontOriginAccessIdentitiesInput) ListCloudFrontOriginAccessIdentitiesRequest {
	op := &aws.Operation{
		Name:       opListCloudFrontOriginAccessIdentities,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/origin-access-identity/cloudfront",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"CloudFrontOriginAccessIdentityList.NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "CloudFrontOriginAccessIdentityList.IsTruncated",
		},
	}

	if input == nil {
		input = &types.ListCloudFrontOriginAccessIdentitiesInput{}
	}

	req := c.newRequest(op, input, &types.ListCloudFrontOriginAccessIdentitiesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.ListCloudFrontOriginAccessIdentitiesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListCloudFrontOriginAccessIdentitiesRequest{Request: req, Input: input, Copy: c.ListCloudFrontOriginAccessIdentitiesRequest}
}

// ListCloudFrontOriginAccessIdentitiesRequest is the request type for the
// ListCloudFrontOriginAccessIdentities API operation.
type ListCloudFrontOriginAccessIdentitiesRequest struct {
	*aws.Request
	Input *types.ListCloudFrontOriginAccessIdentitiesInput
	Copy  func(*types.ListCloudFrontOriginAccessIdentitiesInput) ListCloudFrontOriginAccessIdentitiesRequest
}

// Send marshals and sends the ListCloudFrontOriginAccessIdentities API request.
func (r ListCloudFrontOriginAccessIdentitiesRequest) Send(ctx context.Context) (*ListCloudFrontOriginAccessIdentitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCloudFrontOriginAccessIdentitiesResponse{
		ListCloudFrontOriginAccessIdentitiesOutput: r.Request.Data.(*types.ListCloudFrontOriginAccessIdentitiesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCloudFrontOriginAccessIdentitiesRequestPaginator returns a paginator for ListCloudFrontOriginAccessIdentities.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCloudFrontOriginAccessIdentitiesRequest(input)
//   p := cloudfront.NewListCloudFrontOriginAccessIdentitiesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCloudFrontOriginAccessIdentitiesPaginator(req ListCloudFrontOriginAccessIdentitiesRequest) ListCloudFrontOriginAccessIdentitiesPaginator {
	return ListCloudFrontOriginAccessIdentitiesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListCloudFrontOriginAccessIdentitiesInput
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

// ListCloudFrontOriginAccessIdentitiesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCloudFrontOriginAccessIdentitiesPaginator struct {
	aws.Pager
}

func (p *ListCloudFrontOriginAccessIdentitiesPaginator) CurrentPage() *types.ListCloudFrontOriginAccessIdentitiesOutput {
	return p.Pager.CurrentPage().(*types.ListCloudFrontOriginAccessIdentitiesOutput)
}

// ListCloudFrontOriginAccessIdentitiesResponse is the response type for the
// ListCloudFrontOriginAccessIdentities API operation.
type ListCloudFrontOriginAccessIdentitiesResponse struct {
	*types.ListCloudFrontOriginAccessIdentitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCloudFrontOriginAccessIdentities request.
func (r *ListCloudFrontOriginAccessIdentitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
