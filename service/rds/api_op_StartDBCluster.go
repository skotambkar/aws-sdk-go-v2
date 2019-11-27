// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opStartDBCluster = "StartDBCluster"

// StartDBClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Starts an Amazon Aurora DB cluster that was stopped using the AWS console,
// the stop-db-cluster AWS CLI command, or the StopDBCluster action.
//
// For more information, see Stopping and Starting an Aurora Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-cluster-stop-start.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using StartDBClusterRequest.
//    req := client.StartDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/StartDBCluster
func (c *Client) StartDBClusterRequest(input *types.StartDBClusterInput) StartDBClusterRequest {
	op := &aws.Operation{
		Name:       opStartDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartDBClusterInput{}
	}

	req := c.newRequest(op, input, &types.StartDBClusterOutput{})
	return StartDBClusterRequest{Request: req, Input: input, Copy: c.StartDBClusterRequest}
}

// StartDBClusterRequest is the request type for the
// StartDBCluster API operation.
type StartDBClusterRequest struct {
	*aws.Request
	Input *types.StartDBClusterInput
	Copy  func(*types.StartDBClusterInput) StartDBClusterRequest
}

// Send marshals and sends the StartDBCluster API request.
func (r StartDBClusterRequest) Send(ctx context.Context) (*StartDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDBClusterResponse{
		StartDBClusterOutput: r.Request.Data.(*types.StartDBClusterOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDBClusterResponse is the response type for the
// StartDBCluster API operation.
type StartDBClusterResponse struct {
	*types.StartDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDBCluster request.
func (r *StartDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
