CREATE TABLE IF NOT EXISTS public.grade (
    grade_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    scale_id INT NOT NULL,
    grade_grade VARCHAR(3) NOT NULL, 
    CONSTRAINT fk_student FOREIGN KEY(student_id) REFERENCES student(student_id),
    CONSTRAINT fk_course FOREIGN KEY(course_id) REFERENCES course(course_id),
    CONSTRAINT fk_scale FOREIGN KEY(scale_id) REFERENCES grade_scale(scale_id)
);

INSERT INTO public.grade (student_id, course_id, scale_id, grade_grade)
VALUES (1, 1, 1, 50), (1, 2, 1, 80), (1, 3, 1, 10), (1, 4, 1, 100), 
        (2, 1, 2, 1), (2, 2, 2, 3), (2, 3, 2, 2), (2, 4, 2, 5), 
        (3, 1, 3, 20), (3, 2, 3, 17), (3, 3, 3, 12), (3, 4, 3, 11);