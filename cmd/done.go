/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/siuyin/cobratry/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark item as done",
	Run:   doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(args[0], "is not a valid label\n", err)
	}

	if i <= 0 || i > len(items) {
		log.Println(i, "does not match any items")
		return
	}

	items[i-1].Done = true
	fmt.Printf("%q %v\n", items[i-1].Text, "marked done")

	todo.SaveItems(dataFile, items)
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
