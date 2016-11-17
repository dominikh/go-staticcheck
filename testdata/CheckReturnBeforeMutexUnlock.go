package pkg

import "sync"

func fn1() {
	var x sync.Mutex
	x.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn2() {
	var x, y sync.Mutex
	x.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
	y.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	y.Unlock()
}

func fn3() {
	var x sync.Mutex
	x.Lock()
}

func fn4() {
	var x sync.Mutex
	x.Lock()
	defer x.Unlock()
	if true {
		return
	}
}

func fn5() {
	var x sync.Mutex
	x.Lock()
	if true {
	}
	x.Unlock()
	return
}

func fn6() {
	var x, y sync.Mutex
	y.Lock()
	x.Lock()
	y.Unlock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn7() {
	x := &struct {
		sync.Mutex
	}{}

	x.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn8() {
	var x sync.Mutex
	x.Lock()
	defer func() {
		x.Unlock()
	}()
	return
}

func fn9() {
	var x sync.Mutex
	x.Lock()
	if true {

	} else {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn10() {
	var x sync.Mutex
	x.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	} else if false {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn11() {
	type X struct {
		mu sync.Mutex
	}
	var x X
	x.mu.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.mu.Unlock()
}

func fn12() {
	type X struct {
		mu sync.Mutex
	}
	var x X
	x.mu.Lock()
	mu := &x.mu
	mu.Unlock()
	if true {
		return
	}
}

func fn13() {
	var x sync.Mutex
	x.Lock()
	y := &x
	y.Unlock()
	if true {
		return
	}
}

func fn14() {
	var x, y sync.Mutex
	x.Lock()
	y.Unlock()
	if true {
		return // false negative
	}
}

func fn15() { //
	var x, y sync.Mutex
	x.Lock()
	defer x.Unlock()

	x.Unlock()

	y.Lock()
	defer y.Unlock()

	x.Lock()
	return
}

func fn16() {
	fn := func() {
		var x sync.Mutex
		x.Lock()
		if true {
			return // MATCH /return before mutex unlock/
		}
		x.Unlock()
	}
	fn()
}

func fn17() {
	x := struct {
		m1 struct {
			m2 sync.Mutex
		}
	}{}

	x.m1.m2.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.m1.m2.Unlock()
}

func fn18() {
	var x sync.RWMutex
	x.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()

	x.RLock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.RUnlock()
}

func fn19() {
	x := struct {
		m func() *sync.Mutex
	}{
		m: func() *sync.Mutex {
			return new(sync.Mutex)
		},
	}

	x.m().Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.m().Unlock()
}

func fn20() {
	x := &sync.Mutex{}
	x.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn21() {
	x := &struct {
		sync.Mutex
	}{}

	x.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn22() {
	var x sync.Locker
	x = new(sync.Mutex)

	x.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn23() {
	var x sync.Mutex
	x.Lock()
	if true {
		return // MATCH /return before mutex unlock/
	}
	func() {
		x.Unlock()
	}()
}

func fn24() {
	var x sync.Mutex
	x.Lock()
	switch "word" {
	case "word":
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn25() {
	var x sync.Mutex
	fn := func() {
		x.Unlock()
	}
	defer fn()
	x.Lock()
	if true {
		return
	}
}

func fn26() {
	var m map[string]int
	var x sync.Mutex
	x.Lock()
	if _, ok := m["value"]; ok {
		return // MATCH /return before mutex unlock/
	}
	x.Unlock()
}

func fn27() {
	var x sync.Mutex
	x.Lock()
	if true {
		return
	}
}
