// automatically generated by stateify.

//go:build !arm && !mips && !mipsle && !386 && !arm && !mips && !mipsle && !386
// +build !arm,!mips,!mipsle,!386,!arm,!mips,!mipsle,!386

package atomicbitops

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (i *Int32) StateTypeName() string {
	return "pkg/atomicbitops.Int32"
}

func (i *Int32) StateFields() []string {
	return []string{
		"value",
	}
}

func (i *Int32) beforeSave() {}

// +checklocksignore
func (i *Int32) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.value)
}

func (i *Int32) afterLoad() {}

// +checklocksignore
func (i *Int32) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.value)
}

func (u *Uint32) StateTypeName() string {
	return "pkg/atomicbitops.Uint32"
}

func (u *Uint32) StateFields() []string {
	return []string{
		"value",
	}
}

func (u *Uint32) beforeSave() {}

// +checklocksignore
func (u *Uint32) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.value)
}

func (u *Uint32) afterLoad() {}

// +checklocksignore
func (u *Uint32) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.value)
}

func (i *Int64) StateTypeName() string {
	return "pkg/atomicbitops.Int64"
}

func (i *Int64) StateFields() []string {
	return []string{
		"value",
	}
}

func (i *Int64) beforeSave() {}

// +checklocksignore
func (i *Int64) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.value)
}

func (i *Int64) afterLoad() {}

// +checklocksignore
func (i *Int64) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.value)
}

func (u *Uint64) StateTypeName() string {
	return "pkg/atomicbitops.Uint64"
}

func (u *Uint64) StateFields() []string {
	return []string{
		"value",
	}
}

func (u *Uint64) beforeSave() {}

// +checklocksignore
func (u *Uint64) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	stateSinkObject.Save(0, &u.value)
}

func (u *Uint64) afterLoad() {}

// +checklocksignore
func (u *Uint64) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.value)
}

func init() {
	state.Register((*Int32)(nil))
	state.Register((*Uint32)(nil))
	state.Register((*Int64)(nil))
	state.Register((*Uint64)(nil))
}
