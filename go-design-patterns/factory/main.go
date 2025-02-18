package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "Tymee", 50, "The Tymee")
	mag2, _ := newPublication("magazine", "Lyfe", 40, "Lyfe inc")
	news, _ := newPublication("newspaper", "The opinion", 60, "The opinion")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news)
}

func pubDetails(pub iPublication) {
	fmt.Println("--------------------")
	fmt.Printf("%s \n", pub)
	fmt.Printf("Type: %s\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
}

type iPublication interface {
	setName(name string)
	setPages(pages int)
	setPublisher(publisher string)
	getName() string
	getPages() int
	getPublisher() string
}

type publication struct {
	name      string
	pages     int
	publisher string
}

func (f *publication) setName(name string) {
	f.name = name
}

func (f *publication) setPages(pages int) {
	f.pages = pages
}

func (f *publication) setPublisher(publisher string) {
	f.publisher = publisher
}

func (f *publication) getName() string {
	return f.name
}

func (f *publication) getPages() int {
	return f.pages
}

func (f *publication) getPublisher() string {
	return f.publisher
}

func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {

	if pubType == "newspaper" {
		return createNewsPaper(name, pg, pub), nil
	}

	if pubType == "magazine" {
		return createMagazine(name, pg, pub), nil
	}

	return nil, fmt.Errorf("Invalid publication type")
}

func createNewsPaper(name string, pages int, pubType string) iPublication {
	return &newsPaper{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: pubType,
		},
	}
}

func createMagazine(name string, pages int, pubType string) iPublication {
	return &magazine{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: pubType,
		},
	}
}

type magazine struct {
	publication
}

func (m magazine) String() string {
	return fmt.Sprintf("This is a magazin named %s\n", m.name)
}

type newsPaper struct {
	publication
}

func (n newsPaper) String() string {
	return fmt.Sprintf("The is a newspaper name %s\n", n.name)
}
