package t13XX

func angleClock(hour int, minutes int) float64 {
	minuteAngle := float64((minutes % 60) * 6)
	hourAngle := float64((hour%12)*30) + float64(float64(30*minutes)/float64(60))

	diff := float64(0)
	if minuteAngle > hourAngle {
		diff = minuteAngle - hourAngle
	} else {
		diff = hourAngle - minuteAngle
	}

	return minFloat(diff, 360-(diff))
}

func minFloat(x, y float64) float64 {
	if x > y {
		return float64(y)
	}

	return float64(x)
}

// 60 -> 30
// x -> ?
