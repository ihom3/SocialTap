## Sprint 2

### Work Completed 
#### Backend
    For the backend portion of our project, We integerated authentication and authorization by implementing Auth0. Auth0 adds authentication
    to our program and allows our team to avoid the cost, time, and risk that come with building your own solution to authenticate and authorize 
    users. Our users should be able to log in with their social accounts (such as Facebook or Google). You want to retrieve the user's profile 
    after the login so you can customize the UI and apply your authorization policies. Specifically on the backend, we would retrieve the token received
    from the frontend and check if this user is authorized, and can send back the token to the frontend to authenticate that specific user.
    
    Moreover, the database design is finalized. We have three tables total: users, social lists, and unregistered codes. The users table contains the user
    email, firstname and lastname, the sticker code, a bio, their profile picture path, and the socials' list. The social list table has a key associated 
    with it that links it to the users table with a "has a" relationship. It contains the social objects like instagram, facebook, snapchat etc. Those objects
    have each a name, status(displayed or not), and a url associated with it. The unregistered codes of the stickers are stored in the unregistered table. We
    also explored the functionality of the ParseMultipartForm function in go by adding a function to handle upload picture requests.
    
#### Frontend
    For the frontend...

### Unit Tests
### Backend
    Unit tests...
#### Frontend
    Cypress Test...
    Unit tests...
