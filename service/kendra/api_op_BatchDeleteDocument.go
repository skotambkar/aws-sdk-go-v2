// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kendra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDeleteDocumentInput struct {
	_ struct{} `type:"structure"`

	// One or more identifiers for documents to delete from the index.
	//
	// DocumentIdList is a required field
	DocumentIdList []string `min:"1" type:"list" required:"true"`

	// The identifier of the index that contains the documents to delete.
	//
	// IndexId is a required field
	IndexId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s BatchDeleteDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDeleteDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDeleteDocumentInput"}

	if s.DocumentIdList == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentIdList"))
	}
	if s.DocumentIdList != nil && len(s.DocumentIdList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentIdList", 1))
	}

	if s.IndexId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IndexId"))
	}
	if s.IndexId != nil && len(*s.IndexId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDeleteDocumentOutput struct {
	_ struct{} `type:"structure"`

	// A list of documents that could not be removed from the index. Each entry
	// contains an error message that indicates why the document couldn't be removed
	// from the index.
	FailedDocuments []BatchDeleteDocumentResponseFailedDocument `type:"list"`
}

// String returns the string representation
func (s BatchDeleteDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDeleteDocument = "BatchDeleteDocument"

// BatchDeleteDocumentRequest returns a request value for making API operation for
// AWSKendraFrontendService.
//
// Removes one or more documents from an index. The documents must have been
// added with the BatchPutDocument operation.
//
// The documents are deleted asynchronously. You can see the progress of the
// deletion by using AWS CloudWatch. Any error messages releated to the processing
// of the batch are sent to you CloudWatch log.
//
//    // Example sending a request using BatchDeleteDocumentRequest.
//    req := client.BatchDeleteDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kendra-2019-02-03/BatchDeleteDocument
func (c *Client) BatchDeleteDocumentRequest(input *BatchDeleteDocumentInput) BatchDeleteDocumentRequest {
	op := &aws.Operation{
		Name:       opBatchDeleteDocument,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDeleteDocumentInput{}
	}

	req := c.newRequest(op, input, &BatchDeleteDocumentOutput{})
	return BatchDeleteDocumentRequest{Request: req, Input: input, Copy: c.BatchDeleteDocumentRequest}
}

// BatchDeleteDocumentRequest is the request type for the
// BatchDeleteDocument API operation.
type BatchDeleteDocumentRequest struct {
	*aws.Request
	Input *BatchDeleteDocumentInput
	Copy  func(*BatchDeleteDocumentInput) BatchDeleteDocumentRequest
}

// Send marshals and sends the BatchDeleteDocument API request.
func (r BatchDeleteDocumentRequest) Send(ctx context.Context) (*BatchDeleteDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDeleteDocumentResponse{
		BatchDeleteDocumentOutput: r.Request.Data.(*BatchDeleteDocumentOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDeleteDocumentResponse is the response type for the
// BatchDeleteDocument API operation.
type BatchDeleteDocumentResponse struct {
	*BatchDeleteDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDeleteDocument request.
func (r *BatchDeleteDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
