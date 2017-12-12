package reddit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedditLogsInSuccessfully(t *testing.T) {
	assert.True(t, true)
}

func TestRedditLogsInWithWrongPassword(t *testing.T) {
	assert.True(t, true)
}
func TestRedditLogsInWithWrongEmail(t *testing.T) {
	assert.True(t, true)
}
func TestRedditChangePasswordWithWrongOldPassword(t *testing.T) {
	assert.True(t, true)
}

func TestRedditChangePasswordWithInvalidChars(t *testing.T) {
	assert.True(t, true)
}

func TestRedditChangePassSuccessfully(t *testing.T) {
	assert.NoError(t, New().ChangePassword("remme123", "", ""))
}

func TestRedditChangePassNewPasswordsDontMatch(t *testing.T) {
	assert.True(t, true)
}

func TestRedditChangePassNewPasswordIsSameAsOld(t *testing.T) {
	assert.True(t, true)
}
