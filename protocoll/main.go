package protocoll

import ("fmt"
				"time"
				"io/ioutil"
				"os"
			)


func Initialize(destDir string) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	createLogDir(destDir)
	now := time.Now()
	logheader := "File-Copy - Protocol " + now.Format(layout)
	filename := destDir + "/protocol.txt"
	logheaderBytes := []byte(logheader + "\n")

	err := ioutil.WriteFile(filename, logheaderBytes, 0644)

	if err != nil {
		fmt.Println("Error creating protokoll-file: ", err)
	} else {
		fmt.Println("Created protokollfile...")
	}
}

func Success(str string, ) {
	fmt.Println("testlog", str)
}

func Failure(str string) {
	fmt.Println("failed: ", str)
}

func createLogDir(dest string) {
	err := os.MkdirAll(dest, 0755)
	if err != nil {
		fmt.Println("Error creating destination directory: ", dest)
	} else {
		fmt.Println("Created destination directory: ", dest)
	}

}
