package dictionary

import (
	"reflect"
	"strconv"
	"testing"
)

func TestCreateNewDictionary(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Error(err)
	}
}

func TestDefaultDictionaryWords(t *testing.T) {
	dic, err := New()
	if err != nil {
		t.Error(err)
	}

	for _, word := range defaultWords {
		result, err := dic.Find(word)
		if err != nil {
			t.Error(err)
		}
		if result == false {
			t.Errorf("could not find word: %s", word)
		}
	}
}

func TestAddAndFind(t *testing.T) {
	dic, err := New()
	if err != nil {
		t.Error(err)
	}

	err = dic.Add("new-word")
	if err != nil {
		t.Error(err)
	}

	found, err := dic.Find("new-word")
	if err != nil {
		t.Error(err)
	}
	if !found {
		t.Errorf("could not find 'new-word' in the dictionary")
	}
}

func TestCaseInsensitivity(t *testing.T) {
	dic, err := New()
	if err != nil {
		t.Error(err)
	}

	err = dic.Add("CaSe_inSensitiVe")
	if err != nil {
		t.Error(err)
	}

	found, err := dic.Find("cAsE_INsENsitive")
	if err != nil {
		t.Error(err)
	}
	if !found {
		t.Errorf("could not find 'case_insensitive' in the dictionary")
	}
}

func TestDelete(t *testing.T) {
	dic, err := New()
	if err != nil {
		t.Error(err)
	}

	err = dic.Add("new-word")
	if err != nil {
		t.Error(err)
	}

	err = dic.Delete("new-word")
	if err != nil {
		t.Error(err)
	}

	found, err := dic.Find("new-word")
	if err != nil {
		t.Error(err)
	}
	if found {
		t.Errorf("found 'new-word' in the dictionary")
	}
}

func TestMostPopular(t *testing.T) {
	tests := []struct {
		expected map[string]int64
	}{
		{expected: map[string]int64{"hello": 1, "goodbye": 1, "filter": 1, "yes": 1, "no": 1}},
		{expected: map[string]int64{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}},
		{expected: map[string]int64{"hello": 1, "2": 2, "3": 3, "4": 4, "yes": 1}},
	}
	for k, test := range tests {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			testPopular(t, test.expected)
		})
	}
}

func testPopular(t *testing.T, expected map[string]int64) {
	dic, err := New()
	if err != nil {
		t.Error(err)
	}

	for k, v := range expected {
		_ = dic.Add(k)
		for i := 0; i < int(v); i++ {
			_, err := dic.Find(k)
			if err != nil {
				t.Error(err)
			}
		}
	}
	popular, err := dic.MostPopular()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(popular, expected) {
		t.Errorf("expected: %v, got: %v", expected, popular)
	}
}
