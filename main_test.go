package indent

import (
	"testing"
)

func requireEqual(t *testing.T, expectedString, resultString string) {
	t.Helper()
	if expectedString != resultString {
		t.Errorf("expected equal strings:\n  expected: %q\n  result: %q\n", 
		expectedString, resultString)
	}
}

func TestIndentSingleLineInput(t *testing.T) {
	inputData := "some data"
	indentedData := Indent(inputData, "  ", 1)

	expectedLine := "  some data"
	requireEqual(t, expectedLine, indentedData)
}

func TestIndentMultilineInput(t *testing.T) {
	inputData := "some data\nanother data"
	indentedData := Indent(inputData, "  ", 1)

	expectedLine := "  some data\n  another data"
	requireEqual(t, expectedLine, indentedData)
}

func TestIndentEmptyLineInput(t *testing.T) {
	inputData := "some data\n\n"
	indentedData := Indent(inputData, "  ", 1)

	expectedLine := "  some data\n\n"
	requireEqual(t, expectedLine, indentedData)
}

func CountOfTimes(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {
		inputData := "some data"
		indentedData := Indent(inputData, "  ", 0)

		expectedLine := "some data"
		requireEqual(t, expectedLine, indentedData)
	})

	t.Run("Three", func(t *testing.T) {
		inputData := "some data"
		indentedData := Indent(inputData, " ", 3)

		expectedLine := "   some data"
		requireEqual(t, expectedLine, indentedData)
	})
}
