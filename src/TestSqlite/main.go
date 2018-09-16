package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

var db *sql.DB

type EBookInfo struct {
	No      int
	name    string
	ISBN    string
	forsale string
	price   int
}

func Sqlite_Open() error {
	var err error
	//cache=shared&mode=rwc&loc=auto
	var format string = "%s?%s"
	file_name := fmt.Sprintf(format, "../../simpleDB/simpledb.db", "cache=shared&mode=rwc&loc=auto")
	db, err = sql.Open("sqlite3", file_name)
	if nil != err {
		//panic(err)
		fmt.Println(err)
		//get_module_logger.Loggers[get_module_logger.EERROR_LOG].Println(err)
		//return err
	}
	return err
}

func Sqlite_Close() {
	db.Close()
	db = nil
}

func testFunc() {
	var err error
	//var purge string = "purge_seq"
	//var newSeqNo string
	var tempBookList []EBookInfo

	//rows, err := db.Query("SELECT seq_no, seq_name FROM SEQUENCE_NO where seq_name = ? ", purge)

	//rows, err := db.Query("SELECT  name, ISBN, forsale, price FROM ebookInfo where forsale = ? ", "Y")
	rows, err := db.Query("SELECT  no,name, ISBN, forsale, price FROM ebookInfo")
	if err != nil {
		fmt.Println(err)
	}
	var (
		no      int
		name    string
		ISBN    string
		forsale string
		price   int
	)
	//var name string
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&no, &name, &ISBN, &forsale, &price)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		tempBookList = append(tempBookList, EBookInfo{1, name, ISBN, forsale, price})

		//	newSeqNo = seq_no

	}
	fmt.Println(tempBookList)
}
func MakeUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
func InsertTestBookData() {
	var (
		no      int = 1
		name    string
		ISBN    string
		forsale string = "Y"
		price   int    = 10000
	)
	//no++

	//fmt.Printf("UUIDv4: %s\n", u1)
	i := 0
	for i < 90 {
		i++
		no++
		name = "firstbook" + strconv.Itoa(no)
		price += no * 20

		ISBN = MakeUUID()
		_, err := db.Exec("INSERT INTO ebookInfo(no,name, ISBN, forsale, price) values(?,?,?,?,?)", no, name, ISBN, forsale, price)
		if nil != err {
			fmt.Println(err)
		}
	}
	j := 0
	for j < 10 {
		j++
		no++
		name = "notForsale" + strconv.Itoa(no)
		price += no * 20
		forsale = "N"
		ISBN = MakeUUID()
		_, err := db.Exec("INSERT INTO ebookInfo(no,name, ISBN, forsale, price) values(?,?,?,?,?)", no, name, ISBN, forsale, price)
		if nil != err {
			fmt.Println(err)
		}
	}

}
func makeBookandUserTable() {

	rows, err := db.Query("SELECT  no FROM ebookInfo where forsale = 'N'")
	if err != nil {
		fmt.Println(err)
	}
	rows.Close()
	var no int
	for rows.Next() {
		err = rows.Scan(&no)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		_, err := db.Exec("INSERT INTO userAndBook(book_no, user_no ,bought_day) values(?,?,?)", no, 1, "2018-09-01")
		if nil != err {
			fmt.Println(err)
		}
		//	newSeqNo = seq_no

	}

	rows, err = db.Query("SELECT  no FROM ebookInfo where forsale = 'Y' limit 0,3")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&no)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		_, err := db.Exec("INSERT INTO userAndBook(book_no, user_no ,bought_day) values(?,?,?)", no, 1, "2018-09-01")
		if nil != err {
			fmt.Println(err)
		}
		//	newSeqNo = seq_no

	}
	rows.Close()
}

func main() {
	Sqlite_Open()
	InsertTestBookData()
	Sqlite_Close()
	//testFunc()
	Sqlite_Open()
	makeBookandUserTable()
	Sqlite_Close()
}
