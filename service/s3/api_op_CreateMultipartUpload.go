// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opCreateMultipartUpload = "CreateMultipartUpload"

// CreateMultipartUploadRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This operation initiates a multipart upload and returns an upload ID. This
// upload ID is used to associate all of the parts in the specific multipart
// upload. You specify this upload ID in each of your subsequent upload part
// requests (see UploadPart). You also include this upload ID in the final request
// to either complete or abort the multipart upload request.
//
// For more information about multipart uploads, see Multipart Upload Overview
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html).
//
// If you have configured a lifecycle rule to abort incomplete multipart uploads,
// the upload must complete within the number of days specified in the bucket
// lifecycle configuration. Otherwise, the incomplete multipart upload becomes
// eligible for an abort operation and Amazon S3 aborts the multipart upload.
// For more information, see Aborting Incomplete Multipart Uploads Using a Bucket
// Lifecycle Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config).
//
// For information about the permissions required to use the multipart upload
// API, see Multipart Upload API and Permissions (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html).
//
// For request signing, multipart upload is just a series of regular requests.
// You initiate a multipart upload, send one or more requests to upload parts,
// and then complete the multipart upload process. You sign each request individually.
// There is nothing special about signing multipart upload requests. For more
// information about signing, see Authenticating Requests (AWS Signature Version
// 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html).
//
// After you initiate a multipart upload and upload one or more parts, to stop
// being charged for storing the uploaded parts, you must either complete or
// abort the multipart upload. Amazon S3 frees up the space used to store the
// parts and stop charging you for storing them only after you either complete
// or abort a multipart upload.
//
// You can optionally request server-side encryption. For server-side encryption,
// Amazon S3 encrypts your data as it writes it to disks in its data centers
// and decrypts it when you access it. You can provide your own encryption key,
// or use AWS Key Management Service (AWS KMS) customer master keys (CMKs) or
// Amazon S3-managed encryption keys. If you choose to provide your own encryption
// key, the request headers you provide in UploadPart) and UploadPartCopy) requests
// must match the headers you used in the request to initiate the upload by
// using CreateMultipartUpload.
//
// To perform a multipart upload with encryption using an AWS KMS CMK, the requester
// must have permission to the kms:Encrypt, kms:Decrypt, kms:ReEncrypt*, kms:GenerateDataKey*,
// and kms:DescribeKey actions on the key. These permissions are required because
// Amazon S3 must decrypt and read data from the encrypted file parts before
// it completes the multipart upload.
//
// If your AWS Identity and Access Management (IAM) user or role is in the same
// AWS account as the AWS KMS CMK, then you must have these permissions on the
// key policy. If your IAM user or role belongs to a different account than
// the key, then you must have the permissions on both the key policy and your
// IAM user or role.
//
// For more information, see Protecting Data Using Server-Side Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html).
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
// You can optionally tell Amazon S3 to encrypt data at rest using server-side
// encryption. Server-side encryption is for data encryption at rest. Amazon
// S3 encrypts your data as it writes it to disks in its data centers and decrypts
// it when you access it. The option you use depends on whether you want to
// use AWS-managed encryption keys or provide your own encryption key.
//
//    * Use encryption keys managed by Amazon S3 or customer master keys (CMKs)
//    stored in Amazon Key Management Service (KMS) – If you want AWS to manage
//    the keys used to encrypt data, specify the following headers in the request.
//    x-amz-server-side​-encryption x-amz-server-side-encryption-aws-kms-key-id
//    x-amz-server-side-encryption-context If you specify x-amz-server-side-encryption:aws:kms,
//    but don't provide x-amz-server-side- encryption-aws-kms-key-id, Amazon
//    S3 uses the AWS managed CMK in AWS KMS to protect the data. All GET and
//    PUT requests for an object protected by AWS KMS fail if you don't make
//    them with SSL or by using SigV4. For more information on Server-Side Encryption
//    with CMKs Stored in Amazon KMS (SSE-KMS), see Protecting Data Using Server-Side
//    Encryption with CMKs stored in AWS KMS (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html).
//
//    * Use customer-provided encryption keys – If you want to manage your
//    own encryption keys, provide all the following headers in the request.
//    x-amz-server-side​-encryption​-customer-algorithm x-amz-server-side​-encryption​-customer-key
//    x-amz-server-side​-encryption​-customer-key-MD5 For more information
//    on Server-Side Encryption with CMKs stored in AWS KMS (SSE-KMS), see Protecting
//    Data Using Server-Side Encryption with CMKs stored in AWS KMS (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html).
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
// The following operations are related to CreateMultipartUpload:
//
//    * UploadPart
//
//    * CompleteMultipartUpload
//
//    * AbortMultipartUpload
//
//    * ListParts
//
//    * ListMultipartUploads
//
//    // Example sending a request using CreateMultipartUploadRequest.
//    req := client.CreateMultipartUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CreateMultipartUpload
func (c *Client) CreateMultipartUploadRequest(input *types.CreateMultipartUploadInput) CreateMultipartUploadRequest {
	op := &aws.Operation{
		Name:       opCreateMultipartUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}/{Key+}?uploads",
	}

	if input == nil {
		input = &types.CreateMultipartUploadInput{}
	}

	req := c.newRequest(op, input, &types.CreateMultipartUploadOutput{})
	return CreateMultipartUploadRequest{Request: req, Input: input, Copy: c.CreateMultipartUploadRequest}
}

// CreateMultipartUploadRequest is the request type for the
// CreateMultipartUpload API operation.
type CreateMultipartUploadRequest struct {
	*aws.Request
	Input *types.CreateMultipartUploadInput
	Copy  func(*types.CreateMultipartUploadInput) CreateMultipartUploadRequest
}

// Send marshals and sends the CreateMultipartUpload API request.
func (r CreateMultipartUploadRequest) Send(ctx context.Context) (*CreateMultipartUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMultipartUploadResponse{
		CreateMultipartUploadOutput: r.Request.Data.(*types.CreateMultipartUploadOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMultipartUploadResponse is the response type for the
// CreateMultipartUpload API operation.
type CreateMultipartUploadResponse struct {
	*types.CreateMultipartUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMultipartUpload request.
func (r *CreateMultipartUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
