package esepunittests

import "testing"

func TestNewGradeCalculator(t *testing.T) {
	gradeCalculator := NewGradeCalculator()
	if gradeCalculator == nil {
		t.Error("Expected NewGradeCalculator to return a valid instance")
	}

}

func TestAddGrade(t *testing.T) {
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("assignment", 85, Assignment)
	if gradeCalculator.assignments[0].Name != "assignment" || gradeCalculator.assignments[0].Grade != 85 || gradeCalculator.assignments[0].Type != Assignment {
		t.Error("Expected AddGrade to add the grade correctly")
	}
}

func TestGradeTypeString(t *testing.T) {
	cases := map[GradeType]string{
		Assignment: "assignment",
		Exam:       "exam",
		Essay:      "essay",
	}

	for gradeType, expected_value := range cases {
		actual_value := gradeType.String()
		if expected_value != actual_value {
			t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value, actual_value)
		}
	}
}

func TestWeights(t *testing.T) {
	assignment_expected_value := 50
	exam_expected_value := 35
	essay_expected_value := 15

	assignmentWeight := NewGradeCalculator()
	examWeight := NewGradeCalculator()
	essayWeight := NewGradeCalculator()

	assignmentWeight.AddGrade("assignment", 100, Assignment)
	assignmentWeight.AddGrade("exam", 0, Exam)
	assignmentWeight.AddGrade("essay", 0, Essay)

	examWeight.AddGrade("assignment", 0, Assignment)
	examWeight.AddGrade("exam", 100, Exam)
	examWeight.AddGrade("essay", 0, Essay)

	essayWeight.AddGrade("assignment", 0, Assignment)
	essayWeight.AddGrade("exam", 0, Exam)
	essayWeight.AddGrade("essay", 100, Essay)

	if assignmentWeight.calculateNumericalGrade() != assignment_expected_value {
		t.Errorf("Expected GetGrade to return '%d'; got '%d' instead", assignment_expected_value, assignmentWeight.calculateNumericalGrade())
	}
	if examWeight.calculateNumericalGrade() != exam_expected_value {
		t.Errorf("Expected GetGrade to return '%d'; got '%d' instead", exam_expected_value, examWeight.calculateNumericalGrade())
	}
	if essayWeight.calculateNumericalGrade() != essay_expected_value {
		t.Errorf("Expected GetGrade to return '%d'; got '%d' instead", essay_expected_value, essayWeight.calculateNumericalGrade())
	}
}

func TestLowA(t *testing.T) {
	expected_value := A

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 90, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 90, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestHighB(t *testing.T) {
	expected_value := B

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 89, Assignment)
	gradeCalculator.AddGrade("exam 1", 90, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 90, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
func TestLowB(t *testing.T) {
	expected_value := B

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 80, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 80, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestHighC(t *testing.T) {
	expected_value := C

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 79, Assignment)
	gradeCalculator.AddGrade("exam 1", 80, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 80, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestLowC(t *testing.T) {
	expected_value := C

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestHighD(t *testing.T) {
	expected_value := D

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 69, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestLowD(t *testing.T) {
	expected_value := D

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestHighF(t *testing.T) {
	expected_value := F

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 59, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestWeight(t *testing.T) {
	expected_value := F

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 59, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
