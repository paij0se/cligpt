<h1>cligpt</h1>

- Clone the repository.

`$ git clone https://github.com/paij0se/cligpt`

- Build the cli.

`./build.sh`

- Run `./cligpt` for create the config directory.

`$ ./cligpt "create the config directory cuh"`

- You are going to see a error message, so, insert the OpenAI token in your config directory that is located in `$HOME/.config/cligpt/cligpt.yml`

```haskell
auth: token
model: text-davinci-003
max_tokens: 256
```

![image](https://user-images.githubusercontent.com/69026987/209194859-a2456a7d-796f-47e0-8a8e-062848e2cbaf.png)
