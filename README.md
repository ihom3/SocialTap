## Project Name
SocialTap
## Project Description
For our project, we decided to develop an application that allows users to share their social media/professional platforms using an [NFC sticker](https://electronics.howstuffworks.com/nfc-tag.htm) on the back of their phone (or anywhere desired; laptop, wallet, water bottle, etc.). When tapped, the NFC sticker should redirect the person tapping it to a profile configured by the end user. The user will have a choice of what social media and professional platforms they want to display on their page, and they will have the ability to add a photograph and a bio as well through a customization panel.
## Group Members
| Member Name   | Position      |
| ------------- | ------------- |
| Ian Blasko    | Front End     |
| Jason Liang   | Front End     |
| Elsa Osmani   | Back End      |
| Kurtis Gnapp  | Back End      |

## Instructions to Install & Run

1. Clone this github repository into a folder.
2. Navigate to the client folder
3. Execute the following command: ng serve
4. Wait for Angular to download all required dependencies
5. Navigate to the server folder
6. Update 'database/connection.go' with the correct connection string corresponding to your MySQL database and login
7. Execute the following command: go build
8. Run the executable generated titled 'backend'
9. In a seperate terminal, open the client folder
10. Execute the following command: ng serve --host 0.0.0.0
11. Make a POST request to /api/reg-user to create a default Administrator user (email: ian@gmail.com, password: password)
12. Open a browser and navigate to 127.0.0.1:4200/login
13. Confirm you can login with the Administrator account
14. After confirmation, the endpoint '/api/reg-user' should be removed from 'routes/routes.go' for security purposes. The backend should then be recompiled with 'go build' and then the backend should be re-run
