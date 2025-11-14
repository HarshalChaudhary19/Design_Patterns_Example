package designpatterns2

import "fmt"

//Basically Agr do interfaces can't communicate with each other so Adapter works as a bridge between them
//Lets suppose Client wants to use the method from Target interface and the better implementation is in BetterTarget interface
//And Client can't talk to BetterTarget as it can only communicate to Target interface so Adaptr implements the code from the Target interface's method and then
//Adapter internally calls the BetterInterface method

//Target Interface

type FileUpload interface { //Jo call krna pdega
	Upload()
}

//Adaptee Interface
type FileTypeCSV interface {
	CSVUpload()
}

//Adaptee (Jo call krna hai)

type FileCSV struct{}

func (f *FileCSV) CSVUpload() {
	fmt.Println("File Uploaded is CSV")
}

//Adapter

type FileCSVAdapter struct {
	file *FileCSV
}

//Now using the method from the Target interface and assosiating it with the Adapter struct instance
//Iss ke method mein Adaptee ka method call kiya kyuki ye assosiated hai Adaptee wale se
func (f *FileCSVAdapter) Upload() {
	f.file.CSVUpload()
	fmt.Println("Can now upload via CSV")
}

func Adapter() {
	fileCSV := &FileCSV{}                     // filecsv create ki (adaptee ka instance)
	adapter := &FileCSVAdapter{file: fileCSV} // Adapter mein filecsv pass kri
	adapter.Upload()                          //Now Adapter se call krna hai Target wale method k
}
