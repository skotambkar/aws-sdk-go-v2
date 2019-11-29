// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opGetBucketLifecycle = "GetBucketLifecycle"

// GetBucketLifecycleRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
//
// For an updated version of this API, see GetBucketLifecycleConfiguration.
// If you configured a bucket lifecycle using the filter element, you should
// the updated version of this topic. This topic is provided for backward compatibility.
//
// Returns the lifecycle configuration information set on the bucket. For information
// about lifecycle configuration, see Object Lifecycle Management (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html).
//
// To use this operation, you must have permission to perform the s3:GetLifecycleConfiguration
// action. The bucket owner has this permission by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// GetBucketLifecycle has the following special error:
//
//    * Error code: NoSuchLifecycleConfiguration Description: The lifecycle
//    configuration does not exist. HTTP Status Code: 404 Not Found SOAP Fault
//    Code Prefix: Client
//
// The following operations are related to GetBucketLifecycle:
//
//    * GetBucketLifecycleConfiguration
//
//    * PutBucketLifecycle
//
//    * DeleteBucketLifecycle
//
//    // Example sending a request using GetBucketLifecycleRequest.
//    req := client.GetBucketLifecycleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketLifecycle
func (c *Client) GetBucketLifecycleRequest(input *types.GetBucketLifecycleInput) GetBucketLifecycleRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, GetBucketLifecycle, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opGetBucketLifecycle,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &types.GetBucketLifecycleInput{}
	}

	req := c.newRequest(op, input, &types.GetBucketLifecycleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.GetBucketLifecycleMarshaler{Input: input}.GetNamedBuildHandler())

	return GetBucketLifecycleRequest{Request: req, Input: input, Copy: c.GetBucketLifecycleRequest}
}

// GetBucketLifecycleRequest is the request type for the
// GetBucketLifecycle API operation.
type GetBucketLifecycleRequest struct {
	*aws.Request
	Input *types.GetBucketLifecycleInput
	Copy  func(*types.GetBucketLifecycleInput) GetBucketLifecycleRequest
}

// Send marshals and sends the GetBucketLifecycle API request.
func (r GetBucketLifecycleRequest) Send(ctx context.Context) (*GetBucketLifecycleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketLifecycleResponse{
		GetBucketLifecycleOutput: r.Request.Data.(*types.GetBucketLifecycleOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketLifecycleResponse is the response type for the
// GetBucketLifecycle API operation.
type GetBucketLifecycleResponse struct {
	*types.GetBucketLifecycleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketLifecycle request.
func (r *GetBucketLifecycleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
