package commands

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

// ApplicationCommandOptionType enums
const (
	SubCommand      ApplicationCommandOptionType = 1
	SubCommandGroup ApplicationCommandOptionType = 2
	String          ApplicationCommandOptionType = 3
	Integer         ApplicationCommandOptionType = 4
	Boolean         ApplicationCommandOptionType = 5
	User            ApplicationCommandOptionType = 6
	Channel         ApplicationCommandOptionType = 7
	Role            ApplicationCommandOptionType = 8
)
