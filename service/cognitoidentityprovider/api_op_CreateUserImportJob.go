// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opCreateUserImportJob = "CreateUserImportJob"

// CreateUserImportJobRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Creates the user import job.
//
//    // Example sending a request using CreateUserImportJobRequest.
//    req := client.CreateUserImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/CreateUserImportJob
func (c *Client) CreateUserImportJobRequest(input *types.CreateUserImportJobInput) CreateUserImportJobRequest {
	op := &aws.Operation{
		Name:       opCreateUserImportJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateUserImportJobInput{}
	}

	req := c.newRequest(op, input, &types.CreateUserImportJobOutput{})
	return CreateUserImportJobRequest{Request: req, Input: input, Copy: c.CreateUserImportJobRequest}
}

// CreateUserImportJobRequest is the request type for the
// CreateUserImportJob API operation.
type CreateUserImportJobRequest struct {
	*aws.Request
	Input *types.CreateUserImportJobInput
	Copy  func(*types.CreateUserImportJobInput) CreateUserImportJobRequest
}

// Send marshals and sends the CreateUserImportJob API request.
func (r CreateUserImportJobRequest) Send(ctx context.Context) (*CreateUserImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserImportJobResponse{
		CreateUserImportJobOutput: r.Request.Data.(*types.CreateUserImportJobOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserImportJobResponse is the response type for the
// CreateUserImportJob API operation.
type CreateUserImportJobResponse struct {
	*types.CreateUserImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUserImportJob request.
func (r *CreateUserImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
