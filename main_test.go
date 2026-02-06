package main // Change to match your package name

import (
	"testing"
	"time"
)

func TestEmailBody(t *testing.T) {
	tests := []struct {
		name            string
		birthdays       []Birthday
		expectedSubject string
		expectedBody    string
	}{
		{
			name:            "empty slice returns empty strings",
			birthdays:       []Birthday{},
			expectedSubject: "",
			expectedBody:    "",
		},
		{
			name: "single birthday",
			birthdays: []Birthday{
				{Name: "Alice", Birthdate: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)},
			},
			expectedSubject: "Happy 35th Birthday to Alice!", // Adjust age based on current date
			expectedBody:    "Happy Birthday to Alice! Hope you have a wonderful day!",
		},
		{
			name: "two birthdays",
			birthdays: []Birthday{
				{Name: "Alice", Birthdate: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)},
				{Name: "Bob", Birthdate: time.Date(1985, 8, 20, 0, 0, 0, 0, time.UTC)},
			},
			expectedSubject: "Happy Birthday to Alice and Bob!",
			expectedBody:    "Happy Birthday to Alice (35), Bob (40)! Hope you all have a wonderful day!",
		},
		{
			name: "three birthdays",
			birthdays: []Birthday{
				{Name: "Alice", Birthdate: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)},
				{Name: "Bob", Birthdate: time.Date(1985, 8, 20, 0, 0, 0, 0, time.UTC)},
				{Name: "Charlie", Birthdate: time.Date(1995, 12, 10, 0, 0, 0, 0, time.UTC)},
			},
			expectedSubject: "Happy Birthday to Alice, Bob and Charlie!",
			expectedBody:    "Happy Birthday to Alice (35), Bob (40), Charlie (30)! Hope you all have a wonderful day!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subject, body := emailBody(tt.birthdays)

			if subject != tt.expectedSubject {
				t.Errorf("subject mismatch:\ngot:  %q\nwant: %q", subject, tt.expectedSubject)
			}

			if body != tt.expectedBody {
				t.Errorf("body mismatch:\ngot:  %q\nwant: %q", body, tt.expectedBody)
			}
		})
	}
}

// Helper test to verify the age calculation doesn't cause issues
func TestEmailBodyWithFixedDate(t *testing.T) {
	// Using a fixed birthdate that's easier to calculate
	birthday := Birthday{
		Name:      "Test User",
		Birthdate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	subject, body := emailBody([]Birthday{birthday})

	// The age will depend on when you run this test
	// But we can at least verify the format is correct
	if subject == "" {
		t.Error("subject should not be empty for single birthday")
	}

	if body != "Happy Birthday to Test User! Hope you have a wonderful day!" {
		t.Errorf("unexpected body: %q", body)
	}
}
