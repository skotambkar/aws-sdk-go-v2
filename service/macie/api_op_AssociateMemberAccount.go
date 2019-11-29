// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/macie/types"
)

const opAssociateMemberAccount = "AssociateMemberAccount"

// AssociateMemberAccountRequest returns a request value for making API operation for
// Amazon Macie.
//
// Associates a specified AWS account with Amazon Macie as a member account.
//
//    // Example sending a request using AssociateMemberAccountRequest.
//    req := client.AssociateMemberAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie-2017-12-19/AssociateMemberAccount
func (c *Client) AssociateMemberAccountRequest(input *types.AssociateMemberAccountInput) AssociateMemberAccountRequest {
	op := &aws.Operation{
		Name:       opAssociateMemberAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateMemberAccountInput{}
	}

	req := c.newRequest(op, input, &types.AssociateMemberAccountOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AssociateMemberAccountRequest{Request: req, Input: input, Copy: c.AssociateMemberAccountRequest}
}

// AssociateMemberAccountRequest is the request type for the
// AssociateMemberAccount API operation.
type AssociateMemberAccountRequest struct {
	*aws.Request
	Input *types.AssociateMemberAccountInput
	Copy  func(*types.AssociateMemberAccountInput) AssociateMemberAccountRequest
}

// Send marshals and sends the AssociateMemberAccount API request.
func (r AssociateMemberAccountRequest) Send(ctx context.Context) (*AssociateMemberAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateMemberAccountResponse{
		AssociateMemberAccountOutput: r.Request.Data.(*types.AssociateMemberAccountOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateMemberAccountResponse is the response type for the
// AssociateMemberAccount API operation.
type AssociateMemberAccountResponse struct {
	*types.AssociateMemberAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateMemberAccount request.
func (r *AssociateMemberAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
