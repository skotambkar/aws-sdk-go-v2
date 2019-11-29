// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
)

const opAddApplicationReferenceDataSource = "AddApplicationReferenceDataSource"

// AddApplicationReferenceDataSourceRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Adds a reference data source to an existing SQL-based Amazon Kinesis Data
// Analytics application.
//
// Kinesis Data Analytics reads reference data (that is, an Amazon S3 object)
// and creates an in-application table within your application. In the request,
// you provide the source (S3 bucket name and object key name), name of the
// in-application table to create, and the necessary mapping information that
// describes how data in an Amazon S3 object maps to columns in the resulting
// in-application table.
//
//    // Example sending a request using AddApplicationReferenceDataSourceRequest.
//    req := client.AddApplicationReferenceDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/AddApplicationReferenceDataSource
func (c *Client) AddApplicationReferenceDataSourceRequest(input *types.AddApplicationReferenceDataSourceInput) AddApplicationReferenceDataSourceRequest {
	op := &aws.Operation{
		Name:       opAddApplicationReferenceDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AddApplicationReferenceDataSourceInput{}
	}

	req := c.newRequest(op, input, &types.AddApplicationReferenceDataSourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AddApplicationReferenceDataSourceMarshaler{Input: input}.GetNamedBuildHandler())

	return AddApplicationReferenceDataSourceRequest{Request: req, Input: input, Copy: c.AddApplicationReferenceDataSourceRequest}
}

// AddApplicationReferenceDataSourceRequest is the request type for the
// AddApplicationReferenceDataSource API operation.
type AddApplicationReferenceDataSourceRequest struct {
	*aws.Request
	Input *types.AddApplicationReferenceDataSourceInput
	Copy  func(*types.AddApplicationReferenceDataSourceInput) AddApplicationReferenceDataSourceRequest
}

// Send marshals and sends the AddApplicationReferenceDataSource API request.
func (r AddApplicationReferenceDataSourceRequest) Send(ctx context.Context) (*AddApplicationReferenceDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddApplicationReferenceDataSourceResponse{
		AddApplicationReferenceDataSourceOutput: r.Request.Data.(*types.AddApplicationReferenceDataSourceOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddApplicationReferenceDataSourceResponse is the response type for the
// AddApplicationReferenceDataSource API operation.
type AddApplicationReferenceDataSourceResponse struct {
	*types.AddApplicationReferenceDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddApplicationReferenceDataSource request.
func (r *AddApplicationReferenceDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
