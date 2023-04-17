package cookies

type browserType int8

const (
	chromiumBased browserType = iota
	firefox
	safari
)

const AllBrowsers = "Brave, Chrome, Chromium, Edge, Firefox, Opera, Safari, Vivaldi"

func getBrowserType(name string) browserType {
	switch name {
	case "firefox":
		return firefox
	case "safari":
		return safari
	case "chrome", "chromium", "edge", "opera", "vivaldi":
		return chromiumBased
	}
	return -1
}
