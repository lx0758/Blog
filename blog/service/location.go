package service

import (
	"blog/logger"
	"github.com/IncSW/geoip2"
	"net"
	"strings"
)

var (
	geoIp2DbPath = "GeoLite2-City.mmdb"
	cityReader   *geoip2.CityReader
)

type LocationService struct {
	Service
}

func (s *LocationService) OnInitService() {
	reader, err := geoip2.NewCityReaderFromFile(geoIp2DbPath)
	if err != nil {
		logger.Fatalf("failed to load geoip2 database: %s", err)
	}
	cityReader = reader
}

func (s *LocationService) GetLocationFromIp(ip string) string {
	record, err := cityReader.Lookup(net.ParseIP(ip))
	if err != nil {
		return "Unknown"
	}
	regions := make([]string, 0)
	//regions = appendX(regions, record.Continent.Names["zh-CN"])
	regions = appendX(regions, record.Country.Names["zh-CN"])
	if record.Subdivisions != nil {
		regions = appendX(regions, record.Subdivisions[0].Names["zh-CN"])
	}
	regions = appendX(regions, record.City.Names["zh-CN"])
	return strings.Join(regions, ",")
}

func appendX(regions []string, text string) []string {
	if text != "" {
		regions = append(regions, text)
	}
	return regions
}
