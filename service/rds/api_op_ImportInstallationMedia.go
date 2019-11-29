// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opImportInstallationMedia = "ImportInstallationMedia"

// ImportInstallationMediaRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Imports the installation media for a DB engine that requires an on-premises
// customer provided license, such as SQL Server.
//
//    // Example sending a request using ImportInstallationMediaRequest.
//    req := client.ImportInstallationMediaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ImportInstallationMedia
func (c *Client) ImportInstallationMediaRequest(input *types.ImportInstallationMediaInput) ImportInstallationMediaRequest {
	op := &aws.Operation{
		Name:       opImportInstallationMedia,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ImportInstallationMediaInput{}
	}

	req := c.newRequest(op, input, &types.ImportInstallationMediaOutput{})
	return ImportInstallationMediaRequest{Request: req, Input: input, Copy: c.ImportInstallationMediaRequest}
}

// ImportInstallationMediaRequest is the request type for the
// ImportInstallationMedia API operation.
type ImportInstallationMediaRequest struct {
	*aws.Request
	Input *types.ImportInstallationMediaInput
	Copy  func(*types.ImportInstallationMediaInput) ImportInstallationMediaRequest
}

// Send marshals and sends the ImportInstallationMedia API request.
func (r ImportInstallationMediaRequest) Send(ctx context.Context) (*ImportInstallationMediaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportInstallationMediaResponse{
		ImportInstallationMediaOutput: r.Request.Data.(*types.ImportInstallationMediaOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportInstallationMediaResponse is the response type for the
// ImportInstallationMedia API operation.
type ImportInstallationMediaResponse struct {
	*types.ImportInstallationMediaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportInstallationMedia request.
func (r *ImportInstallationMediaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
