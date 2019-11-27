// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
)

const opListPackagingConfigurations = "ListPackagingConfigurations"

// ListPackagingConfigurationsRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Returns a collection of MediaPackage VOD PackagingConfiguration resources.
//
//    // Example sending a request using ListPackagingConfigurationsRequest.
//    req := client.ListPackagingConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListPackagingConfigurations
func (c *Client) ListPackagingConfigurationsRequest(input *types.ListPackagingConfigurationsInput) ListPackagingConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListPackagingConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/packaging_configurations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListPackagingConfigurationsInput{}
	}

	req := c.newRequest(op, input, &types.ListPackagingConfigurationsOutput{})
	return ListPackagingConfigurationsRequest{Request: req, Input: input, Copy: c.ListPackagingConfigurationsRequest}
}

// ListPackagingConfigurationsRequest is the request type for the
// ListPackagingConfigurations API operation.
type ListPackagingConfigurationsRequest struct {
	*aws.Request
	Input *types.ListPackagingConfigurationsInput
	Copy  func(*types.ListPackagingConfigurationsInput) ListPackagingConfigurationsRequest
}

// Send marshals and sends the ListPackagingConfigurations API request.
func (r ListPackagingConfigurationsRequest) Send(ctx context.Context) (*ListPackagingConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPackagingConfigurationsResponse{
		ListPackagingConfigurationsOutput: r.Request.Data.(*types.ListPackagingConfigurationsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPackagingConfigurationsRequestPaginator returns a paginator for ListPackagingConfigurations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPackagingConfigurationsRequest(input)
//   p := mediapackagevod.NewListPackagingConfigurationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPackagingConfigurationsPaginator(req ListPackagingConfigurationsRequest) ListPackagingConfigurationsPaginator {
	return ListPackagingConfigurationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPackagingConfigurationsInput
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

// ListPackagingConfigurationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPackagingConfigurationsPaginator struct {
	aws.Pager
}

func (p *ListPackagingConfigurationsPaginator) CurrentPage() *types.ListPackagingConfigurationsOutput {
	return p.Pager.CurrentPage().(*types.ListPackagingConfigurationsOutput)
}

// ListPackagingConfigurationsResponse is the response type for the
// ListPackagingConfigurations API operation.
type ListPackagingConfigurationsResponse struct {
	*types.ListPackagingConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPackagingConfigurations request.
func (r *ListPackagingConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
