# Sprint2
## Frontend

### Work Completed
Added the following tiles to the user profile page:
1. Tiktok
2. Youtube
3. Discord
4. Twitch
5. Reddit
6. Resume
7. E-mail

- Implemented logging in/logging out using apple, google, and facebook through Auth0
![](https://i.imgur.com/Xl2Ihyj.png)
- Implemented skeleton of dashboard component that retrieves user object from Auth0 and their registration code
![](https://i.imgur.com/ko00MDk.png)
- Implemented authenticated route /dashboard that will redirect the user to the login page if they are not authenticated
- Implemented header component
![](https://i.imgur.com/J85rHYz.png)
- Implemented custom user page that will fetch data from the backend database and use this information to render the user's social page based on the code provided to the URL. 
- Implemented page to register new device codes
![](https://i.imgur.com/hyEitJm.png)

### Unit Tests
![](https://i.imgur.com/ioEVEyT.png)
1. Test the creation of the SocialProviderService
2. Test the creation of the RegisterComponent
3. Test the creation of the SocialTileComponent
4. Test the icon switch statement for the SocialTileComponent
5. Test the creation of the ErrorPageComponent
6. Test that the AppCardComponent is created and initialized with a blank title
7. Test that the general AppComponent is created and functioning

Some of our angular unit tests are failing because of Injection errors in our test files for individual components. We are going to look into resolving these errors in our next sprint. 
### Cypress Tests
- "Home Page Opens": tests whether the home page loads properly 
- "Return Home Button Works": tests whether clicking the "Return home" button on the error screen takes the user back to the home page 
![image](https://user-images.githubusercontent.com/67469516/222322452-f51eb3db-7249-4533-b4b3-0441f2172720.png)

## Backend

### Work Completed
#### Description 
    For the backend portion of our project, We integerated authentication and authorization by implementing Auth0. Auth0 adds authentication
    to our program and allows our team to avoid the cost, time, and risk that come with building your own solution to authenticate and authorize 
    users. Our users should be able to log in with their social accounts (such as Facebook or Google). You want to retrieve the user's profile 
    after the login so you can customize the UI and apply your authorization policies. Specifically on the backend, we would retrieve the token received
    from the frontend and check if this user is authorized, and can send back the token to the frontend to authenticate that specific user.
    For the backend portion of our project, We integerated authentication and authorization by implementing Auth0. Auth0 adds 
    authentication to our program and allows our team to avoid the cost, time, and risk that come with building your own solution 
    to authenticate and authorize users. Our users should be able to log in with their social accounts (such as Facebook or Google). 
    You want to retrieve the user's profile after the login so you can customize the UI and apply your authorization policies. 
    Specifically on the backend, we would retrieve the token received from the frontend and check if this user is authorized, and can
    send back the token to the frontend to authenticate that specific user.

    Moreover, the database design is finalized. We have three tables total: users, social lists, and unregistered codes. The users table contains the user
    email, firstname and lastname, the sticker code, a bio, their profile picture path, and the socials' list. The social list table has a key associated 
    with it that links it to the users table with a "has a" relationship. It contains the social objects like instagram, facebook, snapchat etc. Those objects
    have each a name, status(displayed or not), and a url associated with it. The unregistered codes of the stickers are stored in the unregistered table. We
    also explored the functionality of the ParseMultipartForm function in go by adding a function to handle upload picture requests.
  


- Implemented Auth0 on the back-end to retrieve token from the front-end 
- Denies user access if Bearer token is not authenticated, else user permissions are enabled and can be accessed
## Auth0 Implementation / Back-End Images
- Curl POST request to generate an Auth0 token.
- <img width="1250" alt="Screen Shot 2023-03-01 at 8 03 00 PM" src="https://user-images.githubusercontent.com/73502423/222304370-2add37c7-e50e-4233-93c5-9eee180371c1.png">
- If the Bearer token is valid and accepted, GET USERS is applied and executed
- <img width="1009" alt="Screen Shot 2023-03-01 at 8 05 00 PM" src="https://user-images.githubusercontent.com/73502423/222304517-42d840c7-fbb1-4d55-9832-678f5cdab3fd.png">
- If the Bearer token is invalid, "Bad Request" is printed and GET USERS is not executed.
- <img width="1072" alt="Screen Shot 2023-03-01 at 8 05 10 PM" src="https://user-images.githubusercontent.com/73502423/222304642-848ef0e3-397a-4f81-99db-a9f20ed5695a.png">

### Unit Tests

- TestGetUsersStatusCode - It checks the get users function status code and returns true if the code is OK, else it fails.
- TestUpdateUserStatusCode - It checks the get users function status code and returns true if the code is OK, else it fails.
- TestGetUserByCode - It creates a user and puts in in the test database through the mock server. Then it checks that that is 
        the user which will be returned in the body of the request.
- TestCreateUser - It creates a user and puts in in the test database. It then encodes the user as JSON and make a POST request
        to the handler. It checks the status code and then queries the database to check if the user was created.
- TestInitialMigration - It connects to the test database. Then performs the initial database migration and checks if the users, 
        socials, and unregistered tables were created.
- TestDeleteUser - It sends a delete request through the mock server and first checks the status code and then reads the response
        from the body to see if it matches the expected value.
- TestUpdateProfilePicture - It adds a test picture to the test directory and checks to see if it is saved.
        
 <img width="174" alt="image" src="https://user-images.githubusercontent.com/90483046/222325961-9d199497-7fe2-4025-8e51-62671450b457.png">
The image shows the tests passing.
        
### Backend API Documentation
[Database Structure, Backend Endpoints & Registration Flow](https://ianblasko.notion.site/Software-Engineering-20eed26e5943404e8d357d40bb23a8e1)
- "/users" is a GET METHOD that retrieves our USER Struct Data to request all users stored in database.
- "/users/{id}" is a GET METHOD that retrieves our USER Struct Data to request a specific user with the corresponding ID as a parameter.
- "/users" is a POST METHOD that allows us to store a USER Struct user into our database.
- "/users/{id}" is a PUT METHOD that allows us to update a user inside of our database, such as email, name, etc... with ID as parameter.
- "/users/{id}" is a DELETE METHOD that allows us to remove a specific user based on the id parameter from inside our database
- "/users/{id}/instagram" is a GET METHOD that allows us to retrieve a users social media account, in this case specifically Instagram with id passed in as a parameter.
- "/users/code" is a POST METHOD that allows a user to add a user social to the user specificed.
- "/update-profile-picture" is a POST METHOD that allows for a user to change their profile picture and store it inside our database.

