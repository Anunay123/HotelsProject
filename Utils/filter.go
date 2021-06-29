package Utils

type ProcessedData struct {
	CityID     string
	HotelsData map[string]HotelNode
}

type HotelNode struct {
	Id           string
	Name         string
	HotelPropertyTag string
	Rating       float64
	Amenities  map[string]string
}

type Query struct {
	HotelID 	string
	HotelName 	string
	HotelPropertyTag  string
	Rating       float64
	Amenities    []string
}

func FilterData(query *Query, pd *ProcessedData) int {


	if len(query.HotelID) > 0 {
		return 1
	}

	var count int


	for _, value := range pd.HotelsData {

		if len(query.HotelName) == 0 || query.HotelName == value.Name {
			if len(query.HotelPropertyTag) == 0 || query.HotelPropertyTag == value.HotelPropertyTag {



				if value.Rating >= query.Rating {

					var flag = true

					for _, amenity := range query.Amenities {

						_, isPresent := value.Amenities[amenity]

						if !isPresent {
							flag = false
							break
						}
					}


					if flag {
						count += 1
					}

				}
			}
		}
	}



	return count
}

