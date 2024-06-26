ALTER TABLE "ads"
ALTER COLUMN "image_url" TYPE VARCHAR(500);

INSERT INTO "ads" (company_id, title, text, industry, price, image_url, creation_date, created_by, responsible, is_deleted)
VALUES
    (3, 'Облачное хранилище', '100ТБ хранилища на 1 год', 1, 400000, '', '2023-01-01', 9, 9, false),
    (1, 'Стандартная перевозка', 'До 250 кг, цена за 100 км пути', 3, 25000, '', '2023-02-02', 1, 2, false),
    (1, 'Перевозка чувствительного оборудования', 'До 1 т, цена за 100 км пути', 3, 150000, '', '2023-03-03', 4, 4, false),
    (2, 'Организация деловой встречи', 'Банкетный зал, кухня на выбор. От 75 человек', 5, 450000, '', '2023-04-04', 8, 8, false);

INSERT INTO "responses" (company_id, employee_id, ad_id, creation_date)
VALUES
    (2, 6, 1, '2023-05-05'),
    (2, 7, 2, '2023-06-06'),
    (3, 10, 2, '2023-07-07'),
    (1, 5, 1, '2023-08-08'),
    (2, 7, 3, '2023-09-09'),
    (1, 1, 4, '2023-10-10');
