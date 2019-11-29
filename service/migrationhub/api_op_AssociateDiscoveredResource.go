// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
)

const opAssociateDiscoveredResource = "AssociateDiscoveredResource"

// AssociateDiscoveredResourceRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Associates a discovered resource ID from Application Discovery Service with
// a migration task.
//
//    // Example sending a request using AssociateDiscoveredResourceRequest.
//    req := client.AssociateDiscoveredResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/AssociateDiscoveredResource
func (c *Client) AssociateDiscoveredResourceRequest(input *types.AssociateDiscoveredResourceInput) AssociateDiscoveredResourceRequest {
	op := &aws.Operation{
		Name:       opAssociateDiscoveredResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateDiscoveredResourceInput{}
	}

	req := c.newRequest(op, input, &types.AssociateDiscoveredResourceOutput{})
	return AssociateDiscoveredResourceRequest{Request: req, Input: input, Copy: c.AssociateDiscoveredResourceRequest}
}

// AssociateDiscoveredResourceRequest is the request type for the
// AssociateDiscoveredResource API operation.
type AssociateDiscoveredResourceRequest struct {
	*aws.Request
	Input *types.AssociateDiscoveredResourceInput
	Copy  func(*types.AssociateDiscoveredResourceInput) AssociateDiscoveredResourceRequest
}

// Send marshals and sends the AssociateDiscoveredResource API request.
func (r AssociateDiscoveredResourceRequest) Send(ctx context.Context) (*AssociateDiscoveredResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateDiscoveredResourceResponse{
		AssociateDiscoveredResourceOutput: r.Request.Data.(*types.AssociateDiscoveredResourceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateDiscoveredResourceResponse is the response type for the
// AssociateDiscoveredResource API operation.
type AssociateDiscoveredResourceResponse struct {
	*types.AssociateDiscoveredResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateDiscoveredResource request.
func (r *AssociateDiscoveredResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
