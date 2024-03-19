package maps

import "fmt"

func maps() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}
	fmt.Println(websites["Amazon Web Services"])
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)
}
