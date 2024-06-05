INSERT INTO "notifications" ("company_id", "date", "viewed", "type")
VALUES
    (3, '2023-05-05', TRUE, 'new_lead'),
    (1, '2023-06-06', TRUE, 'new_lead'),
    (1, '2023-07-07', TRUE, 'new_lead'),
    (3, '2023-08-08', TRUE, 'new_lead'),
    (1, '2023-09-09', FALSE, 'new_lead'),
    (2, '2023-10-10', TRUE, 'new_lead'),
    (3, '2023-11-11', TRUE, 'closed_lead'),
    (2, '2023-12-12', FALSE, 'closed_lead');


INSERT INTO "new_lead_notifications" (id, lead_id, client_company)
VALUES
    (1, 1, 2),
    (2, 2, 2),
    (3, 3, 3),
    (4, 4, 1),
    (5, 5, 2),
    (6, 6, 1);

INSERT INTO "closed_lead_notifications" (id, ad_id, producer_company, answered)
VALUES
    (7, 2, 1, TRUE),
    (8, 4, 2, FALSE);
