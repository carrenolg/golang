package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"carrenolg.io/tutorials/myface/models"
)

const TOKEN = "EABD4vyToGPYBAHKn2rDukXlM2bSOKsPlxEMwQkZAo4UYHY2zGVEne2mHV6frnFn8bMC56hOsYeEwJZA39GcMs5jovx4SBH3LpEfBsZARKDgE33uldQF3mBZABMabkCiT5yrur68ZBadc5JZCgAE3ggoqhNRdze8CgjBifTL19KsTvkTp3VjvGyZAdJyLT6R3YQmyqn3hOqvwrX3Qh0QK45ylM72ZBjIvsMbtZCTx5zokAcmpWaIZAZCVY0r"

func main() {
	fmt.Println("welcome to myface app")
	// 1. Getting the posts
	url := fmt.Sprintf("https://graph.facebook.com/v12.0/me/posts?access_token=%s", TOKEN)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// 2. Mapping the posts
	var result models.Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println(err)
	}

	//3. Printing result
	lenData := len(result.Data)
	fmt.Println("Total Post:", lenData)
	for i := 0; i < lenData; i++ {
		fmt.Println(result.Data[i])
	}
}
