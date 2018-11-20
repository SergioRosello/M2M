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

  // Title of the post to import to Medium
  t := os.Args[1]
  // The file's path (Absolute)
  f := os.Args[2]
  // Content format.
  // Can be either:
  //  - markdown
  //  - html
  // cf := os.Args[3]
  // Publish status
  // Can be either:
  //  - public
  //  - draft
  //  - unlisted
  // ps := os.Args[4]

  // Create a new client with the token aquired from
  // https://medium.com/me/settings
  m := medium.NewClientWithAccessToken(os.Getenv("TOKEN"))
  u, err := m.GetUser("")
  check(err)

  // Set necesary parameters necesary to create a post
  // and create the post as a draft
  //TODO: parse arguments passed through parameters to this post options
  p, err := m.CreatePost(medium.CreatePostOptions{
    UserID: u.ID,
    Title: t,
    Content: string(MDFile(f)),
    ContentFormat: "markdown",
    PublishStatus: "draft",
  })
  check(err)

  // Debug log, for development purposes only
  log.Println(u, p, f)
}
