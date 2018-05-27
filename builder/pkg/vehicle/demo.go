package vehicle

// RunDemo - take it for a spin!
func RunDemo(c Actions) error {
	err := c.Check()

	if err != nil {
		return err
	}
	err = c.Drive()
	if err != nil {
		return err
	}

	err = c.Stop()
	if err != nil {
		return err
	}

	return nil
}
