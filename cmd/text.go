/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Отображает текстовые данные сохраненные под указанным именем.",
	Long: `При вызове отображает текстовые данные сохраненные под указанным именем.
	При наличии подключения к интернету, данные будут браться из удаленного сервера`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("text called")

		text, err := getText(args[0])
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
		}
		fmt.Printf("\nText name: %s \n\tData: %s\n",
			text.Name, text.Data)
	},
}

func init() {
	rootCmd.AddCommand(textCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
