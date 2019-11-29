// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDisassociateSubnetCidrBlock = "DisassociateSubnetCidrBlock"

// DisassociateSubnetCidrBlockRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disassociates a CIDR block from a subnet. Currently, you can disassociate
// an IPv6 CIDR block only. You must detach or delete all gateways and resources
// that are associated with the CIDR block before you can disassociate it.
//
//    // Example sending a request using DisassociateSubnetCidrBlockRequest.
//    req := client.DisassociateSubnetCidrBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateSubnetCidrBlock
func (c *Client) DisassociateSubnetCidrBlockRequest(input *types.DisassociateSubnetCidrBlockInput) DisassociateSubnetCidrBlockRequest {
	op := &aws.Operation{
		Name:       opDisassociateSubnetCidrBlock,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisassociateSubnetCidrBlockInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateSubnetCidrBlockOutput{})
	return DisassociateSubnetCidrBlockRequest{Request: req, Input: input, Copy: c.DisassociateSubnetCidrBlockRequest}
}

// DisassociateSubnetCidrBlockRequest is the request type for the
// DisassociateSubnetCidrBlock API operation.
type DisassociateSubnetCidrBlockRequest struct {
	*aws.Request
	Input *types.DisassociateSubnetCidrBlockInput
	Copy  func(*types.DisassociateSubnetCidrBlockInput) DisassociateSubnetCidrBlockRequest
}

// Send marshals and sends the DisassociateSubnetCidrBlock API request.
func (r DisassociateSubnetCidrBlockRequest) Send(ctx context.Context) (*DisassociateSubnetCidrBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateSubnetCidrBlockResponse{
		DisassociateSubnetCidrBlockOutput: r.Request.Data.(*types.DisassociateSubnetCidrBlockOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateSubnetCidrBlockResponse is the response type for the
// DisassociateSubnetCidrBlock API operation.
type DisassociateSubnetCidrBlockResponse struct {
	*types.DisassociateSubnetCidrBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateSubnetCidrBlock request.
func (r *DisassociateSubnetCidrBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
