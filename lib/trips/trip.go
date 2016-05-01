package trips

type Trip struct {
	Id          string
	Name        string
	Image       Image
	Description string
	Slides      SlidesInfo
}

type Image struct {
	FileName string
	AltText  string
}

type SlidesInfo struct {
	IndexFileUrl string
	RootUrl      string
}
