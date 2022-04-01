package utils

import (
	"bytes"
	"fmt"
	"runtime/debug"
	"sync"
)

type GenericFunction func() error

// RunAsync executes a list of functions asynchronously until all are completed,
// or until the first error (even if other background functions are still running).
//
// Example:
//
//   func doSomething() error { ... }
//   func doSomethingElse() error { ... }
//   funcsToRun := []GenericFunction{ doSomething, doSomethingElse }
//   if err := RunAsync(funcsToRun...); err != nil { return nil, err }
//   print getResults.output
//
//
// Functions with return values can be wrapped with a receiver function -
// The receiver struct should define a field to store the calculated value.
// Example:
//
//   type SomeCollector struct{
//     input  int
//     ...
//     output int
//   }
//   func (c *SomeCollector) Run() error {
//     ...
//     c.output = 200
//     ...
//   }
//   collector := SomeCollector{ input: 1 }
//   funcsToRun := []GenericFunction{ doSomething, collector.Run }
//   if err := RunAsync(funcsToRun...); err != nil { return nil, err }
//   print collector.output
//
//
// Multiple return values can also be collected by using closures.
// Notice the collect returns a function in order to avoid all closures being bound to same values.
// (For more info, see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables)
// Example:
//
//   func RunQueriesAsync(ctx context.Context, queries []*Query) ([]*QueryResults, error) {
//     funcsToRun := make([]GenericFunction, len(queries))
//     results := make([]*QueryResults, len(queries))
//     collectQueryResults := func(i int, query *Query) func() error {
//       return func() error {
//         result, err := RunQuery(ctx, query)
//         if err != nil { return err }
//         results[i] = result
//         return nil
//       }
//     }
//     for i, query := range queries { funcsToRun[i] = collectQueryResults(i, query) }
//     if err := RunAsync(funcsToRun...); err != nil { return nil, err }
//     return results, nil
//   }
//
func RunAsync(functions ...GenericFunction) error {
	if errors := asyncImplementation(true, functions...); errors != nil {
		return errors[0]
	}
	return nil
}

// RunAsyncAllowErrors executes a list of functions asynchronously until all are completed.
// Any/all errors that occur in functions being executed are returned as an indexed
// slice of errors. The indices of the functions in the list passed to the function
// correspond to the indices of errors in the slice that is returned. If no errors occur,
// the function will return nil.
//
// Example:
//
//   func doSomething0() error { ... }
//   func doSomething1() error { ... }
//   func doSomething2() error { ... }
//
//   funcsToRun := []GenericFunction{ doSomething1, doSomething2, doSomething3 }
//
//   if errors := RunAsync(funcsToRun...); errors != nil {
//     for idx, err := range errors {
//       if err != nil {
//       	logf("doSomething%v caused error: %s", idx, err)
//       }
//     }
//   }
//
// Refer to `RunAsync` documentation for more usage examples.
//
func RunAsyncAllowErrors(functions ...GenericFunction) []error {
	return asyncImplementation(false, functions...)
}

func asyncImplementation(exitOnError bool, functions ...GenericFunction) []error {
	numFuncs := len(functions)
	errors := make([]error, numFuncs)
	hasEncounteredErr := false
	var wg sync.WaitGroup
	if numFuncs == 0 {
		return nil
	}
	wg.Add(numFuncs)

	// Use buffered channel with enough slots to fit an error per function.
	// Otherwise sending could block the goroutines forever.
	done := make(chan struct{}, numFuncs)

	// Safely deals with an encountered error
	handleError := func(err error, j int) {
		hasEncounteredErr = true
		if exitOnError {
			errors[0] = err
			// Ensure channel stays open if multiple functions error and close
			// channel when all functions complete. RunAsync will still return
			// after first error.
			done <- struct{}{}
		} else {
			errors[j] = err
		}
		wg.Done()
	}
	// Wait group in go routine.  Ensure either all pass or wait forever.
	go func() {
		wg.Wait()
		close(done)
	}()
	for j := range functions {
		go func(j int) {
			defer func() {
				if r := recover(); r != nil {
					// Skip 4 stack frames:
					// 1) debug.Stack()
					// 2) formatStack()
					// 3) this anonymous func
					// 4) runtime/panic
					err := fmt.Errorf("panic in async function: %v\n%s",
						r, formatStack(4))
					handleError(err, j)
				}
			}()
			if err := functions[j](); err != nil {
				handleError(err, j)
				return
			}
			wg.Done()
		}(j)
	}
	<-done

	if hasEncounteredErr {
		return errors
	}
	return nil
}

func formatStack(skip int) string {
	lines := bytes.Split(bytes.TrimSpace(debug.Stack()), []byte("\n"))
	formatted := bytes.Join(lines[1+2*skip:], []byte("\n"))
	return string(formatted)
}
