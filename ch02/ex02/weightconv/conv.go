package weightconv

func KgToLbs(kg Kilogram) Pound {
	return Pound(kg * 2.20462)
}

func LbsToKg(lbs Pound) Kilogram {
	return Kilogram(lbs * 0.453592)
}
