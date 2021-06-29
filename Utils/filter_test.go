package Utils

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestFilterData(t *testing.T) {
	jsonFile, err := os.Open("../test.json")

	if err != nil {
		panic(err)
	}

	fmt.Println("File opened successfully")

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)


	byteValue, _ := ioutil.ReadAll(jsonFile)

	pd := ProcessedData{}
	errr := json.Unmarshal(byteValue, &pd)

	if errr != nil {
		panic(errr)
	}

	query1 := &Query{}
	query2 := &Query{
		HotelPropertyTag: "10",
		Rating: 2,
	}


	amenities := []string{"AC"}
	query3 := &Query{
		HotelPropertyTag: "10",
		Rating:           2,
		Amenities:        amenities,
	}


	count1 := FilterData(query1, &pd)

	count2 := FilterData(query2, &pd)

	count3 := FilterData(query3, &pd)

	//fmt.Println(count1, count2, count3)

	assert.Equal(t, 44,  count1)

	assert.Equal(t, 22, count2)
	assert.Equal(t, 9, count3)




}
