describe('template spec', () => {

  const login = (name = 'ian@gmail.com') => {
    cy.session(name, () => {
        cy.visit('http://127.0.0.1:4200/login')
        cy.get('[id=inputEmail]').type('ian@gmail.com')
        cy.get('[id=inputPassword]').type('password')
        cy.get('button').contains('Sign In').click()
        cy.contains('Dashboard')
        },
      {
        cacheAcrossSpecs: true,
      }
    )
  }

  beforeEach(()=> {
    //logs user in
    login()
    cy.visit('http://127.0.0.1:4200/dashboard')
  })
  
  it('Update Name Works', () => {
    cy.contains('Update Name').click()
    var str1 = makeid(2)
    var str2 = makeid(2)
    cy.get('[formcontrolname="first_name"]').type(str1)
    cy.get('[formcontrolname="last_name"]').type(str2)
    cy.get('button').contains('Update Name').click()
    cy.contains(str1)
    cy.contains('Update Name').click()
    cy.contains(str1)
    cy.contains(str2)
  })

  it('Update Bio Works', () => {
    cy.contains('Update Bio').click()
    var str1 = makeid(100)
    cy.get('[formcontrolname="bio"]').type(str1)
    cy.get('button').contains('Update Bio').click()
    cy.contains('Update Bio').click()
    cy.contains(str1)
  })

  it('Toggle Twitter Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="Twitter"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="Twitter"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=Twitter]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="Twitter"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=Twitter]')
    })   
    
  it('Toggle Instagram Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="Instagram"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="Instagram"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=Instagram]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="Instagram"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=Instagram]')
    })  

  it('Toggle LinkedIn Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="LinkedIn"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="LinkedIn"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=LinkedIn]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="LinkedIn"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=LinkedIn]')
    })  

  it('Toggle Snapchat Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="Snapchat"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="Snapchat"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=Snapchat]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="Snapchat"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=Snapchat]')
    })  
    
  it('Toggle TikTok Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="TikTok"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="TikTok"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=TikTok]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="TikTok"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=TikTok]')
    })  

  it('Toggle YouTube Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="YouTube"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="YouTube"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=YouTube]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="YouTube"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=YouTube]')
    })  

  it('Toggle Twitch Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="Twitch"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="Twitch"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=Twitch]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="Twitch"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=Twitch]')
    }) 

  it('Toggle GitHub Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="GitHub"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="GitHub"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=GitHub]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="GitHub"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=GitHub]')
    }) 

  it('Toggle Spotify Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="Spotify"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="Spotify"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=Spotify]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="Spotify"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=Spotify]')
    }) 

  it('Toggle SoundCloud Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="SoundCloud"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="SoundCloud"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=SoundCloud]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="SoundCloud"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=SoundCloud]')
    }) 

  it('Toggle Discord Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="Discord"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="Discord"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=Discord]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="Discord"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=Discord]')
    }) 

  it('Toggle Email Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="Email"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="Email"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=Email]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="Email"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=Email]')
    }) 

  it('Toggle Reddit Works', () => {
    cy.contains('Update Socials').click()
    cy.get("body").then($body => {
        if ($body.find('[id="Reddit"][ng-reflect-checked=true]').length > 0) {   
          cy.get('[id="Reddit"][ng-reflect-checked=true]').click()
          cy.get('button').contains('Update Socials').click()
        }
        else {
          cy.get('button').contains('Back').click()
        }
    })
    cy.contains('View Page').click()
    cy.get('[id=Reddit]').should('not.exist')
    cy.contains('Dashboard').click()
    cy.contains('Update Socials').click()
    cy.get('[id="Reddit"][ng-reflect-checked=false]').click()
    cy.get('button').contains('Update Socials').click()
    cy.contains('View Page').click()
    cy.get('[id=Reddit]')
    }) 
})

function makeid(length:any) {
  let result = '';
  const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  const charactersLength = characters.length;
  let counter = 0;
  while (counter < length) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
    counter += 1;
  }
  return result;
}