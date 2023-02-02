package db

import (
	"strings"

)

func format_map(str string, param map[string]string) string{
	cnt := strings.Count(str, "{")

	for i := 0; i<cnt ;i++ {
		

		s := strings.Index(str,"{")
		e := strings.Index(str,"}")

		key := str[s+1:e]
	
		str = strings.Replace(str, "{"+key+"}", param[key],1)
	}
	return str
}


func SelectBase(param map[string]string) string{
	sql := `
		select 
			*
		from base
		where description like "%재료%" and description  like "%[%"
	`
	return JsonQuery(format_map(sql, param))
}

func SelectBase2(param map[string]string) map[string]map[string]string{
	sql := `
		select 
			*
		from base
		where description like "%재료%" and description  like "%[%"
	`
	return Query(format_map(sql, param))
}