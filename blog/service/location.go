package service

import (
	"blog/logger"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"strings"
)

var (
	ipDbPath      = "ip2region.xdb"
	ipVectorIndex []byte
)

type LocationService struct {
	Service
}

func (s *LocationService) OnInitService() {
	vIndex, err := xdb.LoadVectorIndexFromFile(ipDbPath)
	if err != nil {
		logger.Fatalf("failed to load vector index: %s", err)
	}
	ipVectorIndex = vIndex
}

func (s *LocationService) GetLocationFromIp(ip string) string {
	searcher, err := xdb.NewWithVectorIndex(ipDbPath, ipVectorIndex)
	if err != nil {
		return "Unknown"
	}
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		return "Unknown"
	}
	regions := make([]string, 0)
	for _, r := range strings.Split(region, "|") {
		if r != "0" {
			regions = append(regions, r)
		}
	}
	return strings.Join(regions, ",")
}
