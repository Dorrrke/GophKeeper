/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// saveusercardCmd represents the saveusercard command
var saveusercardCmd = &cobra.Command{
	Use:   "save_card",
	Short: "Сохраняет данные банковской карты",
	Long: `Сохраняет данные банковской карты пользователя.
	При подключении наличии подключения к сети данные отправляются на хранение на сервере, 
	в ином случае харнятся на личном ПК пользователя.`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("saveusercard called")

		res, err := saveUserCard(args[0], args[1], args[2], args[3])
		if err != nil {
			fmt.Printf("Ошибка при сохранении карты: %s", err.Error())
		}

		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(saveusercardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveusercardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveusercardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
