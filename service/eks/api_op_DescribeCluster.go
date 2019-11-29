// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/eks/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

const opDescribeCluster = "DescribeCluster"

// DescribeClusterRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Returns descriptive information about an Amazon EKS cluster.
//
// The API server endpoint and certificate authority data returned by this operation
// are required for kubelet and kubectl to communicate with your Kubernetes
// API server. For more information, see Create a kubeconfig for Amazon EKS
// (https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html).
//
// The API server endpoint and certificate authority data aren't available until
// the cluster reaches the ACTIVE state.
//
//    // Example sending a request using DescribeClusterRequest.
//    req := client.DescribeClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/DescribeCluster
func (c *Client) DescribeClusterRequest(input *types.DescribeClusterInput) DescribeClusterRequest {
	op := &aws.Operation{
		Name:       opDescribeCluster,
		HTTPMethod: "GET",
		HTTPPath:   "/clusters/{name}",
	}

	if input == nil {
		input = &types.DescribeClusterInput{}
	}

	req := c.newRequest(op, input, &types.DescribeClusterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeClusterMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeClusterRequest{Request: req, Input: input, Copy: c.DescribeClusterRequest}
}

// DescribeClusterRequest is the request type for the
// DescribeCluster API operation.
type DescribeClusterRequest struct {
	*aws.Request
	Input *types.DescribeClusterInput
	Copy  func(*types.DescribeClusterInput) DescribeClusterRequest
}

// Send marshals and sends the DescribeCluster API request.
func (r DescribeClusterRequest) Send(ctx context.Context) (*DescribeClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClusterResponse{
		DescribeClusterOutput: r.Request.Data.(*types.DescribeClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeClusterResponse is the response type for the
// DescribeCluster API operation.
type DescribeClusterResponse struct {
	*types.DescribeClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCluster request.
func (r *DescribeClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
