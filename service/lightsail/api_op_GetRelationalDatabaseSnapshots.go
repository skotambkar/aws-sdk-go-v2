// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetRelationalDatabaseSnapshots = "GetRelationalDatabaseSnapshots"

// GetRelationalDatabaseSnapshotsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about all of your database snapshots in Amazon Lightsail.
//
//    // Example sending a request using GetRelationalDatabaseSnapshotsRequest.
//    req := client.GetRelationalDatabaseSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseSnapshots
func (c *Client) GetRelationalDatabaseSnapshotsRequest(input *types.GetRelationalDatabaseSnapshotsInput) GetRelationalDatabaseSnapshotsRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabaseSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetRelationalDatabaseSnapshotsInput{}
	}

	req := c.newRequest(op, input, &types.GetRelationalDatabaseSnapshotsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetRelationalDatabaseSnapshotsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRelationalDatabaseSnapshotsRequest{Request: req, Input: input, Copy: c.GetRelationalDatabaseSnapshotsRequest}
}

// GetRelationalDatabaseSnapshotsRequest is the request type for the
// GetRelationalDatabaseSnapshots API operation.
type GetRelationalDatabaseSnapshotsRequest struct {
	*aws.Request
	Input *types.GetRelationalDatabaseSnapshotsInput
	Copy  func(*types.GetRelationalDatabaseSnapshotsInput) GetRelationalDatabaseSnapshotsRequest
}

// Send marshals and sends the GetRelationalDatabaseSnapshots API request.
func (r GetRelationalDatabaseSnapshotsRequest) Send(ctx context.Context) (*GetRelationalDatabaseSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabaseSnapshotsResponse{
		GetRelationalDatabaseSnapshotsOutput: r.Request.Data.(*types.GetRelationalDatabaseSnapshotsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabaseSnapshotsResponse is the response type for the
// GetRelationalDatabaseSnapshots API operation.
type GetRelationalDatabaseSnapshotsResponse struct {
	*types.GetRelationalDatabaseSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabaseSnapshots request.
func (r *GetRelationalDatabaseSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
