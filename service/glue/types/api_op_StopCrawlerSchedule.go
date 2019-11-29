// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopCrawlerScheduleInput struct {
	_ struct{} `type:"structure"`

	// Name of the crawler whose schedule state to set.
	//
	// CrawlerName is a required field
	CrawlerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopCrawlerScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopCrawlerScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopCrawlerScheduleInput"}

	if s.CrawlerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CrawlerName"))
	}
	if s.CrawlerName != nil && len(*s.CrawlerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CrawlerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopCrawlerScheduleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopCrawlerScheduleOutput) String() string {
	return awsutil.Prettify(s)
}