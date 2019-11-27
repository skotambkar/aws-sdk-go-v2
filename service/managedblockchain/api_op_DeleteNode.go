// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
)

const opDeleteNode = "DeleteNode"

// DeleteNodeRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Deletes a peer node from a member that your AWS account owns. All data on
// the node is lost and cannot be recovered.
//
//    // Example sending a request using DeleteNodeRequest.
//    req := client.DeleteNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/DeleteNode
func (c *Client) DeleteNodeRequest(input *types.DeleteNodeInput) DeleteNodeRequest {
	op := &aws.Operation{
		Name:       opDeleteNode,
		HTTPMethod: "DELETE",
		HTTPPath:   "/networks/{networkId}/members/{memberId}/nodes/{nodeId}",
	}

	if input == nil {
		input = &types.DeleteNodeInput{}
	}

	req := c.newRequest(op, input, &types.DeleteNodeOutput{})
	return DeleteNodeRequest{Request: req, Input: input, Copy: c.DeleteNodeRequest}
}

// DeleteNodeRequest is the request type for the
// DeleteNode API operation.
type DeleteNodeRequest struct {
	*aws.Request
	Input *types.DeleteNodeInput
	Copy  func(*types.DeleteNodeInput) DeleteNodeRequest
}

// Send marshals and sends the DeleteNode API request.
func (r DeleteNodeRequest) Send(ctx context.Context) (*DeleteNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNodeResponse{
		DeleteNodeOutput: r.Request.Data.(*types.DeleteNodeOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNodeResponse is the response type for the
// DeleteNode API operation.
type DeleteNodeResponse struct {
	*types.DeleteNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNode request.
func (r *DeleteNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
