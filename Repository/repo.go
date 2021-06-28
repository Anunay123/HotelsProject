package Repository

import (
	"awesomeProject/Cache"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Hotel struct {
	Data struct {
		Meta struct {
			CityData struct {
				LocusId string `json:"locus_id"`
				Id      string `json:"_id"`
				T       int    `json:"t"`
				N       string `json:"n"`
			} `json:"city_data"`
			Locations     []interface{} `json:"locations"`
			AmenitiesData map[string]struct {
					Dn  string `json:"dn"`
					Nh  int    `json:"nh"`
					Img string `json:"img"`
			} `json:"amenities_data"`
			Collections struct {
			} `json:"collections"`
			PropertyTypesData map[string]struct {
				Nh      int    `json:"nh"`
				AltAcco bool   `json:"alt_acco"`
				Txt     string `json:"txt"`
				Img     string `json:"img"`
			} `json:"property_types_data"`
			HouseRulesConfig struct {
				Rules struct {
					Field1 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"11"`
					Field2 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"25"`
					Field3 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"39"`
					Field4 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"38"`
					Field5 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"14"`
					Field6 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"55"`
					Field7 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"54"`
					Field8 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"57"`
					Field9 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"30"`
					Field10 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"45"`
					Field11 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"52"`
					Field12 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"5"`
					Field13 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"4"`
					Field14 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"9"`
					Field15 struct {
						Text        string `json:"text"`
						TagCategory string `json:"tag_category"`
						RespType    string `json:"resp_type"`
						Sentiment   string `json:"sentiment"`
						STxt        string `json:"s_txt"`
					} `json:"56"`
				} `json:"rules"`
				Response struct {
					Boolean struct {
						Field1 struct {
							Text string `json:"text"`
						} `json:"1"`
						Field2 struct {
							Text string `json:"text"`
						} `json:"0"`
					} `json:"boolean"`
				} `json:"response"`
			} `json:"house_rules_config"`
		} `json:"meta"`
		HotelsData map[string]struct {
			HotelGeoNodeData HotelGeoNode`json:"hotel_geo_node"`
			HotelNodeData HotelDataNode `json:"hotel_data_node"`
		} `json:"hotels_data"`
	} `json:"data"`
	Success bool `json:"success"`
}

type HotelGeoNode struct {
	Name string `json:"name"`
	Tags struct {
		HotelChainCode   []string `json:"hotel_chain_code"`
		PropertyType     []string `json:"property_type"`
		HotelListingType []string `json:"hotel_listing_type"`
		Others           []string `json:"others"`
	} `json:"tags"`
	RawTags  []interface{} `json:"raw_tags"`
	Location struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	} `json:"location"`
	Id   string `json:"_id"`
	Type int    `json:"type"`
}
type HotelDataNode struct {
	ImgsSrp []struct {
		MJpg   string `json:"m_jpg"`
		SWebp  string `json:"s_webp"`
		Vendor string `json:"vendor"`
		MWebp  string `json:"m_webp"`
		SJpg   string `json:"s_jpg"`
		T      string `json:"t"`
	} `json:"imgs_srp"`
	Loc struct {
		City           string        `json:"city"`
		CntCode        string        `json:"cnt_code"`
		LocusDistricts []interface{} `json:"locus_districts"`
		CityCids       struct {
			Gds   string `json:"gds"`
			Voy   string `json:"voy"`
			Locus string `json:"locus"`
		} `json:"city_cids"`
		Entities   []string      `json:"entities"`
		Country    string        `json:"country"`
		MapLink    string        `json:"map_link"`
		Long       float64       `json:"long"`
		LocusZones []interface{} `json:"locus_zones"`
		State      string        `json:"state"`
		SecLocs    []interface{} `json:"sec_locs"`
		Full       string        `json:"full"`
		Location   string        `json:"location"`
		Pin        string        `json:"pin"`
		Lat        float64       `json:"lat"`
		AltLoc     struct {
		} `json:"alt_loc"`
		Directions   string        `json:"directions"`
		Nhood        []interface{} `json:"nhood"`
		LocusRegions []interface{} `json:"locus_regions"`
	} `json:"loc"`
	ImgProcessed []interface{} `json:"img_processed"`
	Gosafe       int           `json:"gosafe"`
	Name         string        `json:"name"`
	Img          []interface{} `json:"img"`
	Rating       int           `json:"rating"`
	Extra        struct {
		ParsedAttr struct {
			CoupleFriendly int    `json:"couple_friendly"`
			CheckInTime    string `json:"check_in_time"`
			CheckOutTime   string `json:"check_out_time"`
			LocalId        int    `json:"local_id"`
		} `json:"parsed_attr"`
		PolicyInfo struct {
			HouseRuleFilters []interface{} `json:"house_rule_filters"`
			HouseRules       struct {
			} `json:"house_rules"`
		} `json:"policy_info"`
		OrigVendorData struct {
			Default struct {
				CheckIn  string `json:"check_in"`
				CheckOut string `json:"check_out"`
			} `json:"default"`
			Gds struct {
				PersonalizedPrc int    `json:"personalized_prc"`
				CityId          string `json:"city_id"`
				PromoApply      int    `json:"promo_apply"`
				Currency        string `json:"currency"`
				Invoice         int    `json:"invoice"`
				HostelType      string `json:"hostel_type"`
				AltAcco         struct {
					PrebookChatEnabled bool          `json:"prebook_chat_enabled"`
					Hosts              []interface{} `json:"hosts"`
					IsAltAccoProp      bool          `json:"is_alt_acco_prop"`
				} `json:"alt_acco"`
				CleanStay             int           `json:"clean_stay"`
				CheckIn               string        `json:"check_in"`
				BookingAutoConfirm    int           `json:"booking_auto_confirm"`
				Pah                   int           `json:"pah"`
				AutoCheckin           int           `json:"auto_checkin"`
				TaDiscount            int           `json:"ta_discount"`
				PahOnly               int           `json:"pah_only"`
				SpecialTag            []interface{} `json:"special_tag"`
				VendorCode            string        `json:"vendor_code"`
				Pahcc                 int           `json:"pahcc"`
				RelaxationFactor      int           `json:"relaxation_factor"`
				TaxPah                int           `json:"tax_pah"`
				CheckInSplInstruction string        `json:"check_in_spl_instruction"`
				CityName              string        `json:"city_name"`
				ExpressCheckin        int           `json:"express_checkin"`
				CheckOut              string        `json:"check_out"`
				GstnAssured           int           `json:"gstn_assured"`
				LocalId               int           `json:"local_id"`
				ContentOnly           int           `json:"content_only"`
				Mapped                struct {
					Bkg []string `json:"bkg"`
				} `json:"mapped"`
				SlotBooking    int `json:"slot_booking"`
				CoupleFriendly int `json:"couple_friendly"`
				ChatEnabled    int `json:"chat_enabled"`
			} `json:"gds"`
		} `json:"orig_vendor_data"`
	} `json:"extra"`
	RoomInfo []struct {
		Vendor       string   `json:"vendor"`
		Roomsize     string   `json:"roomsize"`
		TypeName     string   `json:"type_name"`
		TypeCodes    []string `json:"type_codes"`
		RoomTypesMap struct {
			BKG []string `json:"BKG"`
			DB  []string `json:"DB"`
		} `json:"room_types_map"`
		St           string `json:"st"`
		MaxChild     int    `json:"max_child"`
		MaxOcc       int    `json:"max_occ"`
		IsSubroom    int    `json:"is_subroom"`
		TypeCode     string `json:"type_code"`
		MaxGuest     int    `json:"max_guest"`
		SellableType string `json:"sellable_type"`
	} `json:"room_info"`
	Ids struct {
		Gds       []string `json:"gds"`
		Chain     string   `json:"chain"`
		ChainCode string   `json:"chain_code"`
	} `json:"ids"`
	St         string `json:"st"`
	Facilities struct {
		NotIncluded []string `json:"not_included"`
		All         []string `json:"all"`
		Ocmapped    []struct {
			Code string `json:"code"`
			N    string `json:"n"`
		} `json:"ocmapped"`
		Cmapped struct {
			BNQTHALL struct {
				Dn string `json:"dn"`
			} `json:"BNQTHALL"`
			FRINTRNT struct {
				Dn string `json:"dn"`
			} `json:"FRINTRNT"`
			AC struct {
				Dn string `json:"dn"`
			} `json:"AC"`
			OACTS struct {
				Dn string `json:"dn"`
			} `json:"OACTS"`
			BONFIRE struct {
				Dn string `json:"dn"`
			} `json:"BONFIRE"`
		} `json:"cmapped"`
		Categorized struct {
			Services []string `json:"Services"`
			Internet []string `json:"Internet"`
			General  []string `json:"General"`
			Room     []string `json:"Room"`
			Others   []string `json:"Others"`
		} `json:"categorized"`
		Mapped []string `json:"mapped"`
	} `json:"facilities"`
	Amenities struct {
		Cmapped map[string]struct {
				Dn string `json:"dn"`
		} `json:"cmapped"`
	} `json:"amenities"`
	ImgSelected struct {
		Fs struct {
			Wpl    string `json:"wpl"`
			Vendor string `json:"vendor"`
			L      string `json:"l"`
		} `json:"fs"`
		Thumb struct {
			Wpl    string `json:"wpl"`
			Vendor string `json:"vendor"`
			L      string `json:"l"`
		} `json:"thumb"`
		G struct {
			Wpl    string `json:"wpl"`
			Vendor string `json:"vendor"`
			L      string `json:"l"`
		} `json:"g"`
		Srp struct {
			Wpl    string `json:"wpl"`
			Vendor string `json:"vendor"`
			L      string `json:"l"`
		} `json:"srp"`
		R struct {
			Wpl    string `json:"wpl"`
			Vendor string `json:"vendor"`
			L      string `json:"l"`
		} `json:"r"`
		Th struct {
			Wpl    string `json:"wpl"`
			Vendor string `json:"vendor"`
			L      string `json:"l"`
		} `json:"th"`
	} `json:"img_selected"`
	HotelPropertyTag string `json:"hotel_property_tag"`
	AltAcco          struct {
		IsAltAccoProp      bool          `json:"is_alt_acco_prop"`
		Hosts              []interface{} `json:"hosts"`
		PrebookChatEnabled bool          `json:"prebook_chat_enabled"`
	} `json:"alt_acco"`
	Id            string `json:"_id"`
	SafetyHygiene int    `json:"safety_hygiene"`
	Desc          struct {
		Default string `json:"default"`
	} `json:"desc"`
}

type ProcessedData struct {
	cityID string
	HotelsData map[string]HotelNode
}

type HotelNode struct {
	Id           string
	Name         string
	HotelPropertyTag string
	Rating       float64
	Amenities  map[string]string
}

var AmenitiesMapping = make(map[string]string)
var PropertyTagsMapping  = make(map[string]string)

func PopulateData(cityId string, pd *ProcessedData) {

	url := "https://voyager.goibibo.com/api/v1/hotels/get_hotels_data_by_city/?params={\"city_id\":" + "\"" + cityId + "\"" + "," + "\"mode\":\"search_page_extended_v2\"}"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("err", err)
	}

	data, _ := ioutil.ReadAll( res.Body )

	var hotel Hotel


	errr := json.Unmarshal(data, &hotel)


	if errr != nil {
		fmt.Println("error 2", errr)
	}


	if _, found := Cache.Get("amenities"); !found {
		for key, value := range hotel.Data.Meta.AmenitiesData {
			//fmt.Println(value.Dn, key)
			AmenitiesMapping[value.Dn] = key
		}

		Cache.Set("amenities", AmenitiesMapping)
	}

	if _, found := Cache.Get("propertyTypes"); !found {
		for key, value := range hotel.Data.Meta.PropertyTypesData {
			PropertyTagsMapping[value.Txt] = key
		}

		Cache.Set("propertyTypes", PropertyTagsMapping)
	}

	pd.cityID = cityId
	pd.HotelsData = make(map[string]HotelNode)

	for key, value := range hotel.Data.HotelsData {
		var amenities = make(map[string]string)

		for key2, value2 := range value.HotelNodeData.Amenities.Cmapped {
			amenities[key2] = value2.Dn
		}

		var temp HotelNode
		temp = HotelNode{
			Id:               value.HotelNodeData.Id,
			Name:             value.HotelNodeData.Name,
			HotelPropertyTag: value.HotelNodeData.HotelPropertyTag,
			Rating: float64(value.HotelNodeData.Rating),
			Amenities:        amenities,
		}

		pd.HotelsData[key] = temp
	}





}





