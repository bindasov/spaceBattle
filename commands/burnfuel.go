package commands

type BurnFuelCommand interface {
	Execute() error
	SetLeftover(leftover int)
}

func NewBurnFuel(leftover int, consumption int) BurnFuelCommand {
	burnFuelCommand := &burnFuelCommand{
		leftover:    leftover,
		consumption: consumption,
	}
	return burnFuelCommand
}

type burnFuelCommand struct {
	leftover    int
	consumption int
}

func (c *burnFuelCommand) Execute() error {
	c.leftover -= c.consumption
	return nil
}

func (c *burnFuelCommand) SetLeftover(leftover int) {
	c.leftover = leftover
}
