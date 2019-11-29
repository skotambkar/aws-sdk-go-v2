// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsmv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
)

const opDeleteCluster = "DeleteCluster"

// DeleteClusterRequest returns a request value for making API operation for
// AWS CloudHSM V2.
//
// Deletes the specified AWS CloudHSM cluster. Before you can delete a cluster,
// you must delete all HSMs in the cluster. To see if the cluster contains any
// HSMs, use DescribeClusters. To delete an HSM, use DeleteHsm.
//
//    // Example sending a request using DeleteClusterRequest.
//    req := client.DeleteClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsmv2-2017-04-28/DeleteCluster
func (c *Client) DeleteClusterRequest(input *types.DeleteClusterInput) DeleteClusterRequest {
	op := &aws.Operation{
		Name:       opDeleteCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteClusterInput{}
	}

	req := c.newRequest(op, input, &types.DeleteClusterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteClusterMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteClusterRequest{Request: req, Input: input, Copy: c.DeleteClusterRequest}
}

// DeleteClusterRequest is the request type for the
// DeleteCluster API operation.
type DeleteClusterRequest struct {
	*aws.Request
	Input *types.DeleteClusterInput
	Copy  func(*types.DeleteClusterInput) DeleteClusterRequest
}

// Send marshals and sends the DeleteCluster API request.
func (r DeleteClusterRequest) Send(ctx context.Context) (*DeleteClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteClusterResponse{
		DeleteClusterOutput: r.Request.Data.(*types.DeleteClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteClusterResponse is the response type for the
// DeleteCluster API operation.
type DeleteClusterResponse struct {
	*types.DeleteClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCluster request.
func (r *DeleteClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
