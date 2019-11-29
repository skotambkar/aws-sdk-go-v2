// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

const opDeleteSnapshot = "DeleteSnapshot"

// DeleteSnapshotRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Deletes an existing snapshot. When you receive a successful response from
// this operation, ElastiCache immediately begins deleting the snapshot; you
// cannot cancel or revert this operation.
//
// This operation is valid for Redis only.
//
//    // Example sending a request using DeleteSnapshotRequest.
//    req := client.DeleteSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DeleteSnapshot
func (c *Client) DeleteSnapshotRequest(input *types.DeleteSnapshotInput) DeleteSnapshotRequest {
	op := &aws.Operation{
		Name:       opDeleteSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteSnapshotInput{}
	}

	req := c.newRequest(op, input, &types.DeleteSnapshotOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteSnapshotMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteSnapshotRequest{Request: req, Input: input, Copy: c.DeleteSnapshotRequest}
}

// DeleteSnapshotRequest is the request type for the
// DeleteSnapshot API operation.
type DeleteSnapshotRequest struct {
	*aws.Request
	Input *types.DeleteSnapshotInput
	Copy  func(*types.DeleteSnapshotInput) DeleteSnapshotRequest
}

// Send marshals and sends the DeleteSnapshot API request.
func (r DeleteSnapshotRequest) Send(ctx context.Context) (*DeleteSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSnapshotResponse{
		DeleteSnapshotOutput: r.Request.Data.(*types.DeleteSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSnapshotResponse is the response type for the
// DeleteSnapshot API operation.
type DeleteSnapshotResponse struct {
	*types.DeleteSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSnapshot request.
func (r *DeleteSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
