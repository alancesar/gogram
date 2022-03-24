package main

import (
	"fmt"
	"github.com/alancesar/gogram/mass"
)

type Data struct {
	Mass mass.Mass `json:"mass"`
}

func main() {

	d := Data{Mass: mass.NewFromGram(100)}

	//marshal, _ := json.Marshal(d)
	fmt.Printf("%v", d)
}
