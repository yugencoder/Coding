package _8XX

import "testing"

func Test_evaluate(t *testing.T) {

	tests := []struct {
		name      string
		s         string
		knowledge [][]string
		want      string
	}{
		//{
		//	name:      "",
		//	s:         "(name)is(age)yearsold",
		//	knowledge: [][]string{{"name", "bob"}, {"age", "two"}},
		//	want:      "bobistwoyearsold",
		//},
		{
			name:      "",
			s:         "hi(name)",
			knowledge: [][]string{{"a", "b"}},
			want:      "hi?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.s, tt.knowledge); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
