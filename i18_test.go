package main

import "testing"

func TestParseJSONi18(t *testing.T) {

	cases := []struct {
		file        string
		keyValueMap map[string]string
	}{
		{
			"",
			make(map[string]string),
		},
		{
			"en_test.json",
			map[string]string{
				"signIn":   "Sign in",
				"username": "Username",
				"password": "Password",
			},
		},
	}

	for _, c := range cases {

		i18Map, err := ParseJSONi18File(c.file)
		if c.file == "" && err == nil {
			t.Error("parsing file should return error")
		}

		for tk, tv := range c.keyValueMap {

			if len(i18Map) == 0 {
				t.Error("map is empty")
			}

			val, ok := i18Map[tk]

			if !ok {
				t.Errorf("expecting to find key: %v", tk)
			}

			if val == "" {
				t.Errorf("map value is empty")
			}

			if val != tv {
				t.Errorf("expecting %v got %v", tv, val)
			}

		}

	}

}
