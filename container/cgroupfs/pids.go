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
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jpra1113/snap-plugin-collector-docker/container"
)

// Pids implements StatGetter interface
type Pids struct{}

// GetStats reads pids metrics from Pids Group
func (p *Pids) GetStats(stats *container.Statistics, opts container.GetStatOpt) error {
	path, err := opts.GetStringValue("cgroup_path")
	if err != nil {
		return err
	}

	pidsCurrentPath := filepath.Join(path, "pids.current")
	if _, err := os.Stat(pidsCurrentPath); os.IsNotExist(err) {
		return nil
	}

	current, err := parseIntValue(pidsCurrentPath)
	if err != nil {
		return err
	}

	pidsMaxPath := filepath.Join(path, "pids.max")
	if _, err := os.Stat(pidsMaxPath); os.IsNotExist(err) {
		return nil
	}

	limit, err := parseStrValue(pidsMaxPath)
	if err != nil {
		return err
	}

	stats.Cgroups.PidsStats.Current = current

	var max uint64
	if limit != "max" {
		max, err = strconv.ParseUint(strings.TrimSpace(string(limit)), 10, 64)
		if err != nil {
			return err
		}
	}
	stats.Cgroups.PidsStats.Limit = max

	return nil
}
