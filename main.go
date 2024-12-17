package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Prompt the user to input the title and link
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter the link: ")
	link, _ := reader.ReadString('\n')
	link = strings.TrimSpace(link)

	// Replace spaces with dashes in the title to create the folder name
	folderName := strings.ReplaceAll(title, " ", "-")
	folderName = strings.ReplaceAll(folderName, ".", "")

	// Create the folder
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Printf("Failed to create folder: %s\n", err)
		return
	}

	// Define the path for the README.md file
	readmePath := filepath.Join(folderName, "README.md")

	// Define the content for the README.md file
	readmeContent := fmt.Sprintf("## %s\n\nChi tiết bài toán xem [tại đây](%s)\n\n### Bài Toán\n\n\n### Yêu Cầu\n\n\n### Phân Tích\n", title, link)

	// Create and write to the README.md file
	err = writeFile(readmePath, readmeContent)
	if err != nil {
		fmt.Printf("Failed to create README.md: %s\n", err)
		return
	}

	// Define the path for the solution.go file
	solutionPath := filepath.Join(folderName, "solution.go")

	// Define the content for the solution.go file
	solutionContent := `package main

import "fmt"

func main() {
	fmt.Println("Solution goes here!")
}
`

	// Create and write to the solution.go file
	err = writeFile(solutionPath, solutionContent)
	if err != nil {
		fmt.Printf("Failed to create solution.go: %s\n", err)
		return
	}

	fmt.Println("Folder, README.md, and solution.go have been successfully created!")
}

// writeFile is a helper function to create and write to a file
func writeFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Failed to close file: %s\n", err)
		}
	}(file)

	_, err = file.WriteString(content)
	return err
}
