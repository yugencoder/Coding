package _8XX

import "testing"

func Test_areSentencesSimilar(t *testing.T) {
	tests := []struct {
		name      string
		sentence1 string
		sentence2 string
		want      bool
	}{
		{name: "", sentence1: "My name is Haley", sentence2: "My Haley", want: true},
		{name: "", sentence1: "Eating right now", sentence2: "Eating", want: true},
		{name: "", sentence1: "Luky", sentence2: "Luccckyy", want: false},
		{name: "", sentence1: "CwFfRo regR", sentence2: "CwFfRo H regR", want: true},
		{name: "", sentence1: "A", sentence2: "a A b A", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areSentencesSimilar(tt.sentence1, tt.sentence2); got != tt.want {
				t.Errorf("areSentencesSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
