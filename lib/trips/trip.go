package trips

import (
    "html/template"
)

type Trip struct {
    Id          string
    Name        string
    Image       Image
    Description template.HTML
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
