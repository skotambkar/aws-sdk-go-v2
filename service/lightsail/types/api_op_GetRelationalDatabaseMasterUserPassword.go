// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/enums"
)

type GetRelationalDatabaseMasterUserPasswordInput struct {
	_ struct{} `type:"structure"`

	// The password version to return.
	//
	// Specifying CURRENT or PREVIOUS returns the current or previous passwords
	// respectively. Specifying PENDING returns the newest version of the password
	// that will rotate to CURRENT. After the PENDING password rotates to CURRENT,
	// the PENDING password is no longer available.
	//
	// Default: CURRENT
	PasswordVersion enums.RelationalDatabasePasswordVersion `locationName:"passwordVersion" type:"string" enum:"true"`

	// The name of your database for which to get the master user password.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRelationalDatabaseMasterUserPasswordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRelationalDatabaseMasterUserPasswordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRelationalDatabaseMasterUserPasswordInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRelationalDatabaseMasterUserPasswordOutput struct {
	_ struct{} `type:"structure"`

	// The timestamp when the specified version of the master user password was
	// created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp"`

	// The master user password for the password version specified.
	MasterUserPassword *string `locationName:"masterUserPassword" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s GetRelationalDatabaseMasterUserPasswordOutput) String() string {
	return awsutil.Prettify(s)
}
