package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// CekPalindrome untuk memeriksa apakah string adalah palindrome
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// CekAnagram untuk memeriksa apakah dua string adalah anagram
func CekAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	counts := make(map[rune]int)
	for _, char := range str1 {
		counts[char]++
	}
	for _, char := range str2 {
		counts[char]--
		if counts[char] < 0 {
			return false
		}
	}
	return true
}

// FormatJSON untuk memformat file JSON
func FormatJSON(inputPath, outputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	data := result["data"].([]interface{})
	formattedData := map[string]interface{}{
		"total": 176,
		"data":  data,
	}
	formattedJSON, _ := json.MarshalIndent(formattedData, "", "    ")
	err = ioutil.WriteFile(outputPath, formattedJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data berhasil diformat dan disimpan ke:", outputPath)
}
