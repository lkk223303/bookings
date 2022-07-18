package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/someurl", nil)
	form := New(r.PostForm)

	if !form.Valid() {
		t.Error("got invalid form when should have been valid")
	}
}

func TestRequired(t *testing.T) {
	r := httptest.NewRequest("POST", "/someurl", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		// required field not filled yet
		t.Error("form shows valid when required fields missing")
	}

	postData := make(url.Values)

	postData.Add("v1", "aa")
	postData.Add("v2", "bb")
	postData.Add("v3", "cc")

	r, _ = http.NewRequest("POST", "/someurl", nil)
	r.PostForm = postData
	form = New(r.PostForm)

	form.Required("v1", "v2", "v3")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
	form.Required("v4")
	err := form.Errors.Get("v4")
	if err == "" {
		t.Error("This field should have error message")
	}

}

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postData = url.Values{}
	postData.Add("a", "a")
	form = New(postData)

	has = form.Has("a")
	if !has {
		t.Error("shows form doed not have field when it should")
	}
}

func TestMinLength(t *testing.T) {
	data := make(url.Values)
	form := New(data)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	data = make(url.Values)
	data.Add("a", "11")
	form = New(data)

	form.MinLength("a", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 valid, but data is shorter")
	}

	data = make(url.Values)
	data.Add("b", "12345678")
	form = New(data)

	form.MinLength("b", 3)
	if !form.Valid() {
		t.Error("shows minlength of 3 unvalid, but data satisfy the minlength")
	}

}

func TestIsEmail(t *testing.T) {
	data := make(url.Values)
	form := New(data)

	data = make(url.Values)
	data.Add("a", "a@a.com")
	data.Add("b", "a@a")
	form.Values = data

	form.IsEmail("a")
	form.IsEmail("b")
	if form.Errors.Get("a") != "" && !form.Valid() {
		t.Error("should not have error in the field, but got one")
	}
	if form.Errors.Get("b") == "" && form.Valid() {
		t.Error("should have error message, but not have")
	}
}
