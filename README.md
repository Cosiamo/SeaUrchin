# What is SeaUrchin?
SeaUrchin is an Open-Source application that lets you search either Google or Bing from the terminal.

***Why is it called SeaUrchin?***

Because it sounds like 'searchin' lol. 

# How to use SeaUrchin
- `go run main.go g` for Google search
- `go run main.go b` for Bing search
- If you run main.go without the `g` or `b` subcommand it will exit with error
    - this is intentional for right now
    - I plan on allowing users to set default search engines so you don't have to type in a flag every time
- If you want to view what the URL that the results were fetched from and how long the backoff time is, use the `-info` flag
- To change which domain (country or region) you want to search in, use the `-url` flag
    - example: `go run main.go g -url uk` will search on google.co.uk 
