package main

import (
	"reflect"
	"testing"
)

func TestCF71A(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"word"}, []string{"word"}},
		{[]string{"localization"}, []string{"l10n"}},
		{[]string{"internationalization", "localization", "a"}, []string{"i18n", "l10n", "a"}},
		{[]string{"pneumonoultramicroscopicsilicovolcanoconiosis"}, []string{"p43s"}},
		{[]string{"hello", "world"}, []string{"hello", "world"}},
		{[]string{"supercalifragilisticexpialidocious"}, []string{"s32s"}},
	}

	for _, test := range tests {
		output := CF71A(test.input)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, output)
		}
	}
}
