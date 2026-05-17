package main

import (
	"github.com/rtnl/blaze/pkg/provider/craft"
	"github.com/samber/mo"
)

const (
	VersionQueryLatest          = "latest"
	VersionDownloadTargetServer = "server"
	VersionDownloadTargetClient = "client"
)

type Logic struct {
}

func (l *Logic) GetVersionDownloadUrl(versionQuery string, versionKind craft.VersionKind, target string) mo.Result[string] {
	var (
		manifest     *craft.VersionManifest
		versionEntry craft.VersionManifestEntry
		version      *craft.VersionData
		download     craft.VersionDataDownloadEntry
		ok           bool
		err          error
	)

	manifest, err = craft.GetLauncherMetaVersionManifest()
	if err != nil {
		return mo.Errf[string]("failed at getting version manifest")
	}

	if versionQuery == VersionQueryLatest {
		versionEntry, ok = manifest.GetLatestVersionEntry(versionKind).Get()
	} else {
		versionEntry, ok = manifest.GetVersion(versionQuery).Get()
	}
	if !ok {
		return mo.Errf[string]("version manifest not found")
	}

	version, err = manifest.DownloadVersionData(versionEntry.Id)
	if err != nil {
		return mo.Errf[string]("failed at getting version data")
	}

	switch target {
	case VersionDownloadTargetServer:
		{
			download, ok = version.GetServerDownload().Get()
			break
		}

	case VersionDownloadTargetClient:
		{
			download, ok = version.GetClientDownload().Get()
			break
		}

	default:
		panic("unimplemented")
	}
	if !ok {
		return mo.Errf[string]("download info not found")
	}

	return mo.Ok(download.Url)
}
