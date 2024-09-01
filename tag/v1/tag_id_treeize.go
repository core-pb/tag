package tag

func (x *TagIDTreeize) Has(tagID uint64) bool {
	if x == nil || tagID == 0 {
		return false
	}
	if tagID == x.Id {
		return true
	}
	for _, v := range x.Child {
		if v.Has(tagID) {
			return true
		}
	}
	return false
}

func (x *TagIDTreeize) Contains(tagID []uint64) bool {
	for _, v := range tagID {
		if !x.Has(v) {
			return false
		}
	}
	return true
}
