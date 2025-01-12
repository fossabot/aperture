package status

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	statusv1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/common/status/v1"
)

// NewStatus creates a new instance of Status to be pushed into status registry. Use this function for creating status instead of by hand.
// It can either have a detail message or a detail error but not both. This is enforced by first checking for detail message to not be nil.
func NewStatus(d proto.Message, e error) *statusv1.Status {
	s := &statusv1.Status{
		Timestamp: timestamppb.Now(),
	}

	if d != nil {
		messageAny, err := anypb.New(d)
		if err != nil {
			return nil
		}
		s.Details = &statusv1.Status_Message{
			Message: messageAny,
		}
		return s
	}

	errorDetails := NewErrorDetails(e)
	s.Details = &statusv1.Status_Error{
		Error: errorDetails,
	}

	return s
}

// NewErrorDetails is a helper function to create a new instance of ErrorDetails.
func NewErrorDetails(e error) *statusv1.ErrorDetails {
	errorDetails := &statusv1.ErrorDetails{}

	if e != nil {
		msg := e.Error()
		if msg != "" {
			errorDetails.Message = e.Error()
		} else {
			errorDetails.Message = "Unknown error"
		}
	}

	return errorDetails
}
