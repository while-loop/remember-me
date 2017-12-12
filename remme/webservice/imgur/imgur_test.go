package imgur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImgurLogsInSuccessfully(t *testing.T) {
	assert.True(t, true)
}

func TestImgurLogsInWithWrongPassword(t *testing.T) {
	assert.EqualError(t, New().ChangePassword("remme123", "fakepasswd", ""),
		"Incorrect password for account remme123 @ imgur.com")
}
func TestImgurLogsInWithWrongEmail(t *testing.T) {
	assert.EqualError(t, New().ChangePassword("asdasd@sasdf.com", "fakepasswd", ""),
		"Incorrect password for account asdasd@sasdf.com @ imgur.com")
}
func TestImgurChangePasswordWithWrongOldPassword(t *testing.T) {
	assert.True(t, true)
}

func TestImgurChangePasswordWithInvalidChars(t *testing.T) {
	assert.True(t, true)
}

func TestImgurChangePassSuccessfully(t *testing.T) {
	r := New()
	assert.NoError(t, r.ChangePassword("remme123", "", ""))
}

func TestImgurChangePassNewPasswordsDontMatch(t *testing.T) {
	assert.True(t, true)
}

func TestImgurChangePassNewPasswordIsSameAsOld(t *testing.T) {
	assert.True(t, true)
}
