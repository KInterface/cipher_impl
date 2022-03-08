# What's this ?

## This is a prototype implementation of token payment.

- Frontend: Typescrypt and WebAssembly(rust)
- Backend: Golang

<div><img src="https://upload.wikimedia.org/wikipedia/commons/d/d5/Rust_programming_language_black_logo.svg" />
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/4c/Typescript_logo_2020.svg/768px-Typescript_logo_2020.svg.png?20210506173343" width="130" />
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="150" />
</div>

## More Precisely...

# What's this?

## This is a prototype implementation of token payment.

- Frontend: Typescrypt and WebAssembly(rust)
- Backend: Golang

## More Precisely...

### We determined to use a stream cipher so-called **XChacha20-Poly1305**.

However, we couldn't find a decent implementation of the cipher by Javascript.

#### Thus We decided to implement encrypt and decrypt method by using rust - wasm.

This is a prototype implementation and we still do not send back any token to the client browser at present.
We are planning to cooperate with our other token calculation microservice,
thus this repository already satisfied our aim to try out implementing payment-flow despite incompleteness.


![Payment - Google Chrome 05 03 2022 23_43_05](https://user-images.githubusercontent.com/100127291/156889072-9b4d43d8-5807-4ac8-a737-4920cde54a03.png)
