# Ricochet Robots Solver

Once Move is implemented, the way I was thinking of doing it (and I may be very wrong) is this:
- For every robot and for every direction
  - Launch a new goroutine that makes a move
  - checks if the move has solved the current position
  - check a channel to see if anyone else has solved the position (if yes, just quit)
  - if my position is solved, signal on channel for the others to stop, send my result to a specific channel
  - if not, for every robot and every direction (Etc etc)


🤖
