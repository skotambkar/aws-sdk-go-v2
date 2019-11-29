// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

const opBatchDetectSyntax = "BatchDetectSyntax"

// BatchDetectSyntaxRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Inspects the text of a batch of documents for the syntax and part of speech
// of the words in the document and returns information about them. For more
// information, see how-syntax.
//
//    // Example sending a request using BatchDetectSyntaxRequest.
//    req := client.BatchDetectSyntaxRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectSyntax
func (c *Client) BatchDetectSyntaxRequest(input *types.BatchDetectSyntaxInput) BatchDetectSyntaxRequest {
	op := &aws.Operation{
		Name:       opBatchDetectSyntax,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchDetectSyntaxInput{}
	}

	req := c.newRequest(op, input, &types.BatchDetectSyntaxOutput{})
	return BatchDetectSyntaxRequest{Request: req, Input: input, Copy: c.BatchDetectSyntaxRequest}
}

// BatchDetectSyntaxRequest is the request type for the
// BatchDetectSyntax API operation.
type BatchDetectSyntaxRequest struct {
	*aws.Request
	Input *types.BatchDetectSyntaxInput
	Copy  func(*types.BatchDetectSyntaxInput) BatchDetectSyntaxRequest
}

// Send marshals and sends the BatchDetectSyntax API request.
func (r BatchDetectSyntaxRequest) Send(ctx context.Context) (*BatchDetectSyntaxResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDetectSyntaxResponse{
		BatchDetectSyntaxOutput: r.Request.Data.(*types.BatchDetectSyntaxOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDetectSyntaxResponse is the response type for the
// BatchDetectSyntax API operation.
type BatchDetectSyntaxResponse struct {
	*types.BatchDetectSyntaxOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDetectSyntax request.
func (r *BatchDetectSyntaxResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
