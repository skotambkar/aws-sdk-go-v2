// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
)

const opListDeployments = "ListDeployments"

// ListDeploymentsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Returns a history of deployments for the group.
//
//    // Example sending a request using ListDeploymentsRequest.
//    req := client.ListDeploymentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListDeployments
func (c *Client) ListDeploymentsRequest(input *types.ListDeploymentsInput) ListDeploymentsRequest {
	op := &aws.Operation{
		Name:       opListDeployments,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/groups/{GroupId}/deployments",
	}

	if input == nil {
		input = &types.ListDeploymentsInput{}
	}

	req := c.newRequest(op, input, &types.ListDeploymentsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListDeploymentsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListDeploymentsRequest{Request: req, Input: input, Copy: c.ListDeploymentsRequest}
}

// ListDeploymentsRequest is the request type for the
// ListDeployments API operation.
type ListDeploymentsRequest struct {
	*aws.Request
	Input *types.ListDeploymentsInput
	Copy  func(*types.ListDeploymentsInput) ListDeploymentsRequest
}

// Send marshals and sends the ListDeployments API request.
func (r ListDeploymentsRequest) Send(ctx context.Context) (*ListDeploymentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeploymentsResponse{
		ListDeploymentsOutput: r.Request.Data.(*types.ListDeploymentsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDeploymentsResponse is the response type for the
// ListDeployments API operation.
type ListDeploymentsResponse struct {
	*types.ListDeploymentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeployments request.
func (r *ListDeploymentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
