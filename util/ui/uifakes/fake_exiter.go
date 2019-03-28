// Code generated by counterfeiter. DO NOT EDIT.
package uifakes

import (
	"sync"

	"code.cloudfoundry.org/cli/util/ui"
)

type FakeExiter struct {
	ExitStub        func(int)
	exitMutex       sync.RWMutex
	exitArgsForCall []struct {
		arg1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExiter) Exit(arg1 int) {
	fake.exitMutex.Lock()
	fake.exitArgsForCall = append(fake.exitArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Exit", []interface{}{arg1})
	fake.exitMutex.Unlock()
	if fake.ExitStub != nil {
		fake.ExitStub(arg1)
	}
}

func (fake *FakeExiter) ExitCallCount() int {
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	return len(fake.exitArgsForCall)
}

func (fake *FakeExiter) ExitCalls(stub func(int)) {
	fake.exitMutex.Lock()
	defer fake.exitMutex.Unlock()
	fake.ExitStub = stub
}

func (fake *FakeExiter) ExitArgsForCall(i int) int {
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	argsForCall := fake.exitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeExiter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExiter) recordInvocation(key string, args []interface{}) {
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

var _ ui.Exiter = new(FakeExiter)
