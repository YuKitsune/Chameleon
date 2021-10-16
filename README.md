> 🏗️ This is a work-in-progress, contributions are welcomed! 🚧

# Chameleon
Chameleon is an email cloaking service written in Go, and React (TypeScript).

# Current Todos
- [ ] API
  - [ ] Users
    - [ ] Sign-up
    - [ ] Log-in
      - [ ] TOTP support 
      - [ ] U2F support
    - [ ] Update details
    - [ ] Migrate account (Moving from cloud to self-hosted instance)
      - [ ] Need to flesh this out
    - [ ] Delete account
  - [ ] Aliases
    - [X] Create
      - [ ] Automatically determine sender on first received message
    - [X] Find by Sender and Recipient
      - [ ] Available to MTD via API 
    - [ ] List all
    - [X] Update
    - [X] Delete
    - [ ] Statistics
      - [ ] What would be useful to know here?
      - [ ] Matches sender addresses (E.g. 50% marketing@xyz.com, 25% security@... etc.)
      - [ ] Failed matches (Potential leaks)
    - [ ] Customisable aliases (Required higher tier for cloud) 
    - [ ] Verification
      - [ ] Can't forward mail unless the recipient address has been verified
    - [ ] Restrict access to owner
  - [ ] Temporary mail storage
    - [ ] Temporarily store mail if it:
      - [ ] fails to deliver
        - [ ] Retry sending every N minutes (default: 1)   
          - [ ] N can be user configurable
          - [ ] Throttle N after X attempts 
          - [ ] Give up after Y attempts
        - [ ] Users can choose to have **all** mail stored by default (Higher tier required for cloud)
    - [ ] Encrypt using S/MIME
      - [ ] Need to flesh this out, S/MIME is weird
  - [ ] Quarantined mail
    - [ ] Restrict access to owner
    - [ ] Retention policies
      - [ ] keep for up-to N days (default: 30)
        - [ ] N can be user configurable
      - [ ] Users can choose to have **all** mail stored by default (Higher tier required for cloud)
  - [ ] Notifications
    - [ ] Cloud only
      - [ ] Approaching/reached maximum number of aliases
      - [ ] Approaching/reached temporary storage limit
      - [ ] Billing reminders
    - [ ] New entry in quarantine
    - [ ] Security
      - [ ] Failed login attempt 
      - [ ] Suspicious login attempt 
  - [ ] Testing
    - [ ] Move mongodb calls behind a repository pattern so testing can be easier
    - [ ] Can re-work some existing tests so that they test the handlers instead of the API itself
- [ ] Frontend
  - [ ] We need one...
  - [ ] Spike out a PoC and work backwards to the API
- [ ] MTD (Mail Transfer Daemon)
  - [ ] Investigate:
    - [ ] Can we rip this out and just use the original go-guerilla? 
  - [ ] Accept incoming mail
    - [ ] Support for TLS
    - [ ] Support for S/MIME
      - [ ] Need to flesh this out, S/MIME is weird
    - [ ] Fetch alias from API based on sender and recipient address
      - [ ] If there's a match:
        - [ ] Forward the message to the intended recipient
          - [ ] If forwarding succeeds
            - [ ] Respond with a success code/message (Whatever SMTP likes)
          - [ ] Otherwise:
            - [ ] Store in the temporary store via the API
      - [ ] Otherwise:
        - [ ] Create a new quarantine entry via the API
        - [ ] Respond with an unsuccessful code/message (Whatever SMTP likes)
  - [ ] Try sending mail that may have previously failed to send
    - [ ]
  - [ ] Testing
    - [ ] Need some stress/smoke testing in place
- [ ] Post MVP
  - [ ] Split `pkg/mediator` into a separate repo
  - [ ] Clean up `cobra`/`viper` integration
  - [ ] Implement Prometheus and Grafana for monitoring
  - [ ] Custom domain support
  - [ ] Reply from alias
  - [ ] Native apps
    - [ ] Electron
    - [ ] iOS/Android
    - [ ] CLI

# Contributing

Contributions are what make the open source community such an amazing place to be, learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`feature/AmazingFeature`)
3. Commit your Changes
4. Push to the Branch
5. Open a Pull Request
