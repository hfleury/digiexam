CREATE TABLE IF NOT EXISTS public.student (
    student_id SERIAL PRIMARY KEY,
    student_name VARCHAR(80) NOT NULL
);

INSERT INTO public.student(student_name) 
VALUES('std1'),('std2'),('std3'),('std4');