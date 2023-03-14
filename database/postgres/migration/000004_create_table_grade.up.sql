CREATE TABLE IF NOT EXISTS public.grade (
    grade_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    scale_id INT NOT NULL,
    grade_grade VARCHAR(3) NOT NULL,
    grade_min VARCHAR(3) NOT NULL,
    grade_gpa INT NOT NULL,
    CONSTRAINT fk_student FOREIGN KEY(student_id) REFERENCES student(student_id),
    CONSTRAINT fk_course FOREIGN KEY(course_id) REFERENCES course(course_id),
    CONSTRAINT fk_scale FOREIGN KEY(scale_id) REFERENCES grade_scale(scale_id)
);

INSERT INTO public.grade (student_id, course_id, scale_id, grade_grade, grade_min, grade_gpa)
VALUES  (2, 1, 1, 'A', 90, 4), (2, 2, 1, 'B', 80, 3), (2, 3, 1, 'C', 70, 2), (2, 4, 1, 'D', 60, 1), (2, 4, 1, 'F', 0, 0), 
        (3, 1, 2, 20, 90, 4), (3, 2, 2, 17, 80, 3), (3, 3, 2, 14, 70, 2), (3, 4, 2, 12, 60, 1), (3, 4, 2, 9, 60, 0);