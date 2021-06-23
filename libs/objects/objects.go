package objects

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
	"wxl/libs/helpers"
	"wxl/memory"
)

var Define = directives.Method{
	Symbol: element.SymbolElement{Value: "def"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		currentContext := *ctx
		firstArgument := *params[1]
		secondArgument := *params[2]

		if !firstArgument.IsSymbol() {
			return helpers.Error(ctx, "First argument must be a symbol.")
		}

		if len(params) < 3 {
			return helpers.Error(ctx, "Not enough arguments.")
		}

		object := *secondArgument.ObjectElementValue().Object()

		currentContext.Declare(*firstArgument.SymbolElementValue(), memory.NewObjectBind(&object))

		return secondArgument
	},
}
