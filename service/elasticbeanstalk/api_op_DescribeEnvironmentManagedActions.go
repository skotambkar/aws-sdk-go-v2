// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
)

const opDescribeEnvironmentManagedActions = "DescribeEnvironmentManagedActions"

// DescribeEnvironmentManagedActionsRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Lists an environment's upcoming and in-progress managed actions.
//
//    // Example sending a request using DescribeEnvironmentManagedActionsRequest.
//    req := client.DescribeEnvironmentManagedActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeEnvironmentManagedActions
func (c *Client) DescribeEnvironmentManagedActionsRequest(input *types.DescribeEnvironmentManagedActionsInput) DescribeEnvironmentManagedActionsRequest {
	op := &aws.Operation{
		Name:       opDescribeEnvironmentManagedActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeEnvironmentManagedActionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEnvironmentManagedActionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DescribeEnvironmentManagedActionsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeEnvironmentManagedActionsRequest{Request: req, Input: input, Copy: c.DescribeEnvironmentManagedActionsRequest}
}

// DescribeEnvironmentManagedActionsRequest is the request type for the
// DescribeEnvironmentManagedActions API operation.
type DescribeEnvironmentManagedActionsRequest struct {
	*aws.Request
	Input *types.DescribeEnvironmentManagedActionsInput
	Copy  func(*types.DescribeEnvironmentManagedActionsInput) DescribeEnvironmentManagedActionsRequest
}

// Send marshals and sends the DescribeEnvironmentManagedActions API request.
func (r DescribeEnvironmentManagedActionsRequest) Send(ctx context.Context) (*DescribeEnvironmentManagedActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEnvironmentManagedActionsResponse{
		DescribeEnvironmentManagedActionsOutput: r.Request.Data.(*types.DescribeEnvironmentManagedActionsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEnvironmentManagedActionsResponse is the response type for the
// DescribeEnvironmentManagedActions API operation.
type DescribeEnvironmentManagedActionsResponse struct {
	*types.DescribeEnvironmentManagedActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEnvironmentManagedActions request.
func (r *DescribeEnvironmentManagedActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
