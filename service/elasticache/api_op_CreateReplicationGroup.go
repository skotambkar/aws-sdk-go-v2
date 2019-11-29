// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

const opCreateReplicationGroup = "CreateReplicationGroup"

// CreateReplicationGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Creates a Redis (cluster mode disabled) or a Redis (cluster mode enabled)
// replication group.
//
// A Redis (cluster mode disabled) replication group is a collection of clusters,
// where one of the clusters is a read/write primary and the others are read-only
// replicas. Writes to the primary are asynchronously propagated to the replicas.
//
// A Redis (cluster mode enabled) replication group is a collection of 1 to
// 90 node groups (shards). Each node group (shard) has one read/write primary
// node and up to 5 read-only replica nodes. Writes to the primary are asynchronously
// propagated to the replicas. Redis (cluster mode enabled) replication groups
// partition the data across node groups (shards).
//
// When a Redis (cluster mode disabled) replication group has been successfully
// created, you can add one or more read replicas to it, up to a total of 5
// read replicas. You cannot alter a Redis (cluster mode enabled) replication
// group after it has been created. However, if you need to increase or decrease
// the number of node groups (console: shards), you can avail yourself of ElastiCache
// for Redis' enhanced backup and restore. For more information, see Restoring
// From a Backup with Cluster Resizing (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-restoring.html)
// in the ElastiCache User Guide.
//
// This operation is valid for Redis only.
//
//    // Example sending a request using CreateReplicationGroupRequest.
//    req := client.CreateReplicationGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateReplicationGroup
func (c *Client) CreateReplicationGroupRequest(input *types.CreateReplicationGroupInput) CreateReplicationGroupRequest {
	op := &aws.Operation{
		Name:       opCreateReplicationGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateReplicationGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateReplicationGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateReplicationGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateReplicationGroupRequest{Request: req, Input: input, Copy: c.CreateReplicationGroupRequest}
}

// CreateReplicationGroupRequest is the request type for the
// CreateReplicationGroup API operation.
type CreateReplicationGroupRequest struct {
	*aws.Request
	Input *types.CreateReplicationGroupInput
	Copy  func(*types.CreateReplicationGroupInput) CreateReplicationGroupRequest
}

// Send marshals and sends the CreateReplicationGroup API request.
func (r CreateReplicationGroupRequest) Send(ctx context.Context) (*CreateReplicationGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReplicationGroupResponse{
		CreateReplicationGroupOutput: r.Request.Data.(*types.CreateReplicationGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReplicationGroupResponse is the response type for the
// CreateReplicationGroup API operation.
type CreateReplicationGroupResponse struct {
	*types.CreateReplicationGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReplicationGroup request.
func (r *CreateReplicationGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
