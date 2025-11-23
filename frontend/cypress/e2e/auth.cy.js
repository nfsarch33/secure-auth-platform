describe('Authentication Flow', () => {
  it('should allow a user to sign up and then sign in', () => {
    const email = `testuser_${Date.now()}@example.com`;
    const password = 'Password123!';

    // Sign Up
    cy.visit('/signup');
    cy.get('input[id="email"]').type(email);
    cy.get('input[id="password"]').type(password);
    cy.get('button[type="submit"]').click();
    
    // Verify success message
    cy.contains('Sign up successful! Please sign in.').should('be.visible');
    
    // Sign In
    cy.visit('/signin');
    cy.get('input[id="email"]').type(email);
    cy.get('input[id="password"]').type(password);
    cy.get('button[type="submit"]').click();
    
    // Verify success message
    cy.contains('Sign in successful!').should('be.visible');
  });
});
