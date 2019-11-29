// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type DescribeMatchmakingInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a matchmaking ticket. You can include up to 10 ID values.
	//
	// TicketIds is a required field
	TicketIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeMatchmakingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMatchmakingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMatchmakingInput"}

	if s.TicketIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("TicketIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type DescribeMatchmakingOutput struct {
	_ struct{} `type:"structure"`

	// Collection of existing matchmaking ticket objects matching the request.
	TicketList []MatchmakingTicket `type:"list"`
}

// String returns the string representation
func (s DescribeMatchmakingOutput) String() string {
	return awsutil.Prettify(s)
}