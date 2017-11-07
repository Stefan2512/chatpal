# Chat Pal


### Project Overview

This is my project for class [__Data Representation and Querying__](https://data-representation.github.io/]). 
It represents my Go implementation of [Eliza](https://en.wikipedia.org/wiki/ELIZA) in a web app with a chat interface
and a simple web server that serves the static files and handles message requests from the user.

### Topics

[Download and Run](#download-and-run)

[Features](#features)

[External Resources](#external-resources)

### Download and Run

To download and start the project, please follow the following instructions:

1. Clone the repository in your src folder. E.G.:

```bash
$ cd ~/go/src
$ git clone https://github.com/Stefan2512/chatpal.git
```

2. Go to the downloaded folder

```bash
$ cd ~/go/src/chatpal
```

3. Build the project (please make sure you are in the right directory as said in step 2) 

```bash
$ go build
```

4. Run the project

```bash
$ ./chatpal
```

Then it will prompt the following message: `Listening to port 4419...` and the app is ready
to be accessed in browser at [http://localhost:4419](http://localhost:4419).

### Features

#### chat bot contexts

Chat Pal is able to respond to simple day to day sentences. If familiar with a context, it will respond accordingly.
Although sending messages can be complex, Chat Pal will be able to treat at most one context (the first one, if any). 
Below are examples of *contexts* Chat Pal is aware of:

* Hi/Hello/Greetings messages
* Possessions - sentences containing *My ...*  (e.g. my day is great)
* Feelings - sentences containing *I feel...* or *I am...*
* Asking for something - *Would you...*
* Going to places - *going to ...*
* If it does something - *do you...?*
* cheap/expensive things

If none of the above situations are "understood" by Chat Pal, a list of *fallback* (default/generic) statements
will be generated.

#### UI
Communicating with Chat Pal is done via a web interface. The UI is a chat
that contains the list of messages send to and received from the Go server.

#### Implementation notes

* The ability of __Chat Pal__  to respond accordingly is based on parsing a few key words. If any of these keywords match,
a list of possible statements is associated with that key. In the end, one possible statement
will be picked at random.
* The page layout was [source](#external-resources) inspired, but the Javascript code and logic was implemented by me
* The Go code contains only built-in libraries

### External Resources

UI - HTML and CSS snippet - [https://bootsnipp.com/snippets/N6BgR](https://bootsnipp.com/snippets/N6BgR) - snippet made by [harunpehlivan](https://bootsnipp.com/harunpehlivan)
