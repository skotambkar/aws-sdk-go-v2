// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opCreateProvisioningArtifact = "CreateProvisioningArtifact"

// CreateProvisioningArtifactRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Creates a provisioning artifact (also known as a version) for the specified
// product.
//
// You cannot create a provisioning artifact for a product that was shared with
// you.
//
//    // Example sending a request using CreateProvisioningArtifactRequest.
//    req := client.CreateProvisioningArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateProvisioningArtifact
func (c *Client) CreateProvisioningArtifactRequest(input *types.CreateProvisioningArtifactInput) CreateProvisioningArtifactRequest {
	op := &aws.Operation{
		Name:       opCreateProvisioningArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateProvisioningArtifactInput{}
	}

	req := c.newRequest(op, input, &types.CreateProvisioningArtifactOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateProvisioningArtifactMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateProvisioningArtifactRequest{Request: req, Input: input, Copy: c.CreateProvisioningArtifactRequest}
}

// CreateProvisioningArtifactRequest is the request type for the
// CreateProvisioningArtifact API operation.
type CreateProvisioningArtifactRequest struct {
	*aws.Request
	Input *types.CreateProvisioningArtifactInput
	Copy  func(*types.CreateProvisioningArtifactInput) CreateProvisioningArtifactRequest
}

// Send marshals and sends the CreateProvisioningArtifact API request.
func (r CreateProvisioningArtifactRequest) Send(ctx context.Context) (*CreateProvisioningArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProvisioningArtifactResponse{
		CreateProvisioningArtifactOutput: r.Request.Data.(*types.CreateProvisioningArtifactOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProvisioningArtifactResponse is the response type for the
// CreateProvisioningArtifact API operation.
type CreateProvisioningArtifactResponse struct {
	*types.CreateProvisioningArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProvisioningArtifact request.
func (r *CreateProvisioningArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
