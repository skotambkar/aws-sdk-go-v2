// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opCopyObject = "CopyObject"

// CopyObjectRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates a copy of an object that is already stored in Amazon S3.
//
// You can store individual objects of up to 5 TB in Amazon S3. You create a
// copy of your object up to 5 GB in size in a single atomic operation using
// this API. However, for copying an object greater than 5 GB, you must use
// the multipart upload Upload Part - Copy API. For conceptual information,
// see Copy Object Using the REST Multipart Upload API (https://docs.aws.amazon.com/AmazonS3/latest/dev/CopyingObjctsUsingRESTMPUapi.html).
//
// When copying an object, you can preserve all metadata (default) or specify
// new metadata. However, the ACL is not preserved and is set to private for
// the user making the request. To override the default ACL setting, specify
// a new ACL when generating a copy request. For more information, see Using
// ACLs (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html).
//
// Amazon S3 Transfer Acceleration does not support cross-region copies. If
// you request a cross-region copy using a Transfer Acceleration endpoint, you
// get a 400 Bad Request error. For more information about transfer acceleration,
// see Transfer Acceleration (https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html).
//
// All copy requests must be authenticated. Additionally, you must have read
// access to the source object and write access to the destination bucket. For
// more information, see REST Authentication (https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html).
// Both the Region that you want to copy the object from and the Region that
// you want to copy the object to must be enabled for your account.
//
// To only copy an object under certain conditions, such as whether the Etag
// matches or whether the object was modified before or after a specified date,
// use the request parameters x-amz-copy-source-if-match, x-amz-copy-source-if-none-match,
// x-amz-copy-source-if-unmodified-since, or x-amz-copy-source-if-modified-since.
//
// All headers with the x-amz- prefix, including x-amz-copy-source, must be
// signed.
//
// You can use this operation to change the storage class of an object that
// is already stored in Amazon S3 using the StorageClass parameter. For more
// information, see Storage Classes (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html).
//
// The source object that you are copying can be encrypted or unencrypted. If
// the source object is encrypted, it can be encrypted by server-side encryption
// using AWS-managed encryption keys or by using a customer-provided encryption
// key. When copying an object, you can request that Amazon S3 encrypt the target
// object by using either the AWS-managed encryption keys or by using your own
// encryption key. You can do this regardless of the form of server-side encryption
// that was used to encrypt the source, or even if the source object was not
// encrypted. For more information about server-side encryption, see Using Server-Side
// Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html).
//
// A copy request might return an error when Amazon S3 receives the copy request
// or while Amazon S3 is copying the files. If the error occurs before the copy
// operation starts, you receive a standard Amazon S3 error. If the error occurs
// during the copy operation, the error response is embedded in the 200 OK response.
// This means that a 200 OK response can contain either a success or an error.
// Design your application to parse the contents of the response and handle
// it appropriately.
//
// If the copy is successful, you receive a response with information about
// the copied object.
//
// If the request is an HTTP 1.1 request, the response is chunk encoded. If
// it were not, it would not contain the content-length, and you would need
// to read the entire body.
//
// Consider the following when using request headers:
//
//    * Consideration 1 – If both the x-amz-copy-source-if-match and x-amz-copy-source-if-unmodified-since
//    headers are present in the request and evaluate as follows, Amazon S3
//    returns 200 OK and copies the data: x-amz-copy-source-if-match condition
//    evaluates to true x-amz-copy-source-if-unmodified-since condition evaluates
//    to false
//
//    * Consideration 2 – If both of the x-amz-copy-source-if-none-match and
//    x-amz-copy-source-if-modified-since headers are present in the request
//    and evaluate as follows, Amazon S3 returns the 412 Precondition Failed
//    response code: x-amz-copy-source-if-none-match condition evaluates to
//    false x-amz-copy-source-if-modified-since condition evaluates to true
//
// The copy request charge is based on the storage class and Region you specify
// for the destination object. For pricing information, see Amazon S3 Pricing
// (https://aws.amazon.com/s3/pricing/).
//
// Following are other considerations when using CopyObject:
//
// Versioning
//
// By default, x-amz-copy-source identifies the current version of an object
// to copy. (If the current version is a delete marker, Amazon S3 behaves as
// if the object was deleted.) To copy a different version, use the versionId
// subresource.
//
// If you enable versioning on the target bucket, Amazon S3 generates a unique
// version ID for the object being copied. This version ID is different from
// the version ID of the source object. Amazon S3 returns the version ID of
// the copied object in the x-amz-version-id response header in the response.
//
// If you do not enable versioning or suspend it on the target bucket, the version
// ID that Amazon S3 generates is always null.
//
// If the source object's storage class is GLACIER, then you must restore a
// copy of this object before you can use it as a source object for the copy
// operation. For more information, see .
//
// Access Permissions
//
// When copying an object, you can optionally specify the accounts or groups
// that should be granted specific permissions on the new object. There are
// two ways to grant the permissions using the request headers:
//
//    * Specify a canned ACL with the x-amz-acl request header. For more information,
//    see Canned ACL (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
//
//    * Specify access permissions explicitly with the x-amz-grant-read, x-amz-grant-read-acp,
//    x-amz-grant-write-acp, and x-amz-grant-full-control headers. These parameters
//    map to the set of permissions that Amazon S3 supports in an ACL. For more
//    information, see Access Control List (ACL) Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html).
//
// You can use either a canned ACL or specify access permissions explicitly.
// You cannot do both.
//
// Server-Side- Encryption-Specific Request Headers
//
// To encrypt the target object, you must provide the appropriate encryption-related
// request headers. The one you use depends on whether you want to use AWS-managed
// encryption keys or provide your own encryption key.
//
//    * To encrypt the target object using server-side encryption with an AWS-managed
//    encryption key, provide the following request headers, as appropriate.
//    x-amz-server-side​-encryption x-amz-server-side-encryption-aws-kms-key-id
//    x-amz-server-side-encryption-context If you specify x-amz-server-side-encryption:aws:kms,
//    but don't provide x-amz-server-side- encryption-aws-kms-key-id, Amazon
//    S3 uses the AWS managed customer master key (CMK) in KMS to protect the
//    data. All GET and PUT requests for an object protected by AWS KMS fail
//    if you don't make them with SSL or by using SigV4. For more information
//    on Server-Side Encryption with CMKs stored in Amazon KMS (SSE-KMS), see
//    Protecting Data Using Server-Side Encryption with CMKs stored in KMS (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html).
//
//    * To encrypt the target object using server-side encryption with an encryption
//    key that you provide, use the following headers. x-amz-server-side​-encryption​-customer-algorithm
//    x-amz-server-side​-encryption​-customer-key x-amz-server-side​-encryption​-customer-key-MD5
//
//    * If the source object is encrypted using server-side encryption with
//    customer-provided encryption keys, you must use the following headers.
//    x-amz-copy-source​-server-side​-encryption​-customer-algorithm x-amz-copy-source​-server-side​-encryption​-customer-key
//    x-amz-copy-source-​server-side​-encryption​-customer-key-MD5 For
//    more information on Server-Side Encryption with CMKs stored in Amazon
//    KMS (SSE-KMS), see Protecting Data Using Server-Side Encryption with CMKs
//    stored in Amazon KMS (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html).
//
// Access-Control-List (ACL)-Specific Request Headers
//
// You also can use the following access control–related headers with this
// operation. By default, all objects are private. Only the owner has full access
// control. When adding a new object, you can grant permissions to individual
// AWS accounts or to predefined groups defined by Amazon S3. These permissions
// are then added to the Access Control List (ACL) on the object. For more information,
// see Using ACLs (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html).
// With this operation, you can grant access permissions using one of the following
// two methods:
//
//    * Specify a canned ACL (x-amz-acl) — Amazon S3 supports a set of predefined
//    ACLs, known as canned ACLs. Each canned ACL has a predefined set of grantees
//    and permissions. For more information, see Canned ACL (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
//
//    * Specify access permissions explicitly — To explicitly grant access
//    permissions to specific AWS accounts or groups, use the following headers.
//    Each header maps to specific permissions that Amazon S3 supports in an
//    ACL. For more information, see Access Control List (ACL) Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html).
//    In the header, you specify a list of grantees who get the specific permission.
//    To grant permissions explicitly use: x-amz-grant-read x-amz-grant-write
//    x-amz-grant-read-acp x-amz-grant-write-acp x-amz-grant-full-control You
//    specify each grantee as a type=value pair, where the type is one of the
//    following: emailAddress – if the value specified is the email address
//    of an AWS account id – if the value specified is the canonical user
//    ID of an AWS account uri – if you are granting permissions to a predefined
//    group For example, the following x-amz-grant-read header grants the AWS
//    accounts identified by email addresses permissions to read object data
//    and its metadata: x-amz-grant-read: emailAddress="xyz@amazon.com", emailAddress="abc@amazon.com"
//
// The following operation are related to CopyObject
//
//    * PutObject
//
//    * GetObject
//
// For more information, see Copying Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/CopyingObjectsExamples.html).
//
//    // Example sending a request using CopyObjectRequest.
//    req := client.CopyObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CopyObject
func (c *Client) CopyObjectRequest(input *types.CopyObjectInput) CopyObjectRequest {
	op := &aws.Operation{
		Name:       opCopyObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &types.CopyObjectInput{}
	}

	req := c.newRequest(op, input, &types.CopyObjectOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.CopyObjectMarshaler{Input: input}.GetNamedBuildHandler())

	return CopyObjectRequest{Request: req, Input: input, Copy: c.CopyObjectRequest}
}

// CopyObjectRequest is the request type for the
// CopyObject API operation.
type CopyObjectRequest struct {
	*aws.Request
	Input *types.CopyObjectInput
	Copy  func(*types.CopyObjectInput) CopyObjectRequest
}

// Send marshals and sends the CopyObject API request.
func (r CopyObjectRequest) Send(ctx context.Context) (*CopyObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyObjectResponse{
		CopyObjectOutput: r.Request.Data.(*types.CopyObjectOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyObjectResponse is the response type for the
// CopyObject API operation.
type CopyObjectResponse struct {
	*types.CopyObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyObject request.
func (r *CopyObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
