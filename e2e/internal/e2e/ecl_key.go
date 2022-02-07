// Copyright (c) Contributors to the Apptainer project, established as
//   Apptainer a Series of LF Projects LLC.
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2021, Sylabs Inc. All rights reserved.
// Copyright (c) 2020, Control Command Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package e2e

import (
	"path/filepath"
	"testing"

	"github.com/apptainer/apptainer/internal/pkg/buildcfg"
	"github.com/apptainer/apptainer/internal/pkg/syecl"
	"github.com/apptainer/apptainer/internal/pkg/util/fs"
	"github.com/apptainer/apptainer/pkg/sypgp"
	"golang.org/x/sys/unix"
)

func SetupSystemECLAndGlobalKeyRing(t *testing.T, testDir string) {
	Privileged(func(t *testing.T) {
		dest := buildcfg.ECL_FILE
		source := filepath.Join(testDir, filepath.Base(dest))

		if err := syecl.PutConfig(syecl.EclConfig{}, source); err != nil {
			t.Fatalf("while generating ECL configuration: %s", err)
		}
		if err := unix.Mount(source, dest, "", unix.MS_BIND, ""); err != nil {
			t.Fatalf("while mounting %s to %s: %s", source, dest, err)
		}

		handle := sypgp.NewHandle(buildcfg.APPTAINER_CONFDIR, sypgp.GlobalHandleOpt())
		dest = handle.PublicPath()
		source = filepath.Join(testDir, filepath.Base(dest))

		if err := fs.Touch(source); err != nil {
			t.Fatalf("while creating %s: %s", source, err)
		}
		if err := unix.Mount(source, dest, "", unix.MS_BIND, ""); err != nil {
			t.Fatalf("while mounting %s to %s: %s", source, dest, err)
		}
	})(t)
}
