// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opRevokeSnapshotAccess = "RevokeSnapshotAccess"

// RevokeSnapshotAccessRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Removes the ability of the specified AWS customer account to restore the
// specified snapshot. If the account is currently restoring the snapshot, the
// restore will run to completion.
//
// For more information about working with snapshots, go to Amazon Redshift
// Snapshots (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using RevokeSnapshotAccessRequest.
//    req := client.RevokeSnapshotAccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/RevokeSnapshotAccess
func (c *Client) RevokeSnapshotAccessRequest(input *types.RevokeSnapshotAccessInput) RevokeSnapshotAccessRequest {
	op := &aws.Operation{
		Name:       opRevokeSnapshotAccess,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RevokeSnapshotAccessInput{}
	}

	req := c.newRequest(op, input, &types.RevokeSnapshotAccessOutput{})
	return RevokeSnapshotAccessRequest{Request: req, Input: input, Copy: c.RevokeSnapshotAccessRequest}
}

// RevokeSnapshotAccessRequest is the request type for the
// RevokeSnapshotAccess API operation.
type RevokeSnapshotAccessRequest struct {
	*aws.Request
	Input *types.RevokeSnapshotAccessInput
	Copy  func(*types.RevokeSnapshotAccessInput) RevokeSnapshotAccessRequest
}

// Send marshals and sends the RevokeSnapshotAccess API request.
func (r RevokeSnapshotAccessRequest) Send(ctx context.Context) (*RevokeSnapshotAccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokeSnapshotAccessResponse{
		RevokeSnapshotAccessOutput: r.Request.Data.(*types.RevokeSnapshotAccessOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokeSnapshotAccessResponse is the response type for the
// RevokeSnapshotAccess API operation.
type RevokeSnapshotAccessResponse struct {
	*types.RevokeSnapshotAccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokeSnapshotAccess request.
func (r *RevokeSnapshotAccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
