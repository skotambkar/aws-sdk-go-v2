// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/rds/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opRestoreDBInstanceFromDBSnapshot = "RestoreDBInstanceFromDBSnapshot"

// RestoreDBInstanceFromDBSnapshotRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new DB instance from a DB snapshot. The target database is created
// from the source database restore point with the most of original configuration
// with the default security group and the default DB parameter group. By default,
// the new DB instance is created as a single-AZ deployment except when the
// instance is a SQL Server instance that has an option group that is associated
// with mirroring; in this case, the instance becomes a mirrored AZ deployment
// and not a single-AZ deployment.
//
// If your intent is to replace your original DB instance with the new, restored
// DB instance, then rename your original DB instance before you call the RestoreDBInstanceFromDBSnapshot
// action. RDS doesn't allow two DB instances with the same name. Once you have
// renamed your original DB instance with a different identifier, then you can
// pass the original name of the DB instance as the DBInstanceIdentifier in
// the call to the RestoreDBInstanceFromDBSnapshot action. The result is that
// you will replace the original DB instance with the DB instance created from
// the snapshot.
//
// If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier
// must be the ARN of the shared DB snapshot.
//
// This command doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterFromSnapshot.
//
//    // Example sending a request using RestoreDBInstanceFromDBSnapshotRequest.
//    req := client.RestoreDBInstanceFromDBSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBInstanceFromDBSnapshot
func (c *Client) RestoreDBInstanceFromDBSnapshotRequest(input *types.RestoreDBInstanceFromDBSnapshotInput) RestoreDBInstanceFromDBSnapshotRequest {
	op := &aws.Operation{
		Name:       opRestoreDBInstanceFromDBSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RestoreDBInstanceFromDBSnapshotInput{}
	}

	req := c.newRequest(op, input, &types.RestoreDBInstanceFromDBSnapshotOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.RestoreDBInstanceFromDBSnapshotMarshaler{Input: input}.GetNamedBuildHandler())

	return RestoreDBInstanceFromDBSnapshotRequest{Request: req, Input: input, Copy: c.RestoreDBInstanceFromDBSnapshotRequest}
}

// RestoreDBInstanceFromDBSnapshotRequest is the request type for the
// RestoreDBInstanceFromDBSnapshot API operation.
type RestoreDBInstanceFromDBSnapshotRequest struct {
	*aws.Request
	Input *types.RestoreDBInstanceFromDBSnapshotInput
	Copy  func(*types.RestoreDBInstanceFromDBSnapshotInput) RestoreDBInstanceFromDBSnapshotRequest
}

// Send marshals and sends the RestoreDBInstanceFromDBSnapshot API request.
func (r RestoreDBInstanceFromDBSnapshotRequest) Send(ctx context.Context) (*RestoreDBInstanceFromDBSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBInstanceFromDBSnapshotResponse{
		RestoreDBInstanceFromDBSnapshotOutput: r.Request.Data.(*types.RestoreDBInstanceFromDBSnapshotOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBInstanceFromDBSnapshotResponse is the response type for the
// RestoreDBInstanceFromDBSnapshot API operation.
type RestoreDBInstanceFromDBSnapshotResponse struct {
	*types.RestoreDBInstanceFromDBSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBInstanceFromDBSnapshot request.
func (r *RestoreDBInstanceFromDBSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
