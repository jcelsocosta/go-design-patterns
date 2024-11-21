package templatemethod

import "strings"

type MessageRetriever interface {
	Message() string
}

type TemplateMethod interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type TemplateMethodImpl struct{}

func (t *TemplateMethodImpl) first() string {
	return ""
}

func (t *TemplateMethodImpl) third() string {
	return ""
}

func (t *TemplateMethodImpl) ExecuteAlgorithm(m MessageRetriever) string {
	return ""
}

type Template struct{}

func (t *Template) first() string {
	return "hello"
}

func (t *Template) third() string {
	return "template"
}

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}
