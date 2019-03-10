package MakeDirsAndFiles


import "os"
import "fmt"


func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os. MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func CreateFile(path string) (error){

	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil { return err} 
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)
	return nil
}
