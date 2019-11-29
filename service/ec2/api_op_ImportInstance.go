// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opImportInstance = "ImportInstance"

// ImportInstanceRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates an import instance task using metadata from the specified disk image.
// ImportInstance only supports single-volume VMs. To import multi-volume VMs,
// use ImportImage. For more information, see Importing a Virtual Machine Using
// the Amazon EC2 CLI (https://docs.aws.amazon.com/AWSEC2/latest/CommandLineReference/ec2-cli-vmimport-export.html).
//
// For information about the import manifest referenced by this API action,
// see VM Import Manifest (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html).
//
//    // Example sending a request using ImportInstanceRequest.
//    req := client.ImportInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ImportInstance
func (c *Client) ImportInstanceRequest(input *types.ImportInstanceInput) ImportInstanceRequest {
	op := &aws.Operation{
		Name:       opImportInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ImportInstanceInput{}
	}

	req := c.newRequest(op, input, &types.ImportInstanceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.ImportInstanceMarshaler{Input: input}.GetNamedBuildHandler())

	return ImportInstanceRequest{Request: req, Input: input, Copy: c.ImportInstanceRequest}
}

// ImportInstanceRequest is the request type for the
// ImportInstance API operation.
type ImportInstanceRequest struct {
	*aws.Request
	Input *types.ImportInstanceInput
	Copy  func(*types.ImportInstanceInput) ImportInstanceRequest
}

// Send marshals and sends the ImportInstance API request.
func (r ImportInstanceRequest) Send(ctx context.Context) (*ImportInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportInstanceResponse{
		ImportInstanceOutput: r.Request.Data.(*types.ImportInstanceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportInstanceResponse is the response type for the
// ImportInstance API operation.
type ImportInstanceResponse struct {
	*types.ImportInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportInstance request.
func (r *ImportInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
