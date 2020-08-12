# Delivery Much Tech Challenge

This is an API made for the Delivery Much backend challenge.

It has 1 endoint that receives a list of comma separated ingredients, and returns a list of recipes from [Recipe Puppy](http://www.recipepuppy.com/about/api) related to those ingredients and a related gif fetched from the [Giphy API](https://developers.giphy.com/docs)

### Build and Run

First build the docker container
`docker build -t api .`

Then, run it:
`docker run -d -p 3000:3000 api`

### Usage

The call must follow the following format:
`http://{HOST}/recipes/?i={ingredient_1},{ingredient_2}`

Exemplo:

`http://127.0.0.1/recipes/?i=onion,tomato`


The response have the following structure:

```
{
	"keywords": [
    "onion",
    "tomato"
    ],
	"recipes": [
    {
      "title": "Greek Omelet with Feta",
      "ingredients": ["eggs", "feta cheese", "garlic", "red onions", "spinach", "tomato", "water"],
      "link": "http://www.kraftfoods.com/kf/recipes/greek-omelet-feta-104508.aspx",
      "gif": "https://media.giphy.com/media/xBRhcST67lI2c/giphy.gif"
      },{
      "title": "Guacamole Dip Recipe",
      "ingredients": ["avocado", "onions", "tomato"],
      "link":"http://cookeatshare.com/recipes/guacamole-dip-2783",
      "gif":"https://media.giphy.com/media/I3eVhMpz8hns4/giphy.gif"
	   }
	]
}
```


### Configuration

All relevant configuration, such as external APIs URLs and PORT, are defined in the main `.env` file in the root of the project.

### Specifications

- Made in Go using Go Modules
- All configuration needed is contained inside the `.env` file in the root of the project
- The gif is obtained by fetching the Giphy API with the title of the recipe
- The ingredients returned by Recipe Puppy are sorted alphabetically in the response
- Correct HTTP Codes and informative messages in case of failure in external service calls
- Dockerized: Build and run with docker
- Separation of layers between routes and handlers
- Use of modules to isolate helper functions from main package files
- Easily add new routes and handlers by editing single packages
