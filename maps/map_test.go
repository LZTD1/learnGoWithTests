package maps

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "test"}

	t.Run("Know word", func(t *testing.T) {
		got, _ := d.Search("test")
		assertStrings(t, got, "test")
	})
	t.Run("unKnow word", func(t *testing.T) {
		_, err := d.Search("test2")

		assertErrors(t, err, ErrKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{"test": "test"}
		d.Add("test2", "test2")

		want := "test2"
		got, err := d.Search("test2")
		assertStrings(t, got, want)
		isNotError(t, err)
	})

	t.Run("existing word", func(t *testing.T) {
		d := Dictionary{"test": "test"}
		err := d.Add("test", "test2")

		assertErrors(t, err, ErrKeyAlreadyExists)
		assertDefinition(t, d, "test", "test")
	})
}
func TestUpdate(t *testing.T) {
	t.Run("Existing key", func(t *testing.T) {
		d := Dictionary{"test": "test"}
		err := d.Update("test", "new_test")

		isNotError(t, err)
		assertDefinition(t, d, "test", "new_test")
	})

	t.Run("New key", func(t *testing.T) {
		d := Dictionary{"test": "test"}
		err := d.Update("test2", "new_test")

		assertErrors(t, err, ErrKeyNotExists)
	})
}
func TestDelete(t *testing.T) {
	t.Run("Existing key", func(t *testing.T) {
		d := Dictionary{"test": "test"}
		err := d.Delete("test")

		isNotError(t, err)

		_, e := d.Search("test")
		assertErrors(t, e, ErrKeyNotFound)
	})
	t.Run("Non-existing key", func(t *testing.T) {
		d := Dictionary{"test": "test"}
		err := d.Delete("test2")

		assertErrors(t, err, ErrKeyNotExists)
	})
}
func assertDefinition(t *testing.T, d Dictionary, s string, s2 string) {
	t.Helper()

	val, err := d.Search(s)
	if err != nil {
		t.Errorf("Error searching %v: %v", s, err)
	}
	assertStrings(t, val, s2)
}
func isNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
func assertErrors(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("Func did not return an error")
	}
	if got.Error() != want.Error() {
		t.Fatalf("got error %v, want %v", got, want)
	}
}
func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
