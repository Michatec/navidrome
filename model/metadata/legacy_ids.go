package metadata

import (
	"cmp"
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/navidrome/navidrome/conf"
	"github.com/navidrome/navidrome/consts"
	"github.com/navidrome/navidrome/model"
)

// These are the legacy ID functions that were used in the original Navidrome ID generation.
// They are kept here for backwards compatibility with existing databases.

func legacyTrackID(mf model.MediaFile) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(mf.Path)))
}

func legacyAlbumID(md Metadata) string {
	releaseDate := legacyReleaseDate(md)
	albumPath := strings.ToLower(fmt.Sprintf("%s\\%s", legacyMapAlbumArtistName(md), legacyMapAlbumName(md)))
	if !conf.Server.Scanner.GroupAlbumReleases {
		if len(releaseDate) != 0 {
			albumPath = fmt.Sprintf("%s\\%s", albumPath, releaseDate)
		}
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(albumPath)))
}

func legacyMapAlbumArtistName(md Metadata) string {
	values := []string{
		md.String(model.TagAlbumArtist),
		"",
		md.String(model.TagTrackArtist),
		consts.UnknownArtist,
	}
	if md.Bool(model.TagCompilation) {
		values[1] = consts.VariousArtists
	}
	return cmp.Or(values...)
}

func legacyMapAlbumName(md Metadata) string {
	return cmp.Or(
		md.String(model.TagAlbum),
		consts.UnknownAlbum,
	)
}

// Keep the TaggedLikePicard logic for backwards compatibility
func legacyReleaseDate(md Metadata) string {
	_, _, releaseDate := md.mapDates()
	return string(releaseDate)
}
