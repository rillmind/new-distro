# New Distro

New distro é um projeto pessoal bem simples que cria um menu interativo
no terminal para escolher entre as opções de scripts que eu uso quando
acabo de instalar uma das distros que eu mais uso. Os scripts estão nos
meus dotfiles. E, basicamente, esse código só funciona para os meus scripts
por causa dos caminhos fixos no código. Esse código pode ser adaptado para
qualquer tipo de script.

Resumindo eu utilizo a biblioteca huh, de @charm_, para gerar o menu interativo,
e lipgloss, também de @charm_, para estilo. E uma outra biblioteca que eu mesmo fiz,
colors, só para não ter que ficar copiando as cores que eu gosto de usar.
Não tem nada além das cores estilizadas.

## O que faz

1. Pergunta qual distro está sendo configurada (Arch ou Fedora) e roda com `bash` o script de
   pós-instalação correspondente em `~/dotfiles/.scripts/linux/` (`ArchInitInstall.sh` ou
   `FedoraInitInstall.sh`).
2. Pergunta se deve instalar os Flatpaks e, se a resposta for "sim", roda
   `~/dotfiles/.scripts/linux/FlatpakEssentials.sh`.
3. Limpa a tela.

Os scripts herdam o terminal, então dá para responder às perguntas deles normalmente.

**Atenção:** nos [dotfiles](https://github.com/rillmind/dotfiles) atuais, os scripts mudaram de
lugar (`arch/ArchInitInstall.sh`, `fedora/InitInstall.sh`, `install/FlatpakEssential.sh`), então
os caminhos fixos em `main.go` não os encontram. Além disso, a etapa de Flatpak chama
`bash ~ dotfiles/...` em vez de `bash ~/dotfiles/...` e falha com "É um diretório" (exit 126).

## Instalação

Pré-requisitos: Go 1.26.1+ e os [dotfiles](https://github.com/rillmind/dotfiles) clonados em
`~/dotfiles`.

```sh
git clone https://github.com/rillmind/new-distro.git
cd new-distro
go build -o new-distro .
./new-distro
```

Os scripts chamados instalam pacotes e mudam o sistema: rode só numa instalação nova.

## Como testar

Não há testes automatizados. `go vet ./...` confere o código. Para testar o menu sem executar os
scripts de verdade, aponte o `HOME` para uma pasta com scripts falsos:

```sh
mkdir -p /tmp/nd/dotfiles/.scripts/linux
for s in ArchInitInstall FedoraInitInstall FlatpakEssentials; do
  echo "echo executou $s" > /tmp/nd/dotfiles/.scripts/linux/$s.sh
done
go build -o new-distro . && HOME=/tmp/nd ./new-distro
```

Escolha "Fedora" e confira que aparece `executou FedoraInitInstall`. Depois responda "não" ao
Flatpak: o programa limpa a tela e sai. Respondendo "sim", hoje ele termina com
`Erro ao executar o script: exit status 126`, por causa do bug citado acima.
