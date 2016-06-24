package pkg

type T struct{}

func fn1(_ T) {}

func fn2() {
	t1 := &T{}
	fn1(&*t1) // MATCH /&*T is ineffective. It will be simplified to T/
	fn1(*&s1) // MATCH /\*&T is ineffective. It will be simplified to T/
}
