package entity

// Meeting an entity with a sponsor, title, start time, end time
// and a list of participators
type Meeting struct {
	Sponsor       string
	Title         string
	StartTime     Date
	EndTime       Date
	Participators []string
}

// NewMeeting *
func NewMeeting(sponsor string, title string, startTime Date, endTime Date, participators []string) *Meeting {
	return &Meeting{
		Sponsor:       sponsor,
		Title:         title,
		StartTime:     startTime,
		EndTime:       endTime,
		Participators: append([]string{}, participators...),
	}
}

func (m Meeting) getSponsor() string {
	return m.Sponsor
}

func (m Meeting) getTitle() string {
	return m.Title
}

func (m Meeting) getStartTime() Date {
	return m.StartTime
}

func (m Meeting) getEndTime() Date {
	return m.EndTime
}

func (m Meeting) getParticipators() []string {
	return m.Participators
}

func (m *Meeting) setSponsor(sponsor string) {
	m.Sponsor = sponsor
}

func (m *Meeting) setTitle(title string) {
	m.Title = title
}

func (m *Meeting) setStartTime(startTime Date) {
	m.StartTime.assign(startTime)
}

func (m *Meeting) setEndTime(endTime Date) {
	m.EndTime.assign(endTime)
}

func (m *Meeting) setParticipators(participators []string) {
	m.Participators = append([]string{}, participators...)
}

func (m *Meeting) addParticipator(participator string) {
	m.Participators = append(m.Participators, participator)
}

func (m *Meeting) removeParticipator(participator string) {
	// find the index of participator
	for i, p := range m.Participators {
		if p == participator {
			m.Participators = append(m.Participators[:i], m.Participators[i+1:]...)
			break
		}
	}
}

func (m *Meeting) isParticipator(username string) bool {
	for _, p := range m.Participators {
		if p == username {
			return true
		}
	}
	return false
}

func (m *Meeting) assign(meeting Meeting) {
	m.Sponsor = meeting.Sponsor
	m.Title = meeting.Title
	m.StartTime.assign(meeting.StartTime)
	m.EndTime.assign(meeting.EndTime)
	m.Participators = append([]string{}, meeting.Participators...)
}
