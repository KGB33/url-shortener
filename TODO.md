# Improve error handling.
  - Wrap lower-level/db error so they can easily be presented to the user (or handled)
  - Identify and handle errors by their type, not just if they exist.
[GopherCon 2020: Jonathan Amsterdam - Working with Errors](https://www.youtube.com/watch?v=IKoSsJFdRtI&list=PL2ntRZ1ySWBfUint2hCE1JRxRWChloasB&index=1)

# Set up golangci-lint
  - On-pre-commit

# Add GH action to build/Publish container to Github container repository
  - Create release branch
  - iff golangci/tests pass then build and publish the container on a push to that branch

# Make the front-end look kinda nice
  - Add colors
  - Add icons to buttons
  - Add Background/border to table

# Reach 85% test coverage
  - Refactor integration tests out of the `app` package

# Post to reddit(?)
  - Gather feedback of some kind
