package main

/*
{
    "name": <value>, e.g. "Tom"
    "age": <value>   e.g. 21
}
*/

import
(
    "strings"
    "net/http"
    "fmt"
    "encoding/json"
    "os"
    //"strconv"
)

type ProjectInfo struct
{
    Owner string    `json:"login"`
    //Langs []string  `json:"age"`
}

func handlerProjectinfo (w http.ResponseWriter, r *http.Request) {
    project := ProjectInfo{""}
    parts := strings.Split(r.URL.Path, "/")
    if (len(parts) != 6) {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }
    resp, err := http.Get("http://api.github.com/repos/" + parts[4] + "/" + parts[5] + "/")
    err2 := json.NewDecoder(resp.Body).Decode(&project)
    defer resp.Body.Close()
    if (err == nil && err2 == nil) {
        fmt.Println(project)
    } else {
        fmt.Println(err)
        fmt.Println(err2)
    }
}

/*func handlerStudent (w http.ResponseWriter, r *http.Request) {
    // --------------
    db := StudentsDB{}
    // --------------
    http.Header.Add(w.Header(), "content-type", "application/json")
    //fmt.Fprintln(w, os.Getenv("PORT"))

    parts := strings.Split(r.URL.Path, "/")

    // 0
    s0 := Student{"Tom", 21}
    // 1
    s1 := Student{"Alice", 24}
    students := []Student { s0, s1, }

    // Error handling
    if (len(parts) != 3) {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }

    // Handle "/student/"
    i, err := strconv.Atoi(parts[2])
    if (err != nil) {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }

    if (parts[2] == "") {
        replyWithAllStudents(&w, &db)
        // Handle "/student/1"
    } else if (parts[2] == "1") {
        replyWithStudent(&w, &db, i)
    }
}*/

func main() {
    //s := Student{ "Tom", 21 }
    fmt.Println(os.Getenv("PORT"))
    http.HandleFunc("/projectinfo/v1/", handlerProjectinfo)
    http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
