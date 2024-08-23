package OrderedForm

import (
	"net/url"
	"testing"
)

const correctOutput = `foo=bar&example=test`

func TestOrderedForm(t *testing.T) {
	form := new(OrderedForm)
	form.Set("foo", "bar")
	form.Set("example", `test`)

	v := form.URLEncode()

	if v != correctOutput {
		t.Fatalf("expected:\n%v, got:\n%v", correctOutput, v)
	}
}

func TestOrderedFormIterate(t *testing.T) {
	form := new(OrderedForm)
	form.Set("foo", "bar")
	form.Set("example", "test")

	expected := map[string]string{
		"foo":     "bar",
		"example": "test",
	}

	form.Iterate(func(k, v string) {
		if val, ok := expected[k]; !ok || val != v {
			t.Errorf("Iterate failed for key %s: expected %s, got %s", k, val, v)
		}
	})
}

func TestOrderedFormUpdateSet(t *testing.T) {
	form := new(OrderedForm)
	form.Set("foo", "bar")
	form.UpdateSet("foo", "baz")   // Update existing key
	form.UpdateSet("new", "value") // Add new key-value pair

	// Check updated value
	if val := form.URLEncode(); val != url.QueryEscape("foo")+"=baz&"+url.QueryEscape("new")+"=value" {
		t.Errorf("UpdateSet failed: expected %s, got %s", "foo=baz&new=value", val)
	}
}
