package craft

import (
	"os"

	"github.com/rtnl/blaze/pkg/client"
	"github.com/samber/mo"
)

type VersionData struct {
	Id        string               `json:"id"`
	Downloads VersionDataDownloads `json:"downloads"`
}

type VersionDataDownloads struct {
	Client *VersionDataDownloadEntry `json:"client"`
	Server *VersionDataDownloadEntry `json:"server"`
}

type VersionDataDownloadEntry struct {
	HashSha1 string `json:"sha1"`
	Size     int    `json:"size"`
	Url      string `json:"url"`
}

func (d *VersionData) GetClientDownload() (result mo.Option[VersionDataDownloadEntry]) {
	if d.Downloads.Client != nil {
		return mo.Some(*d.Downloads.Client)
	}

	return
}

func (d *VersionData) GetServerDownload() (result mo.Option[VersionDataDownloadEntry]) {
	if d.Downloads.Server != nil {
		return mo.Some(*d.Downloads.Server)
	}

	return
}

func (e *VersionDataDownloadEntry) Download(path string) (err error) {
	var (
		file *os.File
	)

	file, err = os.Create(path)
	if err != nil {
		return
	}

	err = client.HttpGetDownload(e.Url, file, &e.HashSha1)
	if err != nil {
		return
	}

	return
}
