// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lexmodelbuildingserviceiface provides an interface to enable mocking the Amazon Lex Model Building Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lexmodelbuildingserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
)

// ClientAPI provides an interface to enable mocking the
// lexmodelbuildingservice.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Lex Model Building Service.
//    func myFunc(svc lexmodelbuildingserviceiface.ClientAPI) bool {
//        // Make svc.CreateBotVersion request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := lexmodelbuildingservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        lexmodelbuildingserviceiface.ClientPI
//    }
//    func (m *mockClientClient) CreateBotVersion(input *types.CreateBotVersionInput) (*types.CreateBotVersionOutput, error) {
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
	CreateBotVersionRequest(*types.CreateBotVersionInput) lexmodelbuildingservice.CreateBotVersionRequest

	CreateIntentVersionRequest(*types.CreateIntentVersionInput) lexmodelbuildingservice.CreateIntentVersionRequest

	CreateSlotTypeVersionRequest(*types.CreateSlotTypeVersionInput) lexmodelbuildingservice.CreateSlotTypeVersionRequest

	DeleteBotRequest(*types.DeleteBotInput) lexmodelbuildingservice.DeleteBotRequest

	DeleteBotAliasRequest(*types.DeleteBotAliasInput) lexmodelbuildingservice.DeleteBotAliasRequest

	DeleteBotChannelAssociationRequest(*types.DeleteBotChannelAssociationInput) lexmodelbuildingservice.DeleteBotChannelAssociationRequest

	DeleteBotVersionRequest(*types.DeleteBotVersionInput) lexmodelbuildingservice.DeleteBotVersionRequest

	DeleteIntentRequest(*types.DeleteIntentInput) lexmodelbuildingservice.DeleteIntentRequest

	DeleteIntentVersionRequest(*types.DeleteIntentVersionInput) lexmodelbuildingservice.DeleteIntentVersionRequest

	DeleteSlotTypeRequest(*types.DeleteSlotTypeInput) lexmodelbuildingservice.DeleteSlotTypeRequest

	DeleteSlotTypeVersionRequest(*types.DeleteSlotTypeVersionInput) lexmodelbuildingservice.DeleteSlotTypeVersionRequest

	DeleteUtterancesRequest(*types.DeleteUtterancesInput) lexmodelbuildingservice.DeleteUtterancesRequest

	GetBotRequest(*types.GetBotInput) lexmodelbuildingservice.GetBotRequest

	GetBotAliasRequest(*types.GetBotAliasInput) lexmodelbuildingservice.GetBotAliasRequest

	GetBotAliasesRequest(*types.GetBotAliasesInput) lexmodelbuildingservice.GetBotAliasesRequest

	GetBotChannelAssociationRequest(*types.GetBotChannelAssociationInput) lexmodelbuildingservice.GetBotChannelAssociationRequest

	GetBotChannelAssociationsRequest(*types.GetBotChannelAssociationsInput) lexmodelbuildingservice.GetBotChannelAssociationsRequest

	GetBotVersionsRequest(*types.GetBotVersionsInput) lexmodelbuildingservice.GetBotVersionsRequest

	GetBotsRequest(*types.GetBotsInput) lexmodelbuildingservice.GetBotsRequest

	GetBuiltinIntentRequest(*types.GetBuiltinIntentInput) lexmodelbuildingservice.GetBuiltinIntentRequest

	GetBuiltinIntentsRequest(*types.GetBuiltinIntentsInput) lexmodelbuildingservice.GetBuiltinIntentsRequest

	GetBuiltinSlotTypesRequest(*types.GetBuiltinSlotTypesInput) lexmodelbuildingservice.GetBuiltinSlotTypesRequest

	GetExportRequest(*types.GetExportInput) lexmodelbuildingservice.GetExportRequest

	GetImportRequest(*types.GetImportInput) lexmodelbuildingservice.GetImportRequest

	GetIntentRequest(*types.GetIntentInput) lexmodelbuildingservice.GetIntentRequest

	GetIntentVersionsRequest(*types.GetIntentVersionsInput) lexmodelbuildingservice.GetIntentVersionsRequest

	GetIntentsRequest(*types.GetIntentsInput) lexmodelbuildingservice.GetIntentsRequest

	GetSlotTypeRequest(*types.GetSlotTypeInput) lexmodelbuildingservice.GetSlotTypeRequest

	GetSlotTypeVersionsRequest(*types.GetSlotTypeVersionsInput) lexmodelbuildingservice.GetSlotTypeVersionsRequest

	GetSlotTypesRequest(*types.GetSlotTypesInput) lexmodelbuildingservice.GetSlotTypesRequest

	GetUtterancesViewRequest(*types.GetUtterancesViewInput) lexmodelbuildingservice.GetUtterancesViewRequest

	PutBotRequest(*types.PutBotInput) lexmodelbuildingservice.PutBotRequest

	PutBotAliasRequest(*types.PutBotAliasInput) lexmodelbuildingservice.PutBotAliasRequest

	PutIntentRequest(*types.PutIntentInput) lexmodelbuildingservice.PutIntentRequest

	PutSlotTypeRequest(*types.PutSlotTypeInput) lexmodelbuildingservice.PutSlotTypeRequest

	StartImportRequest(*types.StartImportInput) lexmodelbuildingservice.StartImportRequest
}

var _ ClientAPI = (*lexmodelbuildingservice.Client)(nil)
