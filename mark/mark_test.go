package mark

import "testing"

func TestSpecifyMark(t *testing.T) {

	mark1 := SpecifyMark(0)
	if mark1 != CORRECT {
		t.Errorf("If difference is 0 then mark should be CORRECT(1), %d is", mark1)
	}

	mark2 := SpecifyMark(1)
	if mark2 != PROPER {
		t.Errorf("If difference is 1 then mark should be PROPER(2), %d is", mark2)
	}

	mark3 := SpecifyMark(2)
	if mark3 != SIMILAR {
		t.Errorf("If difference is 2 then mark should be SIMILAR(2), %d is", mark3)
	}

	mark4 := SpecifyMark(3)
	if mark4 != WRONG {
		t.Errorf("If difference is 3 or more then mark should be WRONG(4), %d is", mark4)
	}

	mark5 := SpecifyMark(5)
	if mark5 != WRONG {
		t.Errorf("If difference is 5 or more then mark should be WRONG(4), %d is", mark5)
	}
}

func TestMark_UpdateMark(t *testing.T) {

	mark1 := InitMark()
	mark1.UpdateMark(0)
	if mark1.value != CORRECT {
		t.Errorf("Init mark updated by 0 difference should be CORRECT but is: %s", mark1.NameMark())
	}
	if mark1.Pass() != true {
		t.Errorf("CORRECT mark should be pass but is not")
	}

	mark2 := InitMark()
	mark2.UpdateMark(1)
	if mark2.value != PROPER {
		t.Errorf("Init mark updated by 1 difference should be PROPER but is: %s", mark2.NameMark())
	}

	mark3 := InitMark()
	mark3.UpdateMark(2)
	if mark3.value != SIMILAR {
		t.Errorf("Init mark updated by 2 difference should be SIMILAR but is: %s", mark3.NameMark())
	}

	mark4 := InitMark()
	mark4.UpdateMark(3)
	if mark4.value != WRONG {
		t.Errorf("Init mark updated by 3 difference should be WRONG but is: %s", mark4.NameMark())
	}

	mark5 := InitMark()
	mark5.UpdateMark(6)
	if mark5.value != WRONG {
		t.Errorf("Init mark updated by 6 difference should be WRONG but is: %s", mark5.NameMark())
	}

	mark6 := InitMark()
	mark6.UpdateMark(2)
	mark6.UpdateMark(6)
	if mark6.value != SIMILAR {
		t.Errorf("SIMILAR mark updated with WRONG state should stay SIMILAR but is: %s", mark6.NameMark())
	}

	mark7 := InitMark()
	mark7.UpdateMark(5)
	mark7.UpdateMark(1)
	if mark7.value != PROPER {
		t.Errorf("WRONG mark updated with PROPER state should change to PROPER but is: %s", mark7.NameMark())
	}
	if mark7.Pass() != false {
		t.Errorf("PROPER mark should not be pass but it is")
	}
	if mark7.Fail() != true {
		t.Errorf("PROPER mark should be fail but is not")
	}
}
