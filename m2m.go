package main

import (
  "os"
  "io/ioutil"
  medium "github.com/medium/medium-sdk-go"
  "log"
)

// Extract the file from memory
// given it's location
func MDFile(s string) []byte {
  dat, err := ioutil.ReadFile(s)
  check(err)
  return dat
}

// Default error handler
func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  // Read first argument,
  // The file's path (Absolute)
  f := os.Args[1]

  // Create a new client with the token aquired from
  // https://medium.com/me/settings
  m := medium.NewClientWithAccessToken(os.Getenv("TOKEN"))
  u, err := m.GetUser("")
  check(err)

  // Set necesary parameters necesary to create a post
  // and create the post as a draft
  p, err := m.CreatePost(medium.CreatePostOptions{
    UserID: u.ID,
    Title: "Post title",
    Content: string(MDFile(f)),
    ContentFormat: "markdown",
    PublishStatus: "draft",
  })
  check(err)

  // Debug log, for development purposes only
  log.Println(u, p, f)
}
