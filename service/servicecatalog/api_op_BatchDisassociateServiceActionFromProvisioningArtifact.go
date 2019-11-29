// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opBatchDisassociateServiceActionFromProvisioningArtifact = "BatchDisassociateServiceActionFromProvisioningArtifact"

// BatchDisassociateServiceActionFromProvisioningArtifactRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Disassociates a batch of self-service actions from the specified provisioning
// artifact.
//
//    // Example sending a request using BatchDisassociateServiceActionFromProvisioningArtifactRequest.
//    req := client.BatchDisassociateServiceActionFromProvisioningArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/BatchDisassociateServiceActionFromProvisioningArtifact
func (c *Client) BatchDisassociateServiceActionFromProvisioningArtifactRequest(input *types.BatchDisassociateServiceActionFromProvisioningArtifactInput) BatchDisassociateServiceActionFromProvisioningArtifactRequest {
	op := &aws.Operation{
		Name:       opBatchDisassociateServiceActionFromProvisioningArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchDisassociateServiceActionFromProvisioningArtifactInput{}
	}

	req := c.newRequest(op, input, &types.BatchDisassociateServiceActionFromProvisioningArtifactOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.BatchDisassociateServiceActionFromProvisioningArtifactMarshaler{Input: input}.GetNamedBuildHandler())

	return BatchDisassociateServiceActionFromProvisioningArtifactRequest{Request: req, Input: input, Copy: c.BatchDisassociateServiceActionFromProvisioningArtifactRequest}
}

// BatchDisassociateServiceActionFromProvisioningArtifactRequest is the request type for the
// BatchDisassociateServiceActionFromProvisioningArtifact API operation.
type BatchDisassociateServiceActionFromProvisioningArtifactRequest struct {
	*aws.Request
	Input *types.BatchDisassociateServiceActionFromProvisioningArtifactInput
	Copy  func(*types.BatchDisassociateServiceActionFromProvisioningArtifactInput) BatchDisassociateServiceActionFromProvisioningArtifactRequest
}

// Send marshals and sends the BatchDisassociateServiceActionFromProvisioningArtifact API request.
func (r BatchDisassociateServiceActionFromProvisioningArtifactRequest) Send(ctx context.Context) (*BatchDisassociateServiceActionFromProvisioningArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDisassociateServiceActionFromProvisioningArtifactResponse{
		BatchDisassociateServiceActionFromProvisioningArtifactOutput: r.Request.Data.(*types.BatchDisassociateServiceActionFromProvisioningArtifactOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDisassociateServiceActionFromProvisioningArtifactResponse is the response type for the
// BatchDisassociateServiceActionFromProvisioningArtifact API operation.
type BatchDisassociateServiceActionFromProvisioningArtifactResponse struct {
	*types.BatchDisassociateServiceActionFromProvisioningArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDisassociateServiceActionFromProvisioningArtifact request.
func (r *BatchDisassociateServiceActionFromProvisioningArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
