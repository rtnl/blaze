package craft

import (
	"fmt"

	"github.com/rtnl/blaze/pkg/client"
	"github.com/samber/mo"
)

const (
	UrlVersionManifest = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
)

type VersionManifest struct {
	Latest   VersionManifestLatest  `json:"latest"`
	Versions []VersionManifestEntry `json:"versions"`
}

type VersionManifestLatest struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

type VersionManifestEntry struct {
	Id          string      `json:"id"`
	Type        VersionKind `json:"type"`
	Url         string      `json:"url"`
	Time        string      `json:"time"`
	ReleaseTime string      `json:"releaseTime"`
}

func GetLauncherMetaVersionManifest() (m *VersionManifest, err error) {
	err = client.HttpGetJson(UrlVersionManifest, &m)
	if err != nil {
		return
	}

	return
}

func (m *VersionManifest) GetLatestVersion(kind VersionKind) string {
	switch kind {
	case VersionKindRelease:
		return m.Latest.Release

	case VersionKindSnapshot:
		return m.Latest.Snapshot

	default:
		panic("unimplemented")
	}
}

func (m *VersionManifest) GetVersion(versionId string) (result mo.Option[VersionManifestEntry]) {
	for _, version := range m.Versions {
		if version.Id == versionId {
			return mo.Some(version)
		}
	}

	return
}

func (m *VersionManifest) GetLatestVersionEntry(kind VersionKind) (result mo.Option[VersionManifestEntry]) {
	return m.GetVersion(m.GetLatestVersion(kind))
}

func (m *VersionManifest) DownloadVersionData(versionId string) (d *VersionData, err error) {
	var (
		version VersionManifestEntry
		ok      bool
	)

	version, ok = m.GetVersion(versionId).Get()
	if !ok {
		err = fmt.Errorf("version entry not found")
		return
	}

	err = client.HttpGetJson(version.Url, &d)
	if err != nil {
		return
	}

	return
}
