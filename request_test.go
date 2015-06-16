package grequests

import "testing"

func TestAddQueryStringParams(t *testing.T) {
	userURL := buildUrlParams("https://www.google.com/", map[string]string{"1": "2", "3": "4"})
	if userURL != "https://www.google.com/?1=2&3=4" {
		t.Error("URL params not properly built", userURL)
	}
}

func TestSortAddQueryStringParams(t *testing.T) {
	userURL := buildUrlParams("https://www.google.com/", map[string]string{"3": "4", "1": "2"})
	if userURL != "https://www.google.com/?1=2&3=4" {
		t.Error("URL params not properly built and sorted", userURL)
	}
}

func TestAddQueryStringParamsExistingParam(t *testing.T) {
	userURL := buildUrlParams("https://www.google.com/?5=6", map[string]string{"3": "4", "1": "2"})
	if userURL != "https://www.google.com/?1=2&3=4&5=6" {
		t.Error("URL params not properly built and sorted", userURL)
	}
}