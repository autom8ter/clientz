package ip2location

import (
	"github.com/ip2location/ip2location-go"
)

func New() {
	ip2location.Open("./IPV6-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ZIPCODE-TIMEZONE-ISP-DOMAIN-NETSPEED-AREACODE-WEATHER-MOBILE-ELEVATION-USAGETYPE.BIN")
}
