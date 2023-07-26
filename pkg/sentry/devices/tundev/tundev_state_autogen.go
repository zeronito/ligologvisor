// automatically generated by stateify.

package tundev

import (
	"github.com/nicocha30/gvisor-ligolo/pkg/state"
)

func (t *tunDevice) StateTypeName() string {
	return "pkg/sentry/devices/tundev.tunDevice"
}

func (t *tunDevice) StateFields() []string {
	return []string{}
}

func (t *tunDevice) beforeSave() {}

// +checklocksignore
func (t *tunDevice) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
}

func (t *tunDevice) afterLoad() {}

// +checklocksignore
func (t *tunDevice) StateLoad(stateSourceObject state.Source) {
}

func (fd *tunFD) StateTypeName() string {
	return "pkg/sentry/devices/tundev.tunFD"
}

func (fd *tunFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"device",
	}
}

func (fd *tunFD) beforeSave() {}

// +checklocksignore
func (fd *tunFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.NoLockFD)
	stateSinkObject.Save(4, &fd.device)
}

func (fd *tunFD) afterLoad() {}

// +checklocksignore
func (fd *tunFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.NoLockFD)
	stateSourceObject.Load(4, &fd.device)
}

func init() {
	state.Register((*tunDevice)(nil))
	state.Register((*tunFD)(nil))
}
