package lib

import (
	"automation/types"
	"encoding/json"
	"fmt"
	"math"
	"sort"
)

func (s *Session)ListCatsAPI()(string,error){
	URL:=types.ListKittyAPI
	fmt.Print("\nAPI:",URL )
	rawResponse, err := RestClient(URL, s.Token, "GET", "")
	if err != nil {
		fmt.Print("\nError while listing cats ...")
		fmt.Print("\nResponse:", rawResponse)
		return "",err
	}
	return rawResponse,nil
}

func (s *Session)ListCats()error{
	rawResponse, err := s.ListCatsAPI()
	if err != nil {
		fmt.Print("\nError while listing cats ...")
		fmt.Print("\nResponse:", rawResponse)
		return err
	}
	var Response types.ListKittyResponse
	err = json.Unmarshal([]byte(rawResponse), &Response)
	if err != nil {
		fmt.Print("\nError while unmarshaling ...")
		return err
	}
	PrettyStruct(Response)
	return nil
}

func (s *Session)Calculate_ASCII_from_Charater() {
	rawResponse, err := s.ListCatsAPI()
	if err != nil {
		fmt.Print("\nResponse:", rawResponse)
		//return 0, err
	}
	var Response types.ListKittyResponse
	err = json.Unmarshal([]byte(rawResponse), &Response)
	if err != nil {
		fmt.Print("\nError while unmarshaling ...")
	//	return err
	}

	var catslist []types.CatRank
	var cats types.CatRank
	for _, i := range Response {
		if i.Name != "James" {
			var sum int
			cats.Name = i.Name
			for _, c := range i.Name {
				ascii := int(c)
				sum = sum + ascii
			}
			cats.Awesomeness = sum
			catslist = append(catslist,cats)
		} else {
			cats.Name = "James"
			cats.Rank = 1
			cats.Awesomeness = math.MaxInt64
			catslist = append(catslist,cats)
		}
	}
	sort.Sort(SortAnim(catslist))
	catlist := Assign_Rank(catslist)
	PrettyStruct(catlist)
}

func Assign_Rank(catslist []types.CatRank)([]types.CatRank){
	var list  []types.CatRank
	for n,t := range catslist{
		t.Rank=n+1
		list = append(list,t)
	}
	return list
}

type SortAnim []types.CatRank
func (c SortAnim) Len() int           {return len(c)}
func (c SortAnim) Swap(i, j int)      {c[i], c[j] = c[j], c[i]}
func (c SortAnim) Less(i, j int) bool {return c[i].Awesomeness > c[j].Awesomeness}

func (s *Session)DeleteCatsAPI (catname string) error {
	URL:=types.ListKittyAPI+"/"+catname
	fmt.Print("\nAPI:",URL )
	rawResponse, err := RestClient(URL, s.Token, "DELETE", "")
	if err != nil {
		fmt.Print("\nResponse:", rawResponse)
		return err
	}
	return nil
}

func (s *Session)DeleteCatsAndCheck (catname string)error{
	err := s.DeleteCatsAPI(catname)
	if err != nil {
		fmt.Print("\nError while deleting cats name:",catname)
		return err
	}
	rawResponse, err := s.ListCatsAPI()
	if err != nil {
		fmt.Print("\nError while listing cats ...")
		fmt.Print("\nResponse:", rawResponse)
		return err
	}
	var Response types.ListKittyResponse
	err = json.Unmarshal([]byte(rawResponse), &Response)
	if err != nil {
		fmt.Print("\nError while unmarshaling ...")
		return err
	}
	PrettyStruct(Response)
	for _, i := range Response {
		if i.Name == catname {
			fmt.Errorf("Cat is not deleted", catname)
		}
	}
	fmt.Print("cat is deleted successfully")
	return nil
}

func (s *Session)ModifyCatsAPI (catname_original string,catname_tomodify string) error {
	URL:=types.ListKittyAPI+"/"+catname_original
	loginPayload := ToJSONString(types.CatNewName{NewName:catname_tomodify})
	fmt.Print("\nAPI:",URL )
	rawResponse, err := RestClient(URL, s.Token, "PUT", loginPayload)
	if err != nil {
		fmt.Print("\nResponse:", rawResponse)
		return err
	}
	return nil
}

func (s *Session)ModifyCatsNameAndCheck (catname_original string,catname_tomodify string)error{
	err := s.ModifyCatsAPI(catname_original,catname_tomodify)
	if err != nil {
		fmt.Print("\nError while modify cats name:",catname_original)
		return err
	}
	rawResponse, err := s.ListCatsAPI()
	if err != nil {
		fmt.Print("\nError while listing cats ...")
		fmt.Print("\nResponse:", rawResponse)
		return err
	}
	var Response types.ListKittyResponse
	err = json.Unmarshal([]byte(rawResponse), &Response)
	if err != nil {
		fmt.Print("\nError while unmarshaling ...")
		return err
	}
	PrettyStruct(Response)

	if !FindCat(Response,catname_original){
		fmt.Print(catname_original," catname is modified successfully to ", catname_tomodify)
	}else {
		fmt.Print(catname_original," catname is not modified successfully to ", catname_tomodify)
		return fmt.Errorf(catname_original," catname is not modified successfully to ",catname_tomodify)
	}
	return nil
}

func FindCat(slice types.ListKittyResponse, catname_original string) bool {
	for _, item := range slice {
		if item.Name == catname_original {
			return true
		}
	}
	return false
}