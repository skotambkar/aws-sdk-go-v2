// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
)

const opPromoteReadReplicaDBCluster = "PromoteReadReplicaDBCluster"

// PromoteReadReplicaDBClusterRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Not supported.
//
//    // Example sending a request using PromoteReadReplicaDBClusterRequest.
//    req := client.PromoteReadReplicaDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/PromoteReadReplicaDBCluster
func (c *Client) PromoteReadReplicaDBClusterRequest(input *types.PromoteReadReplicaDBClusterInput) PromoteReadReplicaDBClusterRequest {
	op := &aws.Operation{
		Name:       opPromoteReadReplicaDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PromoteReadReplicaDBClusterInput{}
	}

	req := c.newRequest(op, input, &types.PromoteReadReplicaDBClusterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.PromoteReadReplicaDBClusterMarshaler{Input: input}.GetNamedBuildHandler())

	return PromoteReadReplicaDBClusterRequest{Request: req, Input: input, Copy: c.PromoteReadReplicaDBClusterRequest}
}

// PromoteReadReplicaDBClusterRequest is the request type for the
// PromoteReadReplicaDBCluster API operation.
type PromoteReadReplicaDBClusterRequest struct {
	*aws.Request
	Input *types.PromoteReadReplicaDBClusterInput
	Copy  func(*types.PromoteReadReplicaDBClusterInput) PromoteReadReplicaDBClusterRequest
}

// Send marshals and sends the PromoteReadReplicaDBCluster API request.
func (r PromoteReadReplicaDBClusterRequest) Send(ctx context.Context) (*PromoteReadReplicaDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PromoteReadReplicaDBClusterResponse{
		PromoteReadReplicaDBClusterOutput: r.Request.Data.(*types.PromoteReadReplicaDBClusterOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PromoteReadReplicaDBClusterResponse is the response type for the
// PromoteReadReplicaDBCluster API operation.
type PromoteReadReplicaDBClusterResponse struct {
	*types.PromoteReadReplicaDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PromoteReadReplicaDBCluster request.
func (r *PromoteReadReplicaDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
