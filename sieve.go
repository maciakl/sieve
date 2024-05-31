package main

import (
    "encoding/csv"
	"bufio"
    "strings"
    "strconv"
	"flag"
	"log"
    "fmt"
	"os"
)

const version = "0.2.0"


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
    flag.StringVar(&values, "filter", "", "comma separated list of values to be filtered on")

    var ver bool
    flag.BoolVar(&ver, "version", false, "display version number and exit")

    var peek bool
    flag.BoolVar(&peek, "peek", false, "peek at the first 10 lines, ignore all other options, make output pretty, show column numbers")

    var partial bool
    flag.BoolVar(&partial, "partial", false, "use partial search instead of exact match")

    var csv bool
    flag.BoolVar(&csv, "csv", false, "assume that the file is a csv file as defined in RFC 4180 (handles quoted fields)")
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


    file, err := os.Open(path)
    if err != nil { log.Fatal(err) }
    defer file.Close()


    if peek {
        if csv {
            peekcsv(file)
        } else {
            peak(file, delimiter)
        }
        os.Exit(0)
    }

    if values != "" {
        value_list := strings.Split(values, ",")

        if csv {
            filttercsv(file, column, value_list, limit, partial)
        } else {
            filter(file, column, value_list, delimiter, limit, partial)
        }
        os.Exit(0)
    }


    if values == "" && limit > 0 {
        head(file, limit)
    }

}

func filter(file *os.File, column int, values []string, delimiter string, limit int, partial bool) {

    scanner := bufio.NewScanner(file)

    i := 1
    for scanner.Scan() {

        if limit != 0 && i >= limit { break }

        line := scanner.Text()
        cols := strings.Split(line, delimiter)
        
        if column > len(cols) { log.Fatal("the input file has ", len(cols), " column(s), you wanted to filter on column ", column) }

        if contains(values, cols[column], partial) {
            fmt.Println(line)
            i++
        }
    }
}

// filter assuming input file is in csv format as defined in RFC 4180
func filttercsv(file *os.File, column int, values []string, limit int, partial bool) {

    reader := csv.NewReader(file)

    i := 1
    for {
        cols, err := reader.Read()
        if err != nil { break }

        if column > len(cols) { log.Fatal("the input file has ", len(cols), " column(s), you wanted to filter on column ", column) }

        if contains(values, cols[column], partial) {
            fmt.Println(strings.Join(cols, ","))
            i++
        }

        if limit != 0 && i >= limit { break }
    }
}


// print the first n lines of the file (no filtering)
func head(file *os.File, n int) {
    scanner := bufio.NewScanner(file)
    i := 1
    for scanner.Scan() {
        if n != 0 && i >= n { break }
        fmt.Println(scanner.Text())
        i++
    }
}

// peak at the first 10 lines of the file
func peak(file *os.File, delimiter string) {

    scanner := bufio.NewScanner(file)

    for i := 0; i < 10; i++ {
        if scanner.Scan() {
            line := scanner.Text()
            cols := strings.Split(line, delimiter)

            if i == 0 { print_colmns(len(cols)) }

            for _,col := range cols {
                if len(col) > 18 { col = col[:15]+"..." }
                fmt.Printf("%-18s", col)
                fmt.Print("\t")
            }
            fmt.Print("\n")
        }
    }
}




// peak at the first 10 lines of the file (csv version)
func peekcsv(file *os.File) {

    reader := csv.NewReader(file)

    for i := 0; i < 10; i++ {
        cols, err := reader.Read()
        if err != nil { break }

        if i == 0 { print_colmns(len(cols)) }

        for _,col := range cols {
            if len(col) > 18 { col = col[:15]+"..." }
            fmt.Printf("%-18s", col)
            fmt.Print("\t")
        }
        fmt.Print("\n")
    }
}

// print the column headers for the peak function
func print_colmns(number int) {
    for i := 0; i < number; i++ {
        fmt.Printf("%-18s", "COLUMN " +strconv.Itoa(i))
        fmt.Print("\t")
    }
    fmt.Print("\n")
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
