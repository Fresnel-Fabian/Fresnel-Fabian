package school

import (
    "sort"
)

// Grade represents a grade and the students in it.
type Grade struct {
    Grade    int
    Students []string
}

type School struct {
    roster map[int][]string
    seen   map[string]bool
}

func New() *School {
    return &School{
        roster: make(map[int][]string),
        seen:   make(map[string]bool),
    }
}

func (s *School) Add(student string, grade int) {
    if s.seen[student] {
        return
    }

    s.roster[grade] = append(s.roster[grade], student)
    s.seen[student] = true

    sort.Strings(s.roster[grade])
}

func (s *School) Grade(level int) []string {
    students := s.roster[level]
    out := make([]string, len(students))
    copy(out, students)
    return out
}

func (s *School) Enrollment() []Grade {
    var result []Grade

    // sort grade keys
    grades := make([]int, 0, len(s.roster))
    for g := range s.roster {
        grades = append(grades, g)
    }
    sort.Ints(grades)

    // build Grade structs
    for _, g := range grades {
        students := make([]string, len(s.roster[g]))
        copy(students, s.roster[g])

        result = append(result, Grade{
            Grade:    g,
            Students: students,
        })
    }

    return result
}
