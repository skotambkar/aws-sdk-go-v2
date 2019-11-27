// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
)

const opGetNode = "GetNode"

// GetNodeRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns detailed information about a peer node.
//
//    // Example sending a request using GetNodeRequest.
//    req := client.GetNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetNode
func (c *Client) GetNodeRequest(input *types.GetNodeInput) GetNodeRequest {
	op := &aws.Operation{
		Name:       opGetNode,
		HTTPMethod: "GET",
		HTTPPath:   "/networks/{networkId}/members/{memberId}/nodes/{nodeId}",
	}

	if input == nil {
		input = &types.GetNodeInput{}
	}

	req := c.newRequest(op, input, &types.GetNodeOutput{})
	return GetNodeRequest{Request: req, Input: input, Copy: c.GetNodeRequest}
}

// GetNodeRequest is the request type for the
// GetNode API operation.
type GetNodeRequest struct {
	*aws.Request
	Input *types.GetNodeInput
	Copy  func(*types.GetNodeInput) GetNodeRequest
}

// Send marshals and sends the GetNode API request.
func (r GetNodeRequest) Send(ctx context.Context) (*GetNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetNodeResponse{
		GetNodeOutput: r.Request.Data.(*types.GetNodeOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetNodeResponse is the response type for the
// GetNode API operation.
type GetNodeResponse struct {
	*types.GetNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetNode request.
func (r *GetNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
