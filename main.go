package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/rillmind/colors"
)

func main() {
	var (
		distro string
	)

	home, err := os.UserHomeDir()

	fedoraLabel := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.SapphireColor)).Render("Fedora")
	archLabel := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.BlueColor)).Render("Arch")

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Which distro?").
				Options(
					huh.NewOption(fedoraLabel, filepath.Join(home, ".scripts/linux/FedoraInitInstall.sh")),
					huh.NewOption(archLabel, filepath.Join(home, ".scripts/linux/ArchInitInstall.sh")),
				).
				Value(&distro).
				WithTheme(huh.ThemeCatppuccin()),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("bash", distro)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Erro ao executar o script: %v", err)
	}
}
