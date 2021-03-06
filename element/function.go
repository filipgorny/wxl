package element

type FunctionElement struct {
	arguments ListElement
	body      ListElement
	name      string
	hasName   bool
}

func NewFunctionElement(hasName bool, name string, arguments ListElement, body ListElement) FunctionElement {
	return FunctionElement{
		hasName:   hasName,
		name:      name,
		arguments: arguments,
		body:      body,
	}
}

func (f FunctionElement) IsList() bool {
	return false
}

func (s FunctionElement) IsString() bool {
	return false
}

func (e FunctionElement) IsNumber() bool {
	return false
}

func (f FunctionElement) IsSymbol() bool {
	return false
}

func (f FunctionElement) NumberValue() float64 {
	return -1
}

func (f FunctionElement) StringValue() string {
	return "Function " + f.arguments.StringValue()
}

func (f FunctionElement) Children() []*Element {
	return nil
}

func (f FunctionElement) ListElementValue() *ListElement {
	return nil
}

func (f FunctionElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{Value: f.name}
}

func (f FunctionElement) IsFunction() bool {
	return true
}

func (f FunctionElement) FunctionElementValue() *FunctionElement {
	return &f
}

func (f FunctionElement) IsError() bool {
	return false
}

func (f FunctionElement) IsNull() bool {
	return false
}

func (f FunctionElement) BoolValue() bool {
	return false
}

func (f FunctionElement) GetName() StringElement {
	return StringElement{Value: f.name}
}

func (f FunctionElement) GetArguments() ListElement {
	return f.arguments
}

func (f FunctionElement) GetBody() ListElement {
	return f.body
}

// object

func (o FunctionElement) IsObject() bool {
	return false
}

func (o FunctionElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e FunctionElement) StringElementValue() *StringElement {
	return &StringElement{Value: e.name}
}

func (e FunctionElement) NumberElementValue() *NumberElement {
	return &NumberElement{Value: 0}
}

func (e FunctionElement) IsRecord() bool {
	return false
}

func (e FunctionElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e FunctionElement) IsPath() bool {
	return false
}

func (e FunctionElement) PathElementValue() *PathElement {
	return &PathElement{}
}

func (e FunctionElement) IsType() bool {
	return false
}

func (e FunctionElement) TypeElementValue() *TypeElement {
	return &TypeElement{
		Type: TYPE_UNDEFINED,
	}
}

func (e FunctionElement) IsBool() bool {
	return false
}
