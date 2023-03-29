// Tests without needing to log in first

describe('Post-Login Specs', () => {

    beforeEach(() => {
        cy.visit('http://localhost:4200/dashboard') //visit dashboard
        cy.origin('https://dev-6ipuzz0uj7yyefqe.us.auth0.com', () => 
        {
            cy.get('[name=username]').type('Auth0test2@gmail.com')
            cy.get('[name=password]').type('Password123!')
            cy.contains('Continue').click();
        }) 
    })
  
    it('Logs in', () => {
      cy.contains('Dashboard')
    })

    it('Update Name Page Works', () => {
      cy.contains('Update Name').click()
      cy.contains('Update Name')
    })

    it('Update Name Back Button Works', () => {
      cy.contains('Update Name').click()
      cy.contains('Back').click()
      cy.url().should('equal','http://localhost:4200/dashboard')
    })
  
    it('Update Bio Page Works', () => {
      cy.contains('Update Bio').click()
      cy.contains('Update Bio')
    })

    it('Update Bio Back Button Works', () => {
      cy.contains('Update Bio').click()
      cy.contains('Back').click()
      cy.url().should('equal','http://localhost:4200/dashboard')
    })
  
  })

