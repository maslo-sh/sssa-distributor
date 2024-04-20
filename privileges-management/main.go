package main

import (
	"fmt"
	"github.com/SSSaaS/sssa-golang"
	_ "github.com/SSSaaS/sssa-golang"
)

func main() {
	fmt.Println("Hello, World!")

	shares, err := sssa.Create(7, 10, "mzomi_bhl_2023")
	if err != nil {
		fmt.Printf("failed to create secret: %v\n", err)
	}

	fmt.Println(shares)

	recovered, err := sssa.Combine(shares[5:10])
	if err != nil {
		fmt.Printf("failed to combine secrets: %v\n", err)
	}

	fmt.Println(recovered)
}
