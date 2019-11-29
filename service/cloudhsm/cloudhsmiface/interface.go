// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudhsmiface provides an interface to enable mocking the Amazon CloudHSM service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudhsmiface

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm/types"
)

// ClientAPI provides an interface to enable mocking the
// cloudhsm.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // CloudHSM.
//    func myFunc(svc cloudhsmiface.ClientAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cloudhsm.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        cloudhsmiface.ClientPI
//    }
//    func (m *mockClientClient) AddTagsToResource(input *types.AddTagsToResourceInput) (*types.AddTagsToResourceOutput, error) {
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
	AddTagsToResourceRequest(*types.AddTagsToResourceInput) cloudhsm.AddTagsToResourceRequest

	CreateHapgRequest(*types.CreateHapgInput) cloudhsm.CreateHapgRequest

	CreateHsmRequest(*types.CreateHsmInput) cloudhsm.CreateHsmRequest

	CreateLunaClientRequest(*types.CreateLunaClientInput) cloudhsm.CreateLunaClientRequest

	DeleteHapgRequest(*types.DeleteHapgInput) cloudhsm.DeleteHapgRequest

	DeleteHsmRequest(*types.DeleteHsmInput) cloudhsm.DeleteHsmRequest

	DeleteLunaClientRequest(*types.DeleteLunaClientInput) cloudhsm.DeleteLunaClientRequest

	DescribeHapgRequest(*types.DescribeHapgInput) cloudhsm.DescribeHapgRequest

	DescribeHsmRequest(*types.DescribeHsmInput) cloudhsm.DescribeHsmRequest

	DescribeLunaClientRequest(*types.DescribeLunaClientInput) cloudhsm.DescribeLunaClientRequest

	GetConfigRequest(*types.GetConfigInput) cloudhsm.GetConfigRequest

	ListAvailableZonesRequest(*types.ListAvailableZonesInput) cloudhsm.ListAvailableZonesRequest

	ListHapgsRequest(*types.ListHapgsInput) cloudhsm.ListHapgsRequest

	ListHsmsRequest(*types.ListHsmsInput) cloudhsm.ListHsmsRequest

	ListLunaClientsRequest(*types.ListLunaClientsInput) cloudhsm.ListLunaClientsRequest

	ListTagsForResourceRequest(*types.ListTagsForResourceInput) cloudhsm.ListTagsForResourceRequest

	ModifyHapgRequest(*types.ModifyHapgInput) cloudhsm.ModifyHapgRequest

	ModifyHsmRequest(*types.ModifyHsmInput) cloudhsm.ModifyHsmRequest

	ModifyLunaClientRequest(*types.ModifyLunaClientInput) cloudhsm.ModifyLunaClientRequest

	RemoveTagsFromResourceRequest(*types.RemoveTagsFromResourceInput) cloudhsm.RemoveTagsFromResourceRequest
}

var _ ClientAPI = (*cloudhsm.Client)(nil)
