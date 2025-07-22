package summary

type SummaryService struct{}

func NewSummaryService() *SummaryService {
	return &SummaryService{}
}

func (s *SummaryService) GetSummary() string {
	return "summary"
}
