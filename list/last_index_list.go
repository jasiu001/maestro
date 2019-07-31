package list

type LastIndexList struct {
	newer string
	older string
}

func (li *LastIndexList) add(element string) {
	li.older = li.newer
	li.newer = element
}

func (li *LastIndexList) exist(element string) bool {
	if element == li.newer || element == li.older {
		return true
	}

	return false
}
