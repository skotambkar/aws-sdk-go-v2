// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
)

const opDisassociateResolverRule = "DisassociateResolverRule"

// DisassociateResolverRuleRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Removes the association between a specified resolver rule and a specified
// VPC.
//
// If you disassociate a resolver rule from a VPC, Resolver stops forwarding
// DNS queries for the domain name that you specified in the resolver rule.
//
//    // Example sending a request using DisassociateResolverRuleRequest.
//    req := client.DisassociateResolverRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/DisassociateResolverRule
func (c *Client) DisassociateResolverRuleRequest(input *types.DisassociateResolverRuleInput) DisassociateResolverRuleRequest {
	op := &aws.Operation{
		Name:       opDisassociateResolverRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisassociateResolverRuleInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateResolverRuleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DisassociateResolverRuleMarshaler{Input: input}.GetNamedBuildHandler())

	return DisassociateResolverRuleRequest{Request: req, Input: input, Copy: c.DisassociateResolverRuleRequest}
}

// DisassociateResolverRuleRequest is the request type for the
// DisassociateResolverRule API operation.
type DisassociateResolverRuleRequest struct {
	*aws.Request
	Input *types.DisassociateResolverRuleInput
	Copy  func(*types.DisassociateResolverRuleInput) DisassociateResolverRuleRequest
}

// Send marshals and sends the DisassociateResolverRule API request.
func (r DisassociateResolverRuleRequest) Send(ctx context.Context) (*DisassociateResolverRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateResolverRuleResponse{
		DisassociateResolverRuleOutput: r.Request.Data.(*types.DisassociateResolverRuleOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateResolverRuleResponse is the response type for the
// DisassociateResolverRule API operation.
type DisassociateResolverRuleResponse struct {
	*types.DisassociateResolverRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateResolverRule request.
func (r *DisassociateResolverRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
