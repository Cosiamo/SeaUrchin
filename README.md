***Still a work in progress***

# How To Use
- `go run main.go g` for Google search
- `go run main.go b` for Bing search
- If you run main.go without the `g` or `b` subcommand/flag it will exit with error
    - this is intentional for right now
    - I plan on allowing users to set default search engines so you don't have to type in a flag every time

# TODO
- Run program with `sea` keyword
- Allow users to change settings
    - change domain, language, number of requests, and proxy
    - use a flag such as `-s` or `-settings`
- Fix Google's description
    - might have to change the `descTag` and `titleTag` in the parse package
- Improve the format of the results that are being returned