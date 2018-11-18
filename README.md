# API for dialogflow chatbot
Assignment 3 (Project) in the course IMT2681-2018 (Cloud Technologies) at NTNU Gjøvik.

This application is a RESTful API that can querie other public APIs based on the parameter specified.     
This API is used by [dialogflow](https://dialogflow.com/) to post different messaged to a Slack webhook.

***

## API calls:
### Help page
Displays this README in html format.

### dialogflow
Used by dialogflow to querie this API.    
Different public APIs are queried based on the POST sent by dialogflow.
```
POST: /dialogflow - Returns a custum API responce based on a parameter "b" in post body

Correct json request:

"queryResult": {
    "queryText": "joke",
    "action": "actionAndParameters",
    "parameters": {
      "b": "joke"
    },
    "allRequiredParamsPresent": true,
    "fulfillmentMessages": [{
      "text": {
        "text": [""]
      }
    }],

acceptable values for parameter b are:
"joke"    For a chuch Norris Joke
"dad"     For a dad Joke


```

### log
 Returns all logs of all events that has beed logged (mainly errors).  
```
GET: /log
```

### Old posts
Returns all API requests preveausly sent to this API.
```
GET: /OldPosts
```

###  Statistics
Returns statistics for the different API usage.
```
GET: /statistics
```

### Website
A website where you can test the available APIs we have.    
[Link](https://warm-meadow-53471.herokuapp.com/website.html)

***

## How this app is deployed
 * The app runs in Heroku at https://warm-meadow-53471.herokuapp.com/
 * Database ?
 * OpenStack? Docker?
.env file must have 

		MONGO_ADDRESS=<mongodb address:port>
		MONGO_USER=<username>
		MONGO_PASSWORD=<password>
		MONGO_DATABASE=<database name>

Run with: `docker run --name app --rm --env-file ./.env -p 8080:8080 -d celebrian/api:latest`

If you want to build it yourself use: `docker build --rm -t <user>/<image>:<tag> .`
Then: `docker run --name <name> --rm --env-file <path/to/env/file/> -p <port>:8080 -d <user>/<image>:<tag>

## Additional information
Created by Alexander Jakobsen, Martin Brådalen, Mats Ove Mandt Skjærstein with help from Per-Kristian Kongelf Buer
