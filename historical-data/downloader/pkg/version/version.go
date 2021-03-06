package version

import "github.com/coreos/go-semver/semver"

var (
	// VersionMajor is for an API incompatible changes.
	VersionMajor int64 = 0
	// VersionMinor is for functionality in a backwards-compatible manner.
	VersionMinor int64 = 0
	// VersionPatch is for backwards-compatible bug fixes.
	VersionPatch int64 = 1
	// VersionPre indicates pre release.
	VersionPre string
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string = "dev"
)

// Version is the specification version that the package types support.
var Version = semver.Version{
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,
	PreRelease: semver.PreRelease(VersionPre),
	Metadata:   VersionDev,
}
