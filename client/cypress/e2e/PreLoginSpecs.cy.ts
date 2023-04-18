// Tests without needing to log in first

describe('Pre-Login Specs', () => {

  beforeEach(() => {
    //beforeEach() brings the user to the desired site before each test
    cy.visit('http://localhost:4200') // brings the user to the home page 
  })

  it('opens', () => {
    cy.visit('http://localhost:4200')
    cy.contains('Home Page')
  })

  it('displays Auth0 page when navigating to dashboard', () => {
    cy.visit('http://localhost:4200/dashboard') //visit dashboard
    cy.origin('https://dev-6ipuzz0uj7yyefqe.us.auth0.com', () => 
    {
    cy.contains('Welcome')
    cy.contains('Continue with Google')
    })
  })

  it('Auth0 sign up works', () => {
    cy.visit('http://localhost:4200/dashboard') //visit dashboard
    cy.origin('https://dev-6ipuzz0uj7yyefqe.us.auth0.com', () => 
    {

      let x:number = Math.floor(Math.random() * (0 - 10000));
      var str1 = new String(x.toString());
      var str2 = new String('@gmail.com');
      var str3 = str1.concat(str2.toString());

      cy.contains('Sign up').click();
      cy.get('[name=email]').type(str3)
      cy.get('[name=password]').type('Password123!')
      cy.contains('Continue').click();
      cy.contains('Extensibility error').should('not.exist')
      cy.contains('Accept').click();
      
    })
    cy.contains('Log out')
  })

  it('Auth0 existing account sign in works', () => {
    cy.visit('http://localhost:4200/dashboard') //visit dashboard
    cy.origin('https://dev-6ipuzz0uj7yyefqe.us.auth0.com', () => 
    {
      cy.get('[name=username]').type('Auth0test2@gmail.com')
      cy.get('[name=password]').type('Password123!')
      cy.contains('Continue').click();
    })
    cy.contains('Log out')
  })


  it('Auth0 sign in through google works', () => {
    cy.visit('http://localhost:4200/dashboard') //visit dashboard
    cy.origin('https://dev-6ipuzz0uj7yyefqe.us.auth0.com', () => 
    {
      cy.contains('Google').click();
    })
    cy.origin('https://accounts.google.com', () => 
    {
      cy.get('[name=identifier]').type('SocialTapTest@gmail.com')
      cy.contains('Next').click();
      cy.wait(5000)
      cy.get('[name=password]').type('86iSD30Ecqe$')
      cy.contains('Next').click();
    })
    
  })

})
// 86iSD30Ecqe$
//cy.url().should('equal', 'http://localhost:4200/')