describe('the spec to test the-mentalist project', () => {
  beforeEach(() => {
    // Visit the project site before each test
    cy.visit('http://localhost:8080/')
  })

  it('verifies the answer to default question', () => {
    cy.contains('Get Answer').click()
    cy.contains('The CBI stands for California Bureau of Investigation.')
  })

  it('clicks to get the answer', () => {
    cy.get('#question').select(1)
    cy.contains('Get Answer').click()
    cy.contains('Patrick Jane has a keen sense of observation and deduction')
  })
})
