package beam

import (
	"github.com/apache/beam/sdks/go/pkg/beam/util/reflectx"
	"testing"
)

func TestJSONCoder(t *testing.T) {
	tests := []int{43, 12431235, -2, 0, 1}

	for _, test := range tests {
		data, err := jsonEnc(test)
		if err != nil {
			t.Fatalf("Failed to encode %v: %v", tests, err)
		}
		decoded, err := jsonDec(reflectx.Int, data)
		if err != nil {
			t.Fatalf("Failed to decode: %v", err)
		}
		actual := decoded.(int)

		if test != actual {
			t.Errorf("Corrupt coding: %v, want %v", actual, test)
		}
	}
}
