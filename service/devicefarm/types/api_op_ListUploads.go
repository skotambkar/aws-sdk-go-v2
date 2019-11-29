// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/enums"
)

// Represents a request to the list uploads operation.
type ListUploadsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the project for which you want to list
	// uploads.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// The type of upload.
	//
	// Must be one of the following values:
	//
	//    * ANDROID_APP: An Android upload.
	//
	//    * IOS_APP: An iOS upload.
	//
	//    * WEB_APP: A web application upload.
	//
	//    * EXTERNAL_DATA: An external data upload.
	//
	//    * APPIUM_JAVA_JUNIT_TEST_PACKAGE: An Appium Java JUnit test package upload.
	//
	//    * APPIUM_JAVA_TESTNG_TEST_PACKAGE: An Appium Java TestNG test package
	//    upload.
	//
	//    * APPIUM_PYTHON_TEST_PACKAGE: An Appium Python test package upload.
	//
	//    * APPIUM_NODE_TEST_PACKAGE: An Appium Node.js test package upload.
	//
	//    * APPIUM_RUBY_TEST_PACKAGE: An Appium Ruby test package upload.
	//
	//    * APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE: An Appium Java JUnit test package
	//    upload for a web app.
	//
	//    * APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE: An Appium Java TestNG test package
	//    upload for a web app.
	//
	//    * APPIUM_WEB_PYTHON_TEST_PACKAGE: An Appium Python test package upload
	//    for a web app.
	//
	//    * APPIUM_WEB_NODE_TEST_PACKAGE: An Appium Node.js test package upload
	//    for a web app.
	//
	//    * APPIUM_WEB_RUBY_TEST_PACKAGE: An Appium Ruby test package upload for
	//    a web app.
	//
	//    * CALABASH_TEST_PACKAGE: A Calabash test package upload.
	//
	//    * INSTRUMENTATION_TEST_PACKAGE: An instrumentation upload.
	//
	//    * UIAUTOMATION_TEST_PACKAGE: A uiautomation test package upload.
	//
	//    * UIAUTOMATOR_TEST_PACKAGE: A uiautomator test package upload.
	//
	//    * XCTEST_TEST_PACKAGE: An Xcode test package upload.
	//
	//    * XCTEST_UI_TEST_PACKAGE: An Xcode UI test package upload.
	//
	//    * APPIUM_JAVA_JUNIT_TEST_SPEC: An Appium Java JUnit test spec upload.
	//
	//    * APPIUM_JAVA_TESTNG_TEST_SPEC: An Appium Java TestNG test spec upload.
	//
	//    * APPIUM_PYTHON_TEST_SPEC: An Appium Python test spec upload.
	//
	//    * APPIUM_NODE_TEST_SPEC: An Appium Node.js test spec upload.
	//
	//    * APPIUM_RUBY_TEST_SPEC: An Appium Ruby test spec upload.
	//
	//    * APPIUM_WEB_JAVA_JUNIT_TEST_SPEC: An Appium Java JUnit test spec upload
	//    for a web app.
	//
	//    * APPIUM_WEB_JAVA_TESTNG_TEST_SPEC: An Appium Java TestNG test spec upload
	//    for a web app.
	//
	//    * APPIUM_WEB_PYTHON_TEST_SPEC: An Appium Python test spec upload for a
	//    web app.
	//
	//    * APPIUM_WEB_NODE_TEST_SPEC: An Appium Node.js test spec upload for a
	//    web app.
	//
	//    * APPIUM_WEB_RUBY_TEST_SPEC: An Appium Ruby test spec upload for a web
	//    app.
	//
	//    * INSTRUMENTATION_TEST_SPEC: An instrumentation test spec upload.
	//
	//    * XCTEST_UI_TEST_SPEC: An Xcode UI test spec upload.
	Type enums.UploadType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListUploadsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUploadsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUploadsInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a list uploads request.
type ListUploadsOutput struct {
	_ struct{} `type:"structure"`

	// If the number of items that are returned is significantly large, this is
	// an identifier that is also returned, which can be used in a subsequent call
	// to this operation to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// Information about the uploads.
	Uploads []Upload `locationName:"uploads" type:"list"`
}

// String returns the string representation
func (s ListUploadsOutput) String() string {
	return awsutil.Prettify(s)
}
