// Code generated by counterfeiter. DO NOT EDIT.
package metricsfakes

import (
	"sync"

	"code.cloudfoundry.org/eirini/metrics"
)

type FakeForwarder struct {
	ForwardStub        func(metrics.Message)
	forwardMutex       sync.RWMutex
	forwardArgsForCall []struct {
		arg1 metrics.Message
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeForwarder) Forward(arg1 metrics.Message) {
	fake.forwardMutex.Lock()
	fake.forwardArgsForCall = append(fake.forwardArgsForCall, struct {
		arg1 metrics.Message
	}{arg1})
	fake.recordInvocation("Forward", []interface{}{arg1})
	fake.forwardMutex.Unlock()
	if fake.ForwardStub != nil {
		fake.ForwardStub(arg1)
	}
}

func (fake *FakeForwarder) ForwardCallCount() int {
	fake.forwardMutex.RLock()
	defer fake.forwardMutex.RUnlock()
	return len(fake.forwardArgsForCall)
}

func (fake *FakeForwarder) ForwardCalls(stub func(metrics.Message)) {
	fake.forwardMutex.Lock()
	defer fake.forwardMutex.Unlock()
	fake.ForwardStub = stub
}

func (fake *FakeForwarder) ForwardArgsForCall(i int) metrics.Message {
	fake.forwardMutex.RLock()
	defer fake.forwardMutex.RUnlock()
	argsForCall := fake.forwardArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeForwarder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.forwardMutex.RLock()
	defer fake.forwardMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeForwarder) recordInvocation(key string, args []interface{}) {
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

var _ metrics.Forwarder = new(FakeForwarder)
