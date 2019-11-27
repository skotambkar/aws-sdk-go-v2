// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain
// that the Deliverability dashboard is enabled for (PutDeliverabilityDashboardOption
// operation).
type GetDomainDeliverabilityCampaignInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the campaign. Amazon Pinpoint automatically generates
	// and assigns this identifier to a campaign. This value is not the same as
	// the campaign identifier that Amazon Pinpoint assigns to campaigns that you
	// create and manage by using the Amazon Pinpoint API or the Amazon Pinpoint
	// console.
	//
	// CampaignId is a required field
	CampaignId *string `location:"uri" locationName:"CampaignId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDomainDeliverabilityCampaignInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainDeliverabilityCampaignInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainDeliverabilityCampaignInput"}

	if s.CampaignId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CampaignId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An object that contains all the deliverability data for a specific campaign.
// This data is available for a campaign only if the campaign sent email by
// using a domain that the Deliverability dashboard is enabled for (PutDeliverabilityDashboardOption
// operation).
type GetDomainDeliverabilityCampaignOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains the deliverability data for the campaign.
	//
	// DomainDeliverabilityCampaign is a required field
	DomainDeliverabilityCampaign *DomainDeliverabilityCampaign `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetDomainDeliverabilityCampaignOutput) String() string {
	return awsutil.Prettify(s)
}
