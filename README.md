# words
Simple helper package for strings and words.

# English word set
I found this awesome github repo [https://github.com/dwyl/english-words](https://github.com/dwyl/english-words) from which i downloaded words.txt containing huge amount of english words.

## Features
- get text length
- reverse text
- check for palindrome (case-sensitive or case-insensitive) (word that reads same when written backwards, like "bob")
- check for anagram (two different words contain same letters in different order, example: "I am Lord Voldemort" is an anagram of "Tom Marvolo Riddle")
- check for pangram (text contains all 26 letters atleast once)
- check and count vowels or consonants

### Planned features
- find all words that can be written with given letters (for example: with letters "tra" you can write (atleast) words "rat" and "tar"
- check for isogram (word does not contain same letter more than once, example: "dermatoglyphics")
- check for lipogram (text cannot contain selected letters)

## Usage
This package have not been released as go package so follow instructions below on how to use this.

To use this package, download or clone this repo next to golang project where you are going to use this. For example, if you have folder named "go-projects" and in there you have golang project "project-a" and you want to add this package to it you should have this package next to "project-a" like following:

```
├── go-projects
│   ├── words
│   ├── project-a
```

Then at root of project-a folder run command
```
go mod edit -replace github.com/Korppi/words=../words
```
After this run go mod tidy command
```
go mod tidy
```
Your go.mod file should have something similar in it:
```
replace github.com/Korppi/words => ../words

require github.com/Korppi/words v0.0.0-00010101000000-000000000000
```
In case there are problems, you can refer to official documentation [https://go.dev/doc/tutorial/call-module-code](https://go.dev/doc/tutorial/call-module-code)

Example of usage of words package 
```
package main

import (
    "fmt"

    "github.com/Korppi/words"
)

func main() {
    message := words.Reverse("123")
    fmt.Println(message)
}
```
This should output following:
```
321
```
