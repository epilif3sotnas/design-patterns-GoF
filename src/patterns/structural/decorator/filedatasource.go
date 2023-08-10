package decorator


import(
	"bufio"
	"fmt"
	"os"
)


type FileDataSource struct {
	filename string
}


func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{filename: filename}
}

func (self *FileDataSource) WriteData(data []byte) {
	file, err := os.Create(self.filename + ".txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (self *FileDataSource) ReadData() []byte {
	file, err := os.Open(self.filename + ".txt")

	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}

	defer file.Close()

	fileContent := []byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Bytes()...)
	}

	return fileContent
}

func (self *FileDataSource) GetFilename() string {
	return self.filename
}

func (self *FileDataSource) SetFilename(filename string) {
	self.filename = filename
}