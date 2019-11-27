// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDBParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB parameter group.
	//
	// Constraints:
	//
	//    * Must be the name of an existing DB parameter group
	//
	//    * You can't delete a default DB parameter group
	//
	//    * Can't be associated with any DB instances
	//
	// DBParameterGroupName is a required field
	DBParameterGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBParameterGroupInput"}

	if s.DBParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDBParameterGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDBParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}
