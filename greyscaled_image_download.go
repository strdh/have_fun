package main

import (
    "fmt"
    "os"
    "io"
    "time"
    "net/http"
)

func main() {
    start := time.Now()
    fmt.Println("Start at ", start.Format("15:04:05"))

    urls := []string{
        "https://unsplash.com/photos/YbM9RLFHtQg/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8OXx8dGFua3xlbnwwfHx8fDE2Nzc2MDE5MDc&force=true&w=640",
        "https://unsplash.com/photos/1rEZbPBtb9A/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8NDZ8fHRhbmt8ZW58MHx8fHwxNjc3NTQ5OTUx&force=true&w=640",
        "https://unsplash.com/photos/an6PpMhewTY/download?force=true&w=640",
        "https://unsplash.com/photos/2WOXDZaXKQI/download?ixid=MnwxMjA3fDB8MXxhbGx8fHx8fHx8fHwxNjc3NjAyOTI1&force=true&w=640",
        "https://unsplash.com/photos/p6yH8VmGqxo/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8MTJ8fGNhdHxlbnwwfHx8fDE2Nzc1OTkzMzg&force=true&w=640",
        "https://unsplash.com/photos/VvTVkc_p-eg/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8NDJ8fGNhdHxlbnwwfHx8fDE2Nzc1NTQ0OTI&force=true&w=640",
        "https://unsplash.com/photos/uMfSHeycnYQ/download?ixid=MnwxMjA3fDB8MXxhbGx8fHx8fHx8fHwxNjc3NjAzMTA3&force=true&w=640",
        "https://unsplash.com/photos/P86-JPbDnPY/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8Nnx8Zmxvd2Vyc3xlbnwwfHx8fDE2Nzc1ODM0NTc&force=true&w=640",
    }

    ch := make(chan bool)

    for i, url := range urls {
         go func() {
            fmt.Printf("Downloading image %d\r", i+1)
            response, err := http.Get(url)
            if err != nil {
                panic(err)
            }
            defer response.Body.Close()

            file, err := os.Create(fmt.Sprintf("image%d.jpg", i+1))
            if err != nil {
                panic(err)
            }
            defer file.Close()

            _, err = io.Copy(file, response.Body)
            if err != nil {
                panic(err)
            }

            ch <- true
            fmt.Printf("[%d]Finished               \r", i+1)
            fmt.Println("")
         }()
         <- ch
    }
    
    end := time.Now()
    fmt.Println("End at ", end.Format("15:04:05"))
    fmt.Println("Total time: ", end.Sub(start))
}