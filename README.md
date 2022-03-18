***Still a work in progress***

# How To Use
- `go run main.go g` for Google search
- `go run main.go b` for Bing search
- If you run main.go without the `g` or `b` subcommand it will exit with error
    - this is intentional for right now
    - I plan on allowing users to set default search engines so you don't have to type in a flag every time
- You can view what the URL that the results were fetched from by using the `-url` flag
- If you want to see how long the backoff time is, use the `-backoff` flag
    - example: `go run main.go g -url` then once you search what you're looking for the link will be displayed at the end

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