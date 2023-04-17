package extractor

import (
	"time"
)

type LiveStatus int8

const (
	NotLive LiveStatus = iota
	WasLive
	IsLive
)

func (ls LiveStatus) String() string {
	switch ls {
	case NotLive:
		return "not live"
	case WasLive:
		return "was live"
	case IsLive:
		return "is live"
	}
	return ""
}

type Availability int8

const (
	Public Availability = iota
	Unlisted
	NeedsAuth
)

func (a Availability) String() string {
	switch a {
	case Public:
		return "public"
	case Unlisted:
		return "unlisted"
	case NeedsAuth:
		return "needs auth"
	}
	return ""
}

type Info struct {
	Title                string
	Description          string
	Uploader             string
	UploaderURL          string
	UploaderID           string
	UploadDate           time.Time
	Duration             time.Duration
	Availability         Availability
	AgeLimit             int
	ViewCount            int
	LikeCount            int
	CommentCount         int
	ChannelFollowerCount int
	PlayableInEmbed      bool
	ThumbnailURL         string
	Tags                 []string
	Categories           []string
}

type SiteTest struct {
	URL    string
	Md5Sum [16]byte
	Info   Info
	Note   string
	Skip   string
	Ext    string
}

type Site struct {
	ValidURLRe  string
	Name        string
	Description string
	Tests       []SiteTest
	extractor   InfoExtractor
}

type InfoExtractor struct {
	ready bool
}
