// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
)

const opDescribeDatastore = "DescribeDatastore"

// DescribeDatastoreRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves information about a data store.
//
//    // Example sending a request using DescribeDatastoreRequest.
//    req := client.DescribeDatastoreRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribeDatastore
func (c *Client) DescribeDatastoreRequest(input *types.DescribeDatastoreInput) DescribeDatastoreRequest {
	op := &aws.Operation{
		Name:       opDescribeDatastore,
		HTTPMethod: "GET",
		HTTPPath:   "/datastores/{datastoreName}",
	}

	if input == nil {
		input = &types.DescribeDatastoreInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDatastoreOutput{})
	return DescribeDatastoreRequest{Request: req, Input: input, Copy: c.DescribeDatastoreRequest}
}

// DescribeDatastoreRequest is the request type for the
// DescribeDatastore API operation.
type DescribeDatastoreRequest struct {
	*aws.Request
	Input *types.DescribeDatastoreInput
	Copy  func(*types.DescribeDatastoreInput) DescribeDatastoreRequest
}

// Send marshals and sends the DescribeDatastore API request.
func (r DescribeDatastoreRequest) Send(ctx context.Context) (*DescribeDatastoreResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDatastoreResponse{
		DescribeDatastoreOutput: r.Request.Data.(*types.DescribeDatastoreOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDatastoreResponse is the response type for the
// DescribeDatastore API operation.
type DescribeDatastoreResponse struct {
	*types.DescribeDatastoreOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDatastore request.
func (r *DescribeDatastoreResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
