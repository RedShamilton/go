package time

import "runtime"

// WaitUntil pauses the current goroutine until at least the time t.
// A past time causes Sleep to return immediately.
func WaitUntil(t Time) Duration {
  for {
    if Now().Before(t) {
      runtime.Gosched()
    } else {
      break
    }
  }

  return Now().Sub(t)
}
