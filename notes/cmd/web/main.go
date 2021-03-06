package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"os"
)

type application struct{
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr",":4000","Сетевой адрес HTTP")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	srv := &http.Server{
		Addr:		*addr,
		ErrorLog: 	errorLog,
		Handler:	app.routes(),
	}

	infoLog.Printf("Запуск веб-сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

type nFileSystem struct {
	fs http.FileSystem
}

func (nfs nFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}