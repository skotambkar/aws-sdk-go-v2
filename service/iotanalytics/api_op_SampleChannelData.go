// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
)

const opSampleChannelData = "SampleChannelData"

// SampleChannelDataRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves a sample of messages from the specified channel ingested during
// the specified timeframe. Up to 10 messages can be retrieved.
//
//    // Example sending a request using SampleChannelDataRequest.
//    req := client.SampleChannelDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/SampleChannelData
func (c *Client) SampleChannelDataRequest(input *types.SampleChannelDataInput) SampleChannelDataRequest {
	op := &aws.Operation{
		Name:       opSampleChannelData,
		HTTPMethod: "GET",
		HTTPPath:   "/channels/{channelName}/sample",
	}

	if input == nil {
		input = &types.SampleChannelDataInput{}
	}

	req := c.newRequest(op, input, &types.SampleChannelDataOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.SampleChannelDataMarshaler{Input: input}.GetNamedBuildHandler())

	return SampleChannelDataRequest{Request: req, Input: input, Copy: c.SampleChannelDataRequest}
}

// SampleChannelDataRequest is the request type for the
// SampleChannelData API operation.
type SampleChannelDataRequest struct {
	*aws.Request
	Input *types.SampleChannelDataInput
	Copy  func(*types.SampleChannelDataInput) SampleChannelDataRequest
}

// Send marshals and sends the SampleChannelData API request.
func (r SampleChannelDataRequest) Send(ctx context.Context) (*SampleChannelDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SampleChannelDataResponse{
		SampleChannelDataOutput: r.Request.Data.(*types.SampleChannelDataOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SampleChannelDataResponse is the response type for the
// SampleChannelData API operation.
type SampleChannelDataResponse struct {
	*types.SampleChannelDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SampleChannelData request.
func (r *SampleChannelDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
