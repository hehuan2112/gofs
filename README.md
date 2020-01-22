# Go simple File Server

- Having some files to share with others?
- Loading some files from local via http?
- Showing some files in local network?

Using this **gofs**!

It's a very very simple file server!

# Usage

First, build:

    $ go build gofs.go

Then, use the following command to start a http server listening 8000:

    $ ./gofs

or with specified port number:

    $ ./gofs -p 8080

I prefer put the binary file to /usr/bin, then it can be used anywhere!

    $ sudo mv gofs /usr/bin/