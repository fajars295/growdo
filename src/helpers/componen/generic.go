package componen

func CheckIfRoles[T int64](parameter T) T {

	if parameter == 0 {
		return 1
	}
	return parameter

}
