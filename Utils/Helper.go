package Utils

func GetGreeting(name string) string {
	return "Hello " + name
}

func GetTwoNextNumbers(num int) (int, int) {
	return num+1, num+3
}

func UpdateIntValue(num int)  {
	num = 345
}

func UpdateMapValue(dict map[string]int, key string, val int)  {
	dict[key] = val
}