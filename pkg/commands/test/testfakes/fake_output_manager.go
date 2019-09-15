// Code generated by counterfeiter. DO NOT EDIT.
package testfakes

import (
	"sync"

	"github.com/instrumenta/conftest/pkg/commands/test"
)

type FakeOutputManager struct {
	FlushStub        func() error
	flushMutex       sync.RWMutex
	flushArgsForCall []struct {
	}
	flushReturns struct {
		result1 error
	}
	flushReturnsOnCall map[int]struct {
		result1 error
	}
	PutStub        func(string, test.CheckResult) error
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 string
		arg2 test.CheckResult
	}
	putReturns struct {
		result1 error
	}
	putReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOutputManager) Flush() error {
	fake.flushMutex.Lock()
	ret, specificReturn := fake.flushReturnsOnCall[len(fake.flushArgsForCall)]
	fake.flushArgsForCall = append(fake.flushArgsForCall, struct {
	}{})
	fake.recordInvocation("Flush", []interface{}{})
	fake.flushMutex.Unlock()
	if fake.FlushStub != nil {
		return fake.FlushStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.flushReturns
	return fakeReturns.result1
}

func (fake *FakeOutputManager) FlushCallCount() int {
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	return len(fake.flushArgsForCall)
}

func (fake *FakeOutputManager) FlushCalls(stub func() error) {
	fake.flushMutex.Lock()
	defer fake.flushMutex.Unlock()
	fake.FlushStub = stub
}

func (fake *FakeOutputManager) FlushReturns(result1 error) {
	fake.flushMutex.Lock()
	defer fake.flushMutex.Unlock()
	fake.FlushStub = nil
	fake.flushReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputManager) FlushReturnsOnCall(i int, result1 error) {
	fake.flushMutex.Lock()
	defer fake.flushMutex.Unlock()
	fake.FlushStub = nil
	if fake.flushReturnsOnCall == nil {
		fake.flushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.flushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputManager) Put(arg1 string, arg2 test.CheckResult) error {
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 string
		arg2 test.CheckResult
	}{arg1, arg2})
	fake.recordInvocation("Put", []interface{}{arg1, arg2})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.putReturns
	return fakeReturns.result1
}

func (fake *FakeOutputManager) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeOutputManager) PutCalls(stub func(string, test.CheckResult) error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = stub
}

func (fake *FakeOutputManager) PutArgsForCall(i int) (string, test.CheckResult) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	argsForCall := fake.putArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOutputManager) PutReturns(result1 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputManager) PutReturnsOnCall(i int, result1 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOutputManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ test.OutputManager = new(FakeOutputManager)
