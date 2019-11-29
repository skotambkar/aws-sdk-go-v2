// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
)

const opDeleteVocabulary = "DeleteVocabulary"

// DeleteVocabularyRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Deletes a vocabulary from Amazon Transcribe.
//
//    // Example sending a request using DeleteVocabularyRequest.
//    req := client.DeleteVocabularyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/DeleteVocabulary
func (c *Client) DeleteVocabularyRequest(input *types.DeleteVocabularyInput) DeleteVocabularyRequest {
	op := &aws.Operation{
		Name:       opDeleteVocabulary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteVocabularyInput{}
	}

	req := c.newRequest(op, input, &types.DeleteVocabularyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteVocabularyMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVocabularyRequest{Request: req, Input: input, Copy: c.DeleteVocabularyRequest}
}

// DeleteVocabularyRequest is the request type for the
// DeleteVocabulary API operation.
type DeleteVocabularyRequest struct {
	*aws.Request
	Input *types.DeleteVocabularyInput
	Copy  func(*types.DeleteVocabularyInput) DeleteVocabularyRequest
}

// Send marshals and sends the DeleteVocabulary API request.
func (r DeleteVocabularyRequest) Send(ctx context.Context) (*DeleteVocabularyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVocabularyResponse{
		DeleteVocabularyOutput: r.Request.Data.(*types.DeleteVocabularyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVocabularyResponse is the response type for the
// DeleteVocabulary API operation.
type DeleteVocabularyResponse struct {
	*types.DeleteVocabularyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVocabulary request.
func (r *DeleteVocabularyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
