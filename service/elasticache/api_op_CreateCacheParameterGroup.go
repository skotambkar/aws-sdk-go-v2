// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

const opCreateCacheParameterGroup = "CreateCacheParameterGroup"

// CreateCacheParameterGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Creates a new Amazon ElastiCache cache parameter group. An ElastiCache cache
// parameter group is a collection of parameters and their values that are applied
// to all of the nodes in any cluster or replication group using the CacheParameterGroup.
//
// A newly created CacheParameterGroup is an exact duplicate of the default
// parameter group for the CacheParameterGroupFamily. To customize the newly
// created CacheParameterGroup you can change the values of specific parameters.
// For more information, see:
//
//    * ModifyCacheParameterGroup (https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html)
//    in the ElastiCache API Reference.
//
//    * Parameters and Parameter Groups (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/ParameterGroups.html)
//    in the ElastiCache User Guide.
//
//    // Example sending a request using CreateCacheParameterGroupRequest.
//    req := client.CreateCacheParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateCacheParameterGroup
func (c *Client) CreateCacheParameterGroupRequest(input *types.CreateCacheParameterGroupInput) CreateCacheParameterGroupRequest {
	op := &aws.Operation{
		Name:       opCreateCacheParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateCacheParameterGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateCacheParameterGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateCacheParameterGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateCacheParameterGroupRequest{Request: req, Input: input, Copy: c.CreateCacheParameterGroupRequest}
}

// CreateCacheParameterGroupRequest is the request type for the
// CreateCacheParameterGroup API operation.
type CreateCacheParameterGroupRequest struct {
	*aws.Request
	Input *types.CreateCacheParameterGroupInput
	Copy  func(*types.CreateCacheParameterGroupInput) CreateCacheParameterGroupRequest
}

// Send marshals and sends the CreateCacheParameterGroup API request.
func (r CreateCacheParameterGroupRequest) Send(ctx context.Context) (*CreateCacheParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCacheParameterGroupResponse{
		CreateCacheParameterGroupOutput: r.Request.Data.(*types.CreateCacheParameterGroupOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCacheParameterGroupResponse is the response type for the
// CreateCacheParameterGroup API operation.
type CreateCacheParameterGroupResponse struct {
	*types.CreateCacheParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCacheParameterGroup request.
func (r *CreateCacheParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
