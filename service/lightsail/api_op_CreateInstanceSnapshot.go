// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opCreateInstanceSnapshot = "CreateInstanceSnapshot"

// CreateInstanceSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Creates a snapshot of a specific virtual private server, or instance. You
// can use a snapshot to create a new instance that is based on that snapshot.
//
// The create instance snapshot operation supports tag-based access control
// via request tags. For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using CreateInstanceSnapshotRequest.
//    req := client.CreateInstanceSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateInstanceSnapshot
func (c *Client) CreateInstanceSnapshotRequest(input *types.CreateInstanceSnapshotInput) CreateInstanceSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateInstanceSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateInstanceSnapshotInput{}
	}

	req := c.newRequest(op, input, &types.CreateInstanceSnapshotOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateInstanceSnapshotMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateInstanceSnapshotRequest{Request: req, Input: input, Copy: c.CreateInstanceSnapshotRequest}
}

// CreateInstanceSnapshotRequest is the request type for the
// CreateInstanceSnapshot API operation.
type CreateInstanceSnapshotRequest struct {
	*aws.Request
	Input *types.CreateInstanceSnapshotInput
	Copy  func(*types.CreateInstanceSnapshotInput) CreateInstanceSnapshotRequest
}

// Send marshals and sends the CreateInstanceSnapshot API request.
func (r CreateInstanceSnapshotRequest) Send(ctx context.Context) (*CreateInstanceSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInstanceSnapshotResponse{
		CreateInstanceSnapshotOutput: r.Request.Data.(*types.CreateInstanceSnapshotOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInstanceSnapshotResponse is the response type for the
// CreateInstanceSnapshot API operation.
type CreateInstanceSnapshotResponse struct {
	*types.CreateInstanceSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInstanceSnapshot request.
func (r *CreateInstanceSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
