// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opDisassociateConnectionFromLag = "DisassociateConnectionFromLag"

// DisassociateConnectionFromLagRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Disassociates a connection from a link aggregation group (LAG). The connection
// is interrupted and re-established as a standalone connection (the connection
// is not deleted; to delete the connection, use the DeleteConnection request).
// If the LAG has associated virtual interfaces or hosted connections, they
// remain associated with the LAG. A disassociated connection owned by an AWS
// Direct Connect Partner is automatically converted to an interconnect.
//
// If disassociating the connection would cause the LAG to fall below its setting
// for minimum number of operational connections, the request fails, except
// when it's the last member of the LAG. If all connections are disassociated,
// the LAG continues to exist as an empty LAG with no physical connections.
//
//    // Example sending a request using DisassociateConnectionFromLagRequest.
//    req := client.DisassociateConnectionFromLagRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DisassociateConnectionFromLag
func (c *Client) DisassociateConnectionFromLagRequest(input *types.DisassociateConnectionFromLagInput) DisassociateConnectionFromLagRequest {
	op := &aws.Operation{
		Name:       opDisassociateConnectionFromLag,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisassociateConnectionFromLagInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateConnectionFromLagOutput{})
	return DisassociateConnectionFromLagRequest{Request: req, Input: input, Copy: c.DisassociateConnectionFromLagRequest}
}

// DisassociateConnectionFromLagRequest is the request type for the
// DisassociateConnectionFromLag API operation.
type DisassociateConnectionFromLagRequest struct {
	*aws.Request
	Input *types.DisassociateConnectionFromLagInput
	Copy  func(*types.DisassociateConnectionFromLagInput) DisassociateConnectionFromLagRequest
}

// Send marshals and sends the DisassociateConnectionFromLag API request.
func (r DisassociateConnectionFromLagRequest) Send(ctx context.Context) (*DisassociateConnectionFromLagResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateConnectionFromLagResponse{
		DisassociateConnectionFromLagOutput: r.Request.Data.(*types.DisassociateConnectionFromLagOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateConnectionFromLagResponse is the response type for the
// DisassociateConnectionFromLag API operation.
type DisassociateConnectionFromLagResponse struct {
	*types.DisassociateConnectionFromLagOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateConnectionFromLag request.
func (r *DisassociateConnectionFromLagResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
