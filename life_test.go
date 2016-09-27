package life

import (
	"encoding/json"
	"testing"
)

func TestMarshaling(t *testing.T) {
	cases := []struct {
		input, output []byte
	}{
		{
			[]byte(`{"Board":[[true],[false]]}`),
			[]byte(`{"Board":[[true],[false]]}`),
		},
		{
			[]byte(`{"shouldBeNull": "shit"}`),
			[]byte(`{"Board":null}`),
		},
	}

	for _, c := range cases {
		var g Game
		err := json.Unmarshal(c.input, &g)
		if err != nil {
			t.Errorf("Error while unmarshaling: %s", err.Error())
			continue
		}
		unmarshaled, err := json.Marshal(g)
		if err != nil {
			t.Errorf("Error while marshaling: %s", err.Error())
			continue
		}
		if string(unmarshaled) != string(c.output) {
			t.Errorf("Expected %s to be %s", unmarshaled, c.output)
		}
	}
}
