// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opDeleteCertificate = "DeleteCertificate"

// DeleteCertificateRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Deletes the specified certificate.
//
//    // Example sending a request using DeleteCertificateRequest.
//    req := client.DeleteCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DeleteCertificate
func (c *Client) DeleteCertificateRequest(input *types.DeleteCertificateInput) DeleteCertificateRequest {
	op := &aws.Operation{
		Name:       opDeleteCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteCertificateInput{}
	}

	req := c.newRequest(op, input, &types.DeleteCertificateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteCertificateMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteCertificateRequest{Request: req, Input: input, Copy: c.DeleteCertificateRequest}
}

// DeleteCertificateRequest is the request type for the
// DeleteCertificate API operation.
type DeleteCertificateRequest struct {
	*aws.Request
	Input *types.DeleteCertificateInput
	Copy  func(*types.DeleteCertificateInput) DeleteCertificateRequest
}

// Send marshals and sends the DeleteCertificate API request.
func (r DeleteCertificateRequest) Send(ctx context.Context) (*DeleteCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCertificateResponse{
		DeleteCertificateOutput: r.Request.Data.(*types.DeleteCertificateOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCertificateResponse is the response type for the
// DeleteCertificate API operation.
type DeleteCertificateResponse struct {
	*types.DeleteCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCertificate request.
func (r *DeleteCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
