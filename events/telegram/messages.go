package events

const msgHelp = `I can save and keep you pages. Also i can offer you them read.
In order to save the page, just send me all link.
In order to get a random page from your list, send me command /rnd.
Caution! After that, this page will be removed from your list!`

const msgHello = "HI there! \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown Command"
	msgNoSavedPages   = "You have no saved pages"
	msgSaved          = "Saved!"
	msgAlreadyExists  = "You already have this page in your list"
)
