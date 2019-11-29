// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
)

const opCreateApplicationVersion = "CreateApplicationVersion"

// CreateApplicationVersionRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Creates an application version for the specified application. You can create
// an application version from a source bundle in Amazon S3, a commit in AWS
// CodeCommit, or the output of an AWS CodeBuild build as follows:
//
// Specify a commit in an AWS CodeCommit repository with SourceBuildInformation.
//
// Specify a build in an AWS CodeBuild with SourceBuildInformation and BuildConfiguration.
//
// Specify a source bundle in S3 with SourceBundle
//
// Omit both SourceBuildInformation and SourceBundle to use the default sample
// application.
//
// Once you create an application version with a specified Amazon S3 bucket
// and key location, you cannot change that Amazon S3 location. If you change
// the Amazon S3 location, you receive an exception when you attempt to launch
// an environment from the application version.
//
//    // Example sending a request using CreateApplicationVersionRequest.
//    req := client.CreateApplicationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CreateApplicationVersion
func (c *Client) CreateApplicationVersionRequest(input *types.CreateApplicationVersionInput) CreateApplicationVersionRequest {
	op := &aws.Operation{
		Name:       opCreateApplicationVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateApplicationVersionInput{}
	}

	req := c.newRequest(op, input, &types.CreateApplicationVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateApplicationVersionMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateApplicationVersionRequest{Request: req, Input: input, Copy: c.CreateApplicationVersionRequest}
}

// CreateApplicationVersionRequest is the request type for the
// CreateApplicationVersion API operation.
type CreateApplicationVersionRequest struct {
	*aws.Request
	Input *types.CreateApplicationVersionInput
	Copy  func(*types.CreateApplicationVersionInput) CreateApplicationVersionRequest
}

// Send marshals and sends the CreateApplicationVersion API request.
func (r CreateApplicationVersionRequest) Send(ctx context.Context) (*CreateApplicationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationVersionResponse{
		CreateApplicationVersionOutput: r.Request.Data.(*types.CreateApplicationVersionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationVersionResponse is the response type for the
// CreateApplicationVersion API operation.
type CreateApplicationVersionResponse struct {
	*types.CreateApplicationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplicationVersion request.
func (r *CreateApplicationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
