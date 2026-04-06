# candy-coin

`candy-coin` is a private blockchain implementation currently being prototyped for the **Competitive Programming Society** and various other decentralized applications.

---

This project serves as a proof-of-concept for a lightweight, secure, and transparent ledger system. It is designed to demonstrate core blockchain principles including:

* **Proof of Work (PoW):** Secure block mining with adjustable difficulty.
* **Deterministic Hashing:** Consistent SHA-256 integrity checks across different platforms. (Yes, yes, I know SHA-256 is not the best choice here but it's a prototype)
* **Immutable Ledger:** A linked-list structure ensuring data cannot be altered once mined.

The project is in the **prototype phase**. I am currently refining the core Go engine, specifically focusing on:
* Block generation and validation logic.
* Transaction structures for competitive scoring or reward systems.
* Internal networking and consensus rules.

While primarily developed for GUCPS, the architecture of `candy-coin` is being built to support:
* Automated reward distribution for coding challenges.
* Educational demonstrations of distributed systems.

---
*Note: This is a private blockchain prototype and is not intended for use as a public financial instrument.*
