/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// savetextdataCmd represents the savetextdata command
var savetextdataCmd = &cobra.Command{
	Use:   "save_text_data",
	Short: "Сохраняет текстовые данные пользователя.",
	Long: `Сохраняет текстовые данные пользователя введенные при запуске комманды
	или данные из файла при использовании флага file.
	При подключении наличии подключения к сети данные отправляются на хранение на сервере, 
	в ином случае харнятся на личном ПК пользователя.
	Пример использование: 
	1) gophkeeper save_text_data data_name text_data
	2) gophkeeper save_text_data data_name path/to/text/file --file`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("savetextdata called")

		filePath, err := cmd.Flags().GetBool("file")
		if err != nil {
			fmt.Printf("Ошибка при получении флага: %s", err.Error())
		}
		if filePath {
			data, err := parseFromeFile(args[1])
			if err != nil {
				fmt.Printf("Ошибка при получении данных из файла: %s", err.Error())
				return
			}
			res, err := saveTextData(args[0], data)
			if err != nil {
				fmt.Printf("Ошибка при сохранении данных: %s", err.Error())
			}
			fmt.Println(res)
		} else {
			res, err := saveTextData(args[0], args[1])
			if err != nil {
				fmt.Printf("Ошибка при сохранении данных: %s", err.Error())
			}
			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(savetextdataCmd)
	savetextdataCmd.Flags().Bool("file", false, "Файл с текстовыми данными для сохранения")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// savetextdataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// savetextdataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func parseFromeFile(filePath string) (string, error) {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist")
		}
	}
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	dataBuf := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		dataBuf.WriteString(sc.Text())
	}
	return dataBuf.String(), err
}
