describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://127.0.0.1:4200')
    cy.contains('Login').click()
    cy.get('[id=inputEmail]').type('ian@gmail.com')
    cy.get('[id=inputPassword]').type('password')
    cy.get('button').contains('Sign In').click()
    cy.contains('Dashboard')
  })
})