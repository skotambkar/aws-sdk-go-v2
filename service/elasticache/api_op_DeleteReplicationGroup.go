// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

const opDeleteReplicationGroup = "DeleteReplicationGroup"

// DeleteReplicationGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Deletes an existing replication group. By default, this operation deletes
// the entire replication group, including the primary/primaries and all of
// the read replicas. If the replication group has only one primary, you can
// optionally delete only the read replicas, while retaining the primary by
// setting RetainPrimaryCluster=true.
//
// When you receive a successful response from this operation, Amazon ElastiCache
// immediately begins deleting the selected resources; you cannot cancel or
// revert this operation.
//
// This operation is valid for Redis only.
//
//    // Example sending a request using DeleteReplicationGroupRequest.
//    req := client.DeleteReplicationGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteReplicationGroup
func (c *Client) DeleteReplicationGroupRequest(input *types.DeleteReplicationGroupInput) DeleteReplicationGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteReplicationGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteReplicationGroupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteReplicationGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteReplicationGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteReplicationGroupRequest{Request: req, Input: input, Copy: c.DeleteReplicationGroupRequest}
}

// DeleteReplicationGroupRequest is the request type for the
// DeleteReplicationGroup API operation.
type DeleteReplicationGroupRequest struct {
	*aws.Request
	Input *types.DeleteReplicationGroupInput
	Copy  func(*types.DeleteReplicationGroupInput) DeleteReplicationGroupRequest
}

// Send marshals and sends the DeleteReplicationGroup API request.
func (r DeleteReplicationGroupRequest) Send(ctx context.Context) (*DeleteReplicationGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteReplicationGroupResponse{
		DeleteReplicationGroupOutput: r.Request.Data.(*types.DeleteReplicationGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteReplicationGroupResponse is the response type for the
// DeleteReplicationGroup API operation.
type DeleteReplicationGroupResponse struct {
	*types.DeleteReplicationGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteReplicationGroup request.
func (r *DeleteReplicationGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
