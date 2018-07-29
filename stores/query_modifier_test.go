package stores

import "testing"

func TestGenerateWhereStatement(t *testing.T) {
	tables := []struct {
		name         string
		args         []QueryModifier
		resultQuery  string
		resultValues []interface{}
	}{
		{"no args", []QueryModifier{}, "", nil},
		{"single arg", []QueryModifier{QueryMod("id", EQ, "123")}, "WHERE id=$1 ", []interface{}{"123"}},
		{"two args with AND", []QueryModifier{QueryMod("id", EQ, 123), And, QueryMod("val", NE, "abc")}, "WHERE id=$1 AND val!=$2 ", []interface{}{123, "abc"}},
		{"bad args", []QueryModifier{QueryMod("", EQ, "")}, "", nil},
		{"tests floats", []QueryModifier{QueryMod("id", EQ, 1.0)}, "WHERE id=$1 ", []interface{}{1.0}},
		{"tests bools", []QueryModifier{QueryMod("admin", EQ, true)}, "WHERE admin=$1 ", []interface{}{true}},
	}

	for _, tt := range tables {
		tt := tt

		query, vals := generateWhereStatement(&tt.args)
		if query != tt.resultQuery {
			t.Errorf("%s should have query of %s, but is %s", tt.name, tt.resultQuery, query)
		}
		if len(vals) != len(tt.resultValues) {
			t.Errorf("%s should have %d vals, but has %d", tt.name, len(tt.resultValues), len(vals))
		}

		for i, val := range vals {
			switch v := val.(type) {
			case int:
				o, ok := tt.resultValues[i].(int)
				if !ok {
					t.Errorf("%s should have been an %t, but was an int", tt.name, tt.resultValues[i])
				}

				if v != o {
					t.Errorf("%s should have been an %v, but was %v", tt.name, o, v)
				}
				continue
			case string:
				o, ok := tt.resultValues[i].(string)
				if !ok {
					t.Errorf("%s should have been an %t, but was an string", tt.name, tt.resultValues[i])
				}

				if v != o {
					t.Errorf("%s should have been an %v, but was %v", tt.name, o, v)
				}
				continue
			case float64:
				o, ok := tt.resultValues[i].(float64)
				if !ok {
					t.Errorf("%s should have been an %t, but was an float64", tt.name, tt.resultValues[i])
				}

				if v != o {
					t.Errorf("%s should have been an %v, but was %v", tt.name, o, v)
				}
				continue
			case bool:
				o, ok := tt.resultValues[i].(bool)
				if !ok {
					t.Errorf("%s should have been an %t, but was an bool", tt.name, tt.resultValues[i])
				}

				if v != o {
					t.Errorf("%s should have been an %v, but was %v", tt.name, o, v)
				}
				continue
			}
		}
	}
}
