# Photo-Mosaic in Golang

## Usage

eg. `go run main.go -m img/cage.jpg -t img/tiles/ -o img/output.png --mask 180`

```
NAME:
   challenge3 - Get Mosaic from Image and Tiles

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --main value, -m value   Obligatory. path to the main image of the mosaic
   --tiles value, -t value  Obligatory. path to the tiles images directory
   --ch value               number of horizontal splits (default: 80)
   --cv value               number of vertical splits (default: 50)
   --help, -h               show help
   --version, -v            print the version
```

## The Go Challenge 3

### Go Picture This!

#### Preamble

A photographic mosaic, or a photo-mosaic is a picture (usually a photograph) that has been divided into (usually equal sized) rectangular sections, each of which is replaced with another picture (called a tile picture). If we view it from far away or if you squint at it, then the original picture can be seen. If we look closer though, we will see that the picture is in fact made up of many hundreds or thousands of smaller tile pictures.

##### Goals of the challenge

Your mission, should you accept it, is to write a photo-mosaic generating program that:

* Allows the user to select a target picture, which is the picture that will be made into a photo-mosaic
* Allows the user to select a directory containing a set of tile pictures 
* Generates a photo-mosaic of the target picture using the tile pictures

##### Bonus goals (optional, not part of the challenge)

Create a web application that generates the photo-mosaic that:

* Allows the user to log in (can be your own database or log in through a third party like GitHub or Twitter or Facebook, through OAuth2). (Note: if you are authenticating the user through OAuth2, you should use the OAuth2 login flow instead of an external library).
* Allows the user to connect to one or more photo-sharing sites like Instagram or Flickr or Facebook Photos (or any photo-sharing site of your choice) to get tile pictures. Your user doesn't necessarily need to log in, you can use the image search APIs to get the tile pictures
* Allows the user to use a search filter (for e.g. use only pictures with cats in it) to filter out a set of tile pictures
* Allows the user to save the photo-mosaic, either on the site or upload it back to the photo-sharing site

##### Requirements of the challenge

* Use the latest version of Go i.e. version 1.4.2
* Individual tile pictures must be clearly visible when magnified, though it is expected to be smaller.
* You need to write test cases for the main flow. Do submit your test cases. 
* Do [organize your code](https://youtu.be/XCsL89YtqCs).
* Submit a photo-mosaic generated with your program, along with instructions to run the program as part of the submission.

##### Hints

You can find out more about photo-mosaics from this Wikipedia entry - [http://en.wikipedia.org/wiki/Photographic_mosaic](http://en.wikipedia.org/wiki/Photographic_mosaic)

You can also look at some photo-mosaic sites that are already available:

* [http://www.picturemosaics.com](http://www.picturemosaics.com)
* [http://www.easymoza.com](http://www.easymoza.com)
* [http://mosaically.com](http://mosaically.com)

Remember not to use the 'ghosting' technique when creating photo-mosaics, that's a big no-no. 'Ghosting' is when you place a faint picture of the target under your mosaic to create an illusion of a mosaic. It's kind of like cheating!

If you're writing a photo-mosaic web application, you can read this eBook:

[How to Deploy a Go Web App to the Google App Engine 101](https://leanpub.com/howtodeployagowebapptothegoogleappengine101)

Or you can view this YouTube video to learn more - [https://www.youtube.com/watch?v=XCsL89YtqCs](https://www.youtube.com/watch?v=XCsL89YtqCs)

If you find this challenge daunting or find yourself stuck, do go to the [Gophers Slack](http://t.co/n6EesY9Mmv) channel #golang-challenge and chat with me (@sausheong) - I will help you along the way.
