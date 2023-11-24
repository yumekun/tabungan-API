package errs

import (
	"errors"
	"fmt"
	"runtime"
	"sort"

	pkgerrors "github.com/pkg/errors"
)

type Error struct {
	Op Op
	User UserName
	Kind Kind
	Param Parameter
	Code Code
	Realm Realm
	Err error
}

func (e *Error) isZero() bool {
	return e.User == "" && e.Kind == 0 && e.Param == "" && e.Code == "" && e.Err == nil
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func OpStack(err error) []string {
	type o struct {
		Op    string
		Order int
	}

	e := err
	i := 0
	var os []o

	for errors.Unwrap(e) != nil {
		var errsError *Error
		if errors.As(e, &errsError) {
			if errsError.Op != "" {
				op := o{Op: string(errsError.Op), Order: i}
				os = append(os, op)
			}
		}
		e = errors.Unwrap(e)
		i++
	}

	sort.Slice(os, func(i, j int) bool { return os[i].Order > os[j].Order })

	var ops []string
	for _, op := range os {
		ops = append(ops, op.Op)
	}

	return ops
}

func TopError(err error) error {
	currentErr := err
	for errors.Unwrap(currentErr) != nil {
		currentErr = errors.Unwrap(currentErr)
	}

	return currentErr
}

type Op string
type UserName string
type Kind uint8
type Parameter string
type Code string
type Realm string
const (
	Other          Kind = iota // Unclassified error. This value is not printed in the error message.
	Invalid                    // Invalid operation for this type of item.
	IO                         // External I/O error such as network failure.
	Exist                      // Item already exists.
	NotExist                   // Item does not exist.
	Private                    // Information withheld.
	Internal                   // Internal error or inconsistency.
	BrokenLink                 // Link target does not exist.
	Database                   // Error from database.
	Validation                 // Input validation error.
	Unanticipated              // Unanticipated error.
	InvalidRequest             // Invalid Request
	Unauthenticated // Unauthenticated Request
	Unauthorized
)

func (k Kind) String() string {
	switch k {
	case Other:
		return "other error"
	case Invalid:
		return "invalid operation"
	case IO:
		return "I/O error"
	case Exist:
		return "item already exists"
	case NotExist:
		return "item does not exist"
	case BrokenLink:
		return "link target does not exist"
	case Private:
		return "information withheld"
	case Internal:
		return "internal error"
	case Database:
		return "database error"
	case Validation:
		return "input validation error"
	case Unanticipated:
		return "unanticipated error"
	case InvalidRequest:
		return "invalid request error"
	case Unauthenticated:
		return "unauthenticated request"
	case Unauthorized:
		return "unauthorized request"
	}
	return "unknown error kind"
}

func E(args ...interface{}) error {
	type stackTracer interface {
		StackTrace() pkgerrors.StackTrace
	}

	if len(args) == 0 {
		panic("call to errors.E with no arguments")
	}
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case UserName:
			e.User = arg
		case Kind:
			e.Kind = arg
		case string:
			e.Err = Str(arg)
		case *Error:
			// Make a copy
			errorCopy := *arg
			e.Err = &errorCopy
		case error:
			e.Err = arg
		case Code:
			e.Code = arg
		case Parameter:
			e.Param = arg
		case Realm:
			e.Realm = arg
		default:
			_, file, line, _ := runtime.Caller(1)
			return fmt.Errorf("errors.E: bad call from %s:%d: %v, unknown type %T, value %v in error call", file, line, args, arg, arg)
		}
	}

	prev, ok := e.Err.(*Error)
	if !ok {
		return e
	}

	if e.Kind == Other {
		e.Kind = prev.Kind
		prev.Kind = Other
	}

	if prev.Code == e.Code {
		prev.Code = ""
	}
	if e.Code == "" {
		e.Code = prev.Code
		prev.Code = ""
	}

	if prev.Param == e.Param {
		prev.Param = ""
	}
	if e.Param == "" {
		e.Param = prev.Param
		prev.Param = ""
	}

	if prev.Realm == e.Realm {
		prev.Realm = ""
	}
	if e.Realm == "" {
		e.Realm = prev.Realm
		prev.Realm = ""
	}

	return e
}

func Str(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}


func Match(err1, err2 error) bool {
	e1, ok := err1.(*Error)
	if !ok {
		return false
	}
	var e2 *Error
	e2, ok = err2.(*Error)
	if !ok {
		return false
	}
	if e1.User != "" && e2.User != e1.User {
		return false
	}
	if e1.Kind != Other && e2.Kind != e1.Kind {
		return false
	}
	if e1.Param != "" && e2.Param != e1.Param {
		return false
	}
	if e1.Code != "" && e2.Code != e1.Code {
		return false
	}
	if e1.Err != nil {
		if _, k := e1.Err.(*Error); k {
			return Match(e1.Err, e2.Err)
		}
		if e2.Err == nil || e2.Err.Error() != e1.Err.Error() {
			return false
		}
	}
	return true
}

func KindIs(kind Kind, err error) bool {
	var e *Error
	if errors.As(err, &e) {
		if e.Kind != Other {
			return e.Kind == kind
		}
		if e.Err != nil {
			return KindIs(kind, e.Err)
		}
	}
	return false
}