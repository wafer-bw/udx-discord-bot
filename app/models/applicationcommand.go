package models

// https://discord.com/developers/docs/interactions/slash-commands#applicationcommand

// ApplicationCommand - The base commmand model that belongs to an application
type ApplicationCommand struct {
	ID            string                      `json:"id"`
	ApplicationID string                      `json:"application_id"`
	Name          string                      `json:"name"`
	Description   string                      `json:"description"`
	Options       []*ApplicationCommandOption `json:"options"`
}

// ApplicationCommandOption - The parameters for the command
type ApplicationCommandOption struct {
	Type        ApplicationCommandOptionType      `json:"type"`
	Name        string                            `json:"name"`
	Description string                            `json:"description"`
	Required    bool                              `json:"boolean"`
	Choices     []*ApplicationCommandOptionChoice `json:"choices"`
	Options     []*ApplicationCommandOption       `json:"options"`
}

// ApplicationCommandOptionChoice - User choice for `string` and/or `int` type options
type ApplicationCommandOptionChoice struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ApplicationCommandOptionType - Types of command options
type ApplicationCommandOptionType int

// ApplicationCommandOptionTypeEnum - Acts as an enum struct of all `ApplicationCommandOptionType`s
type ApplicationCommandOptionTypeEnum struct {
	SubCommand      ApplicationCommandOptionType
	SubCommandGroup ApplicationCommandOptionType
	String          ApplicationCommandOptionType
	Integer         ApplicationCommandOptionType
	Boolean         ApplicationCommandOptionType
	User            ApplicationCommandOptionType
	Channel         ApplicationCommandOptionType
	Role            ApplicationCommandOptionType
}

// ApplicationCommandOptionTypes - `ApplicationCommandOptionTypeEnum`
var ApplicationCommandOptionTypes = &ApplicationCommandOptionTypeEnum{
	SubCommand:      1,
	SubCommandGroup: 2,
	String:          3,
	Integer:         4,
	Boolean:         5,
	User:            6,
	Channel:         7,
	Role:            8,
}
