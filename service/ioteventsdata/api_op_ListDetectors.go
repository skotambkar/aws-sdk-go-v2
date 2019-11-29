// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ioteventsdata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata/types"
)

const opListDetectors = "ListDetectors"

// ListDetectorsRequest returns a request value for making API operation for
// AWS IoT Events Data.
//
// Lists detectors (the instances of a detector model).
//
//    // Example sending a request using ListDetectorsRequest.
//    req := client.ListDetectorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-data-2018-10-23/ListDetectors
func (c *Client) ListDetectorsRequest(input *types.ListDetectorsInput) ListDetectorsRequest {
	op := &aws.Operation{
		Name:       opListDetectors,
		HTTPMethod: "GET",
		HTTPPath:   "/detectors/{detectorModelName}",
	}

	if input == nil {
		input = &types.ListDetectorsInput{}
	}

	req := c.newRequest(op, input, &types.ListDetectorsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListDetectorsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListDetectorsRequest{Request: req, Input: input, Copy: c.ListDetectorsRequest}
}

// ListDetectorsRequest is the request type for the
// ListDetectors API operation.
type ListDetectorsRequest struct {
	*aws.Request
	Input *types.ListDetectorsInput
	Copy  func(*types.ListDetectorsInput) ListDetectorsRequest
}

// Send marshals and sends the ListDetectors API request.
func (r ListDetectorsRequest) Send(ctx context.Context) (*ListDetectorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDetectorsResponse{
		ListDetectorsOutput: r.Request.Data.(*types.ListDetectorsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDetectorsResponse is the response type for the
// ListDetectors API operation.
type ListDetectorsResponse struct {
	*types.ListDetectorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDetectors request.
func (r *ListDetectorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
