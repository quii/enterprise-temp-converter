# Demo plan

- Demo the two apps
- Show the `cmd/main.go` for both
  - Symmetry
    - Create the "app", which gives you a converter from our domain
    - Wire it into an "adapter" which takes our domain service to do work
    - Start the program
- Then show the test
  - Again, symmetry
  - Build a binary of the program we're testing (and delete it after)
  - Create a "driver", which understands how to "drive" the program for the spec
  - Send the spec into the driver
- Then show the spec
  - Isn't it cool we can use the same spec for both applications
  - It doesn't care **how** the system converts temps, it just cares **what** it does
  - The spec is just a statement of truth, nothing more
  - The fact it works with 2 systems is not so important, the important thing it shows is the test is **decoupled from the implementation**
    - Good tests have this property
- Because of this decoupling, we can also use the spec at different architectural layers of our application
  - Like the unit test
- Show the ergonomics of the tests
  - Nothing fancy, no frameworks or clever business
  - Can just press the green arrow and run the test, **get fast feedback on a whole application in a matter of milliseconds**
- General thoughts
  - By starting with an acceptance test, I could be **really** bold with my refactoring and playing around with the structure of the app. If I broke something, I got immediate feedback and could rollback safely. It felt very liberating
    - No structure at all at first, everything was just in `main.go` because RED/GREEN/REFACTOR. 
      - In particular, as Kent Beck says "commit whatever sins are needed to get into green"
  - I was tempted to use docker to package up the app for testing, but wanted to try just building a binary and seeing. Strong advantage is its very low-tech, very fast.



TODO: need a diagram

And possibly do a form