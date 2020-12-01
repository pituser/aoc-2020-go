package util

import "testing"

func TestReadLinesFromFile(t *testing.T) {
	want := []string{
		"Lorem ipsum dolor sit",
		"amet, consetetur sadipscing elitr,",
		"sed diam nonumy eirmod tempor invidunt.",
		"",
		"1",
		"2",
		"3",
	}

	got, err := ReadLinesFromFile("sample_input.txt")

	if err != nil {
		t.Errorf("unexpected error %q", err)
	}

	if !equalStringList(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestReadLinesFromFileError(t *testing.T) {
	_, err := ReadLinesFromFile("does_not_exist.txt")

	if err == nil {
		t.Error("expected error for non-existing file")
	}
}

func TestParseIntList(t *testing.T) {
	var tests = []struct {
		input []string
		want  []int
	}{
		{[]string{"1", "3", "-99", "42"}, []int{1, 3, -99, 42}},
		{[]string{"0"}, []int{0}},
		{[]string{}, []int{}},
		{[]string{"1 ", " 192541", " -0981 "}, []int{1, 192541, -981}},
	}

	for _, test := range tests {
		t.Run("ParseInList", func(t *testing.T) {
			got, err := ParseIntList(test.input)
			if !equalIntList(got, test.want) {
				t.Errorf("got %v want %v", got, test.want)
			}
			if err != nil {
				t.Errorf("unexpected error %q", err)
			}
		})
	}
}

func TestParseIntListError(t *testing.T) {
	var tests = []struct {
		input []string
	}{
		{[]string{"1", "3", "-99", "42öä"}},
		{[]string{""}},
		{[]string{"1 2 3 4"}},
	}

	for _, test := range tests {
		t.Run("ParseInList", func(t *testing.T) {
			got, err := ParseIntList(test.input)
			if !equalIntList(got, []int{}) {
				t.Errorf("got %v want empty list", got)
			}
			if err == nil {
				t.Error("expected parse error")
			}
		})
	}
}

func equalStringList(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func equalIntList(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
