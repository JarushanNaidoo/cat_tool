WHAT DOES CAT DO:

- The cat utility reads files sequentially, writing them to the standard output.
  The file operands are processed in command-line order. If file is a single
  dash (‘-’) or absent, cat reads from the standard input.

LINK:

- https://codingchallenges.fyi/challenges/challenge-cat/

TODO:

- read input
- split input based on " " (space) and stop in a buffer
- move through buffer and validate
- buffer[0]  = "cat"
- command has to start with '-' (dash)
- if no command take input from std in otherwise take input from file