// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type UpdateGameSessionQueueInput struct {
	_ struct{} `type:"structure"`

	// List of fleets that can be used to fulfill game session placement requests
	// in the queue. Fleets are identified by either a fleet ARN or a fleet alias
	// ARN. Destinations are listed in default preference order. When updating this
	// list, provide a complete list of destinations.
	Destinations []GameSessionQueueDestination `type:"list"`

	// Descriptive label that is associated with game session queue. Queue names
	// must be unique within each region.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Collection of latency policies to apply when processing game sessions placement
	// requests with player latency information. Multiple policies are evaluated
	// in order of the maximum latency value, starting with the lowest latency values.
	// With just one policy, it is enforced at the start of the game session placement
	// for the duration period. With multiple policies, each policy is enforced
	// consecutively for its duration period. For example, a queue might enforce
	// a 60-second policy followed by a 120-second policy, and then no policy for
	// the remainder of the placement. When updating policies, provide a complete
	// collection of policies.
	PlayerLatencyPolicies []PlayerLatencyPolicy `type:"list"`

	// Maximum time, in seconds, that a new game session placement request remains
	// in the queue. When a request exceeds this time, the game session placement
	// changes to a TIMED_OUT status.
	TimeoutInSeconds *int64 `type:"integer"`
}

// String returns the string representation
func (s UpdateGameSessionQueueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGameSessionQueueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGameSessionQueueInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Destinations != nil {
		for i, v := range s.Destinations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Destinations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type UpdateGameSessionQueueOutput struct {
	_ struct{} `type:"structure"`

	// Object that describes the newly updated game session queue.
	GameSessionQueue *GameSessionQueue `type:"structure"`
}

// String returns the string representation
func (s UpdateGameSessionQueueOutput) String() string {
	return awsutil.Prettify(s)
}