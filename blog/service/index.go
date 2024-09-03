package service

type IndexService struct {
	Service
	count int
}

func (s *IndexService) OnInitService() {
	s.count = 0
}

func (s *IndexService) Increment() int {
	s.count++
	return s.count
}
