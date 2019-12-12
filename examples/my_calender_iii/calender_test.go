package my_calender_iii

import "testing"

func TestMyCalendarThree_Book2(t *testing.T) {
	calender := Constructor()
	if expected, actual := 1, calender.Book(10, 20); actual != expected {
		t.Errorf("Step 1 failed, expected = %d, actual = %d", expected, actual)
	}
	if expected, actual := 2, calender.Book(10, 20); actual != expected {
		t.Errorf("Step 2 failed, expected = %d, actual = %d", expected, actual)
	}
	if expected, actual := 3, calender.Book(10, 20); actual != expected {
		t.Errorf("Step 3 failed, expected = %d, actual = %d", expected, actual)
	}
}

func TestMyCalendarThree_Book1(t *testing.T) {
	calender := Constructor()
	if expected, actual := 1, calender.Book(10, 20); actual != expected {
		t.Errorf("Step 1 failed, expected = %d, actual = %d", expected, actual)
	}
	if expected, actual := 1, calender.Book(50, 60); actual != expected {
		t.Errorf("Step 2 failed, expected = %d, actual = %d", expected, actual)
	}
	if expected, actual := 2, calender.Book(10, 40); actual != expected {
		t.Errorf("Step 3 failed, expected = %d, actual = %d", expected, actual)
	}
	if expected, actual := 3, calender.Book(5, 15); actual != expected {
		t.Errorf("Step 4 failed, expected = %d, actual = %d", expected, actual)
	}
	if expected, actual := 3, calender.Book(5, 10); actual != expected {
		t.Errorf("Step 5 failed, expected = %d, actual = %d", expected, actual)
	}
	if expected, actual := 3, calender.Book(25, 55); actual != expected {
		t.Errorf("Step 6 failed, expected = %d, actual = %d", expected, actual)
	}
}
