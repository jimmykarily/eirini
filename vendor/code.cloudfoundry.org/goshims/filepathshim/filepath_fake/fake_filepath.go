// This file was generated by counterfeiter
package filepath_fake

import (
	"path/filepath"
	"sync"

	"code.cloudfoundry.org/goshims/filepathshim"
)

type FakeFilepath struct {
	MatchStub        func(pattern, name string) (matched bool, err error)
	matchMutex       sync.RWMutex
	matchArgsForCall []struct {
		pattern string
		name    string
	}
	matchReturns struct {
		result1 bool
		result2 error
	}
	GlobStub        func(pattern string) (matches []string, err error)
	globMutex       sync.RWMutex
	globArgsForCall []struct {
		pattern string
	}
	globReturns struct {
		result1 []string
		result2 error
	}
	CleanStub        func(path string) string
	cleanMutex       sync.RWMutex
	cleanArgsForCall []struct {
		path string
	}
	cleanReturns struct {
		result1 string
	}
	ToSlashStub        func(path string) string
	toSlashMutex       sync.RWMutex
	toSlashArgsForCall []struct {
		path string
	}
	toSlashReturns struct {
		result1 string
	}
	FromSlashStub        func(path string) string
	fromSlashMutex       sync.RWMutex
	fromSlashArgsForCall []struct {
		path string
	}
	fromSlashReturns struct {
		result1 string
	}
	SplitListStub        func(path string) []string
	splitListMutex       sync.RWMutex
	splitListArgsForCall []struct {
		path string
	}
	splitListReturns struct {
		result1 []string
	}
	SplitStub        func(path string) (dir string, file string)
	splitMutex       sync.RWMutex
	splitArgsForCall []struct {
		path string
	}
	splitReturns struct {
		result1 string
		result2 string
	}
	JoinStub        func(elem ...string) string
	joinMutex       sync.RWMutex
	joinArgsForCall []struct {
		elem []string
	}
	joinReturns struct {
		result1 string
	}
	ExtStub        func(path string) string
	extMutex       sync.RWMutex
	extArgsForCall []struct {
		path string
	}
	extReturns struct {
		result1 string
	}
	EvalSymlinksStub        func(path string) (string, error)
	evalSymlinksMutex       sync.RWMutex
	evalSymlinksArgsForCall []struct {
		path string
	}
	evalSymlinksReturns struct {
		result1 string
		result2 error
	}
	AbsStub        func(path string) (string, error)
	absMutex       sync.RWMutex
	absArgsForCall []struct {
		path string
	}
	absReturns struct {
		result1 string
		result2 error
	}
	RelStub        func(basepath, targpath string) (string, error)
	relMutex       sync.RWMutex
	relArgsForCall []struct {
		basepath string
		targpath string
	}
	relReturns struct {
		result1 string
		result2 error
	}
	WalkStub        func(root string, walkFn filepath.WalkFunc) error
	walkMutex       sync.RWMutex
	walkArgsForCall []struct {
		root   string
		walkFn filepath.WalkFunc
	}
	walkReturns struct {
		result1 error
	}
	BaseStub        func(path string) string
	baseMutex       sync.RWMutex
	baseArgsForCall []struct {
		path string
	}
	baseReturns struct {
		result1 string
	}
	DirStub        func(path string) string
	dirMutex       sync.RWMutex
	dirArgsForCall []struct {
		path string
	}
	dirReturns struct {
		result1 string
	}
	VolumeNameStub        func(path string) string
	volumeNameMutex       sync.RWMutex
	volumeNameArgsForCall []struct {
		path string
	}
	volumeNameReturns struct {
		result1 string
	}
	IsAbsStub        func(path string) bool
	isAbsMutex       sync.RWMutex
	isAbsArgsForCall []struct {
		path string
	}
	isAbsReturns struct {
		result1 bool
	}
	HasPrefixStub        func(p, prefix string) bool
	hasPrefixMutex       sync.RWMutex
	hasPrefixArgsForCall []struct {
		p      string
		prefix string
	}
	hasPrefixReturns struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFilepath) Match(pattern string, name string) (matched bool, err error) {
	fake.matchMutex.Lock()
	fake.matchArgsForCall = append(fake.matchArgsForCall, struct {
		pattern string
		name    string
	}{pattern, name})
	fake.recordInvocation("Match", []interface{}{pattern, name})
	fake.matchMutex.Unlock()
	if fake.MatchStub != nil {
		return fake.MatchStub(pattern, name)
	}
	return fake.matchReturns.result1, fake.matchReturns.result2
}

func (fake *FakeFilepath) MatchCallCount() int {
	fake.matchMutex.RLock()
	defer fake.matchMutex.RUnlock()
	return len(fake.matchArgsForCall)
}

func (fake *FakeFilepath) MatchArgsForCall(i int) (string, string) {
	fake.matchMutex.RLock()
	defer fake.matchMutex.RUnlock()
	return fake.matchArgsForCall[i].pattern, fake.matchArgsForCall[i].name
}

func (fake *FakeFilepath) MatchReturns(result1 bool, result2 error) {
	fake.MatchStub = nil
	fake.matchReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeFilepath) Glob(pattern string) (matches []string, err error) {
	fake.globMutex.Lock()
	fake.globArgsForCall = append(fake.globArgsForCall, struct {
		pattern string
	}{pattern})
	fake.recordInvocation("Glob", []interface{}{pattern})
	fake.globMutex.Unlock()
	if fake.GlobStub != nil {
		return fake.GlobStub(pattern)
	}
	return fake.globReturns.result1, fake.globReturns.result2
}

func (fake *FakeFilepath) GlobCallCount() int {
	fake.globMutex.RLock()
	defer fake.globMutex.RUnlock()
	return len(fake.globArgsForCall)
}

func (fake *FakeFilepath) GlobArgsForCall(i int) string {
	fake.globMutex.RLock()
	defer fake.globMutex.RUnlock()
	return fake.globArgsForCall[i].pattern
}

func (fake *FakeFilepath) GlobReturns(result1 []string, result2 error) {
	fake.GlobStub = nil
	fake.globReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeFilepath) Clean(path string) string {
	fake.cleanMutex.Lock()
	fake.cleanArgsForCall = append(fake.cleanArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Clean", []interface{}{path})
	fake.cleanMutex.Unlock()
	if fake.CleanStub != nil {
		return fake.CleanStub(path)
	}
	return fake.cleanReturns.result1
}

func (fake *FakeFilepath) CleanCallCount() int {
	fake.cleanMutex.RLock()
	defer fake.cleanMutex.RUnlock()
	return len(fake.cleanArgsForCall)
}

func (fake *FakeFilepath) CleanArgsForCall(i int) string {
	fake.cleanMutex.RLock()
	defer fake.cleanMutex.RUnlock()
	return fake.cleanArgsForCall[i].path
}

func (fake *FakeFilepath) CleanReturns(result1 string) {
	fake.CleanStub = nil
	fake.cleanReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilepath) ToSlash(path string) string {
	fake.toSlashMutex.Lock()
	fake.toSlashArgsForCall = append(fake.toSlashArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("ToSlash", []interface{}{path})
	fake.toSlashMutex.Unlock()
	if fake.ToSlashStub != nil {
		return fake.ToSlashStub(path)
	}
	return fake.toSlashReturns.result1
}

func (fake *FakeFilepath) ToSlashCallCount() int {
	fake.toSlashMutex.RLock()
	defer fake.toSlashMutex.RUnlock()
	return len(fake.toSlashArgsForCall)
}

func (fake *FakeFilepath) ToSlashArgsForCall(i int) string {
	fake.toSlashMutex.RLock()
	defer fake.toSlashMutex.RUnlock()
	return fake.toSlashArgsForCall[i].path
}

func (fake *FakeFilepath) ToSlashReturns(result1 string) {
	fake.ToSlashStub = nil
	fake.toSlashReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilepath) FromSlash(path string) string {
	fake.fromSlashMutex.Lock()
	fake.fromSlashArgsForCall = append(fake.fromSlashArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("FromSlash", []interface{}{path})
	fake.fromSlashMutex.Unlock()
	if fake.FromSlashStub != nil {
		return fake.FromSlashStub(path)
	}
	return fake.fromSlashReturns.result1
}

func (fake *FakeFilepath) FromSlashCallCount() int {
	fake.fromSlashMutex.RLock()
	defer fake.fromSlashMutex.RUnlock()
	return len(fake.fromSlashArgsForCall)
}

func (fake *FakeFilepath) FromSlashArgsForCall(i int) string {
	fake.fromSlashMutex.RLock()
	defer fake.fromSlashMutex.RUnlock()
	return fake.fromSlashArgsForCall[i].path
}

func (fake *FakeFilepath) FromSlashReturns(result1 string) {
	fake.FromSlashStub = nil
	fake.fromSlashReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilepath) SplitList(path string) []string {
	fake.splitListMutex.Lock()
	fake.splitListArgsForCall = append(fake.splitListArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("SplitList", []interface{}{path})
	fake.splitListMutex.Unlock()
	if fake.SplitListStub != nil {
		return fake.SplitListStub(path)
	}
	return fake.splitListReturns.result1
}

func (fake *FakeFilepath) SplitListCallCount() int {
	fake.splitListMutex.RLock()
	defer fake.splitListMutex.RUnlock()
	return len(fake.splitListArgsForCall)
}

func (fake *FakeFilepath) SplitListArgsForCall(i int) string {
	fake.splitListMutex.RLock()
	defer fake.splitListMutex.RUnlock()
	return fake.splitListArgsForCall[i].path
}

func (fake *FakeFilepath) SplitListReturns(result1 []string) {
	fake.SplitListStub = nil
	fake.splitListReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeFilepath) Split(path string) (dir string, file string) {
	fake.splitMutex.Lock()
	fake.splitArgsForCall = append(fake.splitArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Split", []interface{}{path})
	fake.splitMutex.Unlock()
	if fake.SplitStub != nil {
		return fake.SplitStub(path)
	}
	return fake.splitReturns.result1, fake.splitReturns.result2
}

func (fake *FakeFilepath) SplitCallCount() int {
	fake.splitMutex.RLock()
	defer fake.splitMutex.RUnlock()
	return len(fake.splitArgsForCall)
}

func (fake *FakeFilepath) SplitArgsForCall(i int) string {
	fake.splitMutex.RLock()
	defer fake.splitMutex.RUnlock()
	return fake.splitArgsForCall[i].path
}

func (fake *FakeFilepath) SplitReturns(result1 string, result2 string) {
	fake.SplitStub = nil
	fake.splitReturns = struct {
		result1 string
		result2 string
	}{result1, result2}
}

func (fake *FakeFilepath) Join(elem ...string) string {
	fake.joinMutex.Lock()
	fake.joinArgsForCall = append(fake.joinArgsForCall, struct {
		elem []string
	}{elem})
	fake.recordInvocation("Join", []interface{}{elem})
	fake.joinMutex.Unlock()
	if fake.JoinStub != nil {
		return fake.JoinStub(elem...)
	}
	return fake.joinReturns.result1
}

func (fake *FakeFilepath) JoinCallCount() int {
	fake.joinMutex.RLock()
	defer fake.joinMutex.RUnlock()
	return len(fake.joinArgsForCall)
}

func (fake *FakeFilepath) JoinArgsForCall(i int) []string {
	fake.joinMutex.RLock()
	defer fake.joinMutex.RUnlock()
	return fake.joinArgsForCall[i].elem
}

func (fake *FakeFilepath) JoinReturns(result1 string) {
	fake.JoinStub = nil
	fake.joinReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilepath) Ext(path string) string {
	fake.extMutex.Lock()
	fake.extArgsForCall = append(fake.extArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Ext", []interface{}{path})
	fake.extMutex.Unlock()
	if fake.ExtStub != nil {
		return fake.ExtStub(path)
	}
	return fake.extReturns.result1
}

func (fake *FakeFilepath) ExtCallCount() int {
	fake.extMutex.RLock()
	defer fake.extMutex.RUnlock()
	return len(fake.extArgsForCall)
}

func (fake *FakeFilepath) ExtArgsForCall(i int) string {
	fake.extMutex.RLock()
	defer fake.extMutex.RUnlock()
	return fake.extArgsForCall[i].path
}

func (fake *FakeFilepath) ExtReturns(result1 string) {
	fake.ExtStub = nil
	fake.extReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilepath) EvalSymlinks(path string) (string, error) {
	fake.evalSymlinksMutex.Lock()
	fake.evalSymlinksArgsForCall = append(fake.evalSymlinksArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("EvalSymlinks", []interface{}{path})
	fake.evalSymlinksMutex.Unlock()
	if fake.EvalSymlinksStub != nil {
		return fake.EvalSymlinksStub(path)
	}
	return fake.evalSymlinksReturns.result1, fake.evalSymlinksReturns.result2
}

func (fake *FakeFilepath) EvalSymlinksCallCount() int {
	fake.evalSymlinksMutex.RLock()
	defer fake.evalSymlinksMutex.RUnlock()
	return len(fake.evalSymlinksArgsForCall)
}

func (fake *FakeFilepath) EvalSymlinksArgsForCall(i int) string {
	fake.evalSymlinksMutex.RLock()
	defer fake.evalSymlinksMutex.RUnlock()
	return fake.evalSymlinksArgsForCall[i].path
}

func (fake *FakeFilepath) EvalSymlinksReturns(result1 string, result2 error) {
	fake.EvalSymlinksStub = nil
	fake.evalSymlinksReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFilepath) Abs(path string) (string, error) {
	fake.absMutex.Lock()
	fake.absArgsForCall = append(fake.absArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Abs", []interface{}{path})
	fake.absMutex.Unlock()
	if fake.AbsStub != nil {
		return fake.AbsStub(path)
	}
	return fake.absReturns.result1, fake.absReturns.result2
}

func (fake *FakeFilepath) AbsCallCount() int {
	fake.absMutex.RLock()
	defer fake.absMutex.RUnlock()
	return len(fake.absArgsForCall)
}

func (fake *FakeFilepath) AbsArgsForCall(i int) string {
	fake.absMutex.RLock()
	defer fake.absMutex.RUnlock()
	return fake.absArgsForCall[i].path
}

func (fake *FakeFilepath) AbsReturns(result1 string, result2 error) {
	fake.AbsStub = nil
	fake.absReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFilepath) Rel(basepath string, targpath string) (string, error) {
	fake.relMutex.Lock()
	fake.relArgsForCall = append(fake.relArgsForCall, struct {
		basepath string
		targpath string
	}{basepath, targpath})
	fake.recordInvocation("Rel", []interface{}{basepath, targpath})
	fake.relMutex.Unlock()
	if fake.RelStub != nil {
		return fake.RelStub(basepath, targpath)
	}
	return fake.relReturns.result1, fake.relReturns.result2
}

func (fake *FakeFilepath) RelCallCount() int {
	fake.relMutex.RLock()
	defer fake.relMutex.RUnlock()
	return len(fake.relArgsForCall)
}

func (fake *FakeFilepath) RelArgsForCall(i int) (string, string) {
	fake.relMutex.RLock()
	defer fake.relMutex.RUnlock()
	return fake.relArgsForCall[i].basepath, fake.relArgsForCall[i].targpath
}

func (fake *FakeFilepath) RelReturns(result1 string, result2 error) {
	fake.RelStub = nil
	fake.relReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFilepath) Walk(root string, walkFn filepath.WalkFunc) error {
	fake.walkMutex.Lock()
	fake.walkArgsForCall = append(fake.walkArgsForCall, struct {
		root   string
		walkFn filepath.WalkFunc
	}{root, walkFn})
	fake.recordInvocation("Walk", []interface{}{root, walkFn})
	fake.walkMutex.Unlock()
	if fake.WalkStub != nil {
		return fake.WalkStub(root, walkFn)
	}
	return fake.walkReturns.result1
}

func (fake *FakeFilepath) WalkCallCount() int {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return len(fake.walkArgsForCall)
}

func (fake *FakeFilepath) WalkArgsForCall(i int) (string, filepath.WalkFunc) {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return fake.walkArgsForCall[i].root, fake.walkArgsForCall[i].walkFn
}

func (fake *FakeFilepath) WalkReturns(result1 error) {
	fake.WalkStub = nil
	fake.walkReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilepath) Base(path string) string {
	fake.baseMutex.Lock()
	fake.baseArgsForCall = append(fake.baseArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Base", []interface{}{path})
	fake.baseMutex.Unlock()
	if fake.BaseStub != nil {
		return fake.BaseStub(path)
	}
	return fake.baseReturns.result1
}

func (fake *FakeFilepath) BaseCallCount() int {
	fake.baseMutex.RLock()
	defer fake.baseMutex.RUnlock()
	return len(fake.baseArgsForCall)
}

func (fake *FakeFilepath) BaseArgsForCall(i int) string {
	fake.baseMutex.RLock()
	defer fake.baseMutex.RUnlock()
	return fake.baseArgsForCall[i].path
}

func (fake *FakeFilepath) BaseReturns(result1 string) {
	fake.BaseStub = nil
	fake.baseReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilepath) Dir(path string) string {
	fake.dirMutex.Lock()
	fake.dirArgsForCall = append(fake.dirArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Dir", []interface{}{path})
	fake.dirMutex.Unlock()
	if fake.DirStub != nil {
		return fake.DirStub(path)
	}
	return fake.dirReturns.result1
}

func (fake *FakeFilepath) DirCallCount() int {
	fake.dirMutex.RLock()
	defer fake.dirMutex.RUnlock()
	return len(fake.dirArgsForCall)
}

func (fake *FakeFilepath) DirArgsForCall(i int) string {
	fake.dirMutex.RLock()
	defer fake.dirMutex.RUnlock()
	return fake.dirArgsForCall[i].path
}

func (fake *FakeFilepath) DirReturns(result1 string) {
	fake.DirStub = nil
	fake.dirReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilepath) VolumeName(path string) string {
	fake.volumeNameMutex.Lock()
	fake.volumeNameArgsForCall = append(fake.volumeNameArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("VolumeName", []interface{}{path})
	fake.volumeNameMutex.Unlock()
	if fake.VolumeNameStub != nil {
		return fake.VolumeNameStub(path)
	}
	return fake.volumeNameReturns.result1
}

func (fake *FakeFilepath) VolumeNameCallCount() int {
	fake.volumeNameMutex.RLock()
	defer fake.volumeNameMutex.RUnlock()
	return len(fake.volumeNameArgsForCall)
}

func (fake *FakeFilepath) VolumeNameArgsForCall(i int) string {
	fake.volumeNameMutex.RLock()
	defer fake.volumeNameMutex.RUnlock()
	return fake.volumeNameArgsForCall[i].path
}

func (fake *FakeFilepath) VolumeNameReturns(result1 string) {
	fake.VolumeNameStub = nil
	fake.volumeNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilepath) IsAbs(path string) bool {
	fake.isAbsMutex.Lock()
	fake.isAbsArgsForCall = append(fake.isAbsArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("IsAbs", []interface{}{path})
	fake.isAbsMutex.Unlock()
	if fake.IsAbsStub != nil {
		return fake.IsAbsStub(path)
	}
	return fake.isAbsReturns.result1
}

func (fake *FakeFilepath) IsAbsCallCount() int {
	fake.isAbsMutex.RLock()
	defer fake.isAbsMutex.RUnlock()
	return len(fake.isAbsArgsForCall)
}

func (fake *FakeFilepath) IsAbsArgsForCall(i int) string {
	fake.isAbsMutex.RLock()
	defer fake.isAbsMutex.RUnlock()
	return fake.isAbsArgsForCall[i].path
}

func (fake *FakeFilepath) IsAbsReturns(result1 bool) {
	fake.IsAbsStub = nil
	fake.isAbsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeFilepath) HasPrefix(p string, prefix string) bool {
	fake.hasPrefixMutex.Lock()
	fake.hasPrefixArgsForCall = append(fake.hasPrefixArgsForCall, struct {
		p      string
		prefix string
	}{p, prefix})
	fake.recordInvocation("HasPrefix", []interface{}{p, prefix})
	fake.hasPrefixMutex.Unlock()
	if fake.HasPrefixStub != nil {
		return fake.HasPrefixStub(p, prefix)
	}
	return fake.hasPrefixReturns.result1
}

func (fake *FakeFilepath) HasPrefixCallCount() int {
	fake.hasPrefixMutex.RLock()
	defer fake.hasPrefixMutex.RUnlock()
	return len(fake.hasPrefixArgsForCall)
}

func (fake *FakeFilepath) HasPrefixArgsForCall(i int) (string, string) {
	fake.hasPrefixMutex.RLock()
	defer fake.hasPrefixMutex.RUnlock()
	return fake.hasPrefixArgsForCall[i].p, fake.hasPrefixArgsForCall[i].prefix
}

func (fake *FakeFilepath) HasPrefixReturns(result1 bool) {
	fake.HasPrefixStub = nil
	fake.hasPrefixReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeFilepath) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.matchMutex.RLock()
	defer fake.matchMutex.RUnlock()
	fake.globMutex.RLock()
	defer fake.globMutex.RUnlock()
	fake.cleanMutex.RLock()
	defer fake.cleanMutex.RUnlock()
	fake.toSlashMutex.RLock()
	defer fake.toSlashMutex.RUnlock()
	fake.fromSlashMutex.RLock()
	defer fake.fromSlashMutex.RUnlock()
	fake.splitListMutex.RLock()
	defer fake.splitListMutex.RUnlock()
	fake.splitMutex.RLock()
	defer fake.splitMutex.RUnlock()
	fake.joinMutex.RLock()
	defer fake.joinMutex.RUnlock()
	fake.extMutex.RLock()
	defer fake.extMutex.RUnlock()
	fake.evalSymlinksMutex.RLock()
	defer fake.evalSymlinksMutex.RUnlock()
	fake.absMutex.RLock()
	defer fake.absMutex.RUnlock()
	fake.relMutex.RLock()
	defer fake.relMutex.RUnlock()
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	fake.baseMutex.RLock()
	defer fake.baseMutex.RUnlock()
	fake.dirMutex.RLock()
	defer fake.dirMutex.RUnlock()
	fake.volumeNameMutex.RLock()
	defer fake.volumeNameMutex.RUnlock()
	fake.isAbsMutex.RLock()
	defer fake.isAbsMutex.RUnlock()
	fake.hasPrefixMutex.RLock()
	defer fake.hasPrefixMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFilepath) recordInvocation(key string, args []interface{}) {
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

var _ filepathshim.Filepath = new(FakeFilepath)
