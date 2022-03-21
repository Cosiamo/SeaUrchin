***Still a work in progress***

# How To Use
- `go run main.go g` for Google search
- `go run main.go b` for Bing search
- If you run main.go without the `g` or `b` subcommand it will exit with error
    - this is intentional for right now
    - I plan on allowing users to set default search engines so you don't have to type in a flag every time
- If you want to view what the URL that the results were fetched from and how long the backoff time is, use the `-info` flag
- To change which domain (country or region) you want to search in, use the `-url` flag
    - example: `go run main.go g -url uk` will search on google.co.uk 

# TODO
- Run program with `sea` keyword
- Allow users to change settings
    - change domain, language, number of requests, and proxy
    - use a subcommand such as `s` or `settings` then use flags for specific settings
- Fix Google's description
    - might have to change the `descTag` and `titleTag` in the parse package
- Improve the format of the results that are being returned
- Allow for a default search engine
- Optional: add colors to the strings to differentiate different things