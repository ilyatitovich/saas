| Feature                      | Why It's Important                               |
| ---------------------------- | ------------------------------------------------ |
| 🕵️ Rate Limiting             | Prevent brute force or spam attacks              |
| ⏳ Timeout & Size Limits     | Prevent slowloris or large payload DoS           |
| 🌍 Unicode normalization     | Avoid homoglyph attacks (e.g. Cyrillic vs Latin) |
| 🔒 Password hashing          | Always hash passwords (use bcrypt/argon2)        |
| 🧼 Trim/Normalize inputs     | Avoid whitespace, casing, duplicate records      |
| ✉️ Email confirmation        | Avoid fake accounts                              |
| 📱 Phone OTP validation      | Validate real numbers, not throwaways            |
| 👮 ReCAPTCHA / Bot filtering | For public APIs                                  |
