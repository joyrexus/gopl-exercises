## 1.10

Find a web site that produces a large amount of data. Investigate caching by running `fetchall` twice in succession to see whether the reported time changes much. Do you get the same content each time? 

Modify `fetchall` to print its output to a file so it can be examined.


## 1.11

Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com. How does the program behave if a web site just doesn’t respond? Section 8.9 describes mechanisms for coping in such cases.

Try `go run server.go` and then `go run fetchall.go http://localhost:8000/noresponse`.  Note how it just waits for the reponse without timing out.  
