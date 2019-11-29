// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
)

const opUpdateDocumentVersion = "UpdateDocumentVersion"

// UpdateDocumentVersionRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Changes the status of the document version to ACTIVE.
//
// Amazon WorkDocs also sets its document container to ACTIVE. This is the last
// step in a document upload, after the client uploads the document to an S3-presigned
// URL returned by InitiateDocumentVersionUpload.
//
//    // Example sending a request using UpdateDocumentVersionRequest.
//    req := client.UpdateDocumentVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/UpdateDocumentVersion
func (c *Client) UpdateDocumentVersionRequest(input *types.UpdateDocumentVersionInput) UpdateDocumentVersionRequest {
	op := &aws.Operation{
		Name:       opUpdateDocumentVersion,
		HTTPMethod: "PATCH",
		HTTPPath:   "/api/v1/documents/{DocumentId}/versions/{VersionId}",
	}

	if input == nil {
		input = &types.UpdateDocumentVersionInput{}
	}

	req := c.newRequest(op, input, &types.UpdateDocumentVersionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateDocumentVersionRequest{Request: req, Input: input, Copy: c.UpdateDocumentVersionRequest}
}

// UpdateDocumentVersionRequest is the request type for the
// UpdateDocumentVersion API operation.
type UpdateDocumentVersionRequest struct {
	*aws.Request
	Input *types.UpdateDocumentVersionInput
	Copy  func(*types.UpdateDocumentVersionInput) UpdateDocumentVersionRequest
}

// Send marshals and sends the UpdateDocumentVersion API request.
func (r UpdateDocumentVersionRequest) Send(ctx context.Context) (*UpdateDocumentVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDocumentVersionResponse{
		UpdateDocumentVersionOutput: r.Request.Data.(*types.UpdateDocumentVersionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDocumentVersionResponse is the response type for the
// UpdateDocumentVersion API operation.
type UpdateDocumentVersionResponse struct {
	*types.UpdateDocumentVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDocumentVersion request.
func (r *UpdateDocumentVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
