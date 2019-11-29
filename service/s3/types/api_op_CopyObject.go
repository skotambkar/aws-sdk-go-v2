// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/s3/enums"
)

type CopyObjectInput struct {
	_ struct{} `type:"structure"`

	// The canned ACL to apply to the object.
	ACL enums.ObjectCannedACL `location:"header" locationName:"x-amz-acl" type:"string" enum:"true"`

	// The name of the destination bucket.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

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

	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The name of the source bucket and key name of the source object, separated
	// by a slash (/). Must be URL-encoded.
	//
	// CopySource is a required field
	CopySource *string `location:"header" locationName:"x-amz-copy-source" type:"string" required:"true"`

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopySourceIfMatch *string `location:"header" locationName:"x-amz-copy-source-if-match" type:"string"`

	// Copies the object if it has been modified since the specified time.
	CopySourceIfModifiedSince *time.Time `location:"header" locationName:"x-amz-copy-source-if-modified-since" type:"timestamp"`

	// Copies the object if its entity tag (ETag) is different than the specified
	// ETag.
	CopySourceIfNoneMatch *string `location:"header" locationName:"x-amz-copy-source-if-none-match" type:"string"`

	// Copies the object if it hasn't been modified since the specified time.
	CopySourceIfUnmodifiedSince *time.Time `location:"header" locationName:"x-amz-copy-source-if-unmodified-since" type:"timestamp"`

	// Specifies the algorithm to use when decrypting the source object (e.g., AES256).
	CopySourceSSECustomerAlgorithm *string `location:"header" locationName:"x-amz-copy-source-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt
	// the source object. The encryption key provided in this header must be one
	// that was used when the source object was created.
	CopySourceSSECustomerKey *string `location:"header" locationName:"x-amz-copy-source-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	CopySourceSSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-copy-source-server-side-encryption-customer-key-MD5" type:"string"`

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time `location:"header" locationName:"Expires" type:"timestamp"`

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`

	// Allows grantee to read the object data and its metadata.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`

	// Allows grantee to read the object ACL.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`

	// Allows grantee to write the ACL for the applicable object.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`

	// The key of the destination object.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// Specifies whether the metadata is copied from the source object or replaced
	// with metadata provided in the request.
	MetadataDirective enums.MetadataDirective `location:"header" locationName:"x-amz-metadata-directive" type:"string" enum:"true"`

	// Specifies whether you want to apply a Legal Hold to the copied object.
	ObjectLockLegalHoldStatus enums.ObjectLockLegalHoldStatus `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"true"`

	// The Object Lock mode that you want to apply to the copied object.
	ObjectLockMode enums.ObjectLockMode `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"true"`

	// The date and time when you want the copied object's Object Lock to expire.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer enums.RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

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

	// Specifies the AWS KMS Encryption Context to use for object encryption. The
	// value of this header is a base64-encoded UTF-8 string holding JSON with the
	// encryption context key-value pairs.
	SSEKMSEncryptionContext *string `location:"header" locationName:"x-amz-server-side-encryption-context" type:"string" sensitive:"true"`

	// Specifies the AWS KMS key ID to use for object encryption. All GET and PUT
	// requests for an object protected by AWS KMS will fail if not made via SSL
	// or using SigV4. Documentation on configuring any of the officially supported
	// AWS SDKs and CLI can be found at http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption enums.ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// The type of storage to use for the object. Defaults to 'STANDARD'.
	StorageClass enums.StorageClass `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"true"`

	// The tag-set for the object destination object this value must be used in
	// conjunction with the TaggingDirective. The tag-set must be encoded as URL
	// Query parameters
	Tagging *string `location:"header" locationName:"x-amz-tagging" type:"string"`

	// Specifies whether the object tag-set are copied from the source object or
	// replaced with tag-set provided in the request.
	TaggingDirective enums.TaggingDirective `location:"header" locationName:"x-amz-tagging-directive" type:"string" enum:"true"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s CopyObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyObjectInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.CopySource == nil {
		invalidParams.Add(aws.NewErrParamRequired("CopySource"))
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

func (s *CopyObjectInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

func (s *CopyObjectInput) GetCopySourceSSECustomerKey() (v string) {
	if s.CopySourceSSECustomerKey == nil {
		return v
	}
	return *s.CopySourceSSECustomerKey
}

func (s *CopyObjectInput) GetSSECustomerKey() (v string) {
	if s.SSECustomerKey == nil {
		return v
	}
	return *s.SSECustomerKey
}

type CopyObjectOutput struct {
	_ struct{} `type:"structure" payload:"CopyObjectResult"`

	// Container for all response elements.
	CopyObjectResult *CopyObjectResult `type:"structure"`

	// Version of the copied object in the destination bucket.
	CopySourceVersionId *string `location:"header" locationName:"x-amz-copy-source-version-id" type:"string"`

	// If the object expiration is configured, the response includes this header.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged enums.RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the AWS KMS Encryption Context to use for object encryption.
	// The value of this header is a base64-encoded UTF-8 string holding JSON with
	// the encryption context key-value pairs.
	SSEKMSEncryptionContext *string `location:"header" locationName:"x-amz-server-side-encryption-context" type:"string" sensitive:"true"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) customer
	// master key (CMK) that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption enums.ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// Version ID of the newly created copy.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s CopyObjectOutput) String() string {
	return awsutil.Prettify(s)
}