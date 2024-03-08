/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// saveauthdataCmd represents the saveauthdata command
var saveauthdataCmd = &cobra.Command{
	Use:   "save_auth_data",
	Short: "Сохраняет пару логин пароль введенные пользователем",
	Long: `Сохраняет пару логин пароль введенные пользователем.
	При подключении наличии подключения к сети данные отправляются на хранение на сервере, 
	в ином случае харнятся на личном ПК пользователя.
	Пример использование: gophkeeper saveauthdata login_name login password`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("saveauthdata called")

		res, err := saveAuthData(args[0], args[1], args[2])
		if err != nil {
			fmt.Printf("Ошибка при сохранении данных: %s", err.Error())
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(saveauthdataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveauthdataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveauthdataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
