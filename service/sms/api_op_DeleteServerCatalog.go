// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
)

const opDeleteServerCatalog = "DeleteServerCatalog"

// DeleteServerCatalogRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Deletes all servers from your server catalog.
//
//    // Example sending a request using DeleteServerCatalogRequest.
//    req := client.DeleteServerCatalogRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/DeleteServerCatalog
func (c *Client) DeleteServerCatalogRequest(input *types.DeleteServerCatalogInput) DeleteServerCatalogRequest {
	op := &aws.Operation{
		Name:       opDeleteServerCatalog,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteServerCatalogInput{}
	}

	req := c.newRequest(op, input, &types.DeleteServerCatalogOutput{})
	return DeleteServerCatalogRequest{Request: req, Input: input, Copy: c.DeleteServerCatalogRequest}
}

// DeleteServerCatalogRequest is the request type for the
// DeleteServerCatalog API operation.
type DeleteServerCatalogRequest struct {
	*aws.Request
	Input *types.DeleteServerCatalogInput
	Copy  func(*types.DeleteServerCatalogInput) DeleteServerCatalogRequest
}

// Send marshals and sends the DeleteServerCatalog API request.
func (r DeleteServerCatalogRequest) Send(ctx context.Context) (*DeleteServerCatalogResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteServerCatalogResponse{
		DeleteServerCatalogOutput: r.Request.Data.(*types.DeleteServerCatalogOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteServerCatalogResponse is the response type for the
// DeleteServerCatalog API operation.
type DeleteServerCatalogResponse struct {
	*types.DeleteServerCatalogOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteServerCatalog request.
func (r *DeleteServerCatalogResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
