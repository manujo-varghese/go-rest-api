# go-rest-api


# Prerequisities 

This REST API is built with GO and containirised with Docker , so the only prerequisite to run the API locally is "Docker" installed !

Steps to run :

1: clone the repo (public)
2: Open your favourate IDE (mine is VS code) and open terminal
3: cd into go-rest-api folder
4: Make sure that your Docker cotainer is up
5: type the command : docker-compose up --build (spins up server on host 8080) change config in docker-compose.yml(if required)
6: Open any API testing tool (Postman) 

Different EndPoints:

There are main 4 endpoints are availabel to api as mentioned in question

1: POST http://localhost:8080/api/article

Lets starts by creating a new Article , successful creation will result in return of created object

simlar body content is 
{
    "Title":"Article one",
    "Body":"This is a test",
    "Tags":["health","test","one","similar"]
}
You should get a similar response like 

{
    "ID": 1,
    "CreatedAt": "2022-01-18T13:10:15.214334+11:00",
    "UpdatedAt": "2022-01-18T13:10:15.214334+11:00",
    "DeletedAt": null,
    "Title": "Article one",
    "Date": "2022-01-18T00:00:00Z",
    "Body": "This is a test",
    "Tags": [
        "health",
        "test",
        "one",
        "similar"
    ]
}
2: GET http://localhost:8080/api/article/1

This will fetch Article detailes of ID 1 and similarly any error if any( like not existing ids)
you should get similar response if the article exists
3: GET http://localhost:8080/api/tags/tag/20220118

You should get a similar response as per the question:
{
    "Tag": "health",
    "Count": 161,
    "Articles": [
        "11",
        "12",
        "14",
        "15",
        "16",
        "17",
        "18",
        "19",
        "20",
        "21"
    ],
    "RelatedTags": [
        "one",
        "test",
        "similar",
        "two"
    ]
}
4: GET http://localhost:8080/api/health

This is to make sure the app is up and running returns a message "I am Alive" if its !!!!
