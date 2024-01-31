package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
    OpenAI_API_Key string
}

func setEnvironmentVariable(key string) error {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", "setx", "OPENAI_API_KEY", key)
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		bashrcPath := filepath.Join(homeDir, ".bashrc")
		f, err := os.OpenFile(bashrcPath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString(fmt.Sprintf("export OPENAI_API_KEY=%s\n", key))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewConfig() (*Config, error) {
    key := os.Getenv("OPENAI_API_KEY")
    if key == "" {
        if _, err := os.Stat("gopt-renamer.conf"); err == nil {
            content, err := os.ReadFile("gopt-renamer.conf")
            if err != nil {
                return nil, err
            }
            key = strings.TrimSpace(string(content))
        } else {
            reader := bufio.NewReader(os.Stdin)
            fmt.Print("An OpenAI API Key is required. Do you have a key? (yes/no): ")
            text, _ := reader.ReadString('\n')
            text = strings.TrimSpace(text)
            if text == "no" {
                return nil, fmt.Errorf("OpenAI API Key not available")
            }
            fmt.Print("Enter your OpenAI API Key (key is only stored locally): ")
            key, _ = reader.ReadString('\n')
            key = strings.TrimSpace(key)
            fmt.Print("How would you like to store the OpenAI API Key for future use? As an enviroment variable, or in a config file? (env/file): ")
            storage, _ := reader.ReadString('\n')
            storage = strings.TrimSpace(storage)
            if storage == "env" {
				err := setEnvironmentVariable(key)
				if err != nil {
					return nil, err
				}
            } else {
                err := os.WriteFile("gopt-renamer.conf", []byte(key), 0644)
                if err != nil {
                    return nil, err
                }
            }
        }
    }
    return &Config{OpenAI_API_Key: key}, nil
}