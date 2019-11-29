// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type InviteMembersInput struct {
	_ struct{} `type:"structure"`

	// A list of account IDs of the accounts that you want to invite to GuardDuty
	// as members.
	//
	// AccountIds is a required field
	AccountIds []string `locationName:"accountIds" min:"1" type:"list" required:"true"`

	// The unique ID of the detector of the GuardDuty account with which you want
	// to invite members.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// A boolean value that specifies whether you want to disable email notification
	// to the accounts that you’re inviting to GuardDuty as members.
	DisableEmailNotification *bool `locationName:"disableEmailNotification" type:"boolean"`

	// The invitation message that you want to send to the accounts that you’re
	// inviting to GuardDuty as members.
	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s InviteMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InviteMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InviteMembersInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}
	if s.AccountIds != nil && len(s.AccountIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountIds", 1))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type InviteMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects containing the unprocessed account and a result string
	// explaining why it was unprocessed.
	//
	// UnprocessedAccounts is a required field
	UnprocessedAccounts []UnprocessedAccount `locationName:"unprocessedAccounts" type:"list" required:"true"`
}

// String returns the string representation
func (s InviteMembersOutput) String() string {
	return awsutil.Prettify(s)
}