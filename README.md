# What is SeaUrchin?
SeaUrchin is an Open-Source application that lets you search either Google or Bing from the terminal.

***Why is it called SeaUrchin?*** Because it sounds like 'searchin' lol.

Check out [my blog post](https://cosiamo.hashnode.dev/search-on-google-or-bing-from-the-terminal-with-seaurchin) for more info.

# How to Install
On Windows click the "code" dropdown menu, then click "download ZIP". Once that's done extract the `sea.exe` file

For Linux, do the same process as Windows except extract the `sea` file.

I don't have a Mac so I'm not able to test if it will work on MacOS or not. If you have Go installed on your Mac you can still build this application, just run the commands:
```
git clone https://github.com/Cosiamo/SeaUrchin.git && cd SeaUrchin
go build
```
I know Go allows you to cross compile applications, however I don't want to have a file in my repo that I haven't tested myself.

# How to use SeaUrchin
- `sea g` for Google search
- `sea b` for Bing search
- If you want to view what the URL that the results were fetched from and how long the backoff time is, use the `-info` flag
- To change which domain (country or region) you want to search in, use the `-url` flag
    - example: `sea g -url uk` will search on google.co.uk
- To view all supported domains/regions, type `sea domains g` for Google or `sea domains b` for Bing
