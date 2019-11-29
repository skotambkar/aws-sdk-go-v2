// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

const opCreateFieldLevelEncryptionConfig = "CreateFieldLevelEncryptionConfig2019_03_26"

// CreateFieldLevelEncryptionConfigRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Create a new field-level encryption configuration.
//
//    // Example sending a request using CreateFieldLevelEncryptionConfigRequest.
//    req := client.CreateFieldLevelEncryptionConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateFieldLevelEncryptionConfig
func (c *Client) CreateFieldLevelEncryptionConfigRequest(input *types.CreateFieldLevelEncryptionConfigInput) CreateFieldLevelEncryptionConfigRequest {
	op := &aws.Operation{
		Name:       opCreateFieldLevelEncryptionConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/2019-03-26/field-level-encryption",
	}

	if input == nil {
		input = &types.CreateFieldLevelEncryptionConfigInput{}
	}

	req := c.newRequest(op, input, &types.CreateFieldLevelEncryptionConfigOutput{})
	return CreateFieldLevelEncryptionConfigRequest{Request: req, Input: input, Copy: c.CreateFieldLevelEncryptionConfigRequest}
}

// CreateFieldLevelEncryptionConfigRequest is the request type for the
// CreateFieldLevelEncryptionConfig API operation.
type CreateFieldLevelEncryptionConfigRequest struct {
	*aws.Request
	Input *types.CreateFieldLevelEncryptionConfigInput
	Copy  func(*types.CreateFieldLevelEncryptionConfigInput) CreateFieldLevelEncryptionConfigRequest
}

// Send marshals and sends the CreateFieldLevelEncryptionConfig API request.
func (r CreateFieldLevelEncryptionConfigRequest) Send(ctx context.Context) (*CreateFieldLevelEncryptionConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFieldLevelEncryptionConfigResponse{
		CreateFieldLevelEncryptionConfigOutput: r.Request.Data.(*types.CreateFieldLevelEncryptionConfigOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFieldLevelEncryptionConfigResponse is the response type for the
// CreateFieldLevelEncryptionConfig API operation.
type CreateFieldLevelEncryptionConfigResponse struct {
	*types.CreateFieldLevelEncryptionConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFieldLevelEncryptionConfig request.
func (r *CreateFieldLevelEncryptionConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
