// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opExportImage = "ExportImage"

// ExportImageRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Exports an Amazon Machine Image (AMI) to a VM file. For more information,
// see Exporting a VM Directory from an Amazon Machine Image (AMI) (https://docs.aws.amazon.com/vm-import/latest/userguide/vmexport_image.html)
// in the VM Import/Export User Guide.
//
//    // Example sending a request using ExportImageRequest.
//    req := client.ExportImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ExportImage
func (c *Client) ExportImageRequest(input *types.ExportImageInput) ExportImageRequest {
	op := &aws.Operation{
		Name:       opExportImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ExportImageInput{}
	}

	req := c.newRequest(op, input, &types.ExportImageOutput{})
	return ExportImageRequest{Request: req, Input: input, Copy: c.ExportImageRequest}
}

// ExportImageRequest is the request type for the
// ExportImage API operation.
type ExportImageRequest struct {
	*aws.Request
	Input *types.ExportImageInput
	Copy  func(*types.ExportImageInput) ExportImageRequest
}

// Send marshals and sends the ExportImage API request.
func (r ExportImageRequest) Send(ctx context.Context) (*ExportImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExportImageResponse{
		ExportImageOutput: r.Request.Data.(*types.ExportImageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExportImageResponse is the response type for the
// ExportImage API operation.
type ExportImageResponse struct {
	*types.ExportImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExportImage request.
func (r *ExportImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
