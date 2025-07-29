package state

// Global state instance
var AppState = &State{}

type State struct {
	LocationNamePage int `json:"locationNamePage"`
}

func (s *State) GetLocationNamePage() int {
	if s.LocationNamePage == 0 {
		s.LocationNamePage = 1
	}
	return s.LocationNamePage
}

func (s *State) IncrementLocationPage() {
	s.LocationNamePage += 1
}

func (s *State) DecrementLocationPage() {
	s.LocationNamePage -= 1
}

func (s *State) ResetLocationPage() {
	s.LocationNamePage = 1
}
