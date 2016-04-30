package trips

type TripRepository struct {
	allTrips []Trip
}

func NewTripRepository() *TripRepository {
	allTrips := make([]Trip, 20)

	t := Trip{
		Id:   "alaska-2005",
		Name: "Alaska 2005",
	}
	allTrips[0] = t

	t = Trip{
		Id:   "southwest-2007",
		Name: "Southwest 2007",
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
