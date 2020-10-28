package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

var booksMap = struct{
	sync.RWMutex
	m map[int]Book

}{m:make(map[int]Book)}

func init(){
	fmt.Print("Loading Books from JSON...")
	bookMap,err :=loadBooks()
	booksMap.m=bookMap
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Loaded books from JSON")

}

func loadBooks() (map[int]Book, error){
	filename :="books.json"
	_, err :=os.Stat(filename)
	if os.IsNotExist(err){
		return nil,fmt.Errorf("file %s does not exist",filename)
	}

	file, _ :=ioutil.ReadFile(filename)
	bookList :=make([]Book, 0)
	err =json.Unmarshal(file,&bookList)
	if err!=nil{
		log.Fatal(err)
		//return nil, fmt.Errorf("Error unmarshaling data from %s file", filename)
	}
	bookMap:=make(map[int]Book)
	for i:=0;i<len(bookList);i++ {
		bookMap[bookList[i].ID] = bookList[i]
	}
	return bookMap,nil

}

func getBook(ID int) *Book{
	booksMap.RLock()
	defer booksMap.RUnlock()
	if book,ok :=booksMap.m[ID]; ok{
		return &book
	}
	return nil
}

func getBookByName(name string) *Book{
	booksMap.RLock()
	defer booksMap.RUnlock()

	for _,value :=range booksMap.m {
		if value.BookName == name {
			return &value
		}
	}
	return nil
}
func getAllBooks() []Book{
	booksMap.RLock()
	defer booksMap.RUnlock()

	bookList :=make([]Book,0)
	for _, book := range booksMap.m{
		bookList = append(bookList, book)
	}
	return bookList
}

func removeBook(ID int){
	booksMap.Lock()
	defer booksMap.Unlock()
	delete(booksMap.m,ID)
}

func getBookIDs()[]int{
	booksMap.RLock()
	defer booksMap.RUnlock()
	bookIDs :=make([]int,0)
	for id:= range booksMap.m {
		bookIDs = append(bookIDs,id)
	}
	sort.Ints(bookIDs)
	return bookIDs
}

func getNextID() int{
	bookIDs :=getBookIDs()
	lastID :=bookIDs[len(bookIDs)-1]
	return lastID +1
}

func addOrUpdateBook(book Book)(int,error){
	addOrUpdateID :=-1
	if book.ID >0{
		oldBook :=getBook(book.ID)
		if oldBook ==nil{
			return addOrUpdateID,fmt.Errorf("book with id %d does not exist",book.ID)
		}
		addOrUpdateID=book.ID
	}else{
		addOrUpdateID=getNextID()
		book.ID=addOrUpdateID
	}
	booksMap.Lock()
	defer booksMap.Unlock()
	booksMap.m[addOrUpdateID]=book
	return addOrUpdateID, nil
}