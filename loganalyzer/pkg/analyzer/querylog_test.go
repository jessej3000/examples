package analyzer

import (
	"reflect"
	"testing"
)

// All test cases also test for GET method request which is really tricky.
// We cannot test for none GET method 100%.
// As of the values we get from the entries, only the status code can give
// us a degree of checking if the request is a GET method or not.
// As of my research status code 201 are most likely a result of POST and PUT request.
// So we are checking if a status is 201 or not.
//
// All results are also averages of each unique request

// TestAverage Test to make sure the normal command brings out expected result
func TestAverage(t *testing.T) {
	a := New(
		"test_average.log",
		"minimal",
		10,
		true,
		true,
		true,
	)

	checker := []string{"/vendor/bootstrap.min.js", "/articles.html?id=4"}

	result := a.Analyze()

	if !reflect.DeepEqual(result, checker) {
		t.Errorf("Invalid result Average.")
	}

}

// TestLowerUpperCase checks case-insensitive urls
func TestLowerUpperCase(t *testing.T) {
	a := New(
		"test_loweruppercase.log",
		"minimal",
		10,
		true,
		true,
		true,
	)

	checker := []string{"/articles_one.html?id=4",
		"/vendor/bootstrap.min.js",
		"/articles.html?id=4",
		"/vendor/bootstrap2.min.js"}

	result := a.Analyze()

	if !reflect.DeepEqual(result, checker) {
		t.Errorf("Invalid result for LowerUpperCase.")
	}
}

// TestGIF test if no gif request are included
func TestGIF(t *testing.T) {
	a := New(
		"test_gif.log",
		"minimal",
		10,
		true,
		true,
		true, // Set this to false allows GIF and fails the test which is expected
	)

	checker := []string{"/articles_one.html?id=4",
		"/vendor/bootstrap.min.js",
		"/articles.html?id=4",
		"/vendor/bootstrap2.min.js"}

	result := a.Analyze()

	if !reflect.DeepEqual(result, checker) {
		t.Errorf("Invalid result for GIF.")
	}
}

// TestOrderAnd10Entries test for the order of the list according to
// duration of processing the request
// It also test that the list should not exceed 10 items
func TestOrderAnd10Entries(t *testing.T) {
	a := New(
		"test_orderand10entries.log",
		"complete",
		10, // Value 11 or higher will fail the test which is expected
		true,
		true,
		true,
	)

	checker := []string{
		"/articles_one.html?id=4 : 21.21s : 8 entries",
		"/articles4.html?id=4 : 20.09s : 5 entries",
		"/vendor/bootstrap.min.js : 12.84s : 38 entries",
		"/articles2.html?id=1 : 12.20s : 1 entries",
		"/articles8.html?id=4 : 11.10s : 6 entries",
		"/articles10.html?id=4 : 10.14s : 1 entries",
		"/articles5.html?id=4 : 8.11s : 5 entries",
		"/articles.html?id=4 : 8.05s : 10 entries",
		"/vendor/bootstrap2.min.js : 6.23s : 8 entries",
		"/articles1.html?id=4 : 5.23s : 4 entries"}

	result := a.Analyze()

	if !reflect.DeepEqual(result, checker) {
		t.Errorf("Invalid result for OrderAnd10Entries.")
	}
}
