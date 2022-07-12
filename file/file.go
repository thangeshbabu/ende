package file

import (
	"aes_encrypt/global"
	"aes_encrypt/utils"
	"os"
)

func Read_file(path string) []byte {
	data, err := os.ReadFile(path)
	utils.Error_check(err)
	return data
}

func Create_file(path string, filename string, data []byte) {
	if utils.If_path_exist(path) {
		str := path + global.PATH_SEPERATOR + filename + ".enc"
		err := os.WriteFile(str, data, 0700)
		utils.Error_check(err)
		global.Infolog.Printf("File created %s", str)

	}
}
