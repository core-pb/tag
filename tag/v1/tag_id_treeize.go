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

type FlatTagID struct {
	ID       uint64 `json:"id"`
	ParentID uint64 `json:"parent_id"`
}

func (x *TagIDTreeize) Flat() []FlatTagID {
	return x.flatten(0)
}

func (x *TagIDTreeize) flatten(parentID uint64) []FlatTagID {
	var items []FlatTagID
	for _, node := range x.Child {
		items = append(items, FlatTagID{ID: node.Id, ParentID: parentID})
		items = append(items, node.flatten(node.Id)...)
	}
	return items
}

func (x *TagIDTreeize) FromFlat(arr []FlatTagID) {
	var (
		tree     []*TagIDTreeize
		childMap = make(map[uint64][]*TagIDTreeize)
	)

	for _, item := range arr {
		node := &TagIDTreeize{Id: item.ID}
		if item.ParentID == 0 {
			tree = append(tree, node)
		} else {
			childMap[item.ParentID] = append(childMap[item.ParentID], node)
		}
	}

	var fn func(nodes []*TagIDTreeize)
	fn = func(nodes []*TagIDTreeize) {
		for i := range nodes {
			if children, ok := childMap[nodes[i].Id]; ok {
				nodes[i].Child = children
				fn(nodes[i].Child)
			}
		}
	}

	fn(tree)
	x.Id = 0
	x.Child = tree
}
