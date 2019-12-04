// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetBucketReplicationInput struct {
	_ struct{} `type:"structure"`

	// The bucket name for which to get the replication information.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketReplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketReplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketReplicationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketReplicationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketReplicationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetBucketReplicationOutput struct {
	_ struct{} `type:"structure" payload:"ReplicationConfiguration"`

	// A container for replication rules. You can add up to 1,000 rules. The maximum
	// size of a replication configuration is 2 MB.
	ReplicationConfiguration *ReplicationConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetBucketReplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketReplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ReplicationConfiguration != nil {
		v := s.ReplicationConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ReplicationConfiguration", v, metadata)
	}
	return nil
}

const opGetBucketReplication = "GetBucketReplication"

// GetBucketReplicationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the replication configuration of a bucket.
//
// It can take a while to propagate the put or delete a replication configuration
// to all Amazon S3 systems. Therefore, a get request soon after put or delete
// can return a wrong result.
//
// For information about replication configuration, see Replication (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html).
//
// This operation requires permissions for the s3:GetReplicationConfiguration
// action. For more information about permissions, see Using Bucket Policies
// and User Policies (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html).
//
// If you include the Filter element in a replication configuration, you must
// also include the DeleteMarkerReplication and Priority elements. The response
// also returns those elements.
//
// GetBucketReplication has the following special error:
//
//    * Error code: NoSuchReplicationConfiguration Description: There is no
//    replication configuration with that name. HTTP Status Code: 404 Not Found
//    SOAP Fault Code Prefix: Client
//
// The following operations are related to GetBucketReplication:
//
//    * PutBucketReplication
//
//    * DeleteBucketReplication
//
//    // Example sending a request using GetBucketReplicationRequest.
//    req := client.GetBucketReplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketReplication
func (c *Client) GetBucketReplicationRequest(input *GetBucketReplicationInput) GetBucketReplicationRequest {
	op := &aws.Operation{
		Name:       opGetBucketReplication,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?replication",
	}

	if input == nil {
		input = &GetBucketReplicationInput{}
	}

	req := c.newRequest(op, input, &GetBucketReplicationOutput{})
	return GetBucketReplicationRequest{Request: req, Input: input, Copy: c.GetBucketReplicationRequest}
}

// GetBucketReplicationRequest is the request type for the
// GetBucketReplication API operation.
type GetBucketReplicationRequest struct {
	*aws.Request
	Input *GetBucketReplicationInput
	Copy  func(*GetBucketReplicationInput) GetBucketReplicationRequest
}

// Send marshals and sends the GetBucketReplication API request.
func (r GetBucketReplicationRequest) Send(ctx context.Context) (*GetBucketReplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketReplicationResponse{
		GetBucketReplicationOutput: r.Request.Data.(*GetBucketReplicationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketReplicationResponse is the response type for the
// GetBucketReplication API operation.
type GetBucketReplicationResponse struct {
	*GetBucketReplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketReplication request.
func (r *GetBucketReplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
