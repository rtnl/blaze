package craft

import (
	"strings"

	"github.com/samber/mo"
)

type VersionKind string

const (
	VersionKindRelease  VersionKind = "release"
	VersionKindSnapshot VersionKind = "snapshot"
)

func VersionKindFromString(input string) mo.Option[VersionKind] {
	switch strings.ToLower(input) {
	case string(VersionKindRelease):
		return mo.Some(VersionKindRelease)

	case string(VersionKindSnapshot):
		return mo.Some(VersionKindSnapshot)

	default:
		return mo.None[VersionKind]()
	}
}
