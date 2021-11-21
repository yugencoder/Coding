package easy

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_areSentencesSimilar(t *testing.T) {


	v, ok := strconv.Atoi("Main_1")
	fmt.Println(v,ok)

	type args struct {
		sentence1    []string
		sentence2    []string
		similarPairs [][]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				sentence1:    []string{"an","extraordinary","meal"},
				sentence2:    []string{"a","good","dinner"},
				similarPairs: [][]string{{"great","good"},{"extraordinary","good"},{"well","good"},{"wonderful","good"},{"excellent","good"},{"fine","good"},{"nice","good"},{"any","one"},{"some","one"},{"unique","one"},{"the","one"},{"an","one"},{"single","one"},{"a","one"},{"truck","car"},{"wagon","car"},{"automobile","car"},{"auto","car"},{"vehicle","car"},{"entertain","have"},{"drink","have"},{"eat","have"},{"take","have"},{"fruits","meal"},{"brunch","meal"},{"breakfast","meal"},{"food","meal"},{"dinner","meal"},{"super","meal"},{"lunch","meal"},{"possess","own"},{"keep","own"},{"have","own"},{"extremely","very"},{"actually","very"},{"really","very"},{"super","very"}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areSentencesSimilar(tt.args.sentence1, tt.args.sentence2, tt.args.similarPairs); got != tt.want {
				t.Errorf("areSentencesSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
