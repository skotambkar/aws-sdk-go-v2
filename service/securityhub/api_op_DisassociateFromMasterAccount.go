// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
)

const opDisassociateFromMasterAccount = "DisassociateFromMasterAccount"

// DisassociateFromMasterAccountRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Disassociates the current Security Hub member account from the associated
// master account.
//
//    // Example sending a request using DisassociateFromMasterAccountRequest.
//    req := client.DisassociateFromMasterAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DisassociateFromMasterAccount
func (c *Client) DisassociateFromMasterAccountRequest(input *types.DisassociateFromMasterAccountInput) DisassociateFromMasterAccountRequest {
	op := &aws.Operation{
		Name:       opDisassociateFromMasterAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/master/disassociate",
	}

	if input == nil {
		input = &types.DisassociateFromMasterAccountInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateFromMasterAccountOutput{})
	return DisassociateFromMasterAccountRequest{Request: req, Input: input, Copy: c.DisassociateFromMasterAccountRequest}
}

// DisassociateFromMasterAccountRequest is the request type for the
// DisassociateFromMasterAccount API operation.
type DisassociateFromMasterAccountRequest struct {
	*aws.Request
	Input *types.DisassociateFromMasterAccountInput
	Copy  func(*types.DisassociateFromMasterAccountInput) DisassociateFromMasterAccountRequest
}

// Send marshals and sends the DisassociateFromMasterAccount API request.
func (r DisassociateFromMasterAccountRequest) Send(ctx context.Context) (*DisassociateFromMasterAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateFromMasterAccountResponse{
		DisassociateFromMasterAccountOutput: r.Request.Data.(*types.DisassociateFromMasterAccountOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateFromMasterAccountResponse is the response type for the
// DisassociateFromMasterAccount API operation.
type DisassociateFromMasterAccountResponse struct {
	*types.DisassociateFromMasterAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateFromMasterAccount request.
func (r *DisassociateFromMasterAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
