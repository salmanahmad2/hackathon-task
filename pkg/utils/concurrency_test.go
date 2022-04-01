package utils

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type ConcurrencyTestSuite struct {
	suite.Suite
}

func doSomething1() error {
	time.Sleep(100 * time.Millisecond)
	return nil
}

func doSomething2() error {
	time.Sleep(100 * time.Millisecond)
	return nil
}

type doSomethingElseWithResults struct {
	input  int // must be a positive number
	output int
}

func (d *doSomethingElseWithResults) Run() error {
	if d.input < 0 {
		return errors.Errorf("doSomethingElseWithResults input must be a positive number.")
	}
	time.Sleep(100 * time.Millisecond)
	d.output = d.input * 2
	return nil
}

//--------------------------------------------------------------------------------------------------
// RunAsync
//--------------------------------------------------------------------------------------------------
func (c *ConcurrencyTestSuite) TestRunAsyncTwoFunctions() {
	err := RunAsync(doSomething1, doSomething2)
	c.Require().NoError(err)
}

func (c *ConcurrencyTestSuite) TestRunAsyncSeveral() {
	get1 := doSomethingElseWithResults{input: 1}
	get2 := doSomethingElseWithResults{input: 20}
	get3 := doSomethingElseWithResults{input: 300}
	get4 := doSomethingElseWithResults{input: 4000}

	funcsToRun := []GenericFunction{get1.Run, get2.Run, get3.Run, get4.Run}
	err := RunAsync(funcsToRun...)
	c.Require().NoError(err)
	c.Require().Equal(2, get1.output)
	c.Require().Equal(40, get2.output)
	c.Require().Equal(600, get3.output)
	c.Require().Equal(8000, get4.output)
}

func (c *ConcurrencyTestSuite) TestRunAsyncReturnError() {
	getErr1 := doSomethingElseWithResults{input: -1}
	getErr2 := doSomethingElseWithResults{input: -1}
	get1 := doSomethingElseWithResults{input: 20}
	get2 := doSomethingElseWithResults{input: 300}
	get3 := doSomethingElseWithResults{input: 4000}

	funcsToRun := []GenericFunction{get1.Run, get2.Run, get3.Run, getErr1.Run, getErr2.Run}
	err := RunAsync(funcsToRun...)
	c.Require().Error(err)
}

var nilPointer *int

func (c *ConcurrencyTestSuite) TestRunAsyncPanic() {
	f1 := func() error {
		return nil
	}
	f2 := func() error {
		*nilPointer = 42
		return nil
	}
	err := RunAsync(f1, f2)
	c.Require().Error(err)
	c.Require().Contains(err.Error(), "panic in async function")
}

//--------------------------------------------------------------------------------------------------
// RunAsyncAllowErrors
//--------------------------------------------------------------------------------------------------
func (c *ConcurrencyTestSuite) TestRunAsyncAllowErrorsTwoFunctions() {
	indexedErrors := RunAsyncAllowErrors(doSomething1, doSomething2)
	c.Require().Nil(indexedErrors)
}

func (c *ConcurrencyTestSuite) TestRunAsyncAllowErrorsSeveral() {
	get1 := doSomethingElseWithResults{input: 1}
	get2 := doSomethingElseWithResults{input: 20}
	get3 := doSomethingElseWithResults{input: 300}
	get4 := doSomethingElseWithResults{input: 4000}

	funcsToRun := []GenericFunction{get1.Run, get2.Run, get3.Run, get4.Run}
	indexedErrors := RunAsyncAllowErrors(funcsToRun...)
	c.Require().Equal(2, get1.output)
	c.Require().Equal(40, get2.output)
	c.Require().Equal(600, get3.output)
	c.Require().Equal(8000, get4.output)
	c.Require().Nil(indexedErrors)
}

func (c *ConcurrencyTestSuite) TestRunAsyncAllowErrorsReturnError() {
	getErr1 := doSomethingElseWithResults{input: -1}
	get1 := doSomethingElseWithResults{input: 1}
	get2 := doSomethingElseWithResults{input: 20}
	get3 := doSomethingElseWithResults{input: 300}
	get4 := doSomethingElseWithResults{input: 4000}
	getErr2 := doSomethingElseWithResults{input: -1}

	funcsToRun := []GenericFunction{
		getErr1.Run, get1.Run, get2.Run, get3.Run, get4.Run, getErr2.Run}
	indexedErrors := RunAsyncAllowErrors(funcsToRun...)
	c.Require().Equal(len(funcsToRun), len(indexedErrors))
	c.Require().Error(indexedErrors[0])
	c.Require().Error(indexedErrors[5])
}

func (c *ConcurrencyTestSuite) TestRunAsyncAllowErrorsPanic() {
	f1 := func() error {
		return nil
	}
	f2 := func() error {
		*nilPointer = 42
		return nil
	}
	indexedErrors := RunAsyncAllowErrors(f1, f2)
	c.Require().Equal(2, len(indexedErrors))
	c.Require().Error(indexedErrors[1])
	c.Require().Contains(indexedErrors[1].Error(), "panic in async function")
}

func (c *ConcurrencyTestSuite) TestRunAsyncAllowErrorsNoFns() {
	funcsToRun := []GenericFunction{}
	indexedErrors := RunAsyncAllowErrors(funcsToRun...)
	c.Require().Nil(indexedErrors)
}

func (s *ConcurrencyTestSuite) SetupSuite() {}

func (s *ConcurrencyTestSuite) TearDownSuite() {}

func (s *ConcurrencyTestSuite) SetupTest() {}

func TestAdapterTestSuite(t *testing.T) {
	suite.Run(t, &ConcurrencyTestSuite{})
}
