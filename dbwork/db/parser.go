package db

import (
	"encoding/json"
	//"log"
	"regexp"
	"strings"
)
	
type Recipe struct {
	Name 			string
	Ingredients		[]Ingredients
	Directions		[]Directions
	Origin			string
}

type Ingredients struct {
	Rownum 	string
	Name 	string
	Dtype 	string
}

type Directions struct {
	Rownum 		string
	Dtype 		string
	Description string

}

func Perser() string{


	dess := SelectBase2(nil)

	recipes := make([]Recipe,0)

	// tmapfoodname := make(map[string][]string)
	// tmapjerou := make(map[string][]string)
	// tmapmake := make(map[string][]string)
	
	//재료 및 만드는법 제목 가져오기
	titlecheck, _ := regexp.Compile("\\[[^a-zA-Z\\]]+\\]")


	//음식명 가져오기
	//re2, _ := regexp.Compile("\\={5,}\\n{1,3}[^a-zA-Z\\n\\[\\(]+")

	//태그 가져오기
	//re3, _ := regexp.Compile("\\#[^a-zA-Z |\n]+")

	for _, des := range dess {
		recipe := new(Recipe)
		recipe.Origin = des["description"]
		ingredients := make([]Ingredients,0)
		directions := make([]Directions,0)


		desarr := strings.Split(des["description"], "\n")
		finditem := ""
		for i, v := range desarr {
			if i < 10 && strings.Contains(v,"======"){
				finditem = "이름"
				continue
			}

			if finditem == "이름" {
				if !blankCheck(v) && v[:1]!= "*" && !strings.Contains(v, "재료")&& !strings.Contains(v, "법"){
					recipe.Name = v
					finditem = ""
				}
			}

			title := titlecheck.FindString(v)
			if title != ""{
				if strings.Contains(title, "재료") && !strings.Contains(title, "법"){
					finditem = "재료"
				}else if strings.Contains(title, "법"){
					finditem = "만드는법"
				}
			
			} 

			if blankCheck(v){
				continue
			}	
			if engCheck(v) {
				finditem = ""
			
			}	
			
			if finditem == "재료" {
					
				bv := strings.ReplaceAll(v, " ", "")
				
				ingredient := new(Ingredients)
				ingredient.Name = v
				if (bv[:1] == "*" || bv[:1] == "[") && blankCheck(desarr[i-1]){
					ingredient.Dtype = "title"
					
				}

				ingredients = append(ingredients,*ingredient)
			

	

			}

			if finditem == "만드는법" {

				bv := strings.ReplaceAll(v, " ", "")					
					
				direction := new(Directions)
				direction.Description = v
				if (bv[:1] == "*" || bv[:1] == "[") && blankCheck(desarr[i-1]){
					direction.Dtype = "title"
				}

				directions = append(directions,*direction)

					

			}
			

		}

		recipe.Ingredients = ingredients
		recipe.Directions = directions
		recipes = append(recipes,*recipe)
		//줄바꿈 카운트~
	

		// title := re1.FindAllString(v["description"], -1)
		// sj := make([]string, 0)
		// sm := make([]string, 0)
		// sf := make([]string, 0)
		// for _, vv := range title {
		
		// 	if strings.Contains(vv, "재료") && !strings.Contains(vv, "법"){
		// 		sj = append(sj, vv)
		// 	}else if strings.Contains(vv, "법"){
		// 		sm = append(sm, vv)	
		// 	}else {
		// 		sf = append(sf, vv)				
		// 	}

		// }
	
		// log.Println(v["url"],sf)
		// iffoodname := re2.FindAllString(v["description"], -1)
		// //log.Println(v["url"],iffoodname)
		//  for _, vv := range iffoodname {
		// 	vv = strings.ReplaceAll(vv, "=", "")
		// 	vv = strings.ReplaceAll(vv, " ", "")
		// 	vv = strings.ReplaceAll(vv, "\n", "")
		// 	log.Println(v["url"],vv)
		//  }



		// 태그
		// tag := re3.FindAllString(v["description"], -1)
		// log.Println(v["url"],tag)
		// for _, vv := range tag {
		// 	log.Println(v["url"],vv)

		// }
		
	}


	byteData, _ := json.Marshal(recipes)
	// "재료" 가 들어가고 "법"이 안들어간거
	// "만드는" 이 들어가고 "법"이 들어간거

	// log.Println(tmapfoodname)
	// log.Println(tmapjerou)
	// log.Println(tmapmake)

	//tinx := re1.FindAllStringIndex(des,-1)
	//log.Println(tinx)
	return string(byteData)
	

}

func blankCheck(str string) bool{
	str = strings.ReplaceAll(str, " ", "")
	if str != ""{
		return false
	}

	return true
}

func engCheck(str string) bool{
	engcheck, _ := regexp.Compile("[가-힣]+")
	han := engcheck.MatchString(str)
	if !han && str != ""{
		return true
	}
	return false
}

func contains(elems []string, v string) bool {
    for _, s := range elems {
        if v == s {
            return true
        }
    }
    return false
}