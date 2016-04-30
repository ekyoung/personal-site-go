package trips

type Trip struct {
	Id     string
	Name   string
	Slides SlidesInfo
}

type SlidesInfo struct {
	IndexFileUrl string
	RootUrl      string
}
