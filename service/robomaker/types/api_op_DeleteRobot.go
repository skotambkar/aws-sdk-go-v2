// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRobotInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the robot.
	//
	// Robot is a required field
	Robot *string `locationName:"robot" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRobotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRobotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRobotInput"}

	if s.Robot == nil {
		invalidParams.Add(aws.NewErrParamRequired("Robot"))
	}
	if s.Robot != nil && len(*s.Robot) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Robot", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRobotOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRobotOutput) String() string {
	return awsutil.Prettify(s)
}
