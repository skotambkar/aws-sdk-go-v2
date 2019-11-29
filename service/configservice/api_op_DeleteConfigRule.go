// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opDeleteConfigRule = "DeleteConfigRule"

// DeleteConfigRuleRequest returns a request value for making API operation for
// AWS Config.
//
// Deletes the specified AWS Config rule and all of its evaluation results.
//
// AWS Config sets the state of a rule to DELETING until the deletion is complete.
// You cannot update a rule while it is in this state. If you make a PutConfigRule
// or DeleteConfigRule request for the rule, you will receive a ResourceInUseException.
//
// You can check the state of a rule by using the DescribeConfigRules request.
//
//    // Example sending a request using DeleteConfigRuleRequest.
//    req := client.DeleteConfigRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteConfigRule
func (c *Client) DeleteConfigRuleRequest(input *types.DeleteConfigRuleInput) DeleteConfigRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteConfigRuleInput{}
	}

	req := c.newRequest(op, input, &types.DeleteConfigRuleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteConfigRuleMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteConfigRuleRequest{Request: req, Input: input, Copy: c.DeleteConfigRuleRequest}
}

// DeleteConfigRuleRequest is the request type for the
// DeleteConfigRule API operation.
type DeleteConfigRuleRequest struct {
	*aws.Request
	Input *types.DeleteConfigRuleInput
	Copy  func(*types.DeleteConfigRuleInput) DeleteConfigRuleRequest
}

// Send marshals and sends the DeleteConfigRule API request.
func (r DeleteConfigRuleRequest) Send(ctx context.Context) (*DeleteConfigRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigRuleResponse{
		DeleteConfigRuleOutput: r.Request.Data.(*types.DeleteConfigRuleOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigRuleResponse is the response type for the
// DeleteConfigRule API operation.
type DeleteConfigRuleResponse struct {
	*types.DeleteConfigRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigRule request.
func (r *DeleteConfigRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
