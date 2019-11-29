// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediastoredataiface provides an interface to enable mocking the AWS Elemental MediaStore Data Plane service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediastoredataiface

import (
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata/types"
)

// ClientAPI provides an interface to enable mocking the
// mediastoredata.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // MediaStore Data.
//    func myFunc(svc mediastoredataiface.ClientAPI) bool {
//        // Make svc.DeleteObject request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := mediastoredata.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        mediastoredataiface.ClientPI
//    }
//    func (m *mockClientClient) DeleteObject(input *types.DeleteObjectInput) (*types.DeleteObjectOutput, error) {
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
	DeleteObjectRequest(*types.DeleteObjectInput) mediastoredata.DeleteObjectRequest

	DescribeObjectRequest(*types.DescribeObjectInput) mediastoredata.DescribeObjectRequest

	GetObjectRequest(*types.GetObjectInput) mediastoredata.GetObjectRequest

	ListItemsRequest(*types.ListItemsInput) mediastoredata.ListItemsRequest

	PutObjectRequest(*types.PutObjectInput) mediastoredata.PutObjectRequest
}

var _ ClientAPI = (*mediastoredata.Client)(nil)
