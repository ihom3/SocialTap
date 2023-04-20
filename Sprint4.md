# Sprint4
## [Group Video]()

## Frontend
### Work Completed
In Sprint 4, we decided to switch from Auth0 to JWT Authentication, and implemented that accordingly, along with an organization overhaul. We added functionality to be able to edit Name, Bio, and Socials on the frontend from the dashboard. Our home page now includes information about the team and what our app does, with some added buttons for improved navigation. We fully implemented new user registration, and have conditional navigation based on whether the user has administrator privileges. 

In the frontend, we implemented an Angular Service to connect to our Golang backend. The angular service has similar functions to those on the backend and it monitors the current status of the user to determine if they’re logged in or not. If they are logged in, the User Service stores the user’s data as well. 

We designed and implemented a new home page:
![homePage](https://i.imgur.com/iSmBAQt.png)

### Unit Tests

From Sprint 2:
1. Test the creation of the SocialProviderService
2. Test the creation of the RegisterComponent
3. Test the creation of the SocialTileComponent
4. Test the icon switch statement for the SocialTileComponent
5. Test the creation of the ErrorPageComponent
6. Test that the AppCardComponent is created and initialized with a blank title
7. Test that the general AppComponent is created and functioning

From Sprint 3:

8.
 ![preLoginSpecs](https://user-images.githubusercontent.com/67469516/228697184-1659b51e-e104-4c88-a304-4797058f3302.PNG)

9.
![postLoginSpecs](https://user-images.githubusercontent.com/67469516/228697197-d27e7fea-f237-40c5-8866-fa853cc7339c.PNG)

Sprint 4:

10.
![image](https://user-images.githubusercontent.com/67469516/233243268-0169b0cd-c546-4050-9204-d86490f49ff5.png)

11.
![image](https://user-images.githubusercontent.com/67469516/233243415-241efc0a-4cef-43cc-8eb4-8ca1ca2d90a7.png)


## Backend

### Work Completed
In Sprint 4, a major decision was made on what form of authentication we were going to use. We switched from Auth0 to JWT Authentication. Also, we found it easier to implement JWT authentication using Fiber instead of GorillaMux. Therefore, we had to rewrite our handler functions. For this sprint we completed the endpoints and tested them using the unit test package in go. We also, made a few changes to our database. Instead of having a struct for our list of socials, we switched it to an array of SocialType struct. That array will hold the information about which social platforms and their links the user wants to display on their homepage. The details about the routes are provided in the API documentation section below.
### Unit Tests

- SPRINT 2
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
- SPRINT 4
13. TestGetProfilePicture - It adds a test picture and checks to see if that the one we get back from the request.
14. TestRegisterNewCode - Checks if the user is of "admin" status and if the code will be properly added in the database.
15. TestUpdateName - Updates the name for a test user and checks if the name is properly updated in the database. It checks three scenarios: with proper request, invalid request body, and unauthenticated request.
16. TestUpdateBio - Updates the bio for a test user and checks if the bio text is properly updated in the database.
17. createTestJWT - Creates a JWT Token that will later be received by a User. If the token fails to generate, we return an empty Token String.
18. TestGetUser - It creates a user along with a JWT token and stores the user inside our database. If the user is not found, it checks first if the wrong user was retrieved, and then checks if no user was found at all.
19. createFormFileUploadRequest - Creates a new file form field and adds the Users ID to the URL path to specify the specific user.
20. TestIDRoute - checks if a Users ID can be found to retrieve their information from inside the database. It checks if the user is not found and/or if the user is not registered.

<img width="239" alt="image" src="https://user-images.githubusercontent.com/90483046/233231352-633d7c76-0173-409b-b4d8-f8319b04740c.png">

### Backend API Documentation
[Database Structure, Backend Endpoints & Registration Flow](https://ianblasko.notion.site/Software-Engineering-20eed26e5943404e8d357d40bb23a8e1)

- "/api/register" is a POST method that checks if the user is unregistered. If they aren't, then you will create the user in the database and remove a code as unregistered.
- "/api/login" is a POST method that checks if the user is already logged in. It then checks the user provided email and password and checks if the user is able to lock in.
- "/api/logout" is a POST method that forces a JWT Token to expire and causes the user to be logged out.
- "/api/register-code" is a POST method that checks the code can be registered only by an admin user, and it cannot be already registered by another user or unregistered. Overall, this function provides basic authentication and authorization checks and ensures that the code being registered is not already in use.
- "/api/update-picture" is a POST METHOD that allows for a user to change their profile picture and store it inside our database.
- "/api/update-bio" is a POST method. It updates a specific users bio.
- "/api/update-email" is a POST method. It updates a specific users email.
- "/api/update-name" is a POST method. It updates a specific users first and last name.
- "/api/update-password" is a POST method. It updates a specific users password associated with their account.
- "/api/update-socials" is a PUT method that updates the Social List information in the database.
- "/api/profile-picture/:id" is a GET method that retrieves a users profile picture that is associated with their sticker code. Else, it's unregistered or it doesn't exist.
