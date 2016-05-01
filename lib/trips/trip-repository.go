package trips

import (
    "html/template"
)

type TripRepository struct {
    allTrips []Trip
}

func NewTripRepository() *TripRepository {
    allTrips := []Trip{}

    allTrips = append(allTrips, Trip{
        Id:   "alaska-2005",
        Name: "Alaska 2005",
        Image: Image{
            FileName: "arctic-circle.jpg",
            AltText:  "My motorcycle at the Arctic Circle at midnight.",
        },
        Description: template.HTML(`This motorcycle trip took me 25,000 miles over about 10 weeks. I visited the Southernmost, Easternmost, Northernmost, and Westernmost points
                     in the lower 48 states, and the Northernmost and Westernmost road accessible points in North America. In addition to these bragging rights
                     points, I stopped at a lot of great sights along the way. I've included some pictures of the (better than expected) road conditions up North
                     for reference by other riders thinking of making the trip.`),
        Slides: SlidesInfo{
            IndexFileUrl: "http://images.ethanyoung.us/Alaska2005/Alaska2005Index.xml",
            RootUrl:      "http://images.ethanyoung.us/Alaska2005",
        },
    })

    allTrips = append(allTrips, Trip{
        Id:   "southwest-2007",
        Name: "Southwest 2007",
        Image: Image{
            FileName: "jumbo-rocks-close-up-2.jpg",
            AltText:  "A rock formation in the Jumbo Rocks campground of Joshua Tree National Park at sunset.",
        },
        Description: template.HTML(`This trip took me through the three national parks in the lower 48 that I hadn't yet been to: Channel Islands, Joshua Tree, and Saguaro.`),
        Slides: SlidesInfo{
            IndexFileUrl: "http://images.ethanyoung.us/Southwest2007/Southwest2007Index.xml",
            RootUrl:      "http://images.ethanyoung.us/Southwest2007",
        },
    })

    allTrips = append(allTrips, Trip{
        Id:   "mexico-2013",
        Name: "Mexico 2013",
        Image: Image{
            FileName: "iguana-on-our-heads.jpg",
            AltText:  "Ethan and Sarah with the Royal Playa del Carmen's pet iguana on their heads.",
        },
        Description: template.HTML(`Sarah and I visited Playa del Carmen, Mexico to attend her friend Anne's wedding. While there we explored the cenote at Rio Secreto, Mayan
                     ruins in Coba and Tulum, and went snorkeling at Gran Arrecife Maya.`),
        Slides: SlidesInfo{
            IndexFileUrl: "http://images.ethanyoung.us/Mexico2013/Mexico2013Index.xml",
            RootUrl:      "http://images.ethanyoung.us/Mexico2013",
        },
    })

    allTrips = append(allTrips, Trip{
        Id:   "moab-may-2006",
        Name: "Moab May 2006",
        Image: Image{
            FileName: "tower-arch-2.jpg",
            AltText:  "Tower Arch in Arches National Park",
        },
        Description: template.HTML(`On this trip we visited Canyonlands and Arches National Parks. We toured around the Needles area of Canyonland on my motorcycle and saw the
                     few sights that can be reached by paved roads and short hikes. To see more requires travelling on moderate and difficult 4x4 trails. In Arches
                     we met up with some friends who were in town for a Nissan Xterra off road convention. After seeing a few of the popular sights from the paved
                     road we all got in the 4x4 vehicles and headed out to Tower Arch. This was one of the best sights and it was nearly deserted because of it's
                     slightly more remote location. I can't wait to go back and explore more of the world famous 4x4 trails in the area.`),
        Slides: SlidesInfo{
            IndexFileUrl: "http://images.ethanyoung.us/MoabMay2006/MoabMay2006Index.xml",
            RootUrl:      "http://images.ethanyoung.us/MoabMay2006",
        },
    })

    allTrips = append(allTrips, Trip{
        Id:   "uncle-bud-hut-2007",
        Name: "Uncle Bud Hut 2007",
        Image: Image{
            FileName: "galena-mountain.jpg",
            AltText:  "Galena Mountain",
        },
        Description: template.HTML(`This backcountry hut is located at 11,380ft, near the Continental Divide, west of Leadville Colorado, on the shores of Turquoise Lake.
                     You can find more information about this and other huts at the <a href="http://www.huts.org" target="_blank">10th Mountain Division</a>
                     web site. There are also several guide books published about the many hut systems in Colorado.`),
        Slides: SlidesInfo{
            IndexFileUrl: "http://images.ethanyoung.us/UncleBudHut2007/UncleBudHut2007Index.xml",
            RootUrl:      "http://images.ethanyoung.us/UncleBudHut2007",
        },
    })

    allTrips = append(allTrips, Trip{
        Id:   "betty-bear-hut-2006",
        Name: "Betty Bear Hut 2006",
        Image: Image{
            FileName: "betty-bear-hut.jpg",
            AltText:  "Betty Bear Hut",
        },
        Description: template.HTML(`This backcountry hut was built by the 10th Mountain Division Hut Association in 1991. It is located at 11,100ft, near the Continental
                     Divide, east along the Fryingpan River from Basalt Colorado. You can find more information about this and other huts at the
                     <a href="http://www.huts.org" target="_blank">10th Mountain Division</a> web site. There are also several guide books
                     published about the many hut systems in Colorado.`),
        Slides: SlidesInfo{
            IndexFileUrl: "http://images.ethanyoung.us/BettyBearHut2006/BettyBearHut2006Index.xml",
            RootUrl:      "http://images.ethanyoung.us/BettyBearHut2006",
        },
    })

    allTrips = append(allTrips, Trip{
        Id:   "bingham-canyon-mine",
        Name: "Bingham Canyon Mine",
        Image: Image{
            FileName: "truck.jpg",
            AltText:  "Big Truck",
        },
        Description: template.HTML(`This open pit copper mine is the largest man-made excavation in the world. In the roughly 100 years since it has been
                     in operation, over 6 billion tons of material has been removed. For more information (and some better pictures) visit the
                     <a href="http://www.kennecott.com/" target="_blank">Kennecott Home Page</a>.`),
        Slides: SlidesInfo{
            IndexFileUrl: "http://images.ethanyoung.us/BinghamCanyonMine/Mine2005Index.xml",
            RootUrl:      "http://images.ethanyoung.us/BinghamCanyonMine",
        },
    })

    allTrips = append(allTrips, Trip{
        Id:   "dinosaur-nm-2005",
        Name: "Dinosaur NM 2005",
        Image: Image{
            FileName: "cliff-and-snow.jpg",
            AltText:  "Cliff And Snow",
        },
        Description: template.HTML(`This park in the Uinta Basin is home to a dinosaur quarry where thousands of fossils
                     have been discovered. Most were removed and placed in museums around the world,
                     but around 2,000 remain encased in stone so that tourists can see them the way they
                     were discovered. Scientists think the bones are located here because dinosaur carcasses
                     were washed down river until they stopped on a sand bar and were burried with more
                     sand and silt to become fossilized.`),
        Slides: SlidesInfo{
            IndexFileUrl: "http://images.ethanyoung.us/DinosaurNM2005/DinosaurNM2005Index.xml",
            RootUrl:      "http://images.ethanyoung.us/DinosaurNM2005",
        },
    })

    return &TripRepository{
        allTrips: allTrips,
    }
}

func (repo *TripRepository) All() []Trip {
    return repo.allTrips
}

func (repo *TripRepository) Lookup(id string) Trip {
    for _, value := range repo.allTrips {
        if value.Id == id {
            return value
        }
    }

    var result Trip
    return result
}
