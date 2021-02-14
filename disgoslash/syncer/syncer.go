package syncer

import (
	"log"

	"github.com/wafer-bw/udx-discord-bot/disgoslash/client"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/config"
	"github.com/wafer-bw/udx-discord-bot/disgoslash/slashcommands"
)

// Deps defines `Syncer` dependencies
type Deps struct{}

// impl implements `Syncer` properties
type impl struct {
	client client.Client
}

// Syncer interfaces `Syncer` methods
type Syncer interface {
	Run(guildIDs []string, slashCommandMap slashcommands.Map) error
}

type unregisterTarget struct {
	guildID   string
	commandID string
	name      string
}

// New returns a new `Syncer` interface
func New() Syncer {
	conf := config.New()
	client := client.New(&client.Deps{}, conf)
	return &impl{client: client}
}

// Run will reregister all of the provided slash commands
func (impl *impl) Run(guildIDs []string, slashCommandMap slashcommands.Map) error {
	if err := impl.unregisterCommands(guildIDs, slashCommandMap); err != nil {
		return err
	}
	if err := impl.registerCommands(slashCommandMap); err != nil {
		return err
	}
	return nil
}

func (impl *impl) registerCommands(commandMap slashcommands.Map) error {
	log.Println("Registering new commands...")
	for _, command := range commandMap {
		for _, guildID := range command.GuildIDs {
			if err := impl.client.CreateApplicationCommand(guildID, command.AppCommand); err != nil {
				return err
			}
			if guildID == "" {
				log.Printf("\t- Registered global command %s\n", command.Name)
			} else {
				log.Printf("\t- Registered command %s in guild %s\n", command.Name, guildID)
			}
		}
	}
	return nil
}

func (impl *impl) getCommandsToUnregister(guildIDs []string, commandMap slashcommands.Map) ([]unregisterTarget, error) {
	uniqueGuildIDs := impl.getUniqueGuildIDs(guildIDs, commandMap)
	unregisterTargets := []unregisterTarget{}
	log.Println("Collecting outdated commands...")
	for _, guildID := range uniqueGuildIDs {
		guildCommands, err := impl.client.ListApplicationCommands(guildID)
		if err != nil {
			return nil, err
		}
		for _, guildCommand := range guildCommands {
			if guildID == "" {
				log.Printf("\t- Collected global command %s\n", guildCommand.Name)
			} else {
				log.Printf("\t- Collected command %s in guild %s\n", guildCommand.Name, guildID)
			}
			unregisterTargets = append(unregisterTargets, unregisterTarget{
				guildID:   guildID,
				commandID: guildCommand.ID,
				name:      guildCommand.Name,
			})
		}
	}
	return unregisterTargets, nil
}

func (impl *impl) unregisterCommands(guildIDs []string, slashCommandMap slashcommands.Map) error {
	unregisterTargets, err := impl.getCommandsToUnregister(guildIDs, slashCommandMap)
	if err != nil {
		return err
	}

	log.Println("Unregistering outdated commands...")
	for _, target := range unregisterTargets {
		if err := impl.client.DeleteApplicationCommand(target.guildID, target.commandID); err != nil {
			return err
		}
		if target.guildID == "" {
			log.Printf("\t- Unregistered global command %s\n", target.name)
		} else {
			log.Printf("\t- Unregistered command %s from guild %s\n", target.name, target.guildID)
		}
	}
	return nil
}

func (impl *impl) getUniqueGuildIDs(guildIDs []string, commands slashcommands.Map) []string {
	uniqueGuildIDsMap := map[string]struct{}{
		"": {}, // include global
	}
	for _, id := range guildIDs {
		if _, ok := uniqueGuildIDsMap[id]; !ok {
			uniqueGuildIDsMap[id] = struct{}{}
		}
	}
	for _, command := range commands {
		for _, guildID := range command.GuildIDs {
			if _, ok := uniqueGuildIDsMap[guildID]; !ok {
				uniqueGuildIDsMap[guildID] = struct{}{}
			}
		}
	}
	uniqueGuildIDs := []string{}
	for id := range uniqueGuildIDsMap {
		uniqueGuildIDs = append(uniqueGuildIDs, id)
	}
	return uniqueGuildIDs
}
