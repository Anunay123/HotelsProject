package main

import (
	"HotelsProject/Cache"
	middlewares "HotelsProject/MiddleWares"
	"HotelsProject/Repository"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)
import "fmt"




type Query struct {
	HotelID 	string
	HotelName 	string
	HotelPropertyTag  string
	Rating       float64
	Amenities    []string
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func getRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger())

	return r
}

func main() {

	setupLogOutput()
	r := getRouter()



	r.GET("/hotels/:cityID", getData())



	fmt.Println("Works")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")


}

func getData() gin.HandlerFunc {
	return func(c *gin.Context) {

		cityId := c.Param("cityID")
		hotelID := c.Query("hotelID")
		hotelName := c.Query("hotelName")
		propertyType := c.Query("type")
		r := c.Query("rating")
		ratings, _ := strconv.ParseFloat(r, 32)
		amenities := c.QueryArray("amenities")

		//fmt.Println(amenities, len(amenities))

		query := Query{
			HotelID:          hotelID,
			HotelName:        hotelName,
			HotelPropertyTag: propertyType,
			Rating:           ratings,
			Amenities:        amenities,
		}

		pd := &Repository.ProcessedData{}

		if PD, found := Cache.Get(cityId); !found {
			Repository.PopulateData(cityId, pd)

			Cache.Set(cityId, pd)
		} else {
			pd = PD.(*Repository.ProcessedData)
		}

		//fmt.Println(Repository.PropertyTagsMapping[propertyType])
		//for _, a := range amenities {
		//	fmt.Println(a, Repository.AmenitiesMapping[a])
		//}

		count := filterData(query, pd)

		c.JSON(200, gin.H{
			"status": "success",
			"count":   count,
		})

	}
}


func filterData(query Query, pd *Repository.ProcessedData) int {
	var count int

	if len(query.HotelID) > 0 {
		return 1
	}

	for _, value := range pd.HotelsData {

		if len(query.HotelName) == 0 || query.HotelName == value.Name {

			if len(query.HotelPropertyTag) == 0 || Repository.PropertyTagsMapping[query.HotelPropertyTag] == value.HotelPropertyTag {

				if value.Rating >= query.Rating {

					var flag bool = true

					for _, amenity := range query.Amenities {

						//temp := value.Amenities[Repository.AmenitiesMapping[amenity]]
						//fmt.Println(temp)

						_, isPresent := value.Amenities[Repository.AmenitiesMapping[amenity]]

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

