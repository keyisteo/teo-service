package controllers

type LogoutController struct {
	ExtendedController // Inherit our, slightly extended controller
}

func (c *LogoutController) Get() {
	// Check if user is logged in
	userData := c.Session.Get("UserData")

	if userData != nil {
		// User already login
		c.Session.Delete("UserData")
		c.Ctx.Redirect(302, "/")
	}
	c.Ctx.Redirect(302, "/")
}
