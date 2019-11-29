// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/emr/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
)

const opListClusters = "ListClusters"

// ListClustersRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Provides the status of all clusters visible to this AWS account. Allows you
// to filter the list of clusters based on certain criteria; for example, filtering
// by cluster creation date and time or by status. This call returns a maximum
// of 50 clusters per call, but returns a marker to track the paging of the
// cluster list across multiple ListClusters calls.
//
//    // Example sending a request using ListClustersRequest.
//    req := client.ListClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListClusters
func (c *Client) ListClustersRequest(input *types.ListClustersInput) ListClustersRequest {
	op := &aws.Operation{
		Name:       opListClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListClustersInput{}
	}

	req := c.newRequest(op, input, &types.ListClustersOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListClustersMarshaler{Input: input}.GetNamedBuildHandler())

	return ListClustersRequest{Request: req, Input: input, Copy: c.ListClustersRequest}
}

// ListClustersRequest is the request type for the
// ListClusters API operation.
type ListClustersRequest struct {
	*aws.Request
	Input *types.ListClustersInput
	Copy  func(*types.ListClustersInput) ListClustersRequest
}

// Send marshals and sends the ListClusters API request.
func (r ListClustersRequest) Send(ctx context.Context) (*ListClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListClustersResponse{
		ListClustersOutput: r.Request.Data.(*types.ListClustersOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListClustersRequestPaginator returns a paginator for ListClusters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListClustersRequest(input)
//   p := emr.NewListClustersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListClustersPaginator(req ListClustersRequest) ListClustersPaginator {
	return ListClustersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListClustersInput
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

// ListClustersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListClustersPaginator struct {
	aws.Pager
}

func (p *ListClustersPaginator) CurrentPage() *types.ListClustersOutput {
	return p.Pager.CurrentPage().(*types.ListClustersOutput)
}

// ListClustersResponse is the response type for the
// ListClusters API operation.
type ListClustersResponse struct {
	*types.ListClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListClusters request.
func (r *ListClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
