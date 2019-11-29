// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

const opDetectDominantLanguage = "DetectDominantLanguage"

// DetectDominantLanguageRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Determines the dominant language of the input text. For a list of languages
// that Amazon Comprehend can detect, see Amazon Comprehend Supported Languages
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html).
//
//    // Example sending a request using DetectDominantLanguageRequest.
//    req := client.DetectDominantLanguageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DetectDominantLanguage
func (c *Client) DetectDominantLanguageRequest(input *types.DetectDominantLanguageInput) DetectDominantLanguageRequest {
	op := &aws.Operation{
		Name:       opDetectDominantLanguage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DetectDominantLanguageInput{}
	}

	req := c.newRequest(op, input, &types.DetectDominantLanguageOutput{})
	return DetectDominantLanguageRequest{Request: req, Input: input, Copy: c.DetectDominantLanguageRequest}
}

// DetectDominantLanguageRequest is the request type for the
// DetectDominantLanguage API operation.
type DetectDominantLanguageRequest struct {
	*aws.Request
	Input *types.DetectDominantLanguageInput
	Copy  func(*types.DetectDominantLanguageInput) DetectDominantLanguageRequest
}

// Send marshals and sends the DetectDominantLanguage API request.
func (r DetectDominantLanguageRequest) Send(ctx context.Context) (*DetectDominantLanguageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetectDominantLanguageResponse{
		DetectDominantLanguageOutput: r.Request.Data.(*types.DetectDominantLanguageOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetectDominantLanguageResponse is the response type for the
// DetectDominantLanguage API operation.
type DetectDominantLanguageResponse struct {
	*types.DetectDominantLanguageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetectDominantLanguage request.
func (r *DetectDominantLanguageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
