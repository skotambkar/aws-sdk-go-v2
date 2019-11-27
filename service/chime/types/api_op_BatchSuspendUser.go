// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchSuspendUserInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The request containing the user IDs to suspend.
	//
	// UserIdList is a required field
	UserIdList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchSuspendUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchSuspendUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchSuspendUserInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UserIdList == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserIdList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchSuspendUserOutput struct {
	_ struct{} `type:"structure"`

	// If the BatchSuspendUser action fails for one or more of the user IDs in the
	// request, a list of the user IDs is returned, along with error codes and error
	// messages.
	UserErrors []UserError `type:"list"`
}

// String returns the string representation
func (s BatchSuspendUserOutput) String() string {
	return awsutil.Prettify(s)
}
