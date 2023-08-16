package connections

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadENV() {

	if godotenv.Load() != nil {
		fmt.Println("Faild to load env")
		return
	}
	fmt.Println("Successfully loaded env")

}
