CREATE TABLE IF NOT EXISTS public.course (
    course_id SERIAL PRIMARY KEY,
    course_name VARCHAR(50)
);

INSERT INTO public.course(course_name)
VALUES ('crs1'), ('crs2'), ('crs3'), ('crs4')