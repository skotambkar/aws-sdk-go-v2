// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/polly/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
)

const opGetLexicon = "GetLexicon"

// GetLexiconRequest returns a request value for making API operation for
// Amazon Polly.
//
// Returns the content of the specified pronunciation lexicon stored in an AWS
// Region. For more information, see Managing Lexicons (https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html).
//
//    // Example sending a request using GetLexiconRequest.
//    req := client.GetLexiconRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/polly-2016-06-10/GetLexicon
func (c *Client) GetLexiconRequest(input *types.GetLexiconInput) GetLexiconRequest {
	op := &aws.Operation{
		Name:       opGetLexicon,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/lexicons/{LexiconName}",
	}

	if input == nil {
		input = &types.GetLexiconInput{}
	}

	req := c.newRequest(op, input, &types.GetLexiconOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetLexiconMarshaler{Input: input}.GetNamedBuildHandler())

	return GetLexiconRequest{Request: req, Input: input, Copy: c.GetLexiconRequest}
}

// GetLexiconRequest is the request type for the
// GetLexicon API operation.
type GetLexiconRequest struct {
	*aws.Request
	Input *types.GetLexiconInput
	Copy  func(*types.GetLexiconInput) GetLexiconRequest
}

// Send marshals and sends the GetLexicon API request.
func (r GetLexiconRequest) Send(ctx context.Context) (*GetLexiconResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLexiconResponse{
		GetLexiconOutput: r.Request.Data.(*types.GetLexiconOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLexiconResponse is the response type for the
// GetLexicon API operation.
type GetLexiconResponse struct {
	*types.GetLexiconOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLexicon request.
func (r *GetLexiconResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
