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
		distro  string
		flatpak string
	)

	home, err := os.UserHomeDir()

	archLabel := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.BlueColor)).Render("Arch")
	fedoraLabel := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.SapphireColor)).Render("Fedora")

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Qual distro está configurando?").
				Options(
					huh.NewOption(archLabel, filepath.Join(home, "dotfiles/.scripts/linux/ArchInitInstall.sh")),
					huh.NewOption(fedoraLabel, filepath.Join(home, "dotfiles/.scripts/linux/FedoraInitInstall.sh")),
				).
				Value(&distro),
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

	form2 := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Deseja instalar pacotes flatpak agora?").
				Options(
					huh.NewOption("sim", "sim"),
					huh.NewOption("não", "não"),
				).
				Value(&flatpak),
		),
	)

	err = form2.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmd2 := exec.Command("bash", home, "dotfiles/.scripts/linux/FlatpakEssentials.sh")

	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	cmd2.Stdin = os.Stdin

	if flatpak == "sim" {
		err = cmd2.Run()
		if err != nil {
			log.Fatalf("Erro ao executar o script: %v", err)
		}
	}

	cmd3 := exec.Command("clear")

	cmd3.Stdout = os.Stdout

	err = cmd3.Run()
	if err != nil {
		log.Fatalf("Erro ao executar o comando: %v", err)
	}
}
