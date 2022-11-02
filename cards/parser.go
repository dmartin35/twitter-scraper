package cards

func GetTwitterCardParserByName(name string) TwitterCardParser {

	if name == "unified_card" {
		return &UnifiedCardParser{}
	}

	if name == "summary" {
		return &SummaryCardParser{}
	}

	if name == "summary_large_image" {
		return &SummaryLargeCardParser{}
	}

	return nil
}
