
import (
	"fmt"
	"os"
)

func main() {

	targetDir := "."
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("Error reading directory '%s': %v\n", targetDir, err)
		os.Exit(1)
	}

	fmt.Printf("Contents of '%s':\n", targetDir)
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  [DIR] %s\n", entry.Name())
		} else {
			fmt.Printf("  [FIL] %s\n", entry.Name())
		}
	}
}
