/// <reference types="cypress" />

describe('spec.cy.js', () => {
  beforeEach(() => {
    // beforeEach() brings the user to the desired site before each test
    cy.visit('/') // brings the user to the home page 
  })

  it('Home Page Opens', () => {
      cy.contains('Home Page')
  })

  it('Return Home Button Works', () => {
      cy.visit('/error') //visit user page 

      cy.get('button').contains('Return home').click()

      cy.url().should('equal', 'http://localhost:4200/')
  })

})