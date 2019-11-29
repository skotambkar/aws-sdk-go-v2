// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsmv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
)

const opInitializeCluster = "InitializeCluster"

// InitializeClusterRequest returns a request value for making API operation for
// AWS CloudHSM V2.
//
// Claims an AWS CloudHSM cluster by submitting the cluster certificate issued
// by your issuing certificate authority (CA) and the CA's root certificate.
// Before you can claim a cluster, you must sign the cluster's certificate signing
// request (CSR) with your issuing CA. To get the cluster's CSR, use DescribeClusters.
//
//    // Example sending a request using InitializeClusterRequest.
//    req := client.InitializeClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/InitializeCluster
func (c *Client) InitializeClusterRequest(input *types.InitializeClusterInput) InitializeClusterRequest {
	op := &aws.Operation{
		Name:       opInitializeCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.InitializeClusterInput{}
	}

	req := c.newRequest(op, input, &types.InitializeClusterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.InitializeClusterMarshaler{Input: input}.GetNamedBuildHandler())

	return InitializeClusterRequest{Request: req, Input: input, Copy: c.InitializeClusterRequest}
}

// InitializeClusterRequest is the request type for the
// InitializeCluster API operation.
type InitializeClusterRequest struct {
	*aws.Request
	Input *types.InitializeClusterInput
	Copy  func(*types.InitializeClusterInput) InitializeClusterRequest
}

// Send marshals and sends the InitializeCluster API request.
func (r InitializeClusterRequest) Send(ctx context.Context) (*InitializeClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InitializeClusterResponse{
		InitializeClusterOutput: r.Request.Data.(*types.InitializeClusterOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InitializeClusterResponse is the response type for the
// InitializeCluster API operation.
type InitializeClusterResponse struct {
	*types.InitializeClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InitializeCluster request.
func (r *InitializeClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
