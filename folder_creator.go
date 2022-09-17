package main

import ("os"; "path/filepath"; "strconv")

func main() {
	var path string
	path = "Path name"
	folder_creator(path)
}

func folder_creator(path string) {
	for i := 1; i < 15; i++ {
		newpath := filepath.Join(path, "Week_" + strconv.Itoa(i))
		if _, err := os.Stat(newpath); os.IsNotExist(err) {
			os.Mkdir(newpath, 0777)
		}
	}
}