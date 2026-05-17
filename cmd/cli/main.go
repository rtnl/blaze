package main

import (
	"github.com/rtnl/blaze/pkg/provider/craft"
)

func main() {
	var (
		err error
	)

	manifest, err := craft.GetLauncherMetaVersionManifest()
	if err != nil {
		panic(err)
	}

	latest, ok := manifest.GetLatestVersionEntry(craft.VersionKindRelease).Get()
	if !ok {
		panic("latest not found")
	}

	version, err := manifest.DownloadVersionData(latest.Id)
	if err != nil {
		panic(err)
	}

	versionServerDownload, ok := version.GetServerDownload().Get()
	if !ok {
		panic("version server download not found")
	}

	err = versionServerDownload.Download("./server.jar")
	if err != nil {
		panic(err)
	}

}
