# sieve

Command line tool to easily filter delimited text files on a couple of values. Reads in a file, outputs to sdtdin.

      Usage:
      -column int
            the column to be used for filtering, starting at 0
      -delimiter string
            the column delimiter character (default ",")
      -file string
            path to the file to be filtered
      -head
            display first 10 lines of the file, ignore all other options
      -limit int
            limit the number of filtered lines that are output (if zero, no limit)
      -partial
            use partial search instead of exact match
      -values string
            comma separated list of values to be filtered on
      -version
            display version number and exit

Use `-head` to peak at the first 10 lines of the file.

## Installing

 On Windows, this tool is distributed via `scoop` (see [scoop.sh](https://scoop.sh)).

 First, you need to add my bucket:

    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update

 Next simply run:
 
    scoop install sieve

If you don't want to use `scoop` you can simply download the executable from the release page and extract it somewhere in your path.
