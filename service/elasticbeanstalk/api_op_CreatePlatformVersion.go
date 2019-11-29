// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
)

const opCreatePlatformVersion = "CreatePlatformVersion"

// CreatePlatformVersionRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Create a new version of your custom platform.
//
//    // Example sending a request using CreatePlatformVersionRequest.
//    req := client.CreatePlatformVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CreatePlatformVersion
func (c *Client) CreatePlatformVersionRequest(input *types.CreatePlatformVersionInput) CreatePlatformVersionRequest {
	op := &aws.Operation{
		Name:       opCreatePlatformVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreatePlatformVersionInput{}
	}

	req := c.newRequest(op, input, &types.CreatePlatformVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreatePlatformVersionMarshaler{Input: input}.GetNamedBuildHandler())

	return CreatePlatformVersionRequest{Request: req, Input: input, Copy: c.CreatePlatformVersionRequest}
}

// CreatePlatformVersionRequest is the request type for the
// CreatePlatformVersion API operation.
type CreatePlatformVersionRequest struct {
	*aws.Request
	Input *types.CreatePlatformVersionInput
	Copy  func(*types.CreatePlatformVersionInput) CreatePlatformVersionRequest
}

// Send marshals and sends the CreatePlatformVersion API request.
func (r CreatePlatformVersionRequest) Send(ctx context.Context) (*CreatePlatformVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePlatformVersionResponse{
		CreatePlatformVersionOutput: r.Request.Data.(*types.CreatePlatformVersionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePlatformVersionResponse is the response type for the
// CreatePlatformVersion API operation.
type CreatePlatformVersionResponse struct {
	*types.CreatePlatformVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePlatformVersion request.
func (r *CreatePlatformVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
