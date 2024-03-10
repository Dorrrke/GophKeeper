/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Отображает данные сохраненой пары логин пароль с указанным именем",
	Long: `При вызове отображает пару логин пароль сохраненных с указаным именем.
	При наличии подключения к интернету, данные будут браться из удаленного сервера`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")

		login, err := getLogin(args[0])
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
		}
		fmt.Printf("\nLogin name: %s \n\tLogin: %s \n\tPassword: %s\n",
			login.Name, login.Login, login.Password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
