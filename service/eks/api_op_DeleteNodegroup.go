// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

const opDeleteNodegroup = "DeleteNodegroup"

// DeleteNodegroupRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Deletes an Amazon EKS node group for a cluster.
//
//    // Example sending a request using DeleteNodegroupRequest.
//    req := client.DeleteNodegroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/DeleteNodegroup
func (c *Client) DeleteNodegroupRequest(input *types.DeleteNodegroupInput) DeleteNodegroupRequest {
	op := &aws.Operation{
		Name:       opDeleteNodegroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/clusters/{name}/node-groups/{nodegroupName}",
	}

	if input == nil {
		input = &types.DeleteNodegroupInput{}
	}

	req := c.newRequest(op, input, &types.DeleteNodegroupOutput{})
	return DeleteNodegroupRequest{Request: req, Input: input, Copy: c.DeleteNodegroupRequest}
}

// DeleteNodegroupRequest is the request type for the
// DeleteNodegroup API operation.
type DeleteNodegroupRequest struct {
	*aws.Request
	Input *types.DeleteNodegroupInput
	Copy  func(*types.DeleteNodegroupInput) DeleteNodegroupRequest
}

// Send marshals and sends the DeleteNodegroup API request.
func (r DeleteNodegroupRequest) Send(ctx context.Context) (*DeleteNodegroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNodegroupResponse{
		DeleteNodegroupOutput: r.Request.Data.(*types.DeleteNodegroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNodegroupResponse is the response type for the
// DeleteNodegroup API operation.
type DeleteNodegroupResponse struct {
	*types.DeleteNodegroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNodegroup request.
func (r *DeleteNodegroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
