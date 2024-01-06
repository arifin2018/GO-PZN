package main

import (
	"encoding/csv"
	"os"
)

func main()  {
	/**
		encode to string
	*/

	// data := []byte("arzfsx")
	// str := base64.StdEncoding.EncodeToString(data)
	// fmt.Println(str)

	// data, err := base64.StdEncoding.DecodeString(str)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Printf("%q\n", data)

	/**
		akhir encode to string
	*/

// 	in := `first_name,last_name,username
// "Rob","Pike",rob
// Ken,Thompson,ken
// "Robert","Griesemer","gri"`

// 	r := csv.NewReader(strings.NewReader(in))

// 	for {
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(record)
// 	}

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"nama","depan","belakang"})
	_ = writer.Write([]string{"azriel","rafiq","pradipta"})
	_ = writer.Write([]string{"khoirunnisa","miftahul","jannah"})

	writer.Flush()
}