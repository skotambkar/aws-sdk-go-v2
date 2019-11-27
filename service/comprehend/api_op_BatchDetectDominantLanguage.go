// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

const opBatchDetectDominantLanguage = "BatchDetectDominantLanguage"

// BatchDetectDominantLanguageRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Determines the dominant language of the input text for a batch of documents.
// For a list of languages that Amazon Comprehend can detect, see Amazon Comprehend
// Supported Languages (https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html).
//
//    // Example sending a request using BatchDetectDominantLanguageRequest.
//    req := client.BatchDetectDominantLanguageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectDominantLanguage
func (c *Client) BatchDetectDominantLanguageRequest(input *types.BatchDetectDominantLanguageInput) BatchDetectDominantLanguageRequest {
	op := &aws.Operation{
		Name:       opBatchDetectDominantLanguage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchDetectDominantLanguageInput{}
	}

	req := c.newRequest(op, input, &types.BatchDetectDominantLanguageOutput{})
	return BatchDetectDominantLanguageRequest{Request: req, Input: input, Copy: c.BatchDetectDominantLanguageRequest}
}

// BatchDetectDominantLanguageRequest is the request type for the
// BatchDetectDominantLanguage API operation.
type BatchDetectDominantLanguageRequest struct {
	*aws.Request
	Input *types.BatchDetectDominantLanguageInput
	Copy  func(*types.BatchDetectDominantLanguageInput) BatchDetectDominantLanguageRequest
}

// Send marshals and sends the BatchDetectDominantLanguage API request.
func (r BatchDetectDominantLanguageRequest) Send(ctx context.Context) (*BatchDetectDominantLanguageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDetectDominantLanguageResponse{
		BatchDetectDominantLanguageOutput: r.Request.Data.(*types.BatchDetectDominantLanguageOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDetectDominantLanguageResponse is the response type for the
// BatchDetectDominantLanguage API operation.
type BatchDetectDominantLanguageResponse struct {
	*types.BatchDetectDominantLanguageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDetectDominantLanguage request.
func (r *BatchDetectDominantLanguageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
