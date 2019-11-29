// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opCreateRule = "CreateRule"

// CreateRuleRequest returns a request value for making API operation for
// AWS WAF.
//
// Creates a Rule, which contains the IPSet objects, ByteMatchSet objects, and
// other predicates that identify the requests that you want to block. If you
// add more than one predicate to a Rule, a request must match all of the specifications
// to be allowed or blocked. For example, suppose that you add the following
// to a Rule:
//
//    * An IPSet that matches the IP address 192.0.2.44/32
//
//    * A ByteMatchSet that matches BadBot in the User-Agent header
//
// You then add the Rule to a WebACL and specify that you want to blocks requests
// that satisfy the Rule. For a request to be blocked, it must come from the
// IP address 192.0.2.44 and the User-Agent header in the request must contain
// the value BadBot.
//
// To create and configure a Rule, perform the following steps:
//
// Create and update the predicates that you want to include in the Rule. For
// more information, see CreateByteMatchSet, CreateIPSet, and CreateSqlInjectionMatchSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateRule request.
//
// Submit a CreateRule request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRule request.
//
// Submit an UpdateRule request to specify the predicates that you want to include
// in the Rule.
//
// Create and update a WebACL that contains the Rule. For more information,
// see CreateWebACL.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateRuleRequest.
//    req := client.CreateRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/CreateRule
func (c *Client) CreateRuleRequest(input *types.CreateRuleInput) CreateRuleRequest {
	op := &aws.Operation{
		Name:       opCreateRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateRuleInput{}
	}

	req := c.newRequest(op, input, &types.CreateRuleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateRuleMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateRuleRequest{Request: req, Input: input, Copy: c.CreateRuleRequest}
}

// CreateRuleRequest is the request type for the
// CreateRule API operation.
type CreateRuleRequest struct {
	*aws.Request
	Input *types.CreateRuleInput
	Copy  func(*types.CreateRuleInput) CreateRuleRequest
}

// Send marshals and sends the CreateRule API request.
func (r CreateRuleRequest) Send(ctx context.Context) (*CreateRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRuleResponse{
		CreateRuleOutput: r.Request.Data.(*types.CreateRuleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRuleResponse is the response type for the
// CreateRule API operation.
type CreateRuleResponse struct {
	*types.CreateRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRule request.
func (r *CreateRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
