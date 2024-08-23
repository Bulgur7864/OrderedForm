package OrderedForm

import (
	"net/url"
	"strings"
)

// OrderedForm is a slice of [2]string, in which the first
// element is the key, and the second element is the value.
// Both the key and value are query escaped when using Set().
type OrderedForm [][2]string

func (o *OrderedForm) Set(k, v string) {
	// Set the key, and url encode the value
	m := [2]string{url.QueryEscape(k), url.QueryEscape(v)}
	*o = append(*o, m)
}

func (o *OrderedForm) URLEncode() string {
	var b strings.Builder
	for _, v := range *o {
		if b.Len() > 0 {
			b.WriteString("&")
		}
		b.WriteString(v[0])
		b.WriteString("=")
		b.WriteString(v[1])
	}

	return b.String()
}

// Iterate calls the callback function for each key-value pair in the form.
// example:
//	 form.Iterate(func(k, v string) {
//		fmt.Printf("Key: %s, Value: %s\n", k, v)
//	})
func (o *OrderedForm) Iterate(callback func(k, v string)) {
	for _, pair := range *o {
		// Unescape the key and value before passing to the callback
		k, errK := url.QueryUnescape(pair[0])
		v, errV := url.QueryUnescape(pair[1])
		if errK == nil && errV == nil {
			callback(k, v)
		}
	}
}

// UpdateSet updates the value of a key if it exists in the form. If not, it adds the key-value pair.
func (o *OrderedForm) UpdateSet(k, v string) {
    // URL encode the key and value
    encodedKey := url.QueryEscape(k)
    encodedValue := url.QueryEscape(v)
    found := false

    // Iterate over the form to find the key
    for i, pair := range *o {
        if pair[0] == encodedKey {
            // Update the value if the key is found
            (*o)[i][1] = encodedValue
            found = true
            break
        }
    }

    // If the key was not found, add it
    if !found {
        o.Set(k, v)
    }
}
