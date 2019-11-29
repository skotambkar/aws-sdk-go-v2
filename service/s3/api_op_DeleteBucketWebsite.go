// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opDeleteBucketWebsite = "DeleteBucketWebsite"

// DeleteBucketWebsiteRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This operation removes the website configuration for a bucket. Amazon S3
// returns a 200 OK response upon successfully deleting a website configuration
// on the specified bucket. You will get a 200 OK response if the website configuration
// you are trying to delete does not exist on the bucket. Amazon S3 returns
// a 404 response if the bucket specified in the request does not exist.
//
// This DELETE operation requires the S3:DeleteBucketWebsite permission. By
// default, only the bucket owner can delete the website configuration attached
// to a bucket. However, bucket owners can grant other users permission to delete
// the website configuration by writing a bucket policy granting them the S3:DeleteBucketWebsite
// permission.
//
// For more information about hosting websites, see Hosting Websites on Amazon
// S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).
//
// The following operations are related to DeleteBucketWebsite
//
//    * GetBucketWebsite
//
//    * PutBucketWebsite
//
//    // Example sending a request using DeleteBucketWebsiteRequest.
//    req := client.DeleteBucketWebsiteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketWebsite
func (c *Client) DeleteBucketWebsiteRequest(input *types.DeleteBucketWebsiteInput) DeleteBucketWebsiteRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketWebsite,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?website",
	}

	if input == nil {
		input = &types.DeleteBucketWebsiteInput{}
	}

	req := c.newRequest(op, input, &types.DeleteBucketWebsiteOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.DeleteBucketWebsiteMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketWebsiteRequest{Request: req, Input: input, Copy: c.DeleteBucketWebsiteRequest}
}

// DeleteBucketWebsiteRequest is the request type for the
// DeleteBucketWebsite API operation.
type DeleteBucketWebsiteRequest struct {
	*aws.Request
	Input *types.DeleteBucketWebsiteInput
	Copy  func(*types.DeleteBucketWebsiteInput) DeleteBucketWebsiteRequest
}

// Send marshals and sends the DeleteBucketWebsite API request.
func (r DeleteBucketWebsiteRequest) Send(ctx context.Context) (*DeleteBucketWebsiteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketWebsiteResponse{
		DeleteBucketWebsiteOutput: r.Request.Data.(*types.DeleteBucketWebsiteOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketWebsiteResponse is the response type for the
// DeleteBucketWebsite API operation.
type DeleteBucketWebsiteResponse struct {
	*types.DeleteBucketWebsiteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketWebsite request.
func (r *DeleteBucketWebsiteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
