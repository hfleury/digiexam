CREATE TABLE IF NOT EXISTS public.grade_scale (
    scale_id SERIAL PRIMARY KEY,
    scale_type VARCHAR(15) NOT NULL
);

INSERT INTO public.grade_scale(scale_type)
VALUES ('letter'), ('twentypoint');