package entity

// Meeting datamoded of meeting
type Meeting struct {
	Sponsor       string
	Title         string
	StartDate     string
	EndDate       string
	Participators []string
}

func (m *Meeting) initMeeting(tSponsor, tTitle, tStartDate, tEndDate string, tParticipator []string) {
	m.Sponsor = tSponsor
	m.Title = tTitle
	m.StartDate = tStartDate
	m.EndDate = tEndDate
	m.Participators = tParticipator

}

func (m Meeting) getSponsor() string {
	return m.Sponsor
}

func (m *Meeting) setSponsor(s string) {
	m.Sponsor = s
}

func (m Meeting) getParticipators() []string {
	return m.Participators
}

func (m *Meeting) setParticipators(p []string) {
	m.Participators = p
}
func (m Meeting) getStartDate() string {
	return m.StartDate
}

func (m *Meeting) setStartDate(s string) {
	m.StartDate = s
}
func (m Meeting) getEndDate() string {
	return m.EndDate
}

func (m *Meeting) setEndDate(e string) {
	m.EndDate = e
}
func (m Meeting) getTitle() string {
	return m.Title
}

func (m *Meeting) setTitle(t string) {
	m.Title = t
}

// 返回会议成员数目
func (m *Meeting) getMeetingMemberNumber() int {
	return len(m.Participators)
}

func (m Meeting) isParticipator(name string) bool {
	for _, t := range m.Participators {
		if t == name {
			return true
		}
	}
	return false
}
