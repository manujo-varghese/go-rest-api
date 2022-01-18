# go-rest-api


# Prerequisities 

This REST API is built with GO and containirised with Docker , so the only prerequisite to run the API locally is "Docker" installed !

Steps to run :

1: clone the repo (public)<br />
2: Open your favourate IDE (mine is VS code) and open terminal<br />
3: cd into go-rest-api folder<br />
4: Make sure that your Docker cotainer is up<br />
5: type the command : docker-compose up --build (spins up server on host 8080) change config in docker-compose.yml(if required)<br />
6: Open any API testing tool (Postman) <br />

Different EndPoints:<br />

There are main 4 endpoints are availabel to api as mentioned in question

1: POST http://localhost:8080/api/article<br />

Lets starts by creating a new Article , successful creation will result in return of created object

simlar body content is 
{
    "Title":"Article one",<br />
    "Body":"This is a test",<br />
    "Tags":["health","test","one","similar"]<br />
}
You should get a similar response like 

{
    "ID": 1,<br />
    "CreatedAt": "2022-01-18T13:10:15.214334+11:00",<br />
    "UpdatedAt": "2022-01-18T13:10:15.214334+11:00",<br />
    "DeletedAt": null,<br />
    "Title": "Article one",<br />
    "Date": "2022-01-18T00:00:00Z",<br />
    "Body": "This is a test",<br />
    "Tags": [<br />
        "health",<br />
        "test",<br />
        "one",<br />
        "similar"<br />
    ]
}
2: GET http://localhost:8080/api/article/1 <br />

This will fetch Article detailes of ID 1 and similarly any error if any( like not existing ids)
you should get similar response if the article exists
3: GET http://localhost:8080/api/tags/tag/20220118<br />

You should get a similar response as per the question:<br />
{
    "Tag": "health",<br />
    "Count": 161,<br />
    "Articles": [<br />
        "11",<br />
        "12",<br />
        "14",<br />
        "15",<br />
        "16",<br />
        "17",<br />
        "18",<br />
        "19",<br />
        "20",<br />
        "21"<br />
    ],<br />
    "RelatedTags": [<br />
        "one",<br />
        "test",<br />
        "similar",<br />
        "two"<br />
    ]<br />
}
4: GET http://localhost:8080/api/health <br />
 <br />
This is to make sure the app is up and running returns a message "I am Alive" if its !!!!


# Test Cases :

Test cases are writtern inside the folder test<br />
and to run locally you can use the command;<br /> go test  ./... -tags=e2e -v   <br />

# Why Go and GORM

Building API with strongly typed langauge is a cool way and i like GO

# Time i took to complete and what could i have done with more time?

So i took around 3 hours to complete as i browse around some syntax with GO (as Go is not my everyday language as of now),
If i had more time i will spend some more time writing comprehensive test cases to make sure everything is validated
