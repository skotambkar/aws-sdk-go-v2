// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideosignaling

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// A structure for the ICE server connection data.
type IceServer struct {
	_ struct{} `type:"structure"`

	// A password to login to the ICE server.
	Password *string `min:"1" type:"string"`

	// The period of time, in seconds, during which the username and password are
	// valid.
	Ttl *int64 `min:"30" type:"integer"`

	// An array of URIs, in the form specified in the I-D.petithuguenin-behave-turn-uris
	// (https://tools.ietf.org/html/draft-petithuguenin-behave-turn-uris-03) spec.
	// These URIs provide the different addresses and/or protocols that can be used
	// to reach the TURN server.
	Uris []string `type:"list"`

	// A username to login to the ICE server.
	Username *string `min:"1" type:"string"`
}

// String returns the string representation
func (s IceServer) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s IceServer) MarshalFields(e protocol.FieldEncoder) error {
	if s.Password != nil {
		v := *s.Password

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Password", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Ttl != nil {
		v := *s.Ttl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Ttl", protocol.Int64Value(v), metadata)
	}
	if s.Uris != nil {
		v := s.Uris

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Uris", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Username != nil {
		v := *s.Username

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Username", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
