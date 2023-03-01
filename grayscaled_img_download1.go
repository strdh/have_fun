package main

import (
    "fmt"
    "os"
    "time"
    "image"
    "image/jpeg"
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
        "https://unsplash.com/photos/glRqyWJgUeY/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8OTJ8fHRlY2h8ZW58MHx8fHwxNjc3NjE2ODM3&force=true&w=2400",
        "https://unsplash.com/photos/FO7JIlwjOtU/download?force=true&w=2400",
        "https://unsplash.com/photos/FWoq_ldWlNQ/download?ixid=MnwxMjA3fDB8MXxhbGx8fHx8fHx8fHwxNjc3NjU2Njk5&force=true&w=2400",
        "https://unsplash.com/photos/p0j-mE6mGo4/download?ixid=MnwxMjA3fDB8MXxhbGx8fHx8fHx8fHwxNjc3NjU5NDQ1&force=true&w=2400",
        "https://unsplash.com/photos/VvTVkc_p-eg/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8NDJ8fGNhdHxlbnwwfHx8fDE2Nzc1NTQ0OTI&force=true&w=640",
        "https://unsplash.com/photos/uMfSHeycnYQ/download?ixid=MnwxMjA3fDB8MXxhbGx8fHx8fHx8fHwxNjc3NjAzMTA3&force=true&w=640",
        "https://unsplash.com/photos/P86-JPbDnPY/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8Nnx8Zmxvd2Vyc3xlbnwwfHx8fDE2Nzc1ODM0NTc&force=true&w=640",
        "https://unsplash.com/photos/EUsVwEOsblE/download?ixid=MnwxMjA3fDB8MXxhbGx8fHx8fHx8fHwxNjc3NjU1ODk2&force=true&w=2400",
        "https://unsplash.com/photos/Skf7HxARcoc/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8Mjh8fHRlY2h8ZW58MHx8fHwxNjc3NjE0OTMw&force=true&w=2400", 
        "https://unsplash.com/photos/jXd2FSvcRr8/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8NDV8fHRlY2h8ZW58MHx8fHwxNjc3NjUyMzMx&force=true&w=2400",
        "https://unsplash.com/photos/e31ANd1PXUw/download?ixid=MnwxMjA3fDB8MXxzZWFyY2h8MTA1fHx0ZWNofGVufDB8fHx8MTY3NzY1OTY2MQ&force=true&w=2400",
    }

    ch := make(chan bool)

    for i, url := range urls {
        // fmt.Printf("Downloading and converting image to grayscale.. %d\n", i+1)
         go func(iter int, url string) {
            response, err := http.Get(url)
            if err != nil {
                panic(err)
            }
            defer response.Body.Close()
            
            img, _, err := image.Decode(response.Body)
            if err != nil {
                panic(err)
            }

            gray := image.NewGray(img.Bounds())
            for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
                for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
                    gray.Set(x, y, img.At(x, y))
                }
            }

            file, err := os.Create(fmt.Sprintf("images/image%d.jpg", iter+1))
            if err != nil {
                panic(err)
            }
            defer file.Close()

            jpeg.Encode(file, gray, nil)

            ch <- true
         }(i, url)
    }
    
    for i := 0; i < len(urls); i++ {
        <-ch
    }

    fmt.Println("All images downloaded")

    end := time.Now()
    fmt.Println("End at ", end.Format("15:04:05"))
    fmt.Println("Total time: ", end.Sub(start))
}