package gentypes

import "testing"

func TestNotification(t *testing.T) {
	notification := NewNotification()
	notification.Done()
	notification.Wait()
}
