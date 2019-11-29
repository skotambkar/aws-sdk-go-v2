// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/s3/enums"
)

type ListObjectsV2Input struct {
	_ struct{} `type:"structure"`

	// Name of the bucket to list.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// ContinuationToken indicates Amazon S3 that the list is being continued on
	// this bucket with a token. ContinuationToken is obfuscated and is not a real
	// key.
	ContinuationToken *string `location:"querystring" locationName:"continuation-token" type:"string"`

	// A delimiter is a character you use to group keys.
	Delimiter *string `location:"querystring" locationName:"delimiter" type:"string"`

	// Encoding type used by Amazon S3 to encode object keys in the response.
	EncodingType enums.EncodingType `location:"querystring" locationName:"encoding-type" type:"string" enum:"true"`

	// The owner field is not present in listV2 by default, if you want to return
	// owner field with each key in the result then set the fetch owner field to
	// true
	FetchOwner *bool `location:"querystring" locationName:"fetch-owner" type:"boolean"`

	// Sets the maximum number of keys returned in the response. The response might
	// contain fewer keys but will never contain more.
	MaxKeys *int64 `location:"querystring" locationName:"max-keys" type:"integer"`

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string `location:"querystring" locationName:"prefix" type:"string"`

	// Confirms that the requester knows that she or he will be charged for the
	// list objects request in V2 style. Bucket owners need not specify this parameter
	// in their requests.
	RequestPayer enums.RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// StartAfter is where you want Amazon S3 to start listing from. Amazon S3 starts
	// listing after this specified key. StartAfter can be any key in the bucket.
	StartAfter *string `location:"querystring" locationName:"start-after" type:"string"`
}

// String returns the string representation
func (s ListObjectsV2Input) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListObjectsV2Input) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListObjectsV2Input"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *ListObjectsV2Input) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type ListObjectsV2Output struct {
	_ struct{} `type:"structure"`

	// All of the keys rolled up into a common prefix count as a single return when
	// calculating the number of returns.
	//
	// A response can contain CommonPrefixes only if you specify a delimiter.
	//
	// CommonPrefixes contains all (if there are any) keys between Prefix and the
	// next occurrence of the string specified by a delimiter.
	//
	// CommonPrefixes lists keys that act like subdirectories in the directory specified
	// by Prefix.
	//
	// For example, if the prefix is notes/ and the delimiter is a slash (/) as
	// in notes/summer/july, the common prefix is notes/summer/. All of the keys
	// that roll up into a common prefix count as a single return when calculating
	// the number of returns.
	CommonPrefixes []CommonPrefix `type:"list" flattened:"true"`

	// Metadata about each object returned.
	Contents []Object `type:"list" flattened:"true"`

	// If ContinuationToken was sent with the request, it is included in the response.
	ContinuationToken *string `type:"string"`

	// Causes keys that contain the same string between the prefix and the first
	// occurrence of the delimiter to be rolled up into a single result element
	// in the CommonPrefixes collection. These rolled-up keys are not returned elsewhere
	// in the response. Each rolled-up result counts as only one return against
	// the MaxKeys value.
	Delimiter *string `type:"string"`

	// Encoding type used by Amazon S3 to encode object key names in the XML response.
	//
	// If you specify the encoding-type request parameter, Amazon S3 includes this
	// element in the response, and returns encoded key name values in the following
	// response elements:
	//
	// Delimiter, Prefix, Key, and StartAfter.
	EncodingType enums.EncodingType `type:"string" enum:"true"`

	// Set to false if all of the results were returned. Set to true if more keys
	// are available to return. If the number of results exceeds that specified
	// by MaxKeys, all of the results might not be returned.
	IsTruncated *bool `type:"boolean"`

	// KeyCount is the number of keys returned with this request. KeyCount will
	// always be less than equals to MaxKeys field. Say you ask for 50 keys, your
	// result will include less than equals 50 keys
	KeyCount *int64 `type:"integer"`

	// Sets the maximum number of keys returned in the response. The response might
	// contain fewer keys but will never contain more.
	MaxKeys *int64 `type:"integer"`

	// Name of the bucket.
	Name *string `type:"string"`

	// NextContinuationToken is sent when isTruncated is true which means there
	// are more keys in the bucket that can be listed. The next list requests to
	// Amazon S3 can be continued with this NextContinuationToken. NextContinuationToken
	// is obfuscated and is not a real key
	NextContinuationToken *string `type:"string"`

	// Keys that begin with the indicated prefix.
	Prefix *string `type:"string"`

	// If StartAfter was sent with the request, it is included in the response.
	StartAfter *string `type:"string"`
}

// String returns the string representation
func (s ListObjectsV2Output) String() string {
	return awsutil.Prettify(s)
}