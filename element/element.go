package element

import "wxl/resource"

type Element interface {
	resource.Resource

	IsRecord() bool
	IsPath() bool
	IsList() bool
	IsString() bool
	IsNumber() bool
	IsSymbol() bool
	IsError() bool
	IsNull() bool
	IsFunction() bool
	IsObject() bool
	IsType() bool
	IsBool() bool

	NumberValue() float64
	StringValue() string
	BoolValue() bool

	StringElementValue() *StringElement
	NumberElementValue() *NumberElement
	ListElementValue() *ListElement
	FunctionElementValue() *FunctionElement
	SymbolElementValue() *SymbolElement
	ObjectElementValue() *ObjectElement
	RecordElementValue() *RecordElement
	PathElementValue() *PathElement
	TypeElementValue() *TypeElement

	Children() []*Element
}
