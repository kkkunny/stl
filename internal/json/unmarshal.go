package json

import (
	"encoding/json"
	"fmt"
	"io"
	"iter"

	"github.com/kkkunny/stl/container/tuple"
)

func UnmarshalToMapObj[K, V any](r io.Reader) iter.Seq2[tuple.Tuple2[K, V], error] {
	dec := json.NewDecoder(r)

	return func(yield func(tuple.Tuple2[K, V], error) bool) {
		t, err := dec.Token()
		if err != nil {
			yield(tuple.Tuple2[K, V]{}, err)
			return
		}
		if delim, ok := t.(json.Delim); !ok || delim != '{' {
			yield(tuple.Tuple2[K, V]{}, fmt.Errorf("expected JSON object start with '{'"))
			return
		}

		for dec.More() {
			keyToken, err := dec.Token()
			if err != nil {
				yield(tuple.Tuple2[K, V]{}, err)
				return
			}
			key := keyToken.(K)

			var val V
			if err = dec.Decode(&val); err != nil {
				yield(tuple.Tuple2[K, V]{}, err)
				return
			}

			if !yield(tuple.Pack2[K, V](key, val), nil) {
				return
			}
		}

		t, err = dec.Token()
		if err != nil {
			yield(tuple.Tuple2[K, V]{}, err)
			return
		}
		if delim, ok := t.(json.Delim); !ok || delim != '}' {
			yield(tuple.Tuple2[K, V]{}, fmt.Errorf("expected JSON object end with '}'"))
			return
		}
	}
}
