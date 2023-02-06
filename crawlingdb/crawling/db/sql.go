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


func InsertBase(param map[string]string) string{
	sql := `
		insert into base(
			url
			,description
			,title)
		values(
			"{url}"
			,"{description}"
			,"{title}"
		) ON DUPLICATE KEY UPDATE	
		description = "{description}", title = "{title}"
	`
	return Exec(format_map(sql, param))
}

