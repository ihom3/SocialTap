# Sprint4
## [Group Video]()

## Frontend
### Work Completed



### Unit Tests
New for Sprint 3:

![preLoginSpecs](https://user-images.githubusercontent.com/67469516/228697184-1659b51e-e104-4c88-a304-4797058f3302.PNG)
![postLoginSpecs](https://user-images.githubusercontent.com/67469516/228697197-d27e7fea-f237-40c5-8866-fa853cc7339c.PNG)

From Sprint 2:
1. Test the creation of the SocialProviderService
2. Test the creation of the RegisterComponent
3. Test the creation of the SocialTileComponent
4. Test the icon switch statement for the SocialTileComponent
5. Test the creation of the ErrorPageComponent
6. Test that the AppCardComponent is created and initialized with a blank title
7. Test that the general AppComponent is created and functioning

## Backend

### Work Completed

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
- SPRINT 3
7. TestUpdateProfilePicture - It adds a test picture to the test directory and checks to see if it is saved.
8. TestDeleteCode - It checks if a test code is deleted at the unregistered table through the DELETE method.
9. TestAddCode - It checks if the AddCode method is a POST method.
10. TestDashboard - It checks if the Dashboard method is a GET method.
11. TestIDRoute - It checks if the IDRoute method is a GET method.
12. TestGetUserNameByCode - It checks if the user exists or not by using their unique sticker code through a GET method.


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
- SPRINT 3
- "/update-profile-picture" is a POST METHOD that allows for a user to change their profile picture and store it inside our database.
- "/{sticker_code}" is a GET method that checks the unregistered_codes tables first and then the users table. It will return the unregistered code or the user info.
- "/dashboard/{sticker_code}" is a GET method. It is the first endpoint where the users will be directed. If the user has not been registered, it will create an entry in the database for the user with all the fields being empty.If the user is already registered, it will respond with all the data of the user.
- "/update-profile" is a POST method. It updates any data we pass in with the json body.
- "/update-socials/{id}" is a PUT method that updates the Social List information in the database.
- "/unregistered" is a POST method that adds an unregistered code in the database. It carries the sticker code in the json body.
- "/unregistered/{sticker_code}" is a DELETE method that deletes the code from the unregistered table, after a user with that code has been created.
- "/user/{sticker_code}" is a GET method that retrieves a users first and last name that's associated with their sticker code. Else, it's unregistered or it doesn't exist.
- "/picture/{sticker_code}" is a GET method that retrieves a users profile picture that is associated with their sticker code. Else, it's unregistered or it doesn't exist.
