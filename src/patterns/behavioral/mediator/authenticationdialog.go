package mediator


import (
	// std
	"math/rand"
)


type AuthenticationDialog struct {
	title string
	loginOrRegisterChkBx *CheckBox
	loginUsername, loginPassword *TextBox
	registrationUsername, registrationPassword, registrationEmail *TextBox
	okBtn, cancelBtn *Button
}


func NewAuthenticationDialog(title string) *AuthenticationDialog {
	return &AuthenticationDialog{
		title: title,
	}
}

func (self *AuthenticationDialog) Notify(sender *Component, event string) {
	if (sender.GetComponentType() == CHECKBOX && event == "check") {
		if (self.loginOrRegisterChkBx.IsChecked()) {
			self.title = "Log in"
		} else {
			self.title = "Register"
		}
	}

	if (sender.GetComponentType() == BUTTON && event == "click") {
		if (self.loginOrRegisterChkBx.IsChecked()) {
			self.title = "Log in"
			
			if (rand.Intn(2) == 0) {
				self.title += "\tNot Found User"
			}
			self.title += "\tFound User"
			
		} else {
			self.title = "Register"
		}
	}
}

func (self *AuthenticationDialog) GetTitle() string {
	return self.title
}

func (self *AuthenticationDialog) SetTitle(title string) {
	self.title = title
}

func (self *AuthenticationDialog) GetLoginOrRegisterChkBx() *CheckBox {
	return self.loginOrRegisterChkBx
}

func (self *AuthenticationDialog) SetLoginOrRegisterChkBx(loginOrRegisterChkBx *CheckBox) {
	self.loginOrRegisterChkBx = loginOrRegisterChkBx
}

func (self *AuthenticationDialog) GetLoginUsername() *TextBox {
	return self.loginPassword
}

func (self *AuthenticationDialog) SetLoginUsername(loginPassword *TextBox) {
	self.loginPassword = loginPassword
}

func (self *AuthenticationDialog) GetLoginPassword() *TextBox {
	return self.loginPassword
}

func (self *AuthenticationDialog) SetLoginPassword(loginPassword *TextBox) {
	self.loginPassword = loginPassword
}

func (self *AuthenticationDialog) GetRegistrationUsername() *TextBox {
	return self.registrationUsername
}

func (self *AuthenticationDialog) SetRegistrationUsername(registrationUsername *TextBox) {
	self.registrationUsername = registrationUsername
}

func (self *AuthenticationDialog) GetRegistrationPassword() *TextBox {
	return self.registrationPassword
}

func (self *AuthenticationDialog) SetRegistrationPassword(registrationPassword *TextBox) {
	self.registrationPassword = registrationPassword
}

func (self *AuthenticationDialog) GetRegistrationEmail() *TextBox {
	return self.registrationEmail
}

func (self *AuthenticationDialog) SetRegistrationEmail(registrationEmail *TextBox) {
	self.registrationEmail = registrationEmail
}

func (self *AuthenticationDialog) GetOkBtn() *Button {
	return self.okBtn
}

func (self *AuthenticationDialog) SetOkBtn(okBtn *Button) {
	self.okBtn = okBtn
}

func (self *AuthenticationDialog) GetCancelBtn() *Button {
	return self.cancelBtn
}

func (self *AuthenticationDialog) SetCancelBtn(cancelBtn *Button) {
	self.cancelBtn = cancelBtn
}