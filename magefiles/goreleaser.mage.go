package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
	"github.com/pterm/pterm"
	"github.com/sheldonhull/magetools/pkg/magetoolsutils"
	"github.com/sheldonhull/magetools/pkg/req"
)

func checkEnvVar(envVar string, required bool) (string, error) { //nolint:unparam // leaving this as return string value for now
	envVarValue := os.Getenv(envVar)
	if envVarValue == "" && required {
		pterm.Error.Printfln(
			"%s is required and unable to proceed without this being provided. terminating task.",
			envVar,
		)
		return "", fmt.Errorf("%s is required", envVar)
	}
	if envVarValue == "" {
		pterm.Debug.Printfln(
			"checkEnvVar() found no value for: %q, however this is marked as optional, so not exiting task",
			envVar,
		)
	}
	pterm.Debug.Printfln("checkEnvVar() found value: %q=%q", envVar, envVarValue)
	return envVarValue, nil
}

// 🔨 Build builds the project for the current platform.
func Build() error {
	magetoolsutils.CheckPtermDebug()
	binary, err := req.ResolveBinaryByInstall("goreleaser", "github.com/goreleaser/goreleaser@latest")
	if err != nil {
		return err
	}

	releaserArgs := []string{
		"build",
		"--rm-dist",
		"--snapshot",
		"--single-target",
	}
	pterm.Debug.Printfln("goreleaser: %+v", releaserArgs)

	return sh.RunV(binary, releaserArgs...) // "--skip-announce",.
}

// 🔨 BuildAll builds all the binaries defined in the project, for all platforms.
// If there is no additional platforms configured in the task, then basically this will just be the same as `mage build`.
func BuildAll() error {
	magetoolsutils.CheckPtermDebug()
	binary, err := req.ResolveBinaryByInstall("goreleaser", "github.com/goreleaser/goreleaser@latest")
	if err != nil {
		return err
	}

	return sh.RunV(binary,
		"build",
		"--rm-dist",
		"--snapshot",
	)
}

// 🔨 Release generates a release and validates the required environment variables are available.
func Release() error {
	magetoolsutils.CheckPtermDebug()
	binary, err := req.ResolveBinaryByInstall("goreleaser", "github.com/goreleaser/goreleaser@latest")
	if err != nil {
		return err
	}

	if _, err = checkEnvVar("GITHUB_TOKEN", true); err != nil {
		return err
	}
	if _, err = checkEnvVar("GPG_FINGERPRINT", true); err != nil {
		return err
	}

	releaserArgs := []string{
		"release",
		"--rm-dist",
	}
	pterm.Debug.Printfln("goreleaser: %+v", releaserArgs)

	return sh.RunV(binary, releaserArgs...)
}
