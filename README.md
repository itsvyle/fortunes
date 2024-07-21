# fortunes

A reproduction of the Unix `fortune` utility in multiple languages, as a way to compare languages, and create dev container configurations.

It can read fortunes after they have been compiled by my fortune-generator, written in Go: https://github.com/itsvyle/fortune-generator

# General use

Note that this is the base feature set; all languages will not support all the options; checkout the README file in each language directory

| Parameter | Description                                                                                | Required | Default |
| --------- | ------------------------------------------------------------------------------------------ | -------- | ------- |
| -p        | Path to the folder containing fortunes and the `.vyle` file given by the fortune-generator | true     | -       |
| -s        | Print the name of the source file for the fortune (0 or 1)                                 | false    | 0       |
| -max      | Max length of the generated fortune; 0 = no-limit                                          | false    | 0       |
| -min      | Min length of the generated fortune                                                        | false    | 0       |
| -n        | Number of iterations; mostly for testing/benchmarking                                      | false    | 1       |
