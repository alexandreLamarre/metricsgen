package metricsgen

import (
	"fmt"
	"strings"
)

type importDef struct {
	alias      string
	dependency string
}

type importWriter struct {
	sb              strings.Builder
	requiredImports []importDef
}

func NewImportWriter(imports []importDef) *importWriter {
	return &importWriter{
		sb:              strings.Builder{},
		requiredImports: imports,
	}
}

func (i *importWriter) start() {
	i.sb.WriteString("import (\n")
}

func (i *importWriter) writeImports() {
	for _, req := range i.requiredImports {
		if req.dependency == "" {
			continue
		}
		wrapDep := fmt.Sprintf(`"%s"`, req.dependency)
		if req.alias == "" {
			i.sb.WriteString("\t" + wrapDep + "\n")
		} else {
			i.sb.WriteString("\t" + req.alias + " " + wrapDep + "\n")
		}
	}
}

func (i *importWriter) close() {
	i.sb.WriteString(")\n")
}

func (i *importWriter) Generate() string {
	if len(i.requiredImports) == 0 {
		return ""
	}

	i.start()
	i.writeImports()
	i.close()
	return i.sb.String()
}
