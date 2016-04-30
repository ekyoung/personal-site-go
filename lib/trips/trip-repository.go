package trips

type TripRepository struct {
	allTrips []Trip
}

func NewTripRepository() *TripRepository {
	allTrips := make([]Trip, 20)

	t := Trip{
		Id:   "alaska-2005",
		Name: "Alaska 2005",
		Slides: SlidesInfo{
			IndexFileUrl: "http://images.ethanyoung.us/Alaska2005/Alaska2005Index.xml",
			RootUrl:      "http://images.ethanyoung.us/Alaska2005",
		},
	}
	allTrips[0] = t

	t = Trip{
		Id:   "southwest-2007",
		Name: "Southwest 2007",
		Slides: SlidesInfo{
			IndexFileUrl: "http://images.ethanyoung.us/Southwest2007/Southwest2007Index.xml",
			RootUrl:      "http://images.ethanyoung.us/Southwest2007",
		},
	}
	allTrips[1] = t

	return &TripRepository{
		allTrips: allTrips,
	}
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
