package utils

var ColorSlice = []string{"#CC0033", "#99FF33", "#6600FF", "#999933", "#996699", "#990033", "#66FF33", "#000033", "#0033CC", "#CC3333", "#CCFF99", "#CCFFCC", "#FFCCCC", "#FFFF33", "#990000", "#66CCFF", "#660000", "#333333", "#330000", "#000000"}

func SumSlice(slice []int) int {
	var b int
	for _, val := range slice {
		b += val
	}
	return b
}

func GetColor(i int) string {
	if i < len(ColorSlice) {
		return ColorSlice[i]
	}
	return ""
}
