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

// GetSponsor *
func (m Meeting) GetSponsor() string {
	return m.Sponsor
}

// GetTitle *
func (m Meeting) GetTitle() string {
	return m.Title
}

// GetStartTime *
func (m Meeting) GetStartTime() Date {
	return m.StartTime
}

// GetEndTime *
func (m Meeting) GetEndTime() Date {
	return m.EndTime
}

// GetParticipators *
func (m Meeting) GetParticipators() []string {
	return m.Participators
}

// SetSponsor *
func (m *Meeting) SetSponsor(sponsor string) {
	m.Sponsor = sponsor
}

// SetTitle *
func (m *Meeting) SetTitle(title string) {
	m.Title = title
}

// SetStartTime *
func (m *Meeting) SetStartTime(startTime Date) {
	m.StartTime.Assign(startTime)
}

// SetEndTime *
func (m *Meeting) SetEndTime(endTime Date) {
	m.EndTime.Assign(endTime)
}

// SetParticipators *
func (m *Meeting) SetParticipators(participators []string) {
	m.Participators = append([]string{}, participators...)
}

// AddParticipator *
func (m *Meeting) AddParticipator(participator string) {
	m.Participators = append(m.Participators, participator)
}

// RemoveParticipator *
func (m *Meeting) RemoveParticipator(participator string) {
	// find the index of participator
	for i, p := range m.Participators {
		if p == participator {
			m.Participators = append(m.Participators[:i], m.Participators[i+1:]...)
			break
		}
	}
}

// IsParticipator *
func (m *Meeting) IsParticipator(username string) bool {
	for _, p := range m.Participators {
		if p == username {
			return true
		}
	}
	return false
}

// Assign *
func (m *Meeting) Assign(meeting Meeting) {
	m.Sponsor = meeting.Sponsor
	m.Title = meeting.Title
	m.StartTime.Assign(meeting.StartTime)
	m.EndTime.Assign(meeting.EndTime)
	m.Participators = append([]string{}, meeting.Participators...)
}
