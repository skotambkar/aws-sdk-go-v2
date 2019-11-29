// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opBatchAssociateServiceActionWithProvisioningArtifact = "BatchAssociateServiceActionWithProvisioningArtifact"

// BatchAssociateServiceActionWithProvisioningArtifactRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Associates multiple self-service actions with provisioning artifacts.
//
//    // Example sending a request using BatchAssociateServiceActionWithProvisioningArtifactRequest.
//    req := client.BatchAssociateServiceActionWithProvisioningArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/BatchAssociateServiceActionWithProvisioningArtifact
func (c *Client) BatchAssociateServiceActionWithProvisioningArtifactRequest(input *types.BatchAssociateServiceActionWithProvisioningArtifactInput) BatchAssociateServiceActionWithProvisioningArtifactRequest {
	op := &aws.Operation{
		Name:       opBatchAssociateServiceActionWithProvisioningArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchAssociateServiceActionWithProvisioningArtifactInput{}
	}

	req := c.newRequest(op, input, &types.BatchAssociateServiceActionWithProvisioningArtifactOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.BatchAssociateServiceActionWithProvisioningArtifactMarshaler{Input: input}.GetNamedBuildHandler())

	return BatchAssociateServiceActionWithProvisioningArtifactRequest{Request: req, Input: input, Copy: c.BatchAssociateServiceActionWithProvisioningArtifactRequest}
}

// BatchAssociateServiceActionWithProvisioningArtifactRequest is the request type for the
// BatchAssociateServiceActionWithProvisioningArtifact API operation.
type BatchAssociateServiceActionWithProvisioningArtifactRequest struct {
	*aws.Request
	Input *types.BatchAssociateServiceActionWithProvisioningArtifactInput
	Copy  func(*types.BatchAssociateServiceActionWithProvisioningArtifactInput) BatchAssociateServiceActionWithProvisioningArtifactRequest
}

// Send marshals and sends the BatchAssociateServiceActionWithProvisioningArtifact API request.
func (r BatchAssociateServiceActionWithProvisioningArtifactRequest) Send(ctx context.Context) (*BatchAssociateServiceActionWithProvisioningArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchAssociateServiceActionWithProvisioningArtifactResponse{
		BatchAssociateServiceActionWithProvisioningArtifactOutput: r.Request.Data.(*types.BatchAssociateServiceActionWithProvisioningArtifactOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchAssociateServiceActionWithProvisioningArtifactResponse is the response type for the
// BatchAssociateServiceActionWithProvisioningArtifact API operation.
type BatchAssociateServiceActionWithProvisioningArtifactResponse struct {
	*types.BatchAssociateServiceActionWithProvisioningArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchAssociateServiceActionWithProvisioningArtifact request.
func (r *BatchAssociateServiceActionWithProvisioningArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
