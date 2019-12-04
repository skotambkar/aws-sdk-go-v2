// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhubconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetHomeRegionInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetHomeRegionInput) String() string {
	return awsutil.Prettify(s)
}

type GetHomeRegionOutput struct {
	_ struct{} `type:"structure"`

	// The name of the home region of the calling account.
	HomeRegion *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetHomeRegionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetHomeRegion = "GetHomeRegion"

// GetHomeRegionRequest returns a request value for making API operation for
// AWS Migration Hub Config.
//
// Returns the calling account’s home region, if configured. This API is used
// by other AWS services to determine the regional endpoint for calling AWS
// Application Discovery Service and Migration Hub. You must call GetHomeRegion
// at least once before you call any other AWS Application Discovery Service
// and AWS Migration Hub APIs, to obtain the account's Migration Hub home region.
//
//    // Example sending a request using GetHomeRegionRequest.
//    req := client.GetHomeRegionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/migrationhub-config-2019-06-30/GetHomeRegion
func (c *Client) GetHomeRegionRequest(input *GetHomeRegionInput) GetHomeRegionRequest {
	op := &aws.Operation{
		Name:       opGetHomeRegion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetHomeRegionInput{}
	}

	req := c.newRequest(op, input, &GetHomeRegionOutput{})
	return GetHomeRegionRequest{Request: req, Input: input, Copy: c.GetHomeRegionRequest}
}

// GetHomeRegionRequest is the request type for the
// GetHomeRegion API operation.
type GetHomeRegionRequest struct {
	*aws.Request
	Input *GetHomeRegionInput
	Copy  func(*GetHomeRegionInput) GetHomeRegionRequest
}

// Send marshals and sends the GetHomeRegion API request.
func (r GetHomeRegionRequest) Send(ctx context.Context) (*GetHomeRegionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetHomeRegionResponse{
		GetHomeRegionOutput: r.Request.Data.(*GetHomeRegionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetHomeRegionResponse is the response type for the
// GetHomeRegion API operation.
type GetHomeRegionResponse struct {
	*GetHomeRegionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetHomeRegion request.
func (r *GetHomeRegionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
