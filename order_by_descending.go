package lq

type orderByDescending[TValue any, TBy Ordered] struct {
	selector func(v TValue) TBy
}

func OrderByDescending[TValue any, TBy Ordered](selector func(v TValue) TBy) Orderer[TValue, TBy] {
	return orderByDescending[TValue, TBy]{
		selector: selector,
	}
}

func (o orderByDescending[TValue, TBy]) Compare(a, b TValue) int {
	va := o.selector(a)
	vb := o.selector(b)

	switch {
	case va == vb:
		return 0

	case va < vb:
		return 1

	case va > vb:
		return -1

	default:
		panic("no compare")
	}
}
