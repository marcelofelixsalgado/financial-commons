package responses_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/marcelofelixsalgado/financial-commons/api/responses"
	"github.com/marcelofelixsalgado/financial-commons/api/responses/faults"
)

const fieldDoesNotSupportDecimals = "Field value does not support decimals"

func TestGetMessages(t *testing.T) {

	expectedMessage := ResponseMessage{
		ErrorCode:      "INVALID_REQUEST_SYNTAX",
		Message:        "Request is not well-formed, syntactically incorrect, or violates schema",
		HttpStatusCode: 400,
		Details: []ResponseMessageDetail{
			{
				Issue:       "DECIMALS_NOT_SUPPORTED",
				Description: fieldDoesNotSupportDecimals,
				Location:    "body",
				Field:       "field3",
				Value:       "value3",
			},
		},
	}

	actualMessage := NewResponseMessage()
	actualMessage.AddMessageByIssue(faults.DecimalsNotSupported, Body, "field3", "value3")

	if !reflect.DeepEqual(actualMessage.GetMessage(), expectedMessage) {
		t.Errorf(formatMessageDiff(expectedMessage, actualMessage))
	}
}

func TestGetMessagesMultipleDetais(t *testing.T) {

	expectedMessage := ResponseMessage{
		ErrorCode:      "INVALID_REQUEST_SYNTAX",
		Message:        "Request is not well-formed, syntactically incorrect, or violates schema",
		HttpStatusCode: 400,
		Details: []ResponseMessageDetail{
			{
				Issue:       "DECIMALS_NOT_SUPPORTED",
				Description: fieldDoesNotSupportDecimals,
				Location:    "body",
				Field:       "field1",
				Value:       "value1",
			},
			{
				Issue:       "DECIMALS_NOT_SUPPORTED",
				Description: fieldDoesNotSupportDecimals,
				Location:    "body",
				Field:       "field2",
				Value:       "value2",
			},
		},
	}

	actualMessage := NewResponseMessage()
	actualMessage.AddMessageByIssue(faults.DecimalsNotSupported, Body, "field1", "value1")
	actualMessage.AddMessageByIssue(faults.DecimalsNotSupported, Body, "field2", "value2")

	if !reflect.DeepEqual(actualMessage.GetMessage(), expectedMessage) {
		t.Errorf(formatMessageDiff(expectedMessage, actualMessage))
	}
}

func TestGetMessagesReplacementSuccess(t *testing.T) {

	expectedMessage := ResponseMessage{
		ErrorCode:      "UNPROCESSABLE_ENTITY",
		Message:        "The request is semantically incorrect or fails business validation",
		HttpStatusCode: 422,
		Details: []ResponseMessageDetail{
			{
				Issue:       "CONDITIONAL_FIELD_NOT_ALLOWED",
				Description: "field1 is not allowed when field field2 is set to value2",
				Location:    "body",
				Field:       "field1",
				Value:       "value1",
			},
		},
	}

	actualMessage := NewResponseMessage()
	actualMessage.AddMessageByIssue(faults.ConditionalFieldNotAllowed, Body, "field1", "value1", "field1", "field2", "value2")

	if !reflect.DeepEqual(actualMessage.GetMessage(), expectedMessage) {
		t.Errorf(formatMessageDiff(expectedMessage, actualMessage))
	}
}

func formatMessageDiff(expectedMessage ResponseMessage, actualMessage *ResponseMessage) string {
	return fmt.Sprintf("Expected message: [%+v] is not equal Returned Message: [%+v]", expectedMessage, actualMessage)
}
