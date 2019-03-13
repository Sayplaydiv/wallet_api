package string_manage

import "strings"

func Json01(Str string) string{

	json_0 := []string{"id",Str}
	json_1 := strings.Join(json_0,"\":\"")
	json_2 := "\""+ json_1 + "\""

	return json_2
}


func Json02(Str string) string{

	json_0 := "\""+ Str + "\""

	return json_0
}