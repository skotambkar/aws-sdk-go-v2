// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opDisassociateApprovalRuleTemplateFromRepository = "DisassociateApprovalRuleTemplateFromRepository"

// DisassociateApprovalRuleTemplateFromRepositoryRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Removes the association between a template and a repository so that approval
// rules based on the template are not automatically created when pull requests
// are created in the specified repository. This does not delete any approval
// rules previously created for pull requests through the template association.
//
//    // Example sending a request using DisassociateApprovalRuleTemplateFromRepositoryRequest.
//    req := client.DisassociateApprovalRuleTemplateFromRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DisassociateApprovalRuleTemplateFromRepository
func (c *Client) DisassociateApprovalRuleTemplateFromRepositoryRequest(input *types.DisassociateApprovalRuleTemplateFromRepositoryInput) DisassociateApprovalRuleTemplateFromRepositoryRequest {
	op := &aws.Operation{
		Name:       opDisassociateApprovalRuleTemplateFromRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisassociateApprovalRuleTemplateFromRepositoryInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateApprovalRuleTemplateFromRepositoryOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DisassociateApprovalRuleTemplateFromRepositoryRequest{Request: req, Input: input, Copy: c.DisassociateApprovalRuleTemplateFromRepositoryRequest}
}

// DisassociateApprovalRuleTemplateFromRepositoryRequest is the request type for the
// DisassociateApprovalRuleTemplateFromRepository API operation.
type DisassociateApprovalRuleTemplateFromRepositoryRequest struct {
	*aws.Request
	Input *types.DisassociateApprovalRuleTemplateFromRepositoryInput
	Copy  func(*types.DisassociateApprovalRuleTemplateFromRepositoryInput) DisassociateApprovalRuleTemplateFromRepositoryRequest
}

// Send marshals and sends the DisassociateApprovalRuleTemplateFromRepository API request.
func (r DisassociateApprovalRuleTemplateFromRepositoryRequest) Send(ctx context.Context) (*DisassociateApprovalRuleTemplateFromRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateApprovalRuleTemplateFromRepositoryResponse{
		DisassociateApprovalRuleTemplateFromRepositoryOutput: r.Request.Data.(*types.DisassociateApprovalRuleTemplateFromRepositoryOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateApprovalRuleTemplateFromRepositoryResponse is the response type for the
// DisassociateApprovalRuleTemplateFromRepository API operation.
type DisassociateApprovalRuleTemplateFromRepositoryResponse struct {
	*types.DisassociateApprovalRuleTemplateFromRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateApprovalRuleTemplateFromRepository request.
func (r *DisassociateApprovalRuleTemplateFromRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
