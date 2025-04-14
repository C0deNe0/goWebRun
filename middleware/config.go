package main

import "os"

var JWT_SECRET = []byte(getEnv("JWT_SECRET", "supersecret"))

func getEnv(key string, fallback string) string {
	if err := os.Getenv(key); err !=""{
	return err
}
	return fallback
}