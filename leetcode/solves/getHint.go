package solves

import "fmt"

func GetHint(secret string, guess string) string {
	bulls, cows := 0, 0
	secretMap := make(map[byte]int)
	guessMap := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretMap[secret[i]]++
			guessMap[guess[i]]++
		}
	}
	for k, v := range secretMap {
		cows += min(v, guessMap[k])
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
