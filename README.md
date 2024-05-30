## sieve

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
