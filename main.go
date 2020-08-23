package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("file .env gagal dimuat :(")
	}

	env := os.Getenv(key)
	return env
}

func main() {
	token := getEnv("DISCORD_TOKEN")

	dc, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	dc.AddHandler(onReady)
	dc.AddHandler(onMessageCreate)
	dc.AddHandler(onMessageCreateCommandHandler)

	err = dc.Open()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dc.Close()
}

func onReady(dc *discordgo.Session, r *discordgo.Ready) {
	log.Println("Login Sebagai " + r.User.Username + "#" + r.User.Discriminator)

	err := dc.UpdateStatus(0, "Ok")
	if err != nil {
		log.Println("Error saat ingin mengubah status")
	}
}
