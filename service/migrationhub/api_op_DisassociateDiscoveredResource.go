// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
)

const opDisassociateDiscoveredResource = "DisassociateDiscoveredResource"

// DisassociateDiscoveredResourceRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Disassociate an Application Discovery Service discovered resource from a
// migration task.
//
//    // Example sending a request using DisassociateDiscoveredResourceRequest.
//    req := client.DisassociateDiscoveredResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/DisassociateDiscoveredResource
func (c *Client) DisassociateDiscoveredResourceRequest(input *types.DisassociateDiscoveredResourceInput) DisassociateDiscoveredResourceRequest {
	op := &aws.Operation{
		Name:       opDisassociateDiscoveredResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisassociateDiscoveredResourceInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateDiscoveredResourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DisassociateDiscoveredResourceMarshaler{Input: input}.GetNamedBuildHandler())

	return DisassociateDiscoveredResourceRequest{Request: req, Input: input, Copy: c.DisassociateDiscoveredResourceRequest}
}

// DisassociateDiscoveredResourceRequest is the request type for the
// DisassociateDiscoveredResource API operation.
type DisassociateDiscoveredResourceRequest struct {
	*aws.Request
	Input *types.DisassociateDiscoveredResourceInput
	Copy  func(*types.DisassociateDiscoveredResourceInput) DisassociateDiscoveredResourceRequest
}

// Send marshals and sends the DisassociateDiscoveredResource API request.
func (r DisassociateDiscoveredResourceRequest) Send(ctx context.Context) (*DisassociateDiscoveredResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateDiscoveredResourceResponse{
		DisassociateDiscoveredResourceOutput: r.Request.Data.(*types.DisassociateDiscoveredResourceOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateDiscoveredResourceResponse is the response type for the
// DisassociateDiscoveredResource API operation.
type DisassociateDiscoveredResourceResponse struct {
	*types.DisassociateDiscoveredResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateDiscoveredResource request.
func (r *DisassociateDiscoveredResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
