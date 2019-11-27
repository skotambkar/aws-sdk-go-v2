// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
)

const opCreateDatasetContent = "CreateDatasetContent"

// CreateDatasetContentRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Creates the content of a data set by applying a "queryAction" (a SQL query)
// or a "containerAction" (executing a containerized application).
//
//    // Example sending a request using CreateDatasetContentRequest.
//    req := client.CreateDatasetContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/CreateDatasetContent
func (c *Client) CreateDatasetContentRequest(input *types.CreateDatasetContentInput) CreateDatasetContentRequest {
	op := &aws.Operation{
		Name:       opCreateDatasetContent,
		HTTPMethod: "POST",
		HTTPPath:   "/datasets/{datasetName}/content",
	}

	if input == nil {
		input = &types.CreateDatasetContentInput{}
	}

	req := c.newRequest(op, input, &types.CreateDatasetContentOutput{})
	return CreateDatasetContentRequest{Request: req, Input: input, Copy: c.CreateDatasetContentRequest}
}

// CreateDatasetContentRequest is the request type for the
// CreateDatasetContent API operation.
type CreateDatasetContentRequest struct {
	*aws.Request
	Input *types.CreateDatasetContentInput
	Copy  func(*types.CreateDatasetContentInput) CreateDatasetContentRequest
}

// Send marshals and sends the CreateDatasetContent API request.
func (r CreateDatasetContentRequest) Send(ctx context.Context) (*CreateDatasetContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDatasetContentResponse{
		CreateDatasetContentOutput: r.Request.Data.(*types.CreateDatasetContentOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDatasetContentResponse is the response type for the
// CreateDatasetContent API operation.
type CreateDatasetContentResponse struct {
	*types.CreateDatasetContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDatasetContent request.
func (r *CreateDatasetContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
