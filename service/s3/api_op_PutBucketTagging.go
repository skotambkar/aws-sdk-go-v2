// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type PutBucketTaggingInput struct {
	_ struct{} `type:"structure" payload:"Tagging"`

	// The bucket name.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Container for the TagSet and Tag elements.
	//
	// Tagging is a required field
	Tagging *Tagging `locationName:"Tagging" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketTaggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketTaggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketTaggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Tagging == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tagging"))
	}
	if s.Tagging != nil {
		if err := s.Tagging.Validate(); err != nil {
			invalidParams.AddNested("Tagging", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketTaggingInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketTaggingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Tagging != nil {
		v := s.Tagging

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "Tagging", v, metadata)
	}
	return nil
}

type PutBucketTaggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketTaggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketTaggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketTagging = "PutBucketTagging"

// PutBucketTaggingRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets the tags for a bucket.
//
// Use tags to organize your AWS bill to reflect your own cost structure. To
// do this, sign up to get your AWS account bill with tag key values included.
// Then, to see the cost of combined resources, organize your billing information
// according to resources with the same tag key values. For example, you can
// tag several resources with a specific application name, and then organize
// your billing information to see the total cost of that application across
// several services. For more information, see Cost Allocation and Tagging (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html).
//
// Within a bucket, if you add a tag that has the same key as an existing tag,
// the new value overwrites the old value. For more information, see Using Cost
// Allocation in Amazon S3 Bucket Tags (https://docs.aws.amazon.com/AmazonS3/latest/dev/CostAllocTagging.html).
//
// To use this operation, you must have permissions to perform the s3:PutBucketTagging
// action. The bucket owner has this permission by default and can grant this
// permission to others. For more information about permissions, see Permissions
// Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// PutBucketTagging has the following special errors:
//
//    * Error code: InvalidTagError Description: The tag provided was not a
//    valid tag. This error can occur if the tag did not pass input validation.
//    For information about tag restrictions, see User-Defined Tag Restrictions
//    (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2//allocation-tag-restrictions.html)
//    and AWS-Generated Cost Allocation Tag Restrictions (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2//aws-tag-restrictions.html).
//
//    * Error code: MalformedXMLError Description: The XML provided does not
//    match the schema.
//
//    * Error code: OperationAbortedError Description: A conflicting conditional
//    operation is currently in progress against this resource. Please try again.
//
//    * Error code: InternalError Description: The service was unable to apply
//    the provided tag to the bucket.
//
// The following operations are related to PutBucketTagging:
//
//    * GetBucketTagging
//
//    * DeleteBucketTagging
//
//    // Example sending a request using PutBucketTaggingRequest.
//    req := client.PutBucketTaggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketTagging
func (c *Client) PutBucketTaggingRequest(input *PutBucketTaggingInput) PutBucketTaggingRequest {
	op := &aws.Operation{
		Name:       opPutBucketTagging,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?tagging",
	}

	if input == nil {
		input = &PutBucketTaggingInput{}
	}

	req := c.newRequest(op, input, &PutBucketTaggingOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketTaggingRequest{Request: req, Input: input, Copy: c.PutBucketTaggingRequest}
}

// PutBucketTaggingRequest is the request type for the
// PutBucketTagging API operation.
type PutBucketTaggingRequest struct {
	*aws.Request
	Input *PutBucketTaggingInput
	Copy  func(*PutBucketTaggingInput) PutBucketTaggingRequest
}

// Send marshals and sends the PutBucketTagging API request.
func (r PutBucketTaggingRequest) Send(ctx context.Context) (*PutBucketTaggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketTaggingResponse{
		PutBucketTaggingOutput: r.Request.Data.(*PutBucketTaggingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketTaggingResponse is the response type for the
// PutBucketTagging API operation.
type PutBucketTaggingResponse struct {
	*PutBucketTaggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketTagging request.
func (r *PutBucketTaggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
