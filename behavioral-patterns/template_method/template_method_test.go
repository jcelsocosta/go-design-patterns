package templatemethod

import (
	"strings"
	"testing"
)

type TestStruct struct {
	Method TemplateMethod
}

func (m *TestStruct) Message() string {
	return "world"
}

func TestTemplateMethod_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{
			Method: &Template{},
		}
		res := s.Method.ExecuteAlgorithm(s)
		expected := "world"

		if !strings.Contains(res, expected) {
			t.Errorf("Expected string '%s' wasn't found on returned string \n", expected)
		}
	})
}
