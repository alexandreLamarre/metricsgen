package metricsgen

import (
	"fmt"
	"strings"
)

type attributeDef struct {
	field    string
	attrType string
	pointer  bool
}

type structWriter struct {
	structName        string
	structDescription string
	attrs             []attributeDef
	sb                strings.Builder
}

func NewStructWriter(
	name string,
	desc string,
	attrs []attributeDef,
) *structWriter {
	return &structWriter{
		sb:                strings.Builder{},
		structName:        name,
		structDescription: desc,
		attrs:             attrs,
	}
}

func (s *structWriter) startStruct() {
	if s.structDescription != "" {
		s.sb.WriteString(
			fmt.Sprintf("//%s %s\n", s.structName, s.structDescription),
		)
	}
	s.sb.WriteString(
		fmt.Sprintf("type %s struct {\n", s.structName),
	)
}

func (s *structWriter) closeStruct() {
	s.sb.WriteString("}\n")
}

func (s *structWriter) writeAttrs() {
	for _, attr := range s.attrs {
		s.sb.WriteString("\t")
		if attr.field != "" {
			s.sb.WriteString(attr.field)
			s.sb.WriteString(" ")
		}
		if attr.pointer {
			s.sb.WriteString("*")
		}
		s.sb.WriteString(attr.attrType)
		s.sb.WriteString("\n")
	}
}

func (s *structWriter) Generate() string {
	s.startStruct()
	s.writeAttrs()
	s.closeStruct()
	return s.sb.String()
}
