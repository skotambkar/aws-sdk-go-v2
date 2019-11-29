// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opCreateDocument = "CreateDocument"

// CreateDocumentRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Creates a Systems Manager document.
//
// After you create a document, you can use CreateAssociation to associate it
// with one or more running instances.
//
//    // Example sending a request using CreateDocumentRequest.
//    req := client.CreateDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/CreateDocument
func (c *Client) CreateDocumentRequest(input *types.CreateDocumentInput) CreateDocumentRequest {
	op := &aws.Operation{
		Name:       opCreateDocument,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDocumentInput{}
	}

	req := c.newRequest(op, input, &types.CreateDocumentOutput{})
	return CreateDocumentRequest{Request: req, Input: input, Copy: c.CreateDocumentRequest}
}

// CreateDocumentRequest is the request type for the
// CreateDocument API operation.
type CreateDocumentRequest struct {
	*aws.Request
	Input *types.CreateDocumentInput
	Copy  func(*types.CreateDocumentInput) CreateDocumentRequest
}

// Send marshals and sends the CreateDocument API request.
func (r CreateDocumentRequest) Send(ctx context.Context) (*CreateDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDocumentResponse{
		CreateDocumentOutput: r.Request.Data.(*types.CreateDocumentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDocumentResponse is the response type for the
// CreateDocument API operation.
type CreateDocumentResponse struct {
	*types.CreateDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDocument request.
func (r *CreateDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
