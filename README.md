# fortunes

A reproduction of the Unix `fortune` utility in multiple languages, as a way to compare languages, and create dev container configurations, and to have cool stuff in my terminal

![image](https://github.com/user-attachments/assets/629c9281-5774-441d-b0a6-8649b0cd5a0d)


It can read fortunes after they have been compiled by my fortune-generator, written in Go: https://github.com/itsvyle/fortune-generator

By default, it'll look for a `fortunes.vyle` folder and it's associated fortunes in ~/.config/fortunes

# General use

Note that this is the base feature set; all languages will not support all the options; checkout the README file in each language directory

| Parameter | Description                                                                                | Required | Default |
| --------- | ------------------------------------------------------------------------------------------ | -------- | ------- |
| -path     | Path to the folder containing fortunes and the `.vyle` file given by the fortune-generator | true     | -       |
| -s        | Print the name of the source file for the fortune (0 or 1)                                 | false    | 0       |
| -max      | Max length of the generated fortune; 0 = no-limit                                          | false    | 0       |
| -min      | Min length of the generated fortune                                                        | false    | 0       |
| -n        | Number of iterations; mostly for testing/benchmarking                                      | false    | 1       |
