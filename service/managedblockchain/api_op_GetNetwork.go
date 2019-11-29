// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
)

const opGetNetwork = "GetNetwork"

// GetNetworkRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns detailed information about a network.
//
//    // Example sending a request using GetNetworkRequest.
//    req := client.GetNetworkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/GetNetwork
func (c *Client) GetNetworkRequest(input *types.GetNetworkInput) GetNetworkRequest {
	op := &aws.Operation{
		Name:       opGetNetwork,
		HTTPMethod: "GET",
		HTTPPath:   "/networks/{networkId}",
	}

	if input == nil {
		input = &types.GetNetworkInput{}
	}

	req := c.newRequest(op, input, &types.GetNetworkOutput{})
	return GetNetworkRequest{Request: req, Input: input, Copy: c.GetNetworkRequest}
}

// GetNetworkRequest is the request type for the
// GetNetwork API operation.
type GetNetworkRequest struct {
	*aws.Request
	Input *types.GetNetworkInput
	Copy  func(*types.GetNetworkInput) GetNetworkRequest
}

// Send marshals and sends the GetNetwork API request.
func (r GetNetworkRequest) Send(ctx context.Context) (*GetNetworkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetNetworkResponse{
		GetNetworkOutput: r.Request.Data.(*types.GetNetworkOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetNetworkResponse is the response type for the
// GetNetwork API operation.
type GetNetworkResponse struct {
	*types.GetNetworkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetNetwork request.
func (r *GetNetworkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
