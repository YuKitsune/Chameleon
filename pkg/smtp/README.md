# pkg/SMTP
The vast majority of this code is based off of [go-guerrilla](https://github.com/flashmob/go-guerrilla) (as of commit `aa54b3ac4a0b4b34232fd29239422d024ad9395e`),
with slight modifications here and there to better accommodate Chameleons needs.

Thank you to those who have worked so hard to make go-guerrilla possible!

## Todos
- ~~Review [go-guerrilla workers and processors](https://github.com/flashmob/go-guerrilla/wiki/Backends,-configuring-and-extending) for thread safety~~

## Notes
`go-guerrilla` implemented [workers and processors](https://github.com/flashmob/go-guerrilla/wiki/Backends,-configuring-and-extending) as part of handling mail.
I don't think we will require such a system, but if we do at some point in the future, then we will copy/paste the implementation
or just import `go-guerrilla` directly.
