// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Provides options for downloading output of an Amazon S3 Glacier job.
type GetJobOutputInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The job ID whose data is downloaded.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" type:"string" required:"true"`

	// The range of bytes to retrieve from the output. For example, if you want
	// to download the first 1,048,576 bytes, specify the range as bytes=0-1048575.
	// By default, this operation downloads the entire output.
	//
	// If the job output is large, then you can use a range to retrieve a portion
	// of the output. This allows you to download the entire output in smaller chunks
	// of bytes. For example, suppose you have 1 GB of job output you want to download
	// and you decide to download 128 MB chunks of data at a time, which is a total
	// of eight Get Job Output requests. You use the following process to download
	// the job output:
	//
	// Download a 128 MB chunk of output by specifying the appropriate byte range.
	// Verify that all 128 MB of data was received.
	//
	// Along with the data, the response includes a SHA256 tree hash of the payload.
	// You compute the checksum of the payload on the client and compare it with
	// the checksum you received in the response to ensure you received all the
	// expected data.
	//
	// Repeat steps 1 and 2 for all the eight 128 MB chunks of output data, each
	// time specifying the appropriate byte range.
	//
	// After downloading all the parts of the job output, you have a list of eight
	// checksum values. Compute the tree hash of these values to find the checksum
	// of the entire output. Using the DescribeJob API, obtain job information of
	// the job that provided you the output. The response includes the checksum
	// of the entire archive stored in Amazon S3 Glacier. You compare this value
	// with the checksum you computed to ensure you have downloaded the entire archive
	// content with no errors.
	Range *string `location:"header" locationName:"Range" type:"string"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJobOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJobOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJobOutputInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the Amazon S3 Glacier response to your request.
type GetJobOutputOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Indicates the range units accepted. For more information, see RFC2616 (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	AcceptRanges *string `location:"header" locationName:"Accept-Ranges" type:"string"`

	// The description of an archive.
	ArchiveDescription *string `location:"header" locationName:"x-amz-archive-description" type:"string"`

	// The job data, either archive data or inventory data.
	Body io.ReadCloser `locationName:"body" type:"blob"`

	// The checksum of the data in the response. This header is returned only when
	// retrieving the output for an archive retrieval job. Furthermore, this header
	// appears only under the following conditions:
	//
	//    * You get the entire range of the archive.
	//
	//    * You request a range to return of the archive that starts and ends on
	//    a multiple of 1 MB. For example, if you have an 3.1 MB archive and you
	//    specify a range to return that starts at 1 MB and ends at 2 MB, then the
	//    x-amz-sha256-tree-hash is returned as a response header.
	//
	//    * You request a range of the archive to return that starts on a multiple
	//    of 1 MB and goes to the end of the archive. For example, if you have a
	//    3.1 MB archive and you specify a range that starts at 2 MB and ends at
	//    3.1 MB (the end of the archive), then the x-amz-sha256-tree-hash is returned
	//    as a response header.
	Checksum *string `location:"header" locationName:"x-amz-sha256-tree-hash" type:"string"`

	// The range of bytes returned by Amazon S3 Glacier. If only partial output
	// is downloaded, the response provides the range of bytes Amazon S3 Glacier
	// returned. For example, bytes 0-1048575/8388608 returns the first 1 MB from
	// 8 MB.
	ContentRange *string `location:"header" locationName:"Content-Range" type:"string"`

	// The Content-Type depends on whether the job output is an archive or a vault
	// inventory. For archive data, the Content-Type is application/octet-stream.
	// For vault inventory, if you requested CSV format when you initiated the job,
	// the Content-Type is text/csv. Otherwise, by default, vault inventory is returned
	// as JSON, and the Content-Type is application/json.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The HTTP response code for a job output request. The value depends on whether
	// a range was specified in the request.
	Status *int64 `location:"statusCode" locationName:"status" type:"integer"`
}

// String returns the string representation
func (s GetJobOutputOutput) String() string {
	return awsutil.Prettify(s)
}
