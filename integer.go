package math4g

func Norm(value Scala) Scala {
	if value > 0 {
		return 1.0
	} else if value < 0 {
		return -1.0
	}
	return 0
}

func Floor(x Scala) Scala {
	xi := Scala(int(x))
	if x < xi {
		return xi - 1
	}
	return xi
}

func Ceil(x Scala) Scala {
	xi := Scala(int(-x))
	if -x < xi {
		return 1 - xi
	}
	return -xi
}

func Trunc(x Scala) Scala {
	return Scala(int(x))
}
