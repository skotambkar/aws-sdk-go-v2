// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

const opCreateFunction = "CreateFunction"

// CreateFunctionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Creates a Lambda function. To create a function, you need a deployment package
// (https://docs.aws.amazon.com/lambda/latest/dg/deployment-package-v2.html)
// and an execution role (https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role).
// The deployment package contains your function code. The execution role grants
// the function permission to use AWS services, such as Amazon CloudWatch Logs
// for log streaming and AWS X-Ray for request tracing.
//
// A function has an unpublished version, and can have published versions and
// aliases. The unpublished version changes when you update your function's
// code and configuration. A published version is a snapshot of your function
// code and configuration that can't be changed. An alias is a named resource
// that maps to a version, and can be changed to map to a different version.
// Use the Publish parameter to create version 1 of your function from its initial
// configuration.
//
// The other parameters let you configure version-specific and function-level
// settings. You can modify version-specific settings later with UpdateFunctionConfiguration.
// Function-level settings apply to both the unpublished and published versions
// of the function, and include tags (TagResource) and per-function concurrency
// limits (PutFunctionConcurrency).
//
// If another account or an AWS service invokes your function, use AddPermission
// to grant permission by creating a resource-based IAM policy. You can grant
// permissions at the function level, on a version, or on an alias.
//
// To invoke your function directly, use Invoke. To invoke your function in
// response to events in other AWS services, create an event source mapping
// (CreateEventSourceMapping), or configure a function trigger in the other
// service. For more information, see Invoking Functions (https://docs.aws.amazon.com/lambda/latest/dg/invoking-lambda-functions.html).
//
//    // Example sending a request using CreateFunctionRequest.
//    req := client.CreateFunctionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/CreateFunction
func (c *Client) CreateFunctionRequest(input *types.CreateFunctionInput) CreateFunctionRequest {
	op := &aws.Operation{
		Name:       opCreateFunction,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-03-31/functions",
	}

	if input == nil {
		input = &types.CreateFunctionInput{}
	}

	req := c.newRequest(op, input, &types.CreateFunctionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateFunctionMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateFunctionRequest{Request: req, Input: input, Copy: c.CreateFunctionRequest}
}

// CreateFunctionRequest is the request type for the
// CreateFunction API operation.
type CreateFunctionRequest struct {
	*aws.Request
	Input *types.CreateFunctionInput
	Copy  func(*types.CreateFunctionInput) CreateFunctionRequest
}

// Send marshals and sends the CreateFunction API request.
func (r CreateFunctionRequest) Send(ctx context.Context) (*CreateFunctionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFunctionResponse{
		CreateFunctionOutput: r.Request.Data.(*types.CreateFunctionOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFunctionResponse is the response type for the
// CreateFunction API operation.
type CreateFunctionResponse struct {
	*types.CreateFunctionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFunction request.
func (r *CreateFunctionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
