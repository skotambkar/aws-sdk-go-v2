// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
)

const opGetResolverRuleAssociation = "GetResolverRuleAssociation"

// GetResolverRuleAssociationRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Gets information about an association between a specified resolver rule and
// a VPC. You associate a resolver rule and a VPC using AssociateResolverRule.
//
//    // Example sending a request using GetResolverRuleAssociationRequest.
//    req := client.GetResolverRuleAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/GetResolverRuleAssociation
func (c *Client) GetResolverRuleAssociationRequest(input *types.GetResolverRuleAssociationInput) GetResolverRuleAssociationRequest {
	op := &aws.Operation{
		Name:       opGetResolverRuleAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetResolverRuleAssociationInput{}
	}

	req := c.newRequest(op, input, &types.GetResolverRuleAssociationOutput{})
	return GetResolverRuleAssociationRequest{Request: req, Input: input, Copy: c.GetResolverRuleAssociationRequest}
}

// GetResolverRuleAssociationRequest is the request type for the
// GetResolverRuleAssociation API operation.
type GetResolverRuleAssociationRequest struct {
	*aws.Request
	Input *types.GetResolverRuleAssociationInput
	Copy  func(*types.GetResolverRuleAssociationInput) GetResolverRuleAssociationRequest
}

// Send marshals and sends the GetResolverRuleAssociation API request.
func (r GetResolverRuleAssociationRequest) Send(ctx context.Context) (*GetResolverRuleAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResolverRuleAssociationResponse{
		GetResolverRuleAssociationOutput: r.Request.Data.(*types.GetResolverRuleAssociationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetResolverRuleAssociationResponse is the response type for the
// GetResolverRuleAssociation API operation.
type GetResolverRuleAssociationResponse struct {
	*types.GetResolverRuleAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResolverRuleAssociation request.
func (r *GetResolverRuleAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
