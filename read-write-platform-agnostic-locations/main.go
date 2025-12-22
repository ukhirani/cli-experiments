import (
	"fmt"
	"os"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(homeDir)
	}
}
