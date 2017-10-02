package webservice

import (
	"testing"
	_ "github.com/stretchr/testify/assert"
)

// Test cases
// Logged in successfully
// Logged in wrong passwd
// Logged in wrong email
// Logged in get login item to validate TODO?

// Change password wrong password
// Change passwd inv format
// change passws not matching
// change pass success
// change pass same pass
// change pass prev pass

func TestLogsInSuccessfully(t *testing.T)                 {}
func TestLogsInWithWrongPassword(t *testing.T)            {}
func TestLogsInWithWrongEmail(t *testing.T)               {}
func TestChangePasswordWithWrongOldPassword(t *testing.T) {}
func TestChangePasswordWithInvalidChars(t *testing.T)     {}
func TestChangePassSuccessfully(t *testing.T)             {}
func TestChangePassNewPasswordsDontMatch(t *testing.T)    {}
func TestChangePassNewPasswordIsSameAsOld(t *testing.T)   {}
