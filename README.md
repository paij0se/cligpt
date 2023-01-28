<h1>cligpt</h1>

<h1>Installation</h1>

Go 1.17+
```
go install github.com/paij0se/cligpt@latest
```

test

```bash
cligpt "Do a poem about Golang in Latin"

Nunc est tempus, ut in Golang gradimur
Mox nostris verbis linguam didicimus
Dum in codice nos ponimus,
Cuncta sic apte nos iungimus.

Verba nova nos scrutamur
Et in praxi illa provamus
Ut optimus sic codex fiat,
Hic scimus nos invenire quod quaerimus.

Golang saepe gloriosa est,
Mentem nostram stimulat et excitat
Nobis efficit ut mira creemus
Et magna in mundo fata sequamur.
```

<h1>Configuration</h1>

- If is the first time you run the cli, You are going to see a error message, so, insert the OpenAI token in your config directory that is located in `$HOME/.config/cligpt/cligpt.yml` And in Windows is located in `C:\Users\user\AppData\Roaming\cligpt`

```haskell
auth: token
model: text-davinci-003
max_tokens: 2000
```

<h1>Building</h1>

- Clone the repository.

`git clone https://github.com/paij0se/cligpt`

- Build the cli.

```bash
go build
```
