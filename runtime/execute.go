package runtime

import (
	"log"
	"rpl/core"
	"rpl/element"
	"rpl/environment"
	"rpl/flow"
)

func Run(env environment.Environment, codeTree core.ExprList) element.Element {
	resultList := element.ListElement{}
	ctx := flow.NewContext(env)

	for _, element := range codeTree.Elements {
		if element.IsList() {
			result := evaluateList(spec, element.ListElementValue(), ctx)

			resultList.Push(result)

			log.Println("Result: " + result.StringValue())
		}
	}

	if len(resultList.Elements) == 1 {
		return *resultList.First()
	}

	return resultList
}
