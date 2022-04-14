# What is SeaUrchin?
SeaUrchin is an Open-Source application that lets you search either Google or Bing from the terminal.

***Why is it called SeaUrchin?***

Because it sounds like 'searchin' lol.

# How to Install
On Windows click the "code" dropdown menu, then click "download ZIP". Once that's done extract the `.exe` file

Will have Linux and MacOS downloads available for next commit. For now, if you have Golang installed, run:
```
git clone https://github.com/Cosiamo/SeaUrchin.git && cd SeaUrchin
go build
```

# How to use SeaUrchin
- `sea g` for Google search
- `sea b` for Bing search
- If you want to view what the URL that the results were fetched from and how long the backoff time is, use the `-info` flag
- To change which domain (country or region) you want to search in, use the `-url` flag
    - example: `sea g -url uk` will search on google.co.uk
- To view all supported domains/regions, type `sea domains g` for Google or `sea domains b` for Bing
