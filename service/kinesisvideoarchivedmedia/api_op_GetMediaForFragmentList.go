// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia/types"
)

const opGetMediaForFragmentList = "GetMediaForFragmentList"

// GetMediaForFragmentListRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams Archived Media.
//
// Gets media for a list of fragments (specified by fragment number) from the
// archived data in an Amazon Kinesis video stream.
//
// You must first call the GetDataEndpoint API to get an endpoint. Then send
// the GetMediaForFragmentList requests to this endpoint using the --endpoint-url
// parameter (https://docs.aws.amazon.com/cli/latest/reference/).
//
// The following limits apply when using the GetMediaForFragmentList API:
//
//    * A client can call GetMediaForFragmentList up to five times per second
//    per stream.
//
//    * Kinesis Video Streams sends media data at a rate of up to 25 megabytes
//    per second (or 200 megabits per second) during a GetMediaForFragmentList
//    session.
//
// If an error is thrown after invoking a Kinesis Video Streams archived media
// API, in addition to the HTTP status code and the response body, it includes
// the following pieces of information:
//
//    * x-amz-ErrorType HTTP header – contains a more specific error type
//    in addition to what the HTTP status code provides.
//
//    * x-amz-RequestId HTTP header – if you want to report an issue to AWS,
//    the support team can better diagnose the problem if given the Request
//    Id.
//
// Both the HTTP status code and the ErrorType header can be utilized to make
// programmatic decisions about whether errors are retry-able and under what
// conditions, as well as provide information on what actions the client programmer
// might need to take in order to successfully try again.
//
// For more information, see the Errors section at the bottom of this topic,
// as well as Common Errors (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/CommonErrors.html).
//
//    // Example sending a request using GetMediaForFragmentListRequest.
//    req := client.GetMediaForFragmentListRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetMediaForFragmentList
func (c *Client) GetMediaForFragmentListRequest(input *types.GetMediaForFragmentListInput) GetMediaForFragmentListRequest {
	op := &aws.Operation{
		Name:       opGetMediaForFragmentList,
		HTTPMethod: "POST",
		HTTPPath:   "/getMediaForFragmentList",
	}

	if input == nil {
		input = &types.GetMediaForFragmentListInput{}
	}

	req := c.newRequest(op, input, &types.GetMediaForFragmentListOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetMediaForFragmentListMarshaler{Input: input}.GetNamedBuildHandler())

	return GetMediaForFragmentListRequest{Request: req, Input: input, Copy: c.GetMediaForFragmentListRequest}
}

// GetMediaForFragmentListRequest is the request type for the
// GetMediaForFragmentList API operation.
type GetMediaForFragmentListRequest struct {
	*aws.Request
	Input *types.GetMediaForFragmentListInput
	Copy  func(*types.GetMediaForFragmentListInput) GetMediaForFragmentListRequest
}

// Send marshals and sends the GetMediaForFragmentList API request.
func (r GetMediaForFragmentListRequest) Send(ctx context.Context) (*GetMediaForFragmentListResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMediaForFragmentListResponse{
		GetMediaForFragmentListOutput: r.Request.Data.(*types.GetMediaForFragmentListOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMediaForFragmentListResponse is the response type for the
// GetMediaForFragmentList API operation.
type GetMediaForFragmentListResponse struct {
	*types.GetMediaForFragmentListOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMediaForFragmentList request.
func (r *GetMediaForFragmentListResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
