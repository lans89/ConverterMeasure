package distances

const flow string = "->"

var unitMap map[string]MeasureUnit = map[string]MeasureUnit{
	"m":  {"m", "meters"},
	"cm": {"cm", "centimeters"},
	"Km": {"Km", "Kilometers"},
}

var convertionMap map[string]MeasureCalc = map[string]MeasureCalc{
	"m" + flow + "cm": {"*", 100.00},
	"m" + flow + "Km": {"/", 1000.00},

	"cm" + flow + "m": {"/", 100.00},
	"Km" + flow + "m": {"*", 1000.00},
}
