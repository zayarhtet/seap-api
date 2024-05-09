package lib

import (
	"fmt"
	"html"
	"strings"
)

type HTMLElement struct {
	nodeName string
	children []HTMLElement
	value    string
	classes  string
}

func (h *HTMLElement) SetNodeName(name string) {
	h.nodeName = name
}

func (h *HTMLElement) SetClasses(class string) {
	h.classes = class
}

func (h *HTMLElement) AddChild(child HTMLElement) {
	h.children = append(h.children, child)
}

func (h *HTMLElement) SetValue(value string) {
	h.value = html.EscapeString(value)
}

func (h *HTMLElement) GetHTML() string {
	var htmlString string
	if len(h.nodeName) == 0 {
		htmlString = "{{body}}"
	} else {
		htmlString = fmt.Sprintf("<%s class=\"%s\">{{body}} %s</%s>", h.nodeName, h.classes, h.value, h.nodeName)
	}

	body := ""
	for _, elem := range h.children {
		body += elem.GetHTML()
	}

	return strings.Replace(htmlString, "{{body}}", body, 1)
}
