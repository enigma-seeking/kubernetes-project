/*
给定一个字符串数组，["I", "am", "stupid", "and", "weak"]
用for循环遍历该数组并修改


*/
package main

func main() {
	needToModify := []string{}
	needToModify = append(needToModify, "I")
	needToModify = append(needToModify, "am")
	needToModify = append(needToModify, "stupid")
	needToModify = append(needToModify, "and")
	needToModify = append(needToModify, "weak")

	afterModify := make([]string, 5)

	for index := range needToModify {
		if needToModify[index] == "stupid" {
			afterModify[index] = "smart"
			continue
		}
		if needToModify[index] == "weak" {
			afterModify[index] = "strong"
			continue
		}
		afterModify[index] = needToModify[index]
	}

	for index := range afterModify {
		print(afterModify[index], " ")
	}

}
