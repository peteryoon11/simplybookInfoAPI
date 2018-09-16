package dbConnectModule

import (
	"database/sql"
	"fmt"
	"time"

	"../structModule"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Sqlite_Open() error {
	var err error
	//cache=shared&mode=rwc&loc=auto
	var format string = "%s?%s"
	file_name := fmt.Sprintf(format, "../simpleDB/simpledb.db", "cache=shared&mode=rwc&loc=auto")
	//file_name := fmt.Sprintf(format, "../simpleDB/simpledb.db", "cache=shared")
	db, err = sql.Open("sqlite3", file_name)
	if nil != err {
	}
	return err
}

func Sqlite_Close() {
	db.Close()
	db = nil
}
func GetAPIKeyValiation(valAuth structModule.ValAPIKey) []structModule.ValAPIKey {
	t := time.Now()
	//tempData := strconv.Itoa(t.Year()) + "-" + strconv.Itoa(t.Month()) + "-" + strconv.Itoa(t.Day())
	var timeFormat string = "%s-%s-%s"

	tempData := fmt.Sprintf(timeFormat, t.Year(), t.Month(), t.Day())
	tempTimeStamp := t.Format("2006-01-02150405")[:10]
	fmt.Println(t.Format("20060102150405"))
	fmt.Println(tempTimeStamp)
	fmt.Println("now is =>  " + tempData)
	rows, err := db.Query("SELECT  id, AuthKey FROM userInfo where id = ? and AuthKey = ? and valDate > ? ", valAuth.ID, valAuth.APIKey, tempTimeStamp)
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	var (
		//	no      int
		id      string
		AuthKey string
	)
	//var name string
	defer rows.Close()
	var tempCount []structModule.ValAPIKey

	for rows.Next() {
		err = rows.Scan(&id, &AuthKey)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		tempCount = append(tempCount, structModule.ValAPIKey{id, AuthKey})
		//	newSeqNo = seq_no

	}
	return tempCount
}
func GetAllBookInfo() ([]structModule.EBookInfo, error) {
	var err error
	//var purge string = "purge_seq"
	//var newSeqNo string
	var tempBookList []structModule.EBookInfo

	//rows, err := db.Query("SELECT seq_no, seq_name FROM SEQUENCE_NO where seq_name = ? ", purge)

	rows, err := db.Query("SELECT  name, ISBN, forsale, price FROM ebookInfo ")
	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	var (
		//	no      int
		name    string
		ISBN    string
		forsale string
		price   int
	)
	//var name string
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&name, &ISBN, &forsale, &price)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		tempBookList = append(tempBookList, structModule.EBookInfo{1, name, ISBN, forsale, price})

		//	newSeqNo = seq_no

	}

	return tempBookList, err
}
func GetMyOwnBook(valAuth structModule.ValAPIKey) ([]structModule.EBookInfo, error) {
	var err error
	//var purge string = "purge_seq"
	//var newSeqNo string
	var tempBookList []structModule.EBookInfo

	//rows, err := db.Query("SELECT seq_no, seq_name FROM SEQUENCE_NO where seq_name = ? ", purge)

	//rows, err := db.Query("SELECT  eb.name, eb.ISBN, eb.forsale, eb.price, ua.user_no as user_no FROM ebookInfo as eb left join userAndBook as ua on eb.no=ua.book_no where ua.user_no = ?  ", 1)
	Sqlite_Open()
	rows, err := db.Query("SELECT  eb.name, eb.ISBN, eb.forsale, eb.price FROM ebookInfo as eb left join userAndBook as ua on eb.no=ua.book_no where ua.user_no = (select no from userInfo where id = ?) ", valAuth.ID)

	if err != nil {
		fmt.Println("error in GetBookInfo")
		fmt.Println(err)
	}
	var (
		//	no      int
		name    string
		ISBN    string
		forsale string
		price   int
	)
	//var name string
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&name, &ISBN, &forsale, &price)
		if err != nil {
			fmt.Println("inner rows next")
			fmt.Println(err)
		}
		tempBookList = append(tempBookList, structModule.EBookInfo{1, name, ISBN, forsale, price})

		//	newSeqNo = seq_no

	}
	Sqlite_Close()
	return tempBookList, err
}
