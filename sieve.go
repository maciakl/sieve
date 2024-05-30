package main

import (
	"bufio"
    "strings"
	"flag"
	"log"
    "fmt"
	"os"
)

const version = "0.1.1"


func main() {
 // command line args
    var path string
    flag.StringVar(&path, "file", "", "path to the file to be filtered")

    var delimiter string
    flag.StringVar(&delimiter, "delimiter", ",", "the column delimiter character")
    
    var column int
    flag.IntVar(&column, "column", 0, "the column to be used for filtering, starting at 0")

    var limit int
    flag.IntVar(&limit, "limit", 0, "limit the number of filtered lines that are output (if zero, no limit)")
    
    var values string
    flag.StringVar(&values, "values", "", "comma separated list of values to be filtered on")

    var ver bool
    flag.BoolVar(&ver, "version", false, "display version number and exit")

    var head bool
    flag.BoolVar(&head, "head", false, "display first 10 lines of the file, ignore all other options")

    var partial bool
    flag.BoolVar(&partial, "partial", false, "use partial search instead of exact match")
    flag.Parse()

    // show version and exit
    if ver {
        fmt.Println("sieve version", version)
        os.Exit(0)
    }

    // bail out if there is no file
    if path == "" {
        fmt.Fprintln(os.Stderr, "no file specified")
        os.Exit(1)
    }

    value_list := strings.Split(values, ",")

    file, err := os.Open(path)
    if err != nil { log.Fatal(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    counter := 1
    if head { limit = 10 } // when head is set only display 10 first lines

    for scanner.Scan() {

        if limit != 0 && counter >= limit { break }

        line := scanner.Text()

        cols := strings.Split(line, delimiter)
        
        if head { // if head is set just print out every line
                fmt.Println(line)
                counter++
        } else {

            if column > len(cols) { log.Fatal("the input file has ", len(cols), " column(s), you wanted to filter on column ", column) }

            if contains(value_list, cols[column], partial) {
                fmt.Println(line)
                counter++
            }
        }
    }

}

// returns false if value is not contained in list
// returns true if value is in the list array
// if partial is true, checks if value is a substring of any element of list
func contains(list []string, value string, partial bool) bool {

    for _,s := range list {

        // partial search
        if partial {
            if strings.Contains(value, s) { return true }
        } else { // exact match
            if value == s { return true } 
        }
    }
    return false
}
