package pslint

import (
	"reflect"
	"testing"
)

func TestValid(t *testing.T) {
	var src string
	linter := &Linter{}

	src = `
// aw : http://en.wikipedia.org/wiki/.aw
aw
com.aw

// bd : http://en.wikipedia.org/wiki/.bd
*.bd

// ck : http://en.wikipedia.org/wiki/.ck
*.ck
!www.ck
`

	ps, err := linter.LintString(src)
	if err != nil {
		t.Fatalf("lint() returned an error: %v", err)
	}

	if want, got := 0, len(ps); want != got {
		t.Errorf("Expected %d errors, got %d", want, got)
		t.Fatal(ps)
	}
}

func TestInvalid(t *testing.T) {
	var src string
	linter := NewLinter()
	linter.FailFast = false
	linter.FailFirst = false

	src = `
aw
// invalid : leading space
 com.aw

// bd : http://en.wikipedia.org/wiki/.bd
*.bd

// ck : http://en.wikipedia.org/wiki/.ck
 *.CK
!www.ck
`

	ps, err := linter.LintString(src)
	if err != nil {
		t.Fatalf("lint() returned an error: %v", err)
	}

	if want, got := 3, len(ps); want != got {
		t.Errorf("Expected %d errors, got %d", want, got)
		t.Fatal(ps)
	}

	if want, got := &ps[0], &(Problem{Message: "leading space", Level: LEVEL_WARN, Line: 4, LineSource: " com.aw"}); !reflect.DeepEqual(want, got) {
		t.Fatalf("Problem[%d] is %+v, want %+v", 0, got, want)
	}

	if want, got := &ps[1], &(Problem{Message: "leading space", Level: LEVEL_WARN, Line: 10, LineSource: " *.CK"}); !reflect.DeepEqual(want, got) {
		t.Fatalf("Problem[%d] is %+v, want %+v", 0, got, want)
	}

	if want, got := &ps[2], &(Problem{Message: "non-lowercase suffix", Level: LEVEL_ERROR, Line: 10, LineSource: " *.CK"}); !reflect.DeepEqual(want, got) {
		t.Fatalf("Problem[%d] is %+v, want %+v", 0, got, want)
	}
}

func TestInvalidFailFast(t *testing.T) {
	var src string
	linter := NewLinter()
	linter.FailFast = true
	linter.FailFirst = false

	src = `
aw
// invalid : leading space
 com.aw

// bd : http://en.wikipedia.org/wiki/.bd
*.bd

// ck : http://en.wikipedia.org/wiki/.ck
 *.CK
!www.ck
`

	ps, err := linter.LintString(src)
	if err != nil {
		t.Fatalf("lint() returned an error: %v", err)
	}

	if want, got := 1, len(ps); want != got {
		t.Errorf("Expected %d errors, got %d", want, got)
		t.Fatal(ps)
	}

	if want, got := &ps[0], &(Problem{Message: "leading space", Level: LEVEL_WARN, Line: 4, LineSource: " com.aw"}); !reflect.DeepEqual(want, got) {
		t.Fatalf("Problem[%d] is %+v, want %+v", 0, got, want)
	}
}

func TestInvalidFailFirst(t *testing.T) {
	var src string
	linter := NewLinter()
	linter.FailFast = false
	linter.FailFirst = true

	src = `
aw
// invalid : leading space
 com.aw

// bd : http://en.wikipedia.org/wiki/.bd
*.bd

// ck : http://en.wikipedia.org/wiki/.ck
 *.CK
!www.ck
`

	ps, err := linter.LintString(src)
	if err != nil {
		t.Fatalf("lint() returned an error: %v", err)
	}

	if want, got := 2, len(ps); want != got {
		t.Errorf("Expected %d errors, got %d", want, got)
		t.Fatal(ps)
	}

	if want, got := &ps[0], &(Problem{Message: "leading space", Level: LEVEL_WARN, Line: 4, LineSource: " com.aw"}); !reflect.DeepEqual(want, got) {
		t.Fatalf("Problem[%d] is %+v, want %+v", 0, got, want)
	}

	if want, got := &ps[1], &(Problem{Message: "leading space", Level: LEVEL_WARN, Line: 10, LineSource: " *.CK"}); !reflect.DeepEqual(want, got) {
		t.Fatalf("Problem[%d] is %+v, want %+v", 0, got, want)
	}
}
