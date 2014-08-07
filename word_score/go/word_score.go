package utils

import "math"

func WordScore(word string, pos int) int64 {
  score := float64(0)
  bytes := []byte(word)
  for _, b := range bytes {
    if pos < 0 {
      break
    }
    ns := float64(b) * math.Pow(256, float64(pos))
    score = score + ns
    pos = pos - 1
  }
  return int64(score)
}
