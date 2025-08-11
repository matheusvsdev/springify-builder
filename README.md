# Springify Builder

**syb** é um CLI leve, rápido e interativo que gera arquivos de configuração para perfis de ambiente em projetos Spring Boot.

Ideal para devs que querem automatizar a criação dos arquivos `application.yml` ou `application.properties`, `docker-compose.yml` com suporte a múltiplos perfis (dev, test, prod) e configuração de banco de dados e serviços docker — tudo sem complexidade.

---

## Tecnologias utilizadas

O `springify-builder` foi desenvolvido em **Go (Golang)**, uma linguagem moderna e eficiente voltada para ferramentas robustas e multi-plataforma.

> A escolha pelo Go garante binários leves, rápidos e compatíveis com macOS, Linux e Windows — sem dependências externas.


## Funcionalidades

- Criação do arquivo principal: `application.yml` ou `application.properties`
- Geração de perfis separados: `application-dev.yml`, `application-prod.properties`, etc.
- Escolha interativa de ambiente, tipo de banco, nome da aplicação e formato
- Compatível com PostgreSQL, MySQL, Oracle e H2
- Criação de arquivo `docker/docker-compose.yml` e estrutura de serviços
- Binários prontos para **Linux**, **macOS (M1/M2)** e **Windows**

---

## Instalação

### ➤ macOS / Linux (instalação automática)

```bash
curl -sSL https://raw.githubusercontent.com/matheusvsdev/springify/main/install.sh | bash
```

Este instalador vai:

- Detectar seu sistema e arquitetura
- Baixar o binário correto da versão mais recente
- Mover o executável para /usr/local/bin/syb
- Validar se o comando foi instalado com sucesso

Depois disso, você pode usar syb direto no terminal.

### ➤ Windows (instalação manual)

1. Baixe o binário .zip pela aba Releases ou diretamente:

```bash
Invoke-WebRequest -Uri https://github.com/matheusvsdev/springify/releases/latest/download/syb-windows-amd64.zip -OutFile syb.zip
```
2. Extraia e renomeie:
```bash
Expand-Archive -Path syb.zip -DestinationPath .
Rename-Item -Path .\syb-windows-amd64.exe -NewName syb.exe
```

(Opcional) Adicione a pasta onde está o syb.exe ao seu PATH para usar de qualquer lugar no terminal.

## Como Usar

Execute no terminal:

```bash
syb profile create
```

Você será guiado interativamente para definir:

- Formato do arquivo: .yml ou .properties
- Nome da aplicação
- Ambiente: test, dev ou prod
- Tipo de banco: postgresql, mysql, oracle
- Host, porta e nome do banco (exceto para H2)

O CLI então criará:

- O arquivo principal `application.properties` ou `application.yml`
- O perfil escolhido com configurações específicas de banco

Ou criar `docker-compose.yml` adicinando serviços

```bash
syb compose add mysql
```

Você será guiado interativamente para definir:

- Nome do projeto
- Nome da network
- Nome do serviço
- Nome do container
- Nome do banco e porta

O CLI então criará:

- O arquivo principal `docker/docker-compose.yml`

## Prévia do CLI

Veja como o springify interage com você no terminal:

> O CLI guia você por perguntas rápidas — gerando perfis de ambiente com clareza e agilidade.

![Interface CLI](springify.png)

## Estrutura gerada

Os arquivos serão criados dentro da pasta resources do projeto Java Spring Boot

```plaintext
src/
└── main/
    └── resources/
        ├── application.yml               // principal (ou .properties)
        ├── application-dev.yml           // perfil gerado
        └── application-test.properties   // exemplo de outro perfil
```

O arquivo será criado dentro da pasta docker

```plaintext
project/
    └── docker/
        ├── docker-compose.yml
```

## Roadmap

Planejamos continuar evoluindo o `Springify Builder` com base no uso da comunidade e sugestões recebidas. Aqui estão algumas ideias previstas:

- [x] Geração interativa de perfis `.yml` e `.properties`
- [x] Compatibilidade com múltiplos bancos (PostgreSQL, MySQL, Oracle, H2)
- [x] Binários prontos para macOS, Linux e Windows
- [ ] Suporte multi-idioma no CLI
- [ ] Traduções da documentação (inglês, espanhol, francês...)
- [ ] Comando `syb init` para gerar estrutura completa do projeto
- [x] Geração automatizada de `docker-compose.yml`
- [ ] Customização via flags (sem perguntas interativas)
- [ ] Lançamento no Homebrew e Chocolatey (instalação via gerenciador)
- [ ] Testes automatizados e benchmark de performance

> Sinta-se à vontade para sugerir novas ideias ou abrir Pull Requests.  
> Esse projeto cresce junto com quem usa

## Comunidade

Esse projeto é uma ponte entre a comunidade Go (ferramentas CLI) e a comunidade Spring Boot. Criado para devs que valorizam produtividade, automação e simplicidade na configuração de ambientes.

## Contribuições

Quer colaborar com o projeto?

Confira as diretrizes em [CONTRIBUTING.md](CONTRIBUTING.md)

## Licença

Este projeto está sob a licença **Creative Commons BY-NC 4.0 (Atribuição – Não Comercial)**.

Você pode:

- Usar a ferramenta livremente em projetos pessoais, acadêmicos ou empresariais como auxiliar
- Modificar e estudar o código

Você **não pode**:

- Vender ou redistribuir o código do `springify-builder`
- Atribuir a ferramenta a terceiros como se fosse própria

> O `springify-builder` foi criado para facilitar o trabalho de quem desenvolve com Spring Boot.  
> Pode ser usado como utilitário auxiliar em empresas — exatamente como se usaria `npm install` ou uma imagem Docker.  
> A única exigência é **manter a atribuição** e não **revender ou publicar como se fosse seu**.  
> A licença existe apenas para proteger a autoria, não para limitar o uso legítimo da comunidade.

Saiba mais: [https://creativecommons.org/licenses/by-nc/4.0](https://creativecommons.org/licenses/by-nc/4.0)

Criado por [Matheus Valdevino](https://github.com/matheusvsdev)

---
