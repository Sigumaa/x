package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	i := flag.String("i", "", "𝕚𝕟𝕡𝕦𝕥 𝕤𝕥𝕣𝕚𝕟𝕘")
	f := flag.String("f", "", "𝕚𝕟𝕡𝕦𝕥 𝕗𝕚𝕝𝕖")
	flag.Parse()

	if *i == "" && *f == "" || *i != "" && *f != "" {
		fmt.Println("ℙ𝕝𝕖𝕒𝕤𝕖 𝕤𝕡𝕖𝕔𝕚𝕗𝕪 𝕖𝕚𝕥𝕙𝕖𝕣 𝕚𝕟𝕡𝕦𝕥 𝕤𝕥𝕣𝕚𝕟𝕘 𝕠𝕣 𝕚𝕟𝕡𝕦𝕥 𝕗𝕚𝕝𝕖")
		return
	}

	if *i != "" {
		fmt.Println(replaceCharacters(*i))
		return
	}

	if *f != "" {
		content, err := readFile(*f)
		if err != nil {
			fmt.Println(err)
			return
		}
		res := replaceCharacters(string(content))
		if err := writeFile(*f, res); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("𝔸𝕗𝕥𝕖𝕣 𝕣𝕖𝕡𝕝𝕒𝕔𝕚𝕟𝕘 𝕔𝕙𝕒𝕣𝕒𝕔𝕥𝕖𝕣𝕤, 𝕥𝕙𝕖 𝕗𝕚𝕝𝕖 𝕔𝕠𝕟𝕥𝕖𝕟𝕥 𝕚𝕤:\n")
		fmt.Println(res)
	}
}

func replaceCharacters(input string) string {
	charMap := map[rune]rune{
		'a': '𝕒', 'b': '𝕓', 'c': '𝕔', 'd': '𝕕', 'e': '𝕖', 'f': '𝕗', 'g': '𝕘', 'h': '𝕙', 'i': '𝕚', 'j': '𝕛', 'k': '𝕜', 'l': '𝕝', 'm': '𝕞', 'n': '𝕟', 'o': '𝕠', 'p': '𝕡', 'q': '𝕢', 'r': '𝕣', 's': '𝕤', 't': '𝕥', 'u': '𝕦', 'v': '𝕧', 'w': '𝕨', 'x': '𝕩', 'y': '𝕪', 'z': '𝕫',
		'A': '𝔸', 'B': '𝔹', 'C': 'ℂ', 'D': '𝔻', 'E': '𝔼', 'F': '𝔽', 'G': '𝔾', 'H': 'ℍ', 'I': '𝕀', 'J': '𝕁', 'K': '𝕂', 'L': '𝕃', 'M': '𝕄', 'N': 'ℕ', 'O': '𝕆', 'P': 'ℙ', 'Q': 'ℚ', 'R': 'ℝ', 'S': '𝕊', 'T': '𝕋', 'U': '𝕌', 'V': '𝕍', 'W': '𝕎', 'X': '𝕏', 'Y': '𝕐', 'Z': 'ℤ',
		'0': '𝟘', '1': '𝟙', '2': '𝟚', '3': '𝟛', '4': '𝟜', '5': '𝟝', '6': '𝟞', '7': '𝟟', '8': '𝟠', '9': '𝟡',
	}
	res := ""
	for _, c := range input {
		if val, ok := charMap[c]; ok {
			res += string(val)
		} else {
			res += string(c)
		}
	}
	return res
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("𝕗𝕒𝕚𝕝𝕖𝕕 𝕥𝕠 𝕠𝕡𝕖𝕟 𝕗𝕚𝕝𝕖 %s: %w", filePath, err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("𝕗𝕒𝕚𝕝𝕖𝕕 𝕥𝕠 𝕠𝕡𝕖𝕟 𝕗𝕚𝕝𝕖 %s: %w", filePath, err)
	}
	return content, nil
}

func writeFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("𝕗𝕒𝕚𝕝𝕖𝕕 𝕥𝕠 𝕔𝕣𝕖𝕒𝕥𝕖 𝕗𝕚𝕝𝕖 %s: %w", filePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("𝕗𝕒𝕚𝕝𝕖𝕕 𝕥𝕠 𝕨𝕣𝕚𝕥𝕖 𝕗𝕚𝕝𝕖 %s: %w", filePath, err)
	}
	return nil
}
