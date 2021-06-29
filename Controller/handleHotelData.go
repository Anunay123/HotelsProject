package Controller

import (
	"HotelsProject/Cache"
	"HotelsProject/Services"
	"HotelsProject/Utils"
	"github.com/gin-gonic/gin"
	"strconv"
)



func GetData() gin.HandlerFunc {
	return func(c *gin.Context) {



		cityId := c.Param("cityID")

		pd := &Utils.ProcessedData{}

		if PD, found := Cache.Get(cityId); !found {
			Services.PopulateData(cityId, pd)

			Cache.Set(cityId, pd)
			//writeToFile(pd)
		} else {
			pd = PD.(*Utils.ProcessedData)
		}


		hotelID := c.Query("hotelID")
		hotelName := c.Query("hotelName")
		propertyType := Services.PropertyTagsMapping[c.Query("type")]
		r := c.Query("rating")
		ratings, _ := strconv.ParseFloat(r, 32)
		amenities := c.QueryArray("amenities")

		for index, value := range amenities {
			amenities[index] = Services.AmenitiesMapping[value]
		}



		query := &Utils.Query{
			HotelID:          hotelID,
			HotelName:        hotelName,
			HotelPropertyTag: propertyType,
			Rating:           ratings,
			Amenities:        amenities,
		}






		var count = Utils.FilterData(query, pd)

		c.JSON(200, gin.H{
			"status": "success",
			"count":   count,
		})

	}
}

//func writeToFile(pd interface{})  {
//	file, _ := json.MarshalIndent(pd, "", " ")
//
//	_ = ioutil.WriteFile("test.json", file, 0644)
//
//}
