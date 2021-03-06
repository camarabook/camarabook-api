package models

import (
	"time"

	"github.com/databr/go-popolo"
	"gopkg.in/mgo.v2/bson"
)

// Parliamentarian

type Parliamentarian struct {
	Id               string                 `json:"id"`                // The person's unique identifier
	Name             string                 `json:"name"`              // A person's preferred full name
	OtherNames       []OtherNames           `json:"other_names"`       // Alternate or former names
	Identifiers      []Identifier           `json:"identifiers"`       // Issued identifiers
	FamilyName       string                 `json:"family_name"`       // One or more family names
	GivenName        string                 `json:"given_name"`        // One or more primary given names
	AdditionalName   string                 `json:"additional_name"`   // One or more secondary given names
	HonorificPrefix  string                 `json:"honorific_prefix"`  // One or more honorifics preceding a person's name
	HonorificSuffix  string                 `json:"honorific_suffix"`  // One or more honorifics following a person's name
	PatronymicName   string                 `json:"patronymic_name"`   // One or more patronymic names
	SortName         string                 `json:"sort_name"`         // A name to use in a lexicographically ordered list
	Email            string                 `json:"email"`             // A preferred email address
	Gender           string                 `json:"gender"`            // A gender
	BirthDate        popolo.Date            `json:"birth_date"`        // A date of birth
	DeathDate        popolo.DateTime        `json:"death_date"`        // A date of death
	Image            string                 `json:"image"`             // A URL of a head shot
	Summary          string                 `json:"summary"`           // A one-line account of a person's life
	Biography        string                 `json:"biography"`         // An extended account of a person's life
	NationalIdentify string                 `json:"national_identity"` // A national identity
	ContactDetails   []popolo.ContactDetail `json:"contact_details"`   // Means of contacting the person
	Links            []Link                 `json:"links"`             // URLs to documents about the person
	Memberships      []Membership           `json:"memberships"`       // Memberships
	CreatedAt        time.Time              `json:"created_at"`        // The time at which the resource was created
	UpdatedAt        time.Time              `json:"updated_at"`        // The time at which the resource was last modified
	Sources          []Source               `json:"sources"`           // URLs to documents from which the person is derived
}

type OtherNames struct {
	Name            string       `json:"name"`                       // [required] An alternate or former name
	FamilyName      string       `json:"family_name,omitempty"`      // One or more family names
	GivenName       string       `json:"given_name,omitempty"`       //One or more primary given names
	AdditionalName  string       `json:"additional_name,omitempty"`  // One or more secondary given names
	HonorificPrefix string       `json:"honorific_prefix,omitempty"` // One or more honorifics preceding a person's name
	HonorificSuffix string       `json:"honorific_suffix,omitempty"` // One or more honorifics following a person's name
	PatronymicName  string       `json:"patronymic_name,omitempty"`  // One or more patronymic names
	StartDate       *popolo.Date `json:"start_date,omitempty"`       // The date on which the name was adopted
	EndDate         *popolo.Date `json:"end_date,omitempty"`         // The date on which the name was abandoned
	Note            string       `json:"note,omitempty"`             // A note, e.g. 'Birth name'
}

type Link struct {
	Url  string `json:"url"` // A URL
	Note string `json:"rel"` // A note, e.g. 'Wikipedia page'
}

type Party struct {
	Id               string               `json:"id"`                         // The organization's unique identifier
	Name             string               `json:"name"`                       // A primary name, e.g. a legally recognized name
	OtherNames       []OtherNames         `json:"other_names,omitempty"`      // Alternate or former names
	Identifiers      []*Identifier        `json:"identifiers,omitempty"`      // Issued identifiers
	Classification   string               `json:"classification,omitempty"`   // An organization category, e.g. committee
	ParentId         string               `json:"parent_id,omitempty"`        //The ID of the organization that contains this organization
	Parent           *popolo.Organization `json:"parent,omitempty"`           // The organization that contains this organization
	AreaId           string               `json:"area_id,omitempty"`          // The ID of the geographic area to which this organization is related
	Area             *popolo.Area         `json:"area,omitempty"`             // The geographic area to which this organization is related
	FoundingDate     string               `json:"founding_date,omitempty"`    // A date of founding
	DissoulutionDate string               `json:"dissolution_date,omitempty"` // A date of dissolution
	Image            string               `json:"image,omitempty"`            // A URL of a head shot
	ContactDetails   []ContactDetail      `json:"contact_details,omitempty"`  // Means of contacting the person
	Links            []Link               `json:"link,omitempty"`             // URLs to documents about the person
	Memberships      []popolo.Membership  `json:"memberships,omitempty"`      // Memberships
	Posts            []popolo.Post        `json:"posts,omitempty"`            // Posts within the organization
	CreatedAt        time.Time            `json:"created_at,omitempty"`       // The time at which the resource was created
	UpdatedAt        time.Time            `json:"updated_at,omitempty"`       // The time at which the resource was last modified
	Sources          []Source             `json:"sources,omitempty"`          // URLs to documents from which the person is derived
}

type Company popolo.Organization

type Quota struct {
	Company         string
	Date            time.Time
	Parliamentarian string
	Order           string
	Value           float64

	PassengerName string
	Route         string
	Ticket        string
}

type Membership struct {
	Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Member       Rel           `json:"member"`
	Organization Rel           `json:"organization"`
	Label        string        `json:"label"`
	Role         string        `json:"role"`
	Source       Source        `json:"source"`
}

type Rel struct {
	Id   string `json:"id"`
	Link string `json:"link"`
}

type Source struct {
	Url  string `json:"url"`  // A URL
	Note string `json:"note"` // A note, e.g. 'Wikipedia page'
}

type Identifier struct {
	Identifier string `json:"identifier"` // An issued identifier, e.g. a DUNS number
	Scheme     string `json:"scheme"`     // An identifier scheme, e.g. DUNS
}

type ContactDetail struct {
	Label      string       `json:"label,omitempty"`       // A human-readable label for the contact detail
	Type       string       `json:"type"`                  //  [required] A type of medium, e.g. 'fax' or 'email'
	Value      string       `json:"value"`                 // [required] A value, e.g. a phone number or email address
	Note       string       `json:"note,omitempty"`        // A note, e.g. for grouping contact details by physical location
	ValidFrom  *popolo.Date `json:"valid_from,omitempty"`  // The date from which the contact detail is valid",
	ValidUntil *popolo.Date `json:"valid_until,omitempty"` // The date from which the contact detail is no longer valid",
	CreatedAt  *time.Time   `json:"created_at,omitempty"`  // The time at which the resource was created
	UpdatedAt  *time.Time   `json:"updated_at,omitempty"`  // The time at which the resource was last modified
	Sources    []Source     `json:"sources, omitempty"`    // URLs to documents from which the person is derived
}
