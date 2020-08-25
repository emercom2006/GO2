package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

type TaskPage struct {
	Title string
	Posts Posts
}
type DetailedPage struct {
	Title string
	Post  PostType
}
type PostType struct {
	ID     int
	Header string
	Text   string
}

type Posts []PostType

func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	lists, err := GetAllLists()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}
	file, _ := os.Open("./www/index.html")
	data, _ := ioutil.ReadAll(file)
	tmpl := template.Must(template.New("PostList").Parse(string(data)))
	tmpl.ExecuteTemplate(w, "PostList", lists)

	if err := tmpl.ExecuteTemplate(w, "PostList", lists); err != nil {
		log.Println(err)
	}

}

func (TaskPage *TaskPage) GetDetailHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./les3/www/detailPage.html")
	data, _ := ioutil.ReadAll(file)

	PostID, _ := strconv.Atoi(r.URL.Query().Get("ID"))
	dPost := TaskPage.Posts[PostID-1]
	DPage := DetailedPage{
		Title: dPost.Header,
		Post:  dPost,
	}
	templateMain := template.Must(template.New("Detail").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "Detail", DPage)
}
func (TaskPage *TaskPage) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./les3/www/PostPage.html")
	data, _ := ioutil.ReadAll(file)

	templateMain := template.Must(template.New("Post").Parse(string(data)))
	templateMain.ExecuteTemplate(w, "Post", page)
}

func (TaskPage TaskPage) NewPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("New Post!")
	fmt.Println(r.Form)
	nHead := r.Form["Header"][0]
	nText := r.Form["text"][0]

	newID := len(page.Posts) + 1
	NewPost := PostType{
		ID:     newID,
		Text:   nText,
		Header: nHead,
	}
	page.Posts = append(page.Posts, NewPost)
	http.Redirect(w, r, "/", 301)
}

var page = TaskPage{
	Title: "Новостной портал",
	Posts: Posts{},
}

// GetAllLists — получение всех списков с задачами
func GetAllLists() ([]PostType, error) {
	res := []PostType{}

	rows, err := database.Query("select * from task_tracker.test")
	if err != nil {
		return res, err
	}

	defer rows.Close()
	for rows.Next() {
		list := PostType{}

		err := rows.Scan(&list.ID, &list.Header, &list.Text)
		if err != nil {
			log.Println(err)
			continue
		}

		res = append(res, list)
	}

	return res, nil
}

var database *sql.DB

// GetList — получение списка по id
func GetList(ID int) (PostType, error) {
	list := PostType{}

	row := database.QueryRow(fmt.Sprintf("select * from task_tracker.test where id = %v", ID))
	err := row.Scan(&list.ID, &list.Header, &list.Text)
	if err != nil {
		return list, err
	}

	rows, err := database.Query(fmt.Sprintf("select * from task_tracker.test WHERE id = %v", ID))
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Posts{}

		err := rows.Scan(&task)
		if err != nil {
			log.Println(err)
			continue
		}

		task = append(task, list)
	}

	return list, nil
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/task_tracker")
	if err != nil {
		panic(err)
	}
	database = db
	rows := database.QueryRow(fmt.Sprintf("select Header from task_tracker.test WHERE id = %v", 1))
	defer database.Close()

	fmt.Sprint(rows)

	route := chi.NewRouter()
	route.Route("/", func(r chi.Router) {
		r.Get("/", GetIndexHandler)
		r.Get("/detail/", page.GetDetailHandler)
		r.Get("/post", page.GetPostHandler)
		r.Post("/post/newpost", page.NewPostHandler)
	})

	http.ListenAndServe(":8090", route)

}
