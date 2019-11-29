// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
)

const opModifyDBClusterSnapshotAttribute = "ModifyDBClusterSnapshotAttribute"

// ModifyDBClusterSnapshotAttributeRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Adds an attribute and values to, or removes an attribute and values from,
// a manual DB cluster snapshot.
//
// To share a manual DB cluster snapshot with other AWS accounts, specify restore
// as the AttributeName, and use the ValuesToAdd parameter to add a list of
// IDs of the AWS accounts that are authorized to restore the manual DB cluster
// snapshot. Use the value all to make the manual DB cluster snapshot public,
// which means that it can be copied or restored by all AWS accounts. Do not
// add the all value for any manual DB cluster snapshots that contain private
// information that you don't want available to all AWS accounts. If a manual
// DB cluster snapshot is encrypted, it can be shared, but only by specifying
// a list of authorized AWS account IDs for the ValuesToAdd parameter. You can't
// use all as a value for that parameter in this case.
//
//    // Example sending a request using ModifyDBClusterSnapshotAttributeRequest.
//    req := client.ModifyDBClusterSnapshotAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/ModifyDBClusterSnapshotAttribute
func (c *Client) ModifyDBClusterSnapshotAttributeRequest(input *types.ModifyDBClusterSnapshotAttributeInput) ModifyDBClusterSnapshotAttributeRequest {
	op := &aws.Operation{
		Name:       opModifyDBClusterSnapshotAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyDBClusterSnapshotAttributeInput{}
	}

	req := c.newRequest(op, input, &types.ModifyDBClusterSnapshotAttributeOutput{})
	return ModifyDBClusterSnapshotAttributeRequest{Request: req, Input: input, Copy: c.ModifyDBClusterSnapshotAttributeRequest}
}

// ModifyDBClusterSnapshotAttributeRequest is the request type for the
// ModifyDBClusterSnapshotAttribute API operation.
type ModifyDBClusterSnapshotAttributeRequest struct {
	*aws.Request
	Input *types.ModifyDBClusterSnapshotAttributeInput
	Copy  func(*types.ModifyDBClusterSnapshotAttributeInput) ModifyDBClusterSnapshotAttributeRequest
}

// Send marshals and sends the ModifyDBClusterSnapshotAttribute API request.
func (r ModifyDBClusterSnapshotAttributeRequest) Send(ctx context.Context) (*ModifyDBClusterSnapshotAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBClusterSnapshotAttributeResponse{
		ModifyDBClusterSnapshotAttributeOutput: r.Request.Data.(*types.ModifyDBClusterSnapshotAttributeOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBClusterSnapshotAttributeResponse is the response type for the
// ModifyDBClusterSnapshotAttribute API operation.
type ModifyDBClusterSnapshotAttributeResponse struct {
	*types.ModifyDBClusterSnapshotAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBClusterSnapshotAttribute request.
func (r *ModifyDBClusterSnapshotAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
