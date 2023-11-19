package main

import "fmt"

func getDataMember(name string, address string, age int)(string,string,int)  {
	vName := "name :" +name+ "\n"
	vAddress := "address :" +address+ "\n"
	vAge := age
	return vName, vAddress,vAge
}

func main()  {
	vGetDataMemberName,vGetDataMemberAdress,VGetDataMemberAge := getDataMember("arifin","jagakarsa",23)
	fmt.Println(vGetDataMemberName,vGetDataMemberAdress,VGetDataMemberAge)
	vGetDataMemberName2,vGetDataMemberAdress2,_ := getDataMember("arifin","jagakarsa",23)
	fmt.Println(vGetDataMemberName2,vGetDataMemberAdress2)
	/**
		name :arifin
		address :jagakarsa
		23
	*/
}