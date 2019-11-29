// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/docdb/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
)

const opRestoreDBClusterFromSnapshot = "RestoreDBClusterFromSnapshot"

// RestoreDBClusterFromSnapshotRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
//
// If a DB snapshot is specified, the target DB cluster is created from the
// source DB snapshot with a default configuration and default security group.
//
// If a DB cluster snapshot is specified, the target DB cluster is created from
// the source DB cluster restore point with the same configuration as the original
// source DB cluster, except that the new DB cluster is created with the default
// security group.
//
//    // Example sending a request using RestoreDBClusterFromSnapshotRequest.
//    req := client.RestoreDBClusterFromSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/RestoreDBClusterFromSnapshot
func (c *Client) RestoreDBClusterFromSnapshotRequest(input *types.RestoreDBClusterFromSnapshotInput) RestoreDBClusterFromSnapshotRequest {
	op := &aws.Operation{
		Name:       opRestoreDBClusterFromSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RestoreDBClusterFromSnapshotInput{}
	}

	req := c.newRequest(op, input, &types.RestoreDBClusterFromSnapshotOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.RestoreDBClusterFromSnapshotMarshaler{Input: input}.GetNamedBuildHandler())

	return RestoreDBClusterFromSnapshotRequest{Request: req, Input: input, Copy: c.RestoreDBClusterFromSnapshotRequest}
}

// RestoreDBClusterFromSnapshotRequest is the request type for the
// RestoreDBClusterFromSnapshot API operation.
type RestoreDBClusterFromSnapshotRequest struct {
	*aws.Request
	Input *types.RestoreDBClusterFromSnapshotInput
	Copy  func(*types.RestoreDBClusterFromSnapshotInput) RestoreDBClusterFromSnapshotRequest
}

// Send marshals and sends the RestoreDBClusterFromSnapshot API request.
func (r RestoreDBClusterFromSnapshotRequest) Send(ctx context.Context) (*RestoreDBClusterFromSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBClusterFromSnapshotResponse{
		RestoreDBClusterFromSnapshotOutput: r.Request.Data.(*types.RestoreDBClusterFromSnapshotOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBClusterFromSnapshotResponse is the response type for the
// RestoreDBClusterFromSnapshot API operation.
type RestoreDBClusterFromSnapshotResponse struct {
	*types.RestoreDBClusterFromSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBClusterFromSnapshot request.
func (r *RestoreDBClusterFromSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
