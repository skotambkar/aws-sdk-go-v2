// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opFailoverDBCluster = "FailoverDBCluster"

// FailoverDBClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Forces a failover for a DB cluster.
//
// A failover for a DB cluster promotes one of the Aurora Replicas (read-only
// instances) in the DB cluster to be the primary instance (the cluster writer).
//
// Amazon Aurora will automatically fail over to an Aurora Replica, if one exists,
// when the primary instance fails. You can force a failover when you want to
// simulate a failure of a primary instance for testing. Because each instance
// in a DB cluster has its own endpoint address, you will need to clean up and
// re-establish any existing connections that use those endpoint addresses when
// the failover is complete.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using FailoverDBClusterRequest.
//    req := client.FailoverDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/FailoverDBCluster
func (c *Client) FailoverDBClusterRequest(input *types.FailoverDBClusterInput) FailoverDBClusterRequest {
	op := &aws.Operation{
		Name:       opFailoverDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.FailoverDBClusterInput{}
	}

	req := c.newRequest(op, input, &types.FailoverDBClusterOutput{})
	return FailoverDBClusterRequest{Request: req, Input: input, Copy: c.FailoverDBClusterRequest}
}

// FailoverDBClusterRequest is the request type for the
// FailoverDBCluster API operation.
type FailoverDBClusterRequest struct {
	*aws.Request
	Input *types.FailoverDBClusterInput
	Copy  func(*types.FailoverDBClusterInput) FailoverDBClusterRequest
}

// Send marshals and sends the FailoverDBCluster API request.
func (r FailoverDBClusterRequest) Send(ctx context.Context) (*FailoverDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &FailoverDBClusterResponse{
		FailoverDBClusterOutput: r.Request.Data.(*types.FailoverDBClusterOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// FailoverDBClusterResponse is the response type for the
// FailoverDBCluster API operation.
type FailoverDBClusterResponse struct {
	*types.FailoverDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// FailoverDBCluster request.
func (r *FailoverDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
