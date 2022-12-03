package lq

type Tuple2[A, B any] struct {
	A A
	B B
}

func T2[A, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{
		A: a,
		B: b,
	}
}

func (t Tuple2[A, B]) Unpack() (A, B) {
	return t.A, t.B
}

type Tuple3[A any, B any, C any] struct {
	A A
	B B
	C C
}

func T3[A, B, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{
		A: a,
		B: b,
		C: c,
	}
}

func (t Tuple3[A, B, C]) Unpack() (A, B, C) {
	return t.A, t.B, t.C
}

type Tuple4[A any, B any, C any, D any] struct {
	A A
	B B
	C C
	D D
}

func T4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

func (t Tuple4[A, B, C, D]) Unpack() (A, B, C, D) {
	return t.A, t.B, t.C, t.D
}

type Tuple5[A, B, C, D, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

func T5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
	}
}

func (t Tuple5[A, B, C, D, E]) Unpack() (A, B, C, D, E) {
	return t.A, t.B, t.C, t.D, t.E
}

type Tuple6[A, B, C, D, E, F any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
}

func T6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	return Tuple6[A, B, C, D, E, F]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
	}
}

func (t Tuple6[A, B, C, D, E, F]) Unpack() (A, B, C, D, E, F) {
	return t.A, t.B, t.C, t.D, t.E, t.F
}

type Tuple7[A any, B any, C any, D any, E any, F any, G any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
}

func (t Tuple7[A, B, C, D, E, F, G]) Unpack() (A, B, C, D, E, F, G) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G
}

type Tuple8[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
}

func (t Tuple8[A, B, C, D, E, F, G, H]) Unpack() (A, B, C, D, E, F, G, H) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H
}

type Tuple9[A any, B any, C any, D any, E any, F any, G any, H any, I any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
}

func (t Tuple9[A, B, C, D, E, F, G, H, I]) Unpack() (A, B, C, D, E, F, G, H, I) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H, t.I
}
