package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
    OpenAI_API_Key string
}

func NewConfig() *Config {
    key := os.Getenv("OPENAI_API_KEY")
    if key == "" {
        if _, err := os.Stat("gopt-renamer.conf"); err == nil {
            content, err := os.ReadFile("gopt-renamer.conf")
            if err != nil {
                panic(err)
            }
            key = strings.TrimSpace(string(content))
        } else {
            reader := bufio.NewReader(os.Stdin)
            fmt.Print("You need an OpenAI API Key configured. Do you have a key? (yes/no): ")
            text, _ := reader.ReadString('\n')
            text = strings.TrimSpace(text)
            if text == "no" {
                panic("OpenAI API Key not available")
            }
            fmt.Print("What is your openai api key?: ")
            key, _ = reader.ReadString('\n')
            key = strings.TrimSpace(key)
            fmt.Print("How would you like to store your openai api key as an enviroment variable, or in a config file? (env/file): ")
            storage, _ := reader.ReadString('\n')
            storage = strings.TrimSpace(storage)
            if storage == "env" {
                os.Setenv("OPENAI_API_KEY", key)
            } else {
                err := os.WriteFile("gopt-renamer.conf", []byte(key), 0644)
                if err != nil {
                    panic(err)
                }
            }
        }
    }
    return &Config{
        OpenAI_API_Key: key,
    }
}