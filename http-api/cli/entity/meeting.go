package entity

import (
	"time"
)

type Meeting struct {
	m_sponsor string
	m_participators []string
	m_startDate time.Time
	m_endDate time.Time
	m_title string
}

func (m *Meeting) IsParticipator(t_username string) bool {
	for _, value := range m.m_participators {
		if value == t_username {
			return true
		}
	}
	return false
}
func (m *Meeting) SetSponsor(t_sponsor string) {
	m.m_sponsor = t_sponsor
}
func (m *Meeting) GetSponsor() string {
	return m.m_sponsor
}
func (m *Meeting) GetParticipator() []string{
	return m.m_participators
}
func (m *Meeting) SetParticipator(t_participators []string) {
	m.m_participators = t_participators
}
func (m *Meeting) GetStartDate() time.Time {
	return m.m_startDate
}
func (m *Meeting) SetStartDate(t_startDate time.Time) {
	m.m_startDate = t_startDate
}
func (m *Meeting) GetEndDate() time.Time{
	return m.m_endDate
}
func (m *Meeting) SetEndDate(t_endDate time.Time) {
	m.m_endDate = t_endDate
}
func (m *Meeting) GetTittle() string {
	return m.m_title
}
func (m *Meeting) SetTittle(t_tittle string) {
	m.m_title = t_tittle
}