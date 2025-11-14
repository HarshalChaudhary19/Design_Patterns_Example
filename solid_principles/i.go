package solidprinciples

// Interface Segregation Principle (ISP)
//Clients should not be depend on interfaces that they do not require
//Sab type ke methods ek he interface mein nhi

//In the below example the if we require a read only document we still need to implement all the methods

type Document interface {
	Read() string
	Write(content string)
	Print() string
}

type TextDocument struct {
	content string
}

func (d *TextDocument) Read() string {
	return d.content
}

func (d *TextDocument) Write(content string) {
	d.content = content
}

func (d *TextDocument) Print() string {
	return "Printing: " + d.content
}

//Ye hai read only document
type ReadOnlyDocument struct {
	content string
}

func (d *ReadOnlyDocument) Read() string {
	return d.content
}

func (d *ReadOnlyDocument) Write(content string) {
	// Not supported
}

func (d *ReadOnlyDocument) Print() string {
	// Not supported
}

//So now

type Reader interface {
	Read() string
}

type Writer interface {
	Write(content string)
}

type Printer interface {
	Print() string
}

type TextDocumentt struct {
	content string
}

//Now we can Implement Reader, Writer, and Printer for TextDocument

type ReadOnlyDocumentt struct {
	content string
}
