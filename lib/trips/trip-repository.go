package trips

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
		Description: `This motorcycle trip took me 25,000 miles over about 10 weeks. I visited the Southernmost, Easternmost, Northernmost, and Westernmost points
                    in the lower 48 states, and the Northernmost and Westernmost road accessible points in North America. In addition to these bragging rights
                    points, I stopped at a lot of great sights along the way. I've included some pictures of the (better than expected) road conditions up North
                    for reference by other riders thinking of making the trip.`,
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
		Description: `This trip took me through the three national parks in the lower 48 that I hadn't yet been to: Channel Islands, Joshua Tree, and Saguaro.`,
		Slides: SlidesInfo{
			IndexFileUrl: "http://images.ethanyoung.us/Southwest2007/Southwest2007Index.xml",
			RootUrl:      "http://images.ethanyoung.us/Southwest2007",
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
