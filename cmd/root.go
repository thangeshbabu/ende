/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"aes_encrypt/file"
	"aes_encrypt/global"
	"aes_encrypt/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var encrypt, decrypt bool = false, false

var rootCmd = &cobra.Command{
	Use:   "ende",
	Short: "Command-line tool used to encrypt and decrypt file using AES algorithm.",
	Long:  `Custom Command-line tool to Encrypt/Decrypt files on the go.`,
	Run: func(cmd *cobra.Command, args []string) {

		if !utils.If_path_exist(global.Filepath) {
			fmt.Printf("%t %t\n", encrypt, decrypt)
			global.Errorlog.Printf("Path < %s > doesn't exist !! ", global.Filepath)
			os.Exit(1)
		}

		file_name := filepath.Base(global.Filepath)

		data := file.Read_file(global.Filepath)
		cwd, err := os.Getwd()
		utils.Error_check(err)

		if encrypt {
			cipher_txt := utils.Encrypt(data, global.Passphrase)
			file.Create_file(cwd, file_name, cipher_txt)
		}

		if decrypt {
			clear_txt := utils.Decrypt(data, global.Passphrase)
			file.Create_file(cwd, file_name, clear_txt)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	utils.Logging()

	rootCmd.PersistentFlags().BoolVarP(&encrypt, "encrypt", "e", false, "flag for encryption")
	rootCmd.PersistentFlags().BoolVarP(&decrypt, "decrypt", "d", false, "flag for decryption")
	rootCmd.MarkFlagsMutuallyExclusive("encrypt", "decrypt")

	rootCmd.PersistentFlags().StringVarP(&global.Passphrase, "passphrase", "p", "", "Passphrase for Encryption/Decryption")
	rootCmd.MarkPersistentFlagRequired("passphrase")

	rootCmd.PersistentFlags().StringVarP(&global.Filepath, "filepath", "f", "", "Path to the File")
	rootCmd.MarkPersistentFlagRequired("filepath")

}
