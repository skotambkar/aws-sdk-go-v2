// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetRelationalDatabaseSnapshot = "GetRelationalDatabaseSnapshot"

// GetRelationalDatabaseSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about a specific database snapshot in Amazon Lightsail.
//
//    // Example sending a request using GetRelationalDatabaseSnapshotRequest.
//    req := client.GetRelationalDatabaseSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseSnapshot
func (c *Client) GetRelationalDatabaseSnapshotRequest(input *types.GetRelationalDatabaseSnapshotInput) GetRelationalDatabaseSnapshotRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabaseSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetRelationalDatabaseSnapshotInput{}
	}

	req := c.newRequest(op, input, &types.GetRelationalDatabaseSnapshotOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetRelationalDatabaseSnapshotMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRelationalDatabaseSnapshotRequest{Request: req, Input: input, Copy: c.GetRelationalDatabaseSnapshotRequest}
}

// GetRelationalDatabaseSnapshotRequest is the request type for the
// GetRelationalDatabaseSnapshot API operation.
type GetRelationalDatabaseSnapshotRequest struct {
	*aws.Request
	Input *types.GetRelationalDatabaseSnapshotInput
	Copy  func(*types.GetRelationalDatabaseSnapshotInput) GetRelationalDatabaseSnapshotRequest
}

// Send marshals and sends the GetRelationalDatabaseSnapshot API request.
func (r GetRelationalDatabaseSnapshotRequest) Send(ctx context.Context) (*GetRelationalDatabaseSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabaseSnapshotResponse{
		GetRelationalDatabaseSnapshotOutput: r.Request.Data.(*types.GetRelationalDatabaseSnapshotOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabaseSnapshotResponse is the response type for the
// GetRelationalDatabaseSnapshot API operation.
type GetRelationalDatabaseSnapshotResponse struct {
	*types.GetRelationalDatabaseSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabaseSnapshot request.
func (r *GetRelationalDatabaseSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
