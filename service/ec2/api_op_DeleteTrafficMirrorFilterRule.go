// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDeleteTrafficMirrorFilterRule = "DeleteTrafficMirrorFilterRule"

// DeleteTrafficMirrorFilterRuleRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified Traffic Mirror rule.
//
//    // Example sending a request using DeleteTrafficMirrorFilterRuleRequest.
//    req := client.DeleteTrafficMirrorFilterRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTrafficMirrorFilterRule
func (c *Client) DeleteTrafficMirrorFilterRuleRequest(input *types.DeleteTrafficMirrorFilterRuleInput) DeleteTrafficMirrorFilterRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteTrafficMirrorFilterRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteTrafficMirrorFilterRuleInput{}
	}

	req := c.newRequest(op, input, &types.DeleteTrafficMirrorFilterRuleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DeleteTrafficMirrorFilterRuleMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteTrafficMirrorFilterRuleRequest{Request: req, Input: input, Copy: c.DeleteTrafficMirrorFilterRuleRequest}
}

// DeleteTrafficMirrorFilterRuleRequest is the request type for the
// DeleteTrafficMirrorFilterRule API operation.
type DeleteTrafficMirrorFilterRuleRequest struct {
	*aws.Request
	Input *types.DeleteTrafficMirrorFilterRuleInput
	Copy  func(*types.DeleteTrafficMirrorFilterRuleInput) DeleteTrafficMirrorFilterRuleRequest
}

// Send marshals and sends the DeleteTrafficMirrorFilterRule API request.
func (r DeleteTrafficMirrorFilterRuleRequest) Send(ctx context.Context) (*DeleteTrafficMirrorFilterRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTrafficMirrorFilterRuleResponse{
		DeleteTrafficMirrorFilterRuleOutput: r.Request.Data.(*types.DeleteTrafficMirrorFilterRuleOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTrafficMirrorFilterRuleResponse is the response type for the
// DeleteTrafficMirrorFilterRule API operation.
type DeleteTrafficMirrorFilterRuleResponse struct {
	*types.DeleteTrafficMirrorFilterRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTrafficMirrorFilterRule request.
func (r *DeleteTrafficMirrorFilterRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
