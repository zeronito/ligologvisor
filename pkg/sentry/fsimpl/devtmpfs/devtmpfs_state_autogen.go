// automatically generated by stateify.

package devtmpfs

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (fst *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/devtmpfs.FilesystemType"
}

func (fst *FilesystemType) StateFields() []string {
	return []string{
		"initErr",
		"fs",
		"root",
	}
}

func (fst *FilesystemType) beforeSave() {}

// +checklocksignore
func (fst *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fst.beforeSave()
	stateSinkObject.Save(0, &fst.initErr)
	stateSinkObject.Save(1, &fst.fs)
	stateSinkObject.Save(2, &fst.root)
}

// +checklocksignore
func (fst *FilesystemType) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fst.initErr)
	stateSourceObject.Load(1, &fst.fs)
	stateSourceObject.Load(2, &fst.root)
	stateSourceObject.AfterLoad(fst.afterLoad)
}

func init() {
	state.Register((*FilesystemType)(nil))
}
