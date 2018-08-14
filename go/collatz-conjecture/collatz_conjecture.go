package collatzconjecture

import "errors"

func CollatzConjecture(x int) (int, error) {
  steps := 0

  if x <= 0 {
    return steps, errors.New("Input must be greater than 0.")
  }

  for {
    if x == 1  {
      break
    }

    if x % 2 == 0 {
      x = x / 2
    } else {
      x = (x * 3) + 1
    }
    steps++
  }

  return steps, nil
}
