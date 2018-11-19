# Markdown to Medium importer
This is a very simple script that speeds up my workflow.  
It simply uploads to my account a markdown file.

## Setup

Follow these steps to setup the project in a breeze.

1. `git clone https://github.com/SergioRosello/M2M.git yourRepoName`
2. `cd yourRepoName`
3. `go get github.com/Medium/medium-sdk-go`
4. `export TOKEN=yourPersonalToken`

To get your personal token, go to [profile > settings](https://medium.com/me/settings) > integration tokens. There enter a description for you're token and press on **Get integration token**. Be careful, use this token as if it was a password. Don't share or upload it anywhere.

## Usage

1. In the project root directory, enter `go install` to install the program into your default [$GOPATH](https://golang.org/doc/code.html#GOPATH) 
2. Pass the path of the file to import into [Medium](https://medium.com/) as the first argument to the program. `M2M path/to/file/to/upload.md`
