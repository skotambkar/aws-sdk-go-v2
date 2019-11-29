// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opResizeCluster = "ResizeCluster"

// ResizeClusterRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Changes the size of the cluster. You can change the cluster's type, or change
// the number or type of nodes. The default behavior is to use the elastic resize
// method. With an elastic resize, your cluster is available for read and write
// operations more quickly than with the classic resize method.
//
// Elastic resize operations have the following restrictions:
//
//    * You can only resize clusters of the following types: dc2.large dc2.8xlarge
//    ds2.xlarge ds2.8xlarge
//
//    * The type of nodes that you add must match the node type for the cluster.
//
//    // Example sending a request using ResizeClusterRequest.
//    req := client.ResizeClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ResizeCluster
func (c *Client) ResizeClusterRequest(input *types.ResizeClusterInput) ResizeClusterRequest {
	op := &aws.Operation{
		Name:       opResizeCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResizeClusterInput{}
	}

	req := c.newRequest(op, input, &types.ResizeClusterOutput{})
	return ResizeClusterRequest{Request: req, Input: input, Copy: c.ResizeClusterRequest}
}

// ResizeClusterRequest is the request type for the
// ResizeCluster API operation.
type ResizeClusterRequest struct {
	*aws.Request
	Input *types.ResizeClusterInput
	Copy  func(*types.ResizeClusterInput) ResizeClusterRequest
}

// Send marshals and sends the ResizeCluster API request.
func (r ResizeClusterRequest) Send(ctx context.Context) (*ResizeClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResizeClusterResponse{
		ResizeClusterOutput: r.Request.Data.(*types.ResizeClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResizeClusterResponse is the response type for the
// ResizeCluster API operation.
type ResizeClusterResponse struct {
	*types.ResizeClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResizeCluster request.
func (r *ResizeClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
