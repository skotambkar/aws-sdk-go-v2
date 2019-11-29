// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

const opCreateCacheCluster = "CreateCacheCluster"

// CreateCacheClusterRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Creates a cluster. All nodes in the cluster run the same protocol-compliant
// cache engine software, either Memcached or Redis.
//
// This operation is not supported for Redis (cluster mode enabled) clusters.
//
//    // Example sending a request using CreateCacheClusterRequest.
//    req := client.CreateCacheClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateCacheCluster
func (c *Client) CreateCacheClusterRequest(input *types.CreateCacheClusterInput) CreateCacheClusterRequest {
	op := &aws.Operation{
		Name:       opCreateCacheCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateCacheClusterInput{}
	}

	req := c.newRequest(op, input, &types.CreateCacheClusterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateCacheClusterMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateCacheClusterRequest{Request: req, Input: input, Copy: c.CreateCacheClusterRequest}
}

// CreateCacheClusterRequest is the request type for the
// CreateCacheCluster API operation.
type CreateCacheClusterRequest struct {
	*aws.Request
	Input *types.CreateCacheClusterInput
	Copy  func(*types.CreateCacheClusterInput) CreateCacheClusterRequest
}

// Send marshals and sends the CreateCacheCluster API request.
func (r CreateCacheClusterRequest) Send(ctx context.Context) (*CreateCacheClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCacheClusterResponse{
		CreateCacheClusterOutput: r.Request.Data.(*types.CreateCacheClusterOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCacheClusterResponse is the response type for the
// CreateCacheCluster API operation.
type CreateCacheClusterResponse struct {
	*types.CreateCacheClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCacheCluster request.
func (r *CreateCacheClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
