// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opCreateApprovalRuleTemplate = "CreateApprovalRuleTemplate"

// CreateApprovalRuleTemplateRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Creates a template for approval rules that can then be associated with one
// or more repositories in your AWS account. When you associate a template with
// a repository, AWS CodeCommit creates an approval rule that matches the conditions
// of the template for all pull requests that meet the conditions of the template.
// For more information, see AssociateApprovalRuleTemplateWithRepository.
//
//    // Example sending a request using CreateApprovalRuleTemplateRequest.
//    req := client.CreateApprovalRuleTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/CreateApprovalRuleTemplate
func (c *Client) CreateApprovalRuleTemplateRequest(input *types.CreateApprovalRuleTemplateInput) CreateApprovalRuleTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateApprovalRuleTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateApprovalRuleTemplateInput{}
	}

	req := c.newRequest(op, input, &types.CreateApprovalRuleTemplateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateApprovalRuleTemplateMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateApprovalRuleTemplateRequest{Request: req, Input: input, Copy: c.CreateApprovalRuleTemplateRequest}
}

// CreateApprovalRuleTemplateRequest is the request type for the
// CreateApprovalRuleTemplate API operation.
type CreateApprovalRuleTemplateRequest struct {
	*aws.Request
	Input *types.CreateApprovalRuleTemplateInput
	Copy  func(*types.CreateApprovalRuleTemplateInput) CreateApprovalRuleTemplateRequest
}

// Send marshals and sends the CreateApprovalRuleTemplate API request.
func (r CreateApprovalRuleTemplateRequest) Send(ctx context.Context) (*CreateApprovalRuleTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApprovalRuleTemplateResponse{
		CreateApprovalRuleTemplateOutput: r.Request.Data.(*types.CreateApprovalRuleTemplateOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApprovalRuleTemplateResponse is the response type for the
// CreateApprovalRuleTemplate API operation.
type CreateApprovalRuleTemplateResponse struct {
	*types.CreateApprovalRuleTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApprovalRuleTemplate request.
func (r *CreateApprovalRuleTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
