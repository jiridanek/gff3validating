package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/fcgi"
	"os"
)

import (
	"./services"
	"./view"
)

func main () {
	log.SetPrefix("gff3validating ")
	log.Println("STARTED")
	fcgi.Serve(nil, http.HandlerFunc ( func (w http.ResponseWriter, r *http.Request) {
		log.Println("Got request")
		reader, err := r.MultipartReader()
		if err != nil {
			view.ErrorTmpl.Execute(w, view.Error{err, "The file was not sent or received properly"})
			log.Println(err)
			return
		}
		
		var fname string
		
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				log.Printf("EOF\n")
				break
			}
			if err != nil {
				view.ErrorTmpl.Execute(w, view.Error{err, "The file was not sent or received properly"})
				log.Println(err)
				return
			}
			switch(part.FormName()) {
				case "soubor":
					f, err := ioutil.TempFile("", "gff")
					if err != nil {
						view.ErrorTmpl.Execute(w, view.Error{err, "Could not create a temporary file"})
						log.Println(err)
					}
					fname = f.Name()
					_, err = io.Copy(f, part)
					f.Close()
				default:
					log.Printf("Unexpected FormName %v\n", part.FormName())
			}
			
// 			l/*og.Println("FormName: " + part.FormName())
// 			log.Println*/("FileName: " + part.FileName())
			part.Close()
		}
		
		if fname != "" {
			log.Println("file", fname, "created")
			results, err := services.Validate(fname)
			if err != nil {
				view.ErrorTmpl.Execute(w, view.Error{err, "Running the gt gff3validator command failed"})
			} else {
				view.ResultsTmpl.Execute(w, results)
			}
			
			err = os.Remove(fname)
			if err != nil {
				log.Println(err)
			}
		}
	}))
}

