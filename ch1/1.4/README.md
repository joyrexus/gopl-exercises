# 1.4

Modify `dup2` to print the names of all files in which each duplicated line occurs.

For example, running `go run dup.go a.txt b.txt` should output:

    "beta beta" occurred 2 times
      b.txt: 2 times
    "alpha beta" occurred 3 times
      a.txt: 2 times
      b.txt: 1 times
    "alpha beta gamma" occurred 6 times
      a.txt: 3 times
      b.txt: 3 times

