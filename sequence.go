package rhin

func Seq1[A any, B any](
	value A,
	op1 func(input A) B,
) B {
	return op1(value)
}

func Seq2[A any, B any, C any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
) C {
	return op2(op1(value))
}

func Seq3[A any, B any, C any, D any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	op3 func(input C) D,
) D {
	return op3(op2(op1(value)))
}

func Seq4[A any, B any, C any, D any, E any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	op3 func(input C) D,
	op4 func(input D) E,
) E {
	return op4(op3(op2(op1(value))))
}

func Seq5[A any, B any, C any, D any, E any, F any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	op3 func(input C) D,
	op4 func(input D) E,
	op5 func(input E) F,
) F {
	return op5(op4(op3(op2(op1(value)))))
}

func Seq6[A any, B any, C any, D any, E any, F any, G any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	op3 func(input C) D,
	op4 func(input D) E,
	op5 func(input E) F,
	op6 func(input F) G,
) G {
	return op6(op5(op4(op3(op2(op1(value))))))
}

func Seq7[A any, B any, C any, D any, E any, F any, G any, H any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	op3 func(input C) D,
	op4 func(input D) E,
	op5 func(input E) F,
	op6 func(input F) G,
	op7 func(input G) H,
) H {
	return op7(op6(op5(op4(op3(op2(op1(value)))))))
}

func Seq8[A any, B any, C any, D any, E any, F any, G any, H any, I any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	op3 func(input C) D,
	op4 func(input D) E,
	op5 func(input E) F,
	op6 func(input F) G,
	op7 func(input G) H,
	op8 func(input H) I,
) I {
	return op8(op7(op6(op5(op4(op3(op2(op1(value))))))))
}

func Seq9[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	op3 func(input C) D,
	op4 func(input D) E,
	op5 func(input E) F,
	op6 func(input F) G,
	op7 func(input G) H,
	op8 func(input H) I,
	op9 func(input I) J,
) J {
	return op9(op8(op7(op6(op5(op4(op3(op2(op1(value)))))))))
}

func Seq10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	op3 func(input C) D,
	op4 func(input D) E,
	op5 func(input E) F,
	op6 func(input F) G,
	op7 func(input G) H,
	op8 func(input H) I,
	op9 func(input I) J,
	op10 func(input J) K,
) K {
	return op10(op9(op8(op7(op6(op5(op4(op3(op2(op1(value))))))))))
}
