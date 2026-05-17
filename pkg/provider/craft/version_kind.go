package craft

type VersionKind string

const (
	VersionKindRelease  VersionKind = "release"
	VersionKindSnapshot VersionKind = "snapshot"
)
