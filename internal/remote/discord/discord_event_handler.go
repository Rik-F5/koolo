package discord

import (
	"bytes"
	"context"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/event"
)

func (b *Bot) Handle(_ context.Context, e event.Event) error {
	if b.shouldPublish(e) {

		switch e.(type) {
		case event.GameCreatedEvent, event.GameFinishedEvent, event.RunStartedEvent, event.RunFinishedEvent:
			_, err := b.discordSession.ChannelMessageSend(b.channelID, e.Message())
			return err
		default:
			break
		}

		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, e.Image(), &jpeg.Options{Quality: 80})
		if err != nil {
			return err
		}

		_, err = b.discordSession.ChannelMessageSendComplex(b.channelID, &discordgo.MessageSend{
			File:    &discordgo.File{Name: "Screenshot.jpeg", ContentType: "image/jpeg", Reader: buf},
			Content: e.Message(),
		})

		return err
	}

	return nil
	//_, err := b.discordSession.ChannelMessageSend(b.channelID, e.Message())

	//return err
	return nil
}

func (b *Bot) shouldPublish(e event.Event) bool {

	switch evt := e.(type) {
	case event.GameFinishedEvent:
		if evt.Reason == event.FinishedChicken && !config.Koolo.Discord.EnableDiscordChickenMessages {
			return false
		}
		if evt.Reason == event.FinishedOK && !config.Koolo.Discord.EnableRunFinishMessages {
			return false
		}
		if evt.Reason == event.FinishedError && !config.Koolo.Discord.EnableGameCreatedMessages {
			return false
		}
	case event.RunFinishedEvent:
		if evt.Reason == event.FinishedChicken && !config.Koolo.Discord.EnableDiscordChickenMessages {
			return false
		}
		if evt.Reason == event.FinishedOK {
			return false // supress game finished messages until we add proper option for it
		}
		if evt.Reason == event.FinishedError && !config.Koolo.Discord.EnableGameCreatedMessages {
			return false
		}
		return true
	case event.GameCreatedEvent:
		if !config.Koolo.Discord.EnableGameCreatedMessages {
			return false
		}
	case event.RunStartedEvent:
		if !config.Koolo.Discord.EnableGameCreatedMessages {
			return false
		}
	case event.ItemStashedEvent:
		if evt.Item.Item.Name == "" {
			return false
		}

		// Check if item name is in blacklist
		execPath, err := os.Executable()
		if err != nil {
			return true
		}
		execDir := filepath.Dir(execPath)
		blacklistFile := filepath.Join(execDir, "config", "blacklist.txt")
		if _, err := os.Stat(blacklistFile); err == nil {
			content, err := os.ReadFile(blacklistFile)
			if err == nil {
				blacklistedItems := strings.Split(string(content), "\n")
				itemName := strings.ToLower(string(evt.Item.Item.Name))
				for _, item := range blacklistedItems {
					if strings.Contains(itemName, strings.ToLower(strings.TrimSpace(item))) {
						return false
					}
				}
			}
		}
		return true
	}

	return e.Image() != nil
}
