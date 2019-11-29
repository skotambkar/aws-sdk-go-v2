// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastoredata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata/types"
)

const opGetObject = "GetObject"

// GetObjectRequest returns a request value for making API operation for
// AWS Elemental MediaStore Data Plane.
//
// Downloads the object at the specified path. If the object’s upload availability
// is set to streaming, AWS Elemental MediaStore downloads the object even if
// it’s still uploading the object.
//
//    // Example sending a request using GetObjectRequest.
//    req := client.GetObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-data-2017-09-01/GetObject
func (c *Client) GetObjectRequest(input *types.GetObjectInput) GetObjectRequest {
	op := &aws.Operation{
		Name:       opGetObject,
		HTTPMethod: "GET",
		HTTPPath:   "/{Path+}",
	}

	if input == nil {
		input = &types.GetObjectInput{}
	}

	req := c.newRequest(op, input, &types.GetObjectOutput{})
	return GetObjectRequest{Request: req, Input: input, Copy: c.GetObjectRequest}
}

// GetObjectRequest is the request type for the
// GetObject API operation.
type GetObjectRequest struct {
	*aws.Request
	Input *types.GetObjectInput
	Copy  func(*types.GetObjectInput) GetObjectRequest
}

// Send marshals and sends the GetObject API request.
func (r GetObjectRequest) Send(ctx context.Context) (*GetObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectResponse{
		GetObjectOutput: r.Request.Data.(*types.GetObjectOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectResponse is the response type for the
// GetObject API operation.
type GetObjectResponse struct {
	*types.GetObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObject request.
func (r *GetObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
