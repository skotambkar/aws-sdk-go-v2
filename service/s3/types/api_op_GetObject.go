// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/s3/enums"
)

type GetObjectInput struct {
	_ struct{} `type:"structure"`

	// The bucket name containing the object.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Return the object only if its entity tag (ETag) is the same as the one specified,
	// otherwise return a 412 (precondition failed).
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`

	// Return the object only if it has been modified since the specified time,
	// otherwise return a 304 (not modified).
	IfModifiedSince *time.Time `location:"header" locationName:"If-Modified-Since" type:"timestamp"`

	// Return the object only if its entity tag (ETag) is different from the one
	// specified, otherwise return a 304 (not modified).
	IfNoneMatch *string `location:"header" locationName:"If-None-Match" type:"string"`

	// Return the object only if it has not been modified since the specified time,
	// otherwise return a 412 (precondition failed).
	IfUnmodifiedSince *time.Time `location:"header" locationName:"If-Unmodified-Since" type:"timestamp"`

	// Key of the object to get.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Part number of the object being read. This is a positive integer between
	// 1 and 10,000. Effectively performs a 'ranged' GET request for the part specified.
	// Useful for downloading just a part of an object.
	PartNumber *int64 `location:"querystring" locationName:"partNumber" type:"integer"`

	// Downloads the specified range bytes of an object. For more information about
	// the HTTP Range header, go to http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35.
	Range *string `location:"header" locationName:"Range" type:"string"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer enums.RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Sets the Cache-Control header of the response.
	ResponseCacheControl *string `location:"querystring" locationName:"response-cache-control" type:"string"`

	// Sets the Content-Disposition header of the response
	ResponseContentDisposition *string `location:"querystring" locationName:"response-content-disposition" type:"string"`

	// Sets the Content-Encoding header of the response.
	ResponseContentEncoding *string `location:"querystring" locationName:"response-content-encoding" type:"string"`

	// Sets the Content-Language header of the response.
	ResponseContentLanguage *string `location:"querystring" locationName:"response-content-language" type:"string"`

	// Sets the Content-Type header of the response.
	ResponseContentType *string `location:"querystring" locationName:"response-content-type" type:"string"`

	// Sets the Expires header of the response.
	ResponseExpires *time.Time `location:"querystring" locationName:"response-expires" type:"timestamp"`

	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// does not store the encryption key. The key must be appropriate for use with
	// the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header.
	SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// VersionId used to reference a specific version of the object.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s GetObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

func (s *GetObjectInput) GetSSECustomerKey() (v string) {
	if s.SSECustomerKey == nil {
		return v
	}
	return *s.SSECustomerKey
}

type GetObjectOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Indicates that a range of bytes was specifed.
	AcceptRanges *string `location:"header" locationName:"accept-ranges" type:"string"`

	// Object data.
	Body io.ReadCloser `type:"blob"`

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// Size of the body in bytes.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// The portion of the object returned in the response.
	ContentRange *string `location:"header" locationName:"Content-Range" type:"string"`

	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool `location:"header" locationName:"x-amz-delete-marker" type:"boolean"`

	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// The date and time at which the object is no longer cacheable.
	Expires *string `location:"header" locationName:"Expires" type:"string"`

	// Last modified date of the object
	LastModified *time.Time `location:"header" locationName:"Last-Modified" type:"timestamp"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP,
	// you can create metadata whose values are not legal HTTP headers.
	MissingMeta *int64 `location:"header" locationName:"x-amz-missing-meta" type:"integer"`

	// Indicates whether this object has an active legal hold. This field is only
	// returned if you have permission to view an object's legal hold status.
	ObjectLockLegalHoldStatus enums.ObjectLockLegalHoldStatus `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"true"`

	// The Object Lock mode currently in place for this object.
	ObjectLockMode enums.ObjectLockMode `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"true"`

	// The date and time when this object's Object Lock will expire.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// The count of parts this object has.
	PartsCount *int64 `location:"header" locationName:"x-amz-mp-parts-count" type:"integer"`

	// Amazon S3 can return this if your request involves a bucket that is either
	// a source or destination in a replication rule.
	ReplicationStatus enums.ReplicationStatus `location:"header" locationName:"x-amz-replication-status" type:"string" enum:"true"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged enums.RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// Provides information about object restoration operation and expiration time
	// of the restored object copy.
	Restore *string `location:"header" locationName:"x-amz-restore" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) customer
	// master key (CMK) that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption enums.ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// Provides storage class information of the object. Amazon S3 returns this
	// header for all objects except for Standard storage class objects.
	StorageClass enums.StorageClass `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"true"`

	// The number of tags, if any, on the object.
	TagCount *int64 `location:"header" locationName:"x-amz-tagging-count" type:"integer"`

	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s GetObjectOutput) String() string {
	return awsutil.Prettify(s)
}
