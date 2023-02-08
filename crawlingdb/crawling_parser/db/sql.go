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

func SelectBase(param map[string]string) string{
	sql := `
		select 
			*
		from base
		where description like "%재료%" and description  like "%[%"
	`
	return JsonQuery(format_map(sql, param))
}

func SelectBase2(param map[string]string) []map[string]string{
	sql := `
		select 
			*
		from base
		where description like "%재료%" and description  like "%[%"
	`
	return Query(format_map(sql, param))
}

func InsertRecipe(param map[string]string) string{
	sql := `
		insert into recipe(
			name
			,url
			,imgurl
		)values(
			"{name}"
			,"{url}"
			,"{imgurl}"
		) ON DUPLICATE KEY UPDATE	
		name = "{name}", url = "{url}", imgurl = "{imgurl}"
	`
	return ExecGetLastID(format_map(sql, param))
}

func SelectRecipeID(param map[string]string) []map[string]string{
	sql := `
		select 
			*
		from recipe
		where url = "{url}"
	`
	return Query(format_map(sql, param))
}

func InsertIngredients(param map[string]string) string{
	sql := `
		insert into ingredients(
			recipe_id
			,name
			,rownum
			,type
		)values(
			{recipe_id}
			,"{name}"
			,{rownum}
			,"{type}"
		) ON DUPLICATE KEY UPDATE	
		recipe_id = {recipe_id}, name = "{name}", rownum = {rownum}, type = "{type}"
	`
	return Exec(format_map(sql, param))
}

func InsertDirections(param map[string]string) string{
	sql := `
		insert into directions(
			recipe_id
			,name
			,rownum
			,type
		)values(
			{recipe_id}
			,"{name}"
			,{rownum}
			,"{type}"
		) ON DUPLICATE KEY UPDATE	
		recipe_id = {recipe_id}, name = "{name}", rownum = {rownum}, type = "{type}"
	`
	return Exec(format_map(sql, param))
}
