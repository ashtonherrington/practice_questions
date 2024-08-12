package practice_question

// Solution example from https://medium.com/@nrpotter15/parsing-xml-in-go-14ee87cfa99e)

import (
	"encoding/xml"
	"fmt"
)

type Campground struct {
	Type    string `xml:"type,attr"`
	Name    string `xml:"name"`
	Address struct {
		Street string `xml:"street"`
		City   string `xml:"city"`
		State  string `xml:"state"`
		Zip    int    `xml:"zip"`
	} `xml:"address"`
	Amenities []struct {
		Distance string `xml:"distance"`
		Name     string `xml:"name"`
	} `xml:"amenities>amenity"`
}

func main() {

	input := `<campground type="PUBLIC">
    <name>Pedernales Falls State Park</name>
    <address>
        <street>2585 Park Rd 6026</street>
        <city>Johnson City</city>
        <state>TX</state>
        <zip>78636</zip>
    </address>
    <amenities>
        <amenity>
            <distance>Within Facility</distance>
            <name>Biking</name>
        </amenity>
        <amenity>
            <distance>Within Facility</distance>
            <name>Kayaking</name>
        </amenity>
    </amenities>
</campground>`

	camp := Campground{}
	xml.Unmarshal([]byte(input), &camp)
	fmt.Println(input)
	result, err := xml.Marshal(camp)
	if err != nil {
		fmt.Println(fmt.Errorf("marshaling object into XML: %w", err))
	}
	fmt.Println(string(result))
}
