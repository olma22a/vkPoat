package main

import (
	"github.com/SevereCloud/vksdk/api"
	"log"
)

func main() {
	vk := api.NewVK("")

	params := api.Params{
		"owner_id": "",
		"message":  "blinnica",
	}

	_, err := vk.WallPost(params)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	log.Println("Опубликовано!")
}
