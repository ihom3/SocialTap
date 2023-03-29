# Sprint3
## [Group Video]()

## Frontend

### Work Completed

### Unit Tests

1. Test the creation of the SocialProviderService
2. Test the creation of the RegisterComponent
3. Test the creation of the SocialTileComponent
4. Test the icon switch statement for the SocialTileComponent
5. Test the creation of the ErrorPageComponent
6. Test that the AppCardComponent is created and initialized with a blank title
7. Test that the general AppComponent is created and functioning

## Backend

### Work Completed
#### Description 

### Unit Tests

1. TestGetUsersStatusCode - It checks the get users function status code and returns true if the code is OK, else it fails.
2. TestUpdateUserStatusCode - It checks the get users function status code and returns true if the code is OK, else it fails.
3. TestGetUserByCode - It creates a user and puts in in the test database through the mock server. Then it checks that that is 
        the user which will be returned in the body of the request.
4. TestCreateUser - It creates a user and puts in in the test database. It then encodes the user as JSON and make a POST request
        to the handler. It checks the status code and then queries the database to check if the user was created.
5. TestInitialMigration - It connects to the test database. Then performs the initial database migration and checks if the users, 
        socials, and unregistered tables were created.
6. TestDeleteUser - It sends a delete request through the mock server and first checks the status code and then reads the response
        from the body to see if it matches the expected value.
7. TestUpdateProfilePicture - It adds a test picture to the test directory and checks to see if it is saved.


### Backend API Documentation
[Database Structure, Backend Endpoints & Registration Flow](https://ianblasko.notion.site/Software-Engineering-20eed26e5943404e8d357d40bb23a8e1)

- "/users" is a GET METHOD that retrieves our USER Struct Data to request all users stored in database.
Sample request and response:
GET localhost:9000/users
It gives all users stored, ex. If only two users:
{{
    "user_email": "ian.n.",
    "first_name": "Ian",
    "last_name": "B",
    "sticker_code": "hello",
    "bio_text": "hello world",
    "profile_picture": "1",
    "social_list": {
        "facebook": {
            "name": "Facebook",
            "status": true,
            "url": "/ian"
        },
        "snapchat": {
            "name": "",
            "status": false,
            "url": ""
        },
        "instagram": {
            "name": "Instagram",
            "status": true,
            "url": "/ian"
        }
    }
},
{
    "user_email": "apple@ufl.edu",
    "first_name": "Apple",
    "last_name": "B",
    "sticker_code": "apple",
    "bio_text": "hello Apple",
    "profile_picture": "3",
    "social_list": {
        "facebook": {
            "name": "Facebook",
            "status": true,
            "url": "/ian"
        },
        "snapchat": {
            "name": "Snapchat",
            "status": true,
            "url": "apple.instagram/"
        },
        "instagram": {
            "name": "Instagram",
            "status": true,
            "url": "/ian"
        }
    }
}}
- "/users/{id}" is a GET METHOD that retrieves our USER Struct Data to request a specific user with the corresponding ID as a parameter.
Example with id=1: 
GET localhost:9000/users/1
It gives the user with id=1:
{
    "user_email": "ian.n.",
    "first_name": "Ian",
    "last_name": "B",
    "sticker_code": "hello",
    "bio_text": "hello world",
    "profile_picture": "1",
    "social_list": {
        "facebook": {
            "name": "Facebook",
            "status": true,
            "url": "/ian"
        },
        "snapchat": {
            "name": "",
            "status": false,
            "url": ""
        },
        "instagram": {
            "name": "Instagram",
            "status": true,
            "url": "/ian"
        }
    }
}

- "/users" is a POST METHOD that allows us to store a USER Struct user into our database.
POST localhost:9000/users
{
    "user_email": "ian.n.",
    "first_name": "Ian",
    "last_name": "B",
    "sticker_code": "hello",
    "bio_text": "hello world",
    "profile_picture": "1",
    "social_list": {
        "facebook": {
            "name": "Facebook",
            "status": true,
            "url": "/ian"
        },
        "snapchat": {
            "name": "",
            "status": false,
            "url": ""
        },
        "instagram": {
            "name": "Instagram",
            "status": true,
            "url": "/ian"
        }
    }
}

It will respond with: 
{
    "user_email": "ian.n.",
    "first_name": "Ian",
    "last_name": "B",
    "sticker_code": "hello",
    "bio_text": "hello world",
    "profile_picture": "1",
    "social_list": {
        "facebook": {
            "name": "Facebook",
            "status": true,
            "url": "/ian"
        },
        "snapchat": {
            "name": "",
            "status": false,
            "url": ""
        },
        "instagram": {
            "name": "Instagram",
            "status": true,
            "url": "/ian"
        }
    }
}

- "/users/{id}" is a PUT METHOD that allows us to update a user inside of our database, such as email, name, etc... with ID as parameter.
PUT localhost:9000/users/{id}
Will take id as a parameter and update the value the user decided to change. To save space, the request body should look the same as the previous method, and the request will just have the updated value.
- "/users/{id}" is a DELETE METHOD that allows us to remove a specific user based on the id parameter from inside our database
DELETE localhost:9000/users/{id}
Will take id as a parameter and update the value the user decided to change. To save space, the request body should look the same as the previous method.
Will take id as a parameter and update the value the user decided to change. To save space, the request body should look the same as the previous method. It will return this message if succeded: “The user was deleted successfully.”

- "/users/{id}/instagram" is a GET METHOD that allows us to retrieve a users social media account, in this case specifically Instagram with id passed in as a parameter.
- "/users/code" is a POST METHOD that allows a user to add a user social to the user specificed.
- "/update-profile-picture" is a POST METHOD that allows for a user to change their profile picture and store it inside our database.

