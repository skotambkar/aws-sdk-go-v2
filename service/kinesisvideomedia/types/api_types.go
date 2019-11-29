// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia/enums"
)

// Identifies the chunk on the Kinesis video stream where you want the GetMedia
// API to start returning media data. You have the following options to identify
// the starting chunk:
//
//    * Choose the latest (or oldest) chunk.
//
//    * Identify a specific chunk. You can identify a specific chunk either
//    by providing a fragment number or timestamp (server or producer).
//
//    * Each chunk's metadata includes a continuation token as a Matroska (MKV)
//    tag (AWS_KINESISVIDEO_CONTINUATION_TOKEN). If your previous GetMedia request
//    terminated, you can use this tag value in your next GetMedia request.
//    The API then starts returning chunks starting where the last API ended.
type StartSelector struct {
	_ struct{} `type:"structure"`

	// Specifies the fragment number from where you want the GetMedia API to start
	// returning the fragments.
	AfterFragmentNumber *string `min:"1" type:"string"`

	// Continuation token that Kinesis Video Streams returned in the previous GetMedia
	// response. The GetMedia API then starts with the chunk identified by the continuation
	// token.
	ContinuationToken *string `min:"1" type:"string"`

	// Identifies the fragment on the Kinesis video stream where you want to start
	// getting the data from.
	//
	//    * NOW - Start with the latest chunk on the stream.
	//
	//    * EARLIEST - Start with earliest available chunk on the stream.
	//
	//    * FRAGMENT_NUMBER - Start with the chunk after a specific fragment. You
	//    must also specify the AfterFragmentNumber parameter.
	//
	//    * PRODUCER_TIMESTAMP or SERVER_TIMESTAMP - Start with the chunk containing
	//    a fragment with the specified producer or server timestamp. You specify
	//    the timestamp by adding StartTimestamp.
	//
	//    * CONTINUATION_TOKEN - Read using the specified continuation token.
	//
	// If you choose the NOW, EARLIEST, or CONTINUATION_TOKEN as the startSelectorType,
	// you don't provide any additional information in the startSelector.
	//
	// StartSelectorType is a required field
	StartSelectorType enums.StartSelectorType `type:"string" required:"true" enum:"true"`

	// A timestamp value. This value is required if you choose the PRODUCER_TIMESTAMP
	// or the SERVER_TIMESTAMP as the startSelectorType. The GetMedia API then starts
	// with the chunk containing the fragment that has the specified timestamp.
	StartTimestamp *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s StartSelector) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSelector) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSelector"}
	if s.AfterFragmentNumber != nil && len(*s.AfterFragmentNumber) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AfterFragmentNumber", 1))
	}
	if s.ContinuationToken != nil && len(*s.ContinuationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContinuationToken", 1))
	}
	if len(s.StartSelectorType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("StartSelectorType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
