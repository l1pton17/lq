package lq

type orderBy[TValue any, TBy Ordered] struct {
	selector func(v TValue) TBy
}

func OrderBy[TValue any, TBy Ordered](selector func(v TValue) TBy) Orderer[TValue] {
	return orderBy[TValue, TBy]{
		selector: selector,
	}
}

func (o orderBy[TValue, TBy]) Compare(a, b TValue) int {
	va := o.selector(a)
	vb := o.selector(b)

	switch {
	case va == vb:
		return 0

	case va < vb:
		return -1

	case va > vb:
		return 1

	default:
		panic("no compare")
	}
}
