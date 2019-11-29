// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssignTapePoolInput struct {
	_ struct{} `type:"structure"`

	// The ID of the pool that you want to add your tape to for archiving. The tape
	// in this pool is archived in the S3 storage class that is associated with
	// the pool. When you use your backup application to eject the tape, the tape
	// is archived directly into the storage class (Glacier or Deep Archive) that
	// corresponds to the pool.
	//
	// Valid values: "GLACIER", "DEEP_ARCHIVE"
	//
	// PoolId is a required field
	PoolId *string `min:"1" type:"string" required:"true"`

	// The unique Amazon Resource Name (ARN) of the virtual tape that you want to
	// add to the tape pool.
	//
	// TapeARN is a required field
	TapeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s AssignTapePoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssignTapePoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssignTapePoolInput"}

	if s.PoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PoolId"))
	}
	if s.PoolId != nil && len(*s.PoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PoolId", 1))
	}

	if s.TapeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TapeARN"))
	}
	if s.TapeARN != nil && len(*s.TapeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("TapeARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssignTapePoolOutput struct {
	_ struct{} `type:"structure"`

	// The unique Amazon Resource Names (ARN) of the virtual tape that was added
	// to the tape pool.
	TapeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s AssignTapePoolOutput) String() string {
	return awsutil.Prettify(s)
}