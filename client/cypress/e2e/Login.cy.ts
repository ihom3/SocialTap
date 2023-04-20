describe('Login', () => {
  it('login works with correct password', () => {
    cy.visit('http://127.0.0.1:4200')
    cy.contains('Login').click()
    cy.get('[id=inputEmail]').type('ian@gmail.com')
    cy.get('[id=inputPassword]').type('password')
    cy.get('button').contains('Sign In').click()
    cy.contains('Dashboard')
  })

  it('login fails with incorrect username', () => {
    cy.visit('http://127.0.0.1:4200')
    cy.contains('Login').click()
    cy.get('[id=inputEmail]').type('wronguser')
    cy.get('[id=inputPassword]').type('password')
    cy.get('button').contains('Sign In').click()
    cy.contains('User not found')
  })

  it('login fails with incorrect password', () => {
    cy.visit('http://127.0.0.1:4200')
    cy.contains('Login').click()
    cy.get('[id=inputEmail]').type('ian@gmail.com')
    cy.get('[id=inputPassword]').type('wrongpassword')
    cy.get('button').contains('Sign In').click()
    cy.contains('Incorrect password')
  })
})