/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "Отображает данные сохраненных карт с указанным именем",
	Long: `При вызове отображает данные карты сохраненных с указаным именем.
	При наличии подключения к интернету, данные будут браться из удаленного сервера`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("card called")

		res, err := getCard(args[0])
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
		}
		fmt.Printf("\nCard name: %s \n\tNumber: %s \n\tDate: %s\n\tCVV: %v\n",
			res.Name, res.Number, res.Date, res.CVVCode)
	},
}

func init() {
	rootCmd.AddCommand(cardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
