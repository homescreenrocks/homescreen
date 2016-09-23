package storage_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/homescreenrocks/homescreen/core/backend/storage"
)

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertEquals(t *testing.T, expected, obtained string) {
	if expected != obtained {
		t.Fatalf("%#v != %#v", expected, obtained)
	}
}

func TestAll(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "homescreen-test-")
	assertNoError(t, err)
	defer os.RemoveAll(tempDir)

	db, err := storage.New(path.Join(tempDir, "test.db"))
	assertNoError(t, err)
	defer db.Close()

	err = db.SetRaw("foo/bar", []byte(`"blubber"`))
	assertNoError(t, err)

	err = db.Set("bim/baz", 1337)
	assertNoError(t, err)

	err = db.Set("wut/not", []string{"bish", "bash", "bosh"})
	assertNoError(t, err)

	value, err := db.GetRaw("foo/bar")
	assertNoError(t, err)
	assertEquals(t, `"blubber"`, string(value))

	value, err = db.GetRaw("bim/baz")
	assertNoError(t, err)
	assertEquals(t, `1337`, string(value))

	value, err = db.GetRaw("wut/not")
	assertNoError(t, err)
	assertEquals(t, `["bish","bash","bosh"]`, string(value))

	var foobar string
	err = db.Get("foo/bar", &foobar)
	assertNoError(t, err)
	assertEquals(t, "blubber", foobar)

	var bimbaz int
	err = db.Get("bim/baz", &bimbaz)
	assertNoError(t, err)
	assertEquals(t, "1337", fmt.Sprintf("%#v", bimbaz))

	var wutnot []string
	err = db.Get("wut/not", &wutnot)
	assertNoError(t, err)
	assertEquals(t, `[]string{"bish", "bash", "bosh"}`,
		fmt.Sprintf("%#v", wutnot))
}
