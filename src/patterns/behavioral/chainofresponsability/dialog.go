package chainofresponsability


// std
import (
	"fmt"
)


type Dialog struct {
	Container
	wikiPageURL string
}


func NewDialog(
	container *Container,
	tooltipText string,
	children []*Component,
	wikiPageURL string,
) *Dialog {
	return &Dialog{
		Container: Container{
			Component: Component{
				container: container,
				tooltipText: tooltipText,
			},
			children: children,
		},
		wikiPageURL: wikiPageURL,
	}
}

func (self *Dialog) ShowHelp() {
	if self.wikiPageURL == "" {
		self.Component.ShowHelp()
		return
	}

	fmt.Println("Wiki Page URL: " + self.wikiPageURL)
}

func (self *Dialog) GetWikiPageURL() string {
	return self.wikiPageURL
}

func (self *Dialog) SetWikiPageURL(wikiPageURL string) {
	self.wikiPageURL = wikiPageURL
}