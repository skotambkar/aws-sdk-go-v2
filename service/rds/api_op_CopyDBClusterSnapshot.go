// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opCopyDBClusterSnapshot = "CopyDBClusterSnapshot"

// CopyDBClusterSnapshotRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Copies a snapshot of a DB cluster.
//
// To copy a DB cluster snapshot from a shared manual DB cluster snapshot, SourceDBClusterSnapshotIdentifier
// must be the Amazon Resource Name (ARN) of the shared DB cluster snapshot.
//
// You can copy an encrypted DB cluster snapshot from another AWS Region. In
// that case, the AWS Region where you call the CopyDBClusterSnapshot action
// is the destination AWS Region for the encrypted DB cluster snapshot to be
// copied to. To copy an encrypted DB cluster snapshot from another AWS Region,
// you must provide the following values:
//
//    * KmsKeyId - The AWS Key Management System (AWS KMS) key identifier for
//    the key to use to encrypt the copy of the DB cluster snapshot in the destination
//    AWS Region.
//
//    * PreSignedUrl - A URL that contains a Signature Version 4 signed request
//    for the CopyDBClusterSnapshot action to be called in the source AWS Region
//    where the DB cluster snapshot is copied from. The pre-signed URL must
//    be a valid request for the CopyDBClusterSnapshot API action that can be
//    executed in the source AWS Region that contains the encrypted DB cluster
//    snapshot to be copied. The pre-signed URL request must contain the following
//    parameter values: KmsKeyId - The KMS key identifier for the key to use
//    to encrypt the copy of the DB cluster snapshot in the destination AWS
//    Region. This is the same identifier for both the CopyDBClusterSnapshot
//    action that is called in the destination AWS Region, and the action contained
//    in the pre-signed URL. DestinationRegion - The name of the AWS Region
//    that the DB cluster snapshot will be created in. SourceDBClusterSnapshotIdentifier
//    - The DB cluster snapshot identifier for the encrypted DB cluster snapshot
//    to be copied. This identifier must be in the Amazon Resource Name (ARN)
//    format for the source AWS Region. For example, if you are copying an encrypted
//    DB cluster snapshot from the us-west-2 AWS Region, then your SourceDBClusterSnapshotIdentifier
//    looks like the following example: arn:aws:rds:us-west-2:123456789012:cluster-snapshot:aurora-cluster1-snapshot-20161115.
//    To learn how to generate a Signature Version 4 signed request, see Authenticating
//    Requests: Using Query Parameters (AWS Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
//    and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
//    * TargetDBClusterSnapshotIdentifier - The identifier for the new copy
//    of the DB cluster snapshot in the destination AWS Region.
//
//    * SourceDBClusterSnapshotIdentifier - The DB cluster snapshot identifier
//    for the encrypted DB cluster snapshot to be copied. This identifier must
//    be in the ARN format for the source AWS Region and is the same value as
//    the SourceDBClusterSnapshotIdentifier in the pre-signed URL.
//
// To cancel the copy operation once it is in progress, delete the target DB
// cluster snapshot identified by TargetDBClusterSnapshotIdentifier while that
// DB cluster snapshot is in "copying" status.
//
// For more information on copying encrypted DB cluster snapshots from one AWS
// Region to another, see Copying a Snapshot (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CopySnapshot.html)
// in the Amazon Aurora User Guide.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using CopyDBClusterSnapshotRequest.
//    req := client.CopyDBClusterSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CopyDBClusterSnapshot
func (c *Client) CopyDBClusterSnapshotRequest(input *types.CopyDBClusterSnapshotInput) CopyDBClusterSnapshotRequest {
	op := &aws.Operation{
		Name:       opCopyDBClusterSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CopyDBClusterSnapshotInput{}
	}

	req := c.newRequest(op, input, &types.CopyDBClusterSnapshotOutput{})
	return CopyDBClusterSnapshotRequest{Request: req, Input: input, Copy: c.CopyDBClusterSnapshotRequest}
}

// CopyDBClusterSnapshotRequest is the request type for the
// CopyDBClusterSnapshot API operation.
type CopyDBClusterSnapshotRequest struct {
	*aws.Request
	Input *types.CopyDBClusterSnapshotInput
	Copy  func(*types.CopyDBClusterSnapshotInput) CopyDBClusterSnapshotRequest
}

// Send marshals and sends the CopyDBClusterSnapshot API request.
func (r CopyDBClusterSnapshotRequest) Send(ctx context.Context) (*CopyDBClusterSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyDBClusterSnapshotResponse{
		CopyDBClusterSnapshotOutput: r.Request.Data.(*types.CopyDBClusterSnapshotOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyDBClusterSnapshotResponse is the response type for the
// CopyDBClusterSnapshot API operation.
type CopyDBClusterSnapshotResponse struct {
	*types.CopyDBClusterSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyDBClusterSnapshot request.
func (r *CopyDBClusterSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
