# Talken

Talken is a modular, peer-to-peer messaging. It provides end-to-end encryption, identity management via public/private key pairs, and a peer discovery layer—without any built-in user interface. Talken is intended to be embedded in client applications (CLI tools, web apps, desktop or mobile applications) that handle transport, persistence, and presentation.

> **Project status:** Documentation and scaffolding only. Core implementation is forthcoming.

## Features

- **Key-based identity**  
  Each user generates a public/private key pair. The public key serves as the user’s address; no usernames or passwords are required.

- **End-to-end encryption**  
  All messages are encrypted with the recipient’s public key and signed by the sender’s private key. No intermediaries can read message contents.

- **Decentralized peer discovery (P2P/mesh)**  
  Supports bootstrap nodes or a DHT for initial peer lookup. Subsequent message routing can proceed via direct or multi-hop connections.

- **Protocol-agnostic core**  
  Talken provides only the messaging logic. Transport (TCP, UDP, WebSocket, QUIC, etc.) and UI layers are the responsibility of client applications.

## Contributing

Contributions are welcome at the design and documentation stage:

1. Open an issue if you identify a missing feature or need clarification.
2. Propose changes to project structure, module interfaces, or workflows.
3. Submit pull requests for documentation, templates, or prototype code.

Please follow the templates provided in `ISSUE` and `PR` termplate and adhere to the project’s licensing.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for full terms.
