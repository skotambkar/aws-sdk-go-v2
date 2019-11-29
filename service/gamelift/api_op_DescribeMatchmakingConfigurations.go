// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
)

const opDescribeMatchmakingConfigurations = "DescribeMatchmakingConfigurations"

// DescribeMatchmakingConfigurationsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves the details of FlexMatch matchmaking configurations. With this
// operation, you have the following options: (1) retrieve all existing configurations,
// (2) provide the names of one or more configurations to retrieve, or (3) retrieve
// all configurations that use a specified rule set name. When requesting multiple
// items, use the pagination parameters to retrieve results as a set of sequential
// pages. If successful, a configuration is returned for each requested name.
// When specifying a list of names, only configurations that currently exist
// are returned.
//
// Learn more
//
//  Setting Up FlexMatch Matchmakers (https://docs.aws.amazon.com/gamelift/latest/developerguide/matchmaker-build.html)
//
// Related operations
//
//    * CreateMatchmakingConfiguration
//
//    * DescribeMatchmakingConfigurations
//
//    * UpdateMatchmakingConfiguration
//
//    * DeleteMatchmakingConfiguration
//
//    * CreateMatchmakingRuleSet
//
//    * DescribeMatchmakingRuleSets
//
//    * ValidateMatchmakingRuleSet
//
//    * DeleteMatchmakingRuleSet
//
//    // Example sending a request using DescribeMatchmakingConfigurationsRequest.
//    req := client.DescribeMatchmakingConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeMatchmakingConfigurations
func (c *Client) DescribeMatchmakingConfigurationsRequest(input *types.DescribeMatchmakingConfigurationsInput) DescribeMatchmakingConfigurationsRequest {
	op := &aws.Operation{
		Name:       opDescribeMatchmakingConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeMatchmakingConfigurationsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeMatchmakingConfigurationsOutput{})
	return DescribeMatchmakingConfigurationsRequest{Request: req, Input: input, Copy: c.DescribeMatchmakingConfigurationsRequest}
}

// DescribeMatchmakingConfigurationsRequest is the request type for the
// DescribeMatchmakingConfigurations API operation.
type DescribeMatchmakingConfigurationsRequest struct {
	*aws.Request
	Input *types.DescribeMatchmakingConfigurationsInput
	Copy  func(*types.DescribeMatchmakingConfigurationsInput) DescribeMatchmakingConfigurationsRequest
}

// Send marshals and sends the DescribeMatchmakingConfigurations API request.
func (r DescribeMatchmakingConfigurationsRequest) Send(ctx context.Context) (*DescribeMatchmakingConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMatchmakingConfigurationsResponse{
		DescribeMatchmakingConfigurationsOutput: r.Request.Data.(*types.DescribeMatchmakingConfigurationsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMatchmakingConfigurationsResponse is the response type for the
// DescribeMatchmakingConfigurations API operation.
type DescribeMatchmakingConfigurationsResponse struct {
	*types.DescribeMatchmakingConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMatchmakingConfigurations request.
func (r *DescribeMatchmakingConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
