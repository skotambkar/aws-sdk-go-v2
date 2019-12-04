// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sesv2iface provides an interface to enable mocking the Amazon Simple Email Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sesv2iface

import (
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

// ClientAPI provides an interface to enable mocking the
// sesv2.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon SES V2.
//    func myFunc(svc sesv2iface.ClientAPI) bool {
//        // Make svc.CreateConfigurationSet request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := sesv2.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        sesv2iface.ClientPI
//    }
//    func (m *mockClientClient) CreateConfigurationSet(input *sesv2.CreateConfigurationSetInput) (*sesv2.CreateConfigurationSetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CreateConfigurationSetRequest(*sesv2.CreateConfigurationSetInput) sesv2.CreateConfigurationSetRequest

	CreateConfigurationSetEventDestinationRequest(*sesv2.CreateConfigurationSetEventDestinationInput) sesv2.CreateConfigurationSetEventDestinationRequest

	CreateDedicatedIpPoolRequest(*sesv2.CreateDedicatedIpPoolInput) sesv2.CreateDedicatedIpPoolRequest

	CreateDeliverabilityTestReportRequest(*sesv2.CreateDeliverabilityTestReportInput) sesv2.CreateDeliverabilityTestReportRequest

	CreateEmailIdentityRequest(*sesv2.CreateEmailIdentityInput) sesv2.CreateEmailIdentityRequest

	DeleteConfigurationSetRequest(*sesv2.DeleteConfigurationSetInput) sesv2.DeleteConfigurationSetRequest

	DeleteConfigurationSetEventDestinationRequest(*sesv2.DeleteConfigurationSetEventDestinationInput) sesv2.DeleteConfigurationSetEventDestinationRequest

	DeleteDedicatedIpPoolRequest(*sesv2.DeleteDedicatedIpPoolInput) sesv2.DeleteDedicatedIpPoolRequest

	DeleteEmailIdentityRequest(*sesv2.DeleteEmailIdentityInput) sesv2.DeleteEmailIdentityRequest

	GetAccountRequest(*sesv2.GetAccountInput) sesv2.GetAccountRequest

	GetBlacklistReportsRequest(*sesv2.GetBlacklistReportsInput) sesv2.GetBlacklistReportsRequest

	GetConfigurationSetRequest(*sesv2.GetConfigurationSetInput) sesv2.GetConfigurationSetRequest

	GetConfigurationSetEventDestinationsRequest(*sesv2.GetConfigurationSetEventDestinationsInput) sesv2.GetConfigurationSetEventDestinationsRequest

	GetDedicatedIpRequest(*sesv2.GetDedicatedIpInput) sesv2.GetDedicatedIpRequest

	GetDedicatedIpsRequest(*sesv2.GetDedicatedIpsInput) sesv2.GetDedicatedIpsRequest

	GetDeliverabilityDashboardOptionsRequest(*sesv2.GetDeliverabilityDashboardOptionsInput) sesv2.GetDeliverabilityDashboardOptionsRequest

	GetDeliverabilityTestReportRequest(*sesv2.GetDeliverabilityTestReportInput) sesv2.GetDeliverabilityTestReportRequest

	GetDomainDeliverabilityCampaignRequest(*sesv2.GetDomainDeliverabilityCampaignInput) sesv2.GetDomainDeliverabilityCampaignRequest

	GetDomainStatisticsReportRequest(*sesv2.GetDomainStatisticsReportInput) sesv2.GetDomainStatisticsReportRequest

	GetEmailIdentityRequest(*sesv2.GetEmailIdentityInput) sesv2.GetEmailIdentityRequest

	ListConfigurationSetsRequest(*sesv2.ListConfigurationSetsInput) sesv2.ListConfigurationSetsRequest

	ListDedicatedIpPoolsRequest(*sesv2.ListDedicatedIpPoolsInput) sesv2.ListDedicatedIpPoolsRequest

	ListDeliverabilityTestReportsRequest(*sesv2.ListDeliverabilityTestReportsInput) sesv2.ListDeliverabilityTestReportsRequest

	ListDomainDeliverabilityCampaignsRequest(*sesv2.ListDomainDeliverabilityCampaignsInput) sesv2.ListDomainDeliverabilityCampaignsRequest

	ListEmailIdentitiesRequest(*sesv2.ListEmailIdentitiesInput) sesv2.ListEmailIdentitiesRequest

	ListTagsForResourceRequest(*sesv2.ListTagsForResourceInput) sesv2.ListTagsForResourceRequest

	PutAccountDedicatedIpWarmupAttributesRequest(*sesv2.PutAccountDedicatedIpWarmupAttributesInput) sesv2.PutAccountDedicatedIpWarmupAttributesRequest

	PutAccountSendingAttributesRequest(*sesv2.PutAccountSendingAttributesInput) sesv2.PutAccountSendingAttributesRequest

	PutConfigurationSetDeliveryOptionsRequest(*sesv2.PutConfigurationSetDeliveryOptionsInput) sesv2.PutConfigurationSetDeliveryOptionsRequest

	PutConfigurationSetReputationOptionsRequest(*sesv2.PutConfigurationSetReputationOptionsInput) sesv2.PutConfigurationSetReputationOptionsRequest

	PutConfigurationSetSendingOptionsRequest(*sesv2.PutConfigurationSetSendingOptionsInput) sesv2.PutConfigurationSetSendingOptionsRequest

	PutConfigurationSetTrackingOptionsRequest(*sesv2.PutConfigurationSetTrackingOptionsInput) sesv2.PutConfigurationSetTrackingOptionsRequest

	PutDedicatedIpInPoolRequest(*sesv2.PutDedicatedIpInPoolInput) sesv2.PutDedicatedIpInPoolRequest

	PutDedicatedIpWarmupAttributesRequest(*sesv2.PutDedicatedIpWarmupAttributesInput) sesv2.PutDedicatedIpWarmupAttributesRequest

	PutDeliverabilityDashboardOptionRequest(*sesv2.PutDeliverabilityDashboardOptionInput) sesv2.PutDeliverabilityDashboardOptionRequest

	PutEmailIdentityDkimAttributesRequest(*sesv2.PutEmailIdentityDkimAttributesInput) sesv2.PutEmailIdentityDkimAttributesRequest

	PutEmailIdentityFeedbackAttributesRequest(*sesv2.PutEmailIdentityFeedbackAttributesInput) sesv2.PutEmailIdentityFeedbackAttributesRequest

	PutEmailIdentityMailFromAttributesRequest(*sesv2.PutEmailIdentityMailFromAttributesInput) sesv2.PutEmailIdentityMailFromAttributesRequest

	SendEmailRequest(*sesv2.SendEmailInput) sesv2.SendEmailRequest

	TagResourceRequest(*sesv2.TagResourceInput) sesv2.TagResourceRequest

	UntagResourceRequest(*sesv2.UntagResourceInput) sesv2.UntagResourceRequest

	UpdateConfigurationSetEventDestinationRequest(*sesv2.UpdateConfigurationSetEventDestinationInput) sesv2.UpdateConfigurationSetEventDestinationRequest
}

var _ ClientAPI = (*sesv2.Client)(nil)
