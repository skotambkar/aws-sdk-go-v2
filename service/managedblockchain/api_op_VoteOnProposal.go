// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
)

const opVoteOnProposal = "VoteOnProposal"

// VoteOnProposalRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Casts a vote for a specified ProposalId on behalf of a member. The member
// to vote as, specified by VoterMemberId, must be in the same AWS account as
// the principal that calls the action.
//
//    // Example sending a request using VoteOnProposalRequest.
//    req := client.VoteOnProposalRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/VoteOnProposal
func (c *Client) VoteOnProposalRequest(input *types.VoteOnProposalInput) VoteOnProposalRequest {
	op := &aws.Operation{
		Name:       opVoteOnProposal,
		HTTPMethod: "POST",
		HTTPPath:   "/networks/{networkId}/proposals/{proposalId}/votes",
	}

	if input == nil {
		input = &types.VoteOnProposalInput{}
	}

	req := c.newRequest(op, input, &types.VoteOnProposalOutput{})
	return VoteOnProposalRequest{Request: req, Input: input, Copy: c.VoteOnProposalRequest}
}

// VoteOnProposalRequest is the request type for the
// VoteOnProposal API operation.
type VoteOnProposalRequest struct {
	*aws.Request
	Input *types.VoteOnProposalInput
	Copy  func(*types.VoteOnProposalInput) VoteOnProposalRequest
}

// Send marshals and sends the VoteOnProposal API request.
func (r VoteOnProposalRequest) Send(ctx context.Context) (*VoteOnProposalResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &VoteOnProposalResponse{
		VoteOnProposalOutput: r.Request.Data.(*types.VoteOnProposalOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// VoteOnProposalResponse is the response type for the
// VoteOnProposal API operation.
type VoteOnProposalResponse struct {
	*types.VoteOnProposalOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// VoteOnProposal request.
func (r *VoteOnProposalResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
