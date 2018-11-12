// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeDeleteOrganizationActor struct {
	DeleteOrganizationStub        func(string) (v2action.Warnings, error)
	deleteOrganizationMutex       sync.RWMutex
	deleteOrganizationArgsForCall []struct {
		arg1 string
	}
	deleteOrganizationReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	deleteOrganizationReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganization(arg1 string) (v2action.Warnings, error) {
	fake.deleteOrganizationMutex.Lock()
	ret, specificReturn := fake.deleteOrganizationReturnsOnCall[len(fake.deleteOrganizationArgsForCall)]
	fake.deleteOrganizationArgsForCall = append(fake.deleteOrganizationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteOrganization", []interface{}{arg1})
	fake.deleteOrganizationMutex.Unlock()
	if fake.DeleteOrganizationStub != nil {
		return fake.DeleteOrganizationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteOrganizationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationCallCount() int {
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	return len(fake.deleteOrganizationArgsForCall)
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationCalls(stub func(string) (v2action.Warnings, error)) {
	fake.deleteOrganizationMutex.Lock()
	defer fake.deleteOrganizationMutex.Unlock()
	fake.DeleteOrganizationStub = stub
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationArgsForCall(i int) string {
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	argsForCall := fake.deleteOrganizationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationReturns(result1 v2action.Warnings, result2 error) {
	fake.deleteOrganizationMutex.Lock()
	defer fake.deleteOrganizationMutex.Unlock()
	fake.DeleteOrganizationStub = nil
	fake.deleteOrganizationReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.deleteOrganizationMutex.Lock()
	defer fake.deleteOrganizationMutex.Unlock()
	fake.DeleteOrganizationStub = nil
	if fake.deleteOrganizationReturnsOnCall == nil {
		fake.deleteOrganizationReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.deleteOrganizationReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrganizationActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteOrganizationActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.DeleteOrganizationActor = new(FakeDeleteOrganizationActor)
