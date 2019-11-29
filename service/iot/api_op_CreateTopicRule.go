// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opCreateTopicRule = "CreateTopicRule"

// CreateTopicRuleRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a rule. Creating rules is an administrator-level action. Any user
// who has permission to create rules will be able to access data processed
// by the rule.
//
//    // Example sending a request using CreateTopicRuleRequest.
//    req := client.CreateTopicRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateTopicRuleRequest(input *types.CreateTopicRuleInput) CreateTopicRuleRequest {
	op := &aws.Operation{
		Name:       opCreateTopicRule,
		HTTPMethod: "POST",
		HTTPPath:   "/rules/{ruleName}",
	}

	if input == nil {
		input = &types.CreateTopicRuleInput{}
	}

	req := c.newRequest(op, input, &types.CreateTopicRuleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateTopicRuleMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateTopicRuleRequest{Request: req, Input: input, Copy: c.CreateTopicRuleRequest}
}

// CreateTopicRuleRequest is the request type for the
// CreateTopicRule API operation.
type CreateTopicRuleRequest struct {
	*aws.Request
	Input *types.CreateTopicRuleInput
	Copy  func(*types.CreateTopicRuleInput) CreateTopicRuleRequest
}

// Send marshals and sends the CreateTopicRule API request.
func (r CreateTopicRuleRequest) Send(ctx context.Context) (*CreateTopicRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTopicRuleResponse{
		CreateTopicRuleOutput: r.Request.Data.(*types.CreateTopicRuleOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTopicRuleResponse is the response type for the
// CreateTopicRule API operation.
type CreateTopicRuleResponse struct {
	*types.CreateTopicRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTopicRule request.
func (r *CreateTopicRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
