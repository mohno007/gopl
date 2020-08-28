package main

func rotate(s []int, amount int) {
	amount = amount % len(s)
	if amount < 0 {
		amount = len(s) - amount
	}

	if amount == 0 {
		return
	}

	for isrc := range s {
		itarget := (isrc + amount) % len(s)
		tmp := s[itarget]
		s[itarget] = s[isrc]
		s[isrc] = tmp
	}
}
