// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opGetBucketInventoryConfiguration = "GetBucketInventoryConfiguration"

// GetBucketInventoryConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns an inventory configuration (identified by the inventory configuration
// ID) from the bucket.
//
// To use this operation, you must have permissions to perform the s3:GetInventoryConfiguration
// action. The bucket owner has this permission by default and can grant this
// permission to others. For more information about permissions, see Permissions
// Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// For information about the Amazon S3 inventory feature, see Amazon S3 Inventory
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html).
//
// The following operations are related to GetBucketInventoryConfiguration:
//
//    * DeleteBucketInventoryConfiguration
//
//    * ListBucketInventoryConfigurations
//
//    * PutBucketInventoryConfiguration
//
//    // Example sending a request using GetBucketInventoryConfigurationRequest.
//    req := client.GetBucketInventoryConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketInventoryConfiguration
func (c *Client) GetBucketInventoryConfigurationRequest(input *types.GetBucketInventoryConfigurationInput) GetBucketInventoryConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetBucketInventoryConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?inventory",
	}

	if input == nil {
		input = &types.GetBucketInventoryConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.GetBucketInventoryConfigurationOutput{})
	return GetBucketInventoryConfigurationRequest{Request: req, Input: input, Copy: c.GetBucketInventoryConfigurationRequest}
}

// GetBucketInventoryConfigurationRequest is the request type for the
// GetBucketInventoryConfiguration API operation.
type GetBucketInventoryConfigurationRequest struct {
	*aws.Request
	Input *types.GetBucketInventoryConfigurationInput
	Copy  func(*types.GetBucketInventoryConfigurationInput) GetBucketInventoryConfigurationRequest
}

// Send marshals and sends the GetBucketInventoryConfiguration API request.
func (r GetBucketInventoryConfigurationRequest) Send(ctx context.Context) (*GetBucketInventoryConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketInventoryConfigurationResponse{
		GetBucketInventoryConfigurationOutput: r.Request.Data.(*types.GetBucketInventoryConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketInventoryConfigurationResponse is the response type for the
// GetBucketInventoryConfiguration API operation.
type GetBucketInventoryConfigurationResponse struct {
	*types.GetBucketInventoryConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketInventoryConfiguration request.
func (r *GetBucketInventoryConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
