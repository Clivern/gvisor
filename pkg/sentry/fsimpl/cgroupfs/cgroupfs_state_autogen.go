// automatically generated by stateify.

package cgroupfs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (c *controllerCommon) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.controllerCommon"
}

func (c *controllerCommon) StateFields() []string {
	return []string{
		"ty",
		"fs",
	}
}

func (c *controllerCommon) beforeSave() {}

// +checklocksignore
func (c *controllerCommon) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.ty)
	stateSinkObject.Save(1, &c.fs)
}

func (c *controllerCommon) afterLoad() {}

// +checklocksignore
func (c *controllerCommon) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.ty)
	stateSourceObject.Load(1, &c.fs)
}

func (c *cgroupInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cgroupInode"
}

func (c *cgroupInode) StateFields() []string {
	return []string{
		"dir",
		"ts",
	}
}

func (c *cgroupInode) beforeSave() {}

// +checklocksignore
func (c *cgroupInode) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.dir)
	stateSinkObject.Save(1, &c.ts)
}

func (c *cgroupInode) afterLoad() {}

// +checklocksignore
func (c *cgroupInode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.dir)
	stateSourceObject.Load(1, &c.ts)
}

func (d *cgroupProcsData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cgroupProcsData"
}

func (d *cgroupProcsData) StateFields() []string {
	return []string{
		"cgroupInode",
	}
}

func (d *cgroupProcsData) beforeSave() {}

// +checklocksignore
func (d *cgroupProcsData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cgroupInode)
}

func (d *cgroupProcsData) afterLoad() {}

// +checklocksignore
func (d *cgroupProcsData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cgroupInode)
}

func (d *tasksData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.tasksData"
}

func (d *tasksData) StateFields() []string {
	return []string{
		"cgroupInode",
	}
}

func (d *tasksData) beforeSave() {}

// +checklocksignore
func (d *tasksData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cgroupInode)
}

func (d *tasksData) afterLoad() {}

// +checklocksignore
func (d *tasksData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cgroupInode)
}

func (fsType *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.FilesystemType"
}

func (fsType *FilesystemType) StateFields() []string {
	return []string{}
}

func (fsType *FilesystemType) beforeSave() {}

// +checklocksignore
func (fsType *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fsType.beforeSave()
}

func (fsType *FilesystemType) afterLoad() {}

// +checklocksignore
func (fsType *FilesystemType) StateLoad(stateSourceObject state.Source) {
}

func (i *InternalData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.InternalData"
}

func (i *InternalData) StateFields() []string {
	return []string{
		"DefaultControlValues",
		"InitialCgroupPath",
	}
}

func (i *InternalData) beforeSave() {}

// +checklocksignore
func (i *InternalData) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.DefaultControlValues)
	stateSinkObject.Save(1, &i.InitialCgroupPath)
}

func (i *InternalData) afterLoad() {}

// +checklocksignore
func (i *InternalData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.DefaultControlValues)
	stateSourceObject.Load(1, &i.InitialCgroupPath)
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"Filesystem",
		"devMinor",
		"hierarchyID",
		"controllers",
		"kcontrollers",
		"numCgroups",
		"root",
		"effectiveRoot",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.Filesystem)
	stateSinkObject.Save(1, &fs.devMinor)
	stateSinkObject.Save(2, &fs.hierarchyID)
	stateSinkObject.Save(3, &fs.controllers)
	stateSinkObject.Save(4, &fs.kcontrollers)
	stateSinkObject.Save(5, &fs.numCgroups)
	stateSinkObject.Save(6, &fs.root)
	stateSinkObject.Save(7, &fs.effectiveRoot)
}

func (fs *filesystem) afterLoad() {}

// +checklocksignore
func (fs *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.Filesystem)
	stateSourceObject.Load(1, &fs.devMinor)
	stateSourceObject.Load(2, &fs.hierarchyID)
	stateSourceObject.Load(3, &fs.controllers)
	stateSourceObject.Load(4, &fs.kcontrollers)
	stateSourceObject.Load(5, &fs.numCgroups)
	stateSourceObject.Load(6, &fs.root)
	stateSourceObject.Load(7, &fs.effectiveRoot)
}

func (i *implStatFS) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.implStatFS"
}

func (i *implStatFS) StateFields() []string {
	return []string{}
}

func (i *implStatFS) beforeSave() {}

// +checklocksignore
func (i *implStatFS) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
}

func (i *implStatFS) afterLoad() {}

// +checklocksignore
func (i *implStatFS) StateLoad(stateSourceObject state.Source) {
}

func (d *dir) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.dir"
}

func (d *dir) StateFields() []string {
	return []string{
		"InodeNoopRefCount",
		"InodeAlwaysValid",
		"InodeAttrs",
		"InodeNotSymlink",
		"InodeDirectoryNoNewChildren",
		"OrderedChildren",
		"implStatFS",
		"locks",
		"fs",
		"cgi",
	}
}

func (d *dir) beforeSave() {}

// +checklocksignore
func (d *dir) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.InodeNoopRefCount)
	stateSinkObject.Save(1, &d.InodeAlwaysValid)
	stateSinkObject.Save(2, &d.InodeAttrs)
	stateSinkObject.Save(3, &d.InodeNotSymlink)
	stateSinkObject.Save(4, &d.InodeDirectoryNoNewChildren)
	stateSinkObject.Save(5, &d.OrderedChildren)
	stateSinkObject.Save(6, &d.implStatFS)
	stateSinkObject.Save(7, &d.locks)
	stateSinkObject.Save(8, &d.fs)
	stateSinkObject.Save(9, &d.cgi)
}

func (d *dir) afterLoad() {}

// +checklocksignore
func (d *dir) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.InodeNoopRefCount)
	stateSourceObject.Load(1, &d.InodeAlwaysValid)
	stateSourceObject.Load(2, &d.InodeAttrs)
	stateSourceObject.Load(3, &d.InodeNotSymlink)
	stateSourceObject.Load(4, &d.InodeDirectoryNoNewChildren)
	stateSourceObject.Load(5, &d.OrderedChildren)
	stateSourceObject.Load(6, &d.implStatFS)
	stateSourceObject.Load(7, &d.locks)
	stateSourceObject.Load(8, &d.fs)
	stateSourceObject.Load(9, &d.cgi)
}

func (c *controllerFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.controllerFile"
}

func (c *controllerFile) StateFields() []string {
	return []string{
		"DynamicBytesFile",
	}
}

func (c *controllerFile) beforeSave() {}

// +checklocksignore
func (c *controllerFile) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.DynamicBytesFile)
}

func (c *controllerFile) afterLoad() {}

// +checklocksignore
func (c *controllerFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.DynamicBytesFile)
}

func (s *staticControllerFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.staticControllerFile"
}

func (s *staticControllerFile) StateFields() []string {
	return []string{
		"DynamicBytesFile",
		"StaticData",
	}
}

func (s *staticControllerFile) beforeSave() {}

// +checklocksignore
func (s *staticControllerFile) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.DynamicBytesFile)
	stateSinkObject.Save(1, &s.StaticData)
}

func (s *staticControllerFile) afterLoad() {}

// +checklocksignore
func (s *staticControllerFile) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.DynamicBytesFile)
	stateSourceObject.Load(1, &s.StaticData)
}

func (c *cpuController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuController"
}

func (c *cpuController) StateFields() []string {
	return []string{
		"controllerCommon",
		"cfsPeriod",
		"cfsQuota",
		"shares",
	}
}

func (c *cpuController) beforeSave() {}

// +checklocksignore
func (c *cpuController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.cfsPeriod)
	stateSinkObject.Save(2, &c.cfsQuota)
	stateSinkObject.Save(3, &c.shares)
}

func (c *cpuController) afterLoad() {}

// +checklocksignore
func (c *cpuController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.cfsPeriod)
	stateSourceObject.Load(2, &c.cfsQuota)
	stateSourceObject.Load(3, &c.shares)
}

func (c *cpuacctController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctController"
}

func (c *cpuacctController) StateFields() []string {
	return []string{
		"controllerCommon",
	}
}

func (c *cpuacctController) beforeSave() {}

// +checklocksignore
func (c *cpuacctController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
}

func (c *cpuacctController) afterLoad() {}

// +checklocksignore
func (c *cpuacctController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
}

func (c *cpuacctCgroup) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctCgroup"
}

func (c *cpuacctCgroup) StateFields() []string {
	return []string{
		"cgroupInode",
	}
}

func (c *cpuacctCgroup) beforeSave() {}

// +checklocksignore
func (c *cpuacctCgroup) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.cgroupInode)
}

func (c *cpuacctCgroup) afterLoad() {}

// +checklocksignore
func (c *cpuacctCgroup) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.cgroupInode)
}

func (d *cpuacctStatData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctStatData"
}

func (d *cpuacctStatData) StateFields() []string {
	return []string{
		"cpuacctCgroup",
	}
}

func (d *cpuacctStatData) beforeSave() {}

// +checklocksignore
func (d *cpuacctStatData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cpuacctCgroup)
}

func (d *cpuacctStatData) afterLoad() {}

// +checklocksignore
func (d *cpuacctStatData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctUsageData"
}

func (d *cpuacctUsageData) StateFields() []string {
	return []string{
		"cpuacctCgroup",
	}
}

func (d *cpuacctUsageData) beforeSave() {}

// +checklocksignore
func (d *cpuacctUsageData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageData) afterLoad() {}

// +checklocksignore
func (d *cpuacctUsageData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageUserData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctUsageUserData"
}

func (d *cpuacctUsageUserData) StateFields() []string {
	return []string{
		"cpuacctCgroup",
	}
}

func (d *cpuacctUsageUserData) beforeSave() {}

// +checklocksignore
func (d *cpuacctUsageUserData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageUserData) afterLoad() {}

// +checklocksignore
func (d *cpuacctUsageUserData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageSysData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpuacctUsageSysData"
}

func (d *cpuacctUsageSysData) StateFields() []string {
	return []string{
		"cpuacctCgroup",
	}
}

func (d *cpuacctUsageSysData) beforeSave() {}

// +checklocksignore
func (d *cpuacctUsageSysData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.cpuacctCgroup)
}

func (d *cpuacctUsageSysData) afterLoad() {}

// +checklocksignore
func (d *cpuacctUsageSysData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.cpuacctCgroup)
}

func (c *cpusetController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.cpusetController"
}

func (c *cpusetController) StateFields() []string {
	return []string{
		"controllerCommon",
	}
}

func (c *cpusetController) beforeSave() {}

// +checklocksignore
func (c *cpusetController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
}

func (c *cpusetController) afterLoad() {}

// +checklocksignore
func (c *cpusetController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
}

func (r *dirRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.dirRefs"
}

func (r *dirRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *dirRefs) beforeSave() {}

// +checklocksignore
func (r *dirRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *dirRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (c *jobController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.jobController"
}

func (c *jobController) StateFields() []string {
	return []string{
		"controllerCommon",
		"id",
	}
}

func (c *jobController) beforeSave() {}

// +checklocksignore
func (c *jobController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.id)
}

func (c *jobController) afterLoad() {}

// +checklocksignore
func (c *jobController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.id)
}

func (d *jobIDData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.jobIDData"
}

func (d *jobIDData) StateFields() []string {
	return []string{
		"c",
	}
}

func (d *jobIDData) beforeSave() {}

// +checklocksignore
func (d *jobIDData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.c)
}

func (d *jobIDData) afterLoad() {}

// +checklocksignore
func (d *jobIDData) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.c)
}

func (c *memoryController) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.memoryController"
}

func (c *memoryController) StateFields() []string {
	return []string{
		"controllerCommon",
		"limitBytes",
		"softLimitBytes",
		"moveChargeAtImmigrate",
	}
}

func (c *memoryController) beforeSave() {}

// +checklocksignore
func (c *memoryController) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.controllerCommon)
	stateSinkObject.Save(1, &c.limitBytes)
	stateSinkObject.Save(2, &c.softLimitBytes)
	stateSinkObject.Save(3, &c.moveChargeAtImmigrate)
}

func (c *memoryController) afterLoad() {}

// +checklocksignore
func (c *memoryController) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.controllerCommon)
	stateSourceObject.Load(1, &c.limitBytes)
	stateSourceObject.Load(2, &c.softLimitBytes)
	stateSourceObject.Load(3, &c.moveChargeAtImmigrate)
}

func (d *memoryUsageInBytesData) StateTypeName() string {
	return "pkg/sentry/fsimpl/cgroupfs.memoryUsageInBytesData"
}

func (d *memoryUsageInBytesData) StateFields() []string {
	return []string{}
}

func (d *memoryUsageInBytesData) beforeSave() {}

// +checklocksignore
func (d *memoryUsageInBytesData) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
}

func (d *memoryUsageInBytesData) afterLoad() {}

// +checklocksignore
func (d *memoryUsageInBytesData) StateLoad(stateSourceObject state.Source) {
}

func init() {
	state.Register((*controllerCommon)(nil))
	state.Register((*cgroupInode)(nil))
	state.Register((*cgroupProcsData)(nil))
	state.Register((*tasksData)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*InternalData)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*implStatFS)(nil))
	state.Register((*dir)(nil))
	state.Register((*controllerFile)(nil))
	state.Register((*staticControllerFile)(nil))
	state.Register((*cpuController)(nil))
	state.Register((*cpuacctController)(nil))
	state.Register((*cpuacctCgroup)(nil))
	state.Register((*cpuacctStatData)(nil))
	state.Register((*cpuacctUsageData)(nil))
	state.Register((*cpuacctUsageUserData)(nil))
	state.Register((*cpuacctUsageSysData)(nil))
	state.Register((*cpusetController)(nil))
	state.Register((*dirRefs)(nil))
	state.Register((*jobController)(nil))
	state.Register((*jobIDData)(nil))
	state.Register((*memoryController)(nil))
	state.Register((*memoryUsageInBytesData)(nil))
}
