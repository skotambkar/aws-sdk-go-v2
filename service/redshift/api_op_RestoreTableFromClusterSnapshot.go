// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opRestoreTableFromClusterSnapshot = "RestoreTableFromClusterSnapshot"

// RestoreTableFromClusterSnapshotRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Creates a new table from a table in an Amazon Redshift cluster snapshot.
// You must create the new table within the Amazon Redshift cluster that the
// snapshot was taken from.
//
// You cannot use RestoreTableFromClusterSnapshot to restore a table with the
// same name as an existing table in an Amazon Redshift cluster. That is, you
// cannot overwrite an existing table in a cluster with a restored table. If
// you want to replace your original table with a new, restored table, then
// rename or drop your original table before you call RestoreTableFromClusterSnapshot.
// When you have renamed your original table, then you can pass the original
// name of the table as the NewTableName parameter value in the call to RestoreTableFromClusterSnapshot.
// This way, you can replace the original table with the table created from
// the snapshot.
//
//    // Example sending a request using RestoreTableFromClusterSnapshotRequest.
//    req := client.RestoreTableFromClusterSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/RestoreTableFromClusterSnapshot
func (c *Client) RestoreTableFromClusterSnapshotRequest(input *types.RestoreTableFromClusterSnapshotInput) RestoreTableFromClusterSnapshotRequest {
	op := &aws.Operation{
		Name:       opRestoreTableFromClusterSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RestoreTableFromClusterSnapshotInput{}
	}

	req := c.newRequest(op, input, &types.RestoreTableFromClusterSnapshotOutput{})
	return RestoreTableFromClusterSnapshotRequest{Request: req, Input: input, Copy: c.RestoreTableFromClusterSnapshotRequest}
}

// RestoreTableFromClusterSnapshotRequest is the request type for the
// RestoreTableFromClusterSnapshot API operation.
type RestoreTableFromClusterSnapshotRequest struct {
	*aws.Request
	Input *types.RestoreTableFromClusterSnapshotInput
	Copy  func(*types.RestoreTableFromClusterSnapshotInput) RestoreTableFromClusterSnapshotRequest
}

// Send marshals and sends the RestoreTableFromClusterSnapshot API request.
func (r RestoreTableFromClusterSnapshotRequest) Send(ctx context.Context) (*RestoreTableFromClusterSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreTableFromClusterSnapshotResponse{
		RestoreTableFromClusterSnapshotOutput: r.Request.Data.(*types.RestoreTableFromClusterSnapshotOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreTableFromClusterSnapshotResponse is the response type for the
// RestoreTableFromClusterSnapshot API operation.
type RestoreTableFromClusterSnapshotResponse struct {
	*types.RestoreTableFromClusterSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreTableFromClusterSnapshot request.
func (r *RestoreTableFromClusterSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
