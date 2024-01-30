/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// Moves file to directory
func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName = "dr-who"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}

		if len(args) > 1 && args[1] != "" {
			which := args[1]
			fmt.Println(which)
		}

		URL := "https://raw.githubusercontent.com/scraly/gophers/main/" + gopherName + ".png"
		fmt.Println(URL)
		fmt.Println("Try to get '" + gopherName + "' Gopher...")

		// Get the data
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Create the file
			out, err := os.Create(gopherName + ".png")
			if err != nil {
				fmt.Println(err)
			}

			// Check if a file exists
			if _, err := os.Stat("./harry-gopher.png"); err == nil {
				fmt.Println("File " + gopherName + ".png exists!")
				if err := os.Mkdir("gopher", os.ModePerm); err != nil {
					log.Fatal(err)
				}

				// Move .png file to directory
				sourcePath := "./" + gopherName + ".png"
				destPath := "./gopher"
				// sourcePath := "./harry-gopher.png"
				// destPath := "./gopher"
				fmt.Println(sourcePath + " and " + destPath)
				MoveFile(sourcePath, destPath)

			} else {
				fmt.Println("File does not exists!")
			}

			defer out.Close()

			// Writer the body to file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Perfect! Just saved in " + out.Name() + "!")
		} else {
			fmt.Println("Error: " + gopherName + " not exists! :-(")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
