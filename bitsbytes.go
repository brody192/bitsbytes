package bitsbytes

func KilobytesToBytes[T1 int | int64 | float64, T2 int | int64 | float64](Kilobytes T1) T2 {
	return T2(Kilobytes) * 1024
}

func MegabytesToBytes[T1 int64 | float64, T2 int | int64 | float64](megabytes T1) T2 {
	return T2(megabytes) * 1024 * 1024
}

func MegabitsToBytes[T1 int64 | float64, T2 int | int64 | float64](megabits T1) T2 {
	return T2(megabits) * 131072
}

func BytesToMegabytes[T1 int64 | float64, T2 int | int64 | float64](bytes T1) T2 {
	return T2(bytes) / 1048576
}

func BytesToMegabits[T1 int64 | float64, T2 int | int64 | float64](bytes T1) T2 {
	return T2(bytes) / 131072
}
