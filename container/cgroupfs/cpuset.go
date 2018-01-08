// +build linux

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cgroupfs

import (
	"path/filepath"

	"github.com/jpra1113/snap-plugin-collector-docker/container"
)

// CpuSet implements StatGetter interface
type CpuSet struct{}

// GetStats reads cpuset metrics from Cpuset Group
func (cs *CpuSet) GetStats(stats *container.Statistics, opts container.GetStatOpt) error {
	path, err := opts.GetStringValue("cgroup_path")
	if err != nil {
		return err
	}

	cpus, err := parseStrValue(filepath.Join(path, "cpuset.cpus"))
	if err != nil {
		return err
	}

	mems, err := parseStrValue(filepath.Join(path, "cpuset.mems"))
	if err != nil {
		return err
	}

	memmig, err := parseIntValue(filepath.Join(path, "cpuset.memory_migrate"))
	if err != nil {
		return err
	}

	cpuexc, err := parseIntValue(filepath.Join(path, "cpuset.cpu_exclusive"))
	if err != nil {
		return err
	}

	memexc, err := parseIntValue(filepath.Join(path, "cpuset.mem_exclusive"))
	if err != nil {
		return err
	}

	stats.Cgroups.CpuSetStats.Cpus = cpus
	stats.Cgroups.CpuSetStats.Mems = mems
	stats.Cgroups.CpuSetStats.MemoryMigrate = memmig
	stats.Cgroups.CpuSetStats.CpuExclusive = cpuexc
	stats.Cgroups.CpuSetStats.MemoryExclusive = memexc

	return nil
}
