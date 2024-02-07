# SquareCube Challenge

## Challenge Description
- **Task**: Write a loop that iterates from 1 to `n`, where `n` is a dynamically provided input, not hardcoded.
- **Conditions**:
  - If a number is a perfect square, print `"Square"`.
  - If a number is a perfect cube, print `"Cube"`.
  - If a number is both a perfect square and a perfect cube, print `"SquareCube"`.
  - Otherwise, print the number itself.

### Input
- `n` (integer): The maximum number for the loop iteration. The value of `n` is dynamic.

### Output
- The output should be either `"Square"`, `"Cube"`, `"SquareCube"`, or the number itself based on the conditions described above.

### Notes
- **Perfect Square**: A number is considered a perfect square if it can be expressed as the square of an integer. For example, 16 is a perfect square because it is 4².
- **Perfect Cube**: A number is a perfect cube if it can be expressed as the cube of an integer. For example, 27 is a perfect cube because it is 3³.

## Example
For `n = 10`, the expected output would be:
```
1
2
3
Square
5
6
7
Cube
9
10
```
