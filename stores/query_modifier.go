package stores

import "strconv"

// QueryModifierOp is a enum type to indicate the modifer operation
type QueryModifierOp = string

const (
	EQ  = "="
	LT  = "<"
	LTE = "<="
	NE  = "!="
	GT  = ">"
	GTE = ">="
)

var (
	// And is a query modifier
	And = QueryModifier{"AND", EQ, nil}
	// Or is a query modifier
	Or = QueryModifier{"OR", EQ, nil}
)

// QueryModifier is used in where queries to add selection criteria
type QueryModifier struct {
	Column   string
	Operator QueryModifierOp
	Value    interface{}
}

func QueryMod(col string, operator QueryModifierOp, value interface{}) QueryModifier {
	return QueryModifier{
		Column:   col,
		Operator: operator,
		Value:    value,
	}
}

func generateWhereStatement(modifiers *[]QueryModifier) (string, []interface{}) {
	if len(*modifiers) == 0 {
		return "", nil
	}
	var args []interface{}
	where := "WHERE "

	count := 1
	for _, modifier := range *modifiers {
		if modifier.Column == "AND" || modifier.Column == "OR" {
			where += modifier.Column + " "
			continue
		}

		if modifier.Column == "" || modifier.Value == nil {
			return "", nil
		}

		where += modifier.Column + modifier.Operator + "$" + strconv.Itoa(count) + " "
		args = append(args, modifier.Value)
		count++
	}

	return where, args
}
