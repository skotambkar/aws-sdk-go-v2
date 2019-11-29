// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opAssociateHostedConnection = "AssociateHostedConnection"

// AssociateHostedConnectionRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Associates a hosted connection and its virtual interfaces with a link aggregation
// group (LAG) or interconnect. If the target interconnect or LAG has an existing
// hosted connection with a conflicting VLAN number or IP address, the operation
// fails. This action temporarily interrupts the hosted connection's connectivity
// to AWS as it is being migrated.
//
// Intended for use by AWS Direct Connect Partners only.
//
//    // Example sending a request using AssociateHostedConnectionRequest.
//    req := client.AssociateHostedConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AssociateHostedConnection
func (c *Client) AssociateHostedConnectionRequest(input *types.AssociateHostedConnectionInput) AssociateHostedConnectionRequest {
	op := &aws.Operation{
		Name:       opAssociateHostedConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateHostedConnectionInput{}
	}

	req := c.newRequest(op, input, &types.AssociateHostedConnectionOutput{})
	return AssociateHostedConnectionRequest{Request: req, Input: input, Copy: c.AssociateHostedConnectionRequest}
}

// AssociateHostedConnectionRequest is the request type for the
// AssociateHostedConnection API operation.
type AssociateHostedConnectionRequest struct {
	*aws.Request
	Input *types.AssociateHostedConnectionInput
	Copy  func(*types.AssociateHostedConnectionInput) AssociateHostedConnectionRequest
}

// Send marshals and sends the AssociateHostedConnection API request.
func (r AssociateHostedConnectionRequest) Send(ctx context.Context) (*AssociateHostedConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateHostedConnectionResponse{
		AssociateHostedConnectionOutput: r.Request.Data.(*types.AssociateHostedConnectionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateHostedConnectionResponse is the response type for the
// AssociateHostedConnection API operation.
type AssociateHostedConnectionResponse struct {
	*types.AssociateHostedConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateHostedConnection request.
func (r *AssociateHostedConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
