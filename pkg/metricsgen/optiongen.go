package metricsgen

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

const (
	SuffixOptionStruct = "Options"
	SuffixOptionFunc   = "Option"
)

type optionGen struct {
	name  string
	attrs []optionAttr
	sb    strings.Builder
}

type optionAttr struct {
	attributeDef attributeDef
	defaultValue string
}

func NewOptionWriter(
	name string,
	attrs []optionAttr,
) *optionGen {
	return &optionGen{
		name:  name,
		attrs: attrs,
		sb:    strings.Builder{},
	}
}

func (o *optionGen) writeOptionType() {
	o.sb.WriteString(
		fmt.Sprintf(
			"type %s func(o *%s)\n", o.funcName(), o.structName(),
		),
	)
}

func (o *optionGen) structName() string {
	return o.name + SuffixOptionStruct
}

func (o *optionGen) funcName() string {
	return o.name + SuffixOptionFunc
}

func (o *optionGen) writeDefaultFunction() {
	o.sb.WriteString(
		fmt.Sprintf("func default%s() *%s {\n", o.structName(), o.structName()),
	)
	o.sb.WriteString(
		fmt.Sprintf("\treturn &%s{\n", o.structName()),
	)

	for _, attr := range o.attrs {
		if attr.defaultValue == "" {
			continue
		}

		o.sb.WriteString("\t\t")
		if attr.attributeDef.field == "" {
			o.sb.WriteString(attr.attributeDef.attrType)
		} else {
			o.sb.WriteString(attr.attributeDef.field)
		}
		o.sb.WriteString(" : ")
		o.sb.WriteString(attr.defaultValue)
		o.sb.WriteString(",\n")

	}

	o.sb.WriteString("\t}\n")

	o.sb.WriteString("}\n")
}

func (o *optionGen) writeOptionFunc(attr optionAttr) {
	var attrId string
	if attr.attributeDef.field == "" {
		attrId = attr.attributeDef.attrType
	} else {
		attrId = attr.attributeDef.field
	}
	attrIdC := CapitalizeFirst(attrId)

	var attrType string
	if attr.attributeDef.pointer {
		attrType = "*" + attr.attributeDef.attrType

	} else {
		attrType = attr.attributeDef.attrType
	}

	o.sb.WriteString(fmt.Sprintf("func With%s%s(val %s) %s {\n", o.name, attrIdC, attrType, o.funcName()))
	o.sb.WriteString("\t")
	o.sb.WriteString(fmt.Sprintf("return func(o *%s) {\n", o.structName()))
	o.sb.WriteString("\t\t")
	o.sb.WriteString("o." + attrId + " = val\n")
	o.sb.WriteString("\t}\n")
	o.sb.WriteString("}\n")
}

func (o *optionGen) writeApplyFunc() {
	o.sb.WriteString("\n")
	o.sb.WriteString(fmt.Sprintf("func (o *%s) Apply(opts ...%s) {\n", o.structName(), o.funcName()))
	o.sb.WriteString("\tfor _, opt := range opts {\n")
	o.sb.WriteString("\t\topt(o)\n")
	o.sb.WriteString("\t}\n")
	o.sb.WriteString("}\n")
	o.sb.WriteString("\n")
}

func (o *optionGen) Generate() string {

	attrs := lo.Map(o.attrs, func(oa optionAttr, _ int) attributeDef {
		return oa.attributeDef
	})
	optionStruct := NewStructWriter(
		o.structName(),
		"",
		attrs,
	)

	o.sb.WriteString(optionStruct.Generate())
	o.writeOptionType()
	o.writeApplyFunc()
	o.writeDefaultFunction()
	for _, attr := range o.attrs {
		o.sb.WriteString("\n")
		o.writeOptionFunc(attr)
		o.sb.WriteString("\n")
	}
	return o.sb.String()
}
