# TODO
- Fix user input so that it detects spaces 
    - tried looking up "fried chicken recipe" but it returned definitions for "fried"
    - in `buildGoogleUrls` it should be replacing " " with "+" so that the Google query can detect it
    - possibly need to rework how the variables work in logic and main packages
- Run program with `sea` keyword
- Allow users to change settings
    - change domain, language, number of requests, and proxy
    - use a flag such as `-s` or `-settings`