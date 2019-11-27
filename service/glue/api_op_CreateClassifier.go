// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opCreateClassifier = "CreateClassifier"

// CreateClassifierRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a classifier in the user's account. This can be a GrokClassifier,
// an XMLClassifier, a JsonClassifier, or a CsvClassifier, depending on which
// field of the request is present.
//
//    // Example sending a request using CreateClassifierRequest.
//    req := client.CreateClassifierRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateClassifier
func (c *Client) CreateClassifierRequest(input *types.CreateClassifierInput) CreateClassifierRequest {
	op := &aws.Operation{
		Name:       opCreateClassifier,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateClassifierInput{}
	}

	req := c.newRequest(op, input, &types.CreateClassifierOutput{})
	return CreateClassifierRequest{Request: req, Input: input, Copy: c.CreateClassifierRequest}
}

// CreateClassifierRequest is the request type for the
// CreateClassifier API operation.
type CreateClassifierRequest struct {
	*aws.Request
	Input *types.CreateClassifierInput
	Copy  func(*types.CreateClassifierInput) CreateClassifierRequest
}

// Send marshals and sends the CreateClassifier API request.
func (r CreateClassifierRequest) Send(ctx context.Context) (*CreateClassifierResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClassifierResponse{
		CreateClassifierOutput: r.Request.Data.(*types.CreateClassifierOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClassifierResponse is the response type for the
// CreateClassifier API operation.
type CreateClassifierResponse struct {
	*types.CreateClassifierOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateClassifier request.
func (r *CreateClassifierResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
