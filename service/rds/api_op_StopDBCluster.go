// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/rds/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opStopDBCluster = "StopDBCluster"

// StopDBClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Stops an Amazon Aurora DB cluster. When you stop a DB cluster, Aurora retains
// the DB cluster's metadata, including its endpoints and DB parameter groups.
// Aurora also retains the transaction logs so you can do a point-in-time restore
// if necessary.
//
// For more information, see Stopping and Starting an Aurora Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-cluster-stop-start.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using StopDBClusterRequest.
//    req := client.StopDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/StopDBCluster
func (c *Client) StopDBClusterRequest(input *types.StopDBClusterInput) StopDBClusterRequest {
	op := &aws.Operation{
		Name:       opStopDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopDBClusterInput{}
	}

	req := c.newRequest(op, input, &types.StopDBClusterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.StopDBClusterMarshaler{Input: input}.GetNamedBuildHandler())

	return StopDBClusterRequest{Request: req, Input: input, Copy: c.StopDBClusterRequest}
}

// StopDBClusterRequest is the request type for the
// StopDBCluster API operation.
type StopDBClusterRequest struct {
	*aws.Request
	Input *types.StopDBClusterInput
	Copy  func(*types.StopDBClusterInput) StopDBClusterRequest
}

// Send marshals and sends the StopDBCluster API request.
func (r StopDBClusterRequest) Send(ctx context.Context) (*StopDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopDBClusterResponse{
		StopDBClusterOutput: r.Request.Data.(*types.StopDBClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopDBClusterResponse is the response type for the
// StopDBCluster API operation.
type StopDBClusterResponse struct {
	*types.StopDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopDBCluster request.
func (r *StopDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
