package geos

import (
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {
	const re = `3\.[345]?\.\d+-CAPI-1\.[789]?\.\d+( .+)?$`
	version := Version()
	matched, err := regexp.MatchString(re, version)
	if err != nil {
		t.Fatal("Version regex:", err)
	}
	if !matched {
		t.Errorf("Version(): %q didn't match regex \"%s\"", version, re)
	}
}
